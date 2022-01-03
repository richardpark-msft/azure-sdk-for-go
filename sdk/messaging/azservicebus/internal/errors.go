// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/internal/rpc"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

type NonRetriable interface {
	error
	NonRetriable()
}

// IsNonRetriable indicates an error is fatal. Typically, this means
// the connection or link has been closed.
func IsNonRetriable(err error) bool {
	var nonRetriable NonRetriable

	if errors.As(err, &nonRetriable) {
		return true
	}

	return IsCancelError(err)
}

var sberrClosedPermanently = &ServiceBusError{
	inner:        errors.New("link has been closed permanently"),
	RecoveryKind: recoveryKindNonRetriable,
}

var sberrNoLock = &ServiceBusError{
	inner:        errors.New("lock token on the message is not set, thus cannot send disposition"),
	RecoveryKind: recoveryKindNonRetriable,
}

// Error Conditions
const (
	// Service Bus Errors
	errorServerBusy         amqp.ErrorCondition = "com.microsoft:server-busy"
	errorTimeout            amqp.ErrorCondition = "com.microsoft:timeout"
	errorOperationCancelled amqp.ErrorCondition = "com.microsoft:operation-cancelled"
	errorContainerClose     amqp.ErrorCondition = "com.microsoft:container-close"
)

type recoveryKind string

const recoveryKindNone recoveryKind = ""
const recoveryKindNonRetriable recoveryKind = "fatal"
const recoveryKindLink recoveryKind = "link"
const recoveryKindConnection recoveryKind = "connection"

type ServiceBusError struct {
	inner        error
	RecoveryKind recoveryKind
}

func (sbe *ServiceBusError) String() string {
	return sbe.inner.Error()
}

func (sbe *ServiceBusError) AsError() error {
	return sbe.inner
}

// ToSBE wraps the passed in 'err' with a proper error with one of either:
// - `fatalServiceBusError` if no recovery is possible.
// - `serviceBusError` if the error is recoverable. The `recoveryKind` field contains the
//   type of recovery needed.
func ToSBE(ctx context.Context, err error) *ServiceBusError {
	if err == nil {
		return nil
	}

	sbe := &ServiceBusError{
		inner:        err,
		RecoveryKind: getRecoveryKind(ctx, err),
	}

	return sbe
}

func isDetachError(err error) bool {
	// NOTE: checking of an amqp.Error condition is done
	// by
	if errors.Is(err, amqp.ErrLinkDetached) {
		return true
	}

	var de *amqp.DetachError
	return errors.As(err, &de)
}

func isPermanentNetError(err error) bool {
	var netErr net.Error

	if errors.As(err, &netErr) {
		temp := netErr.Temporary()
		timeout := netErr.Timeout()
		return !temp && !timeout
	}

	return false
}

func isEOF(err error) bool {
	return errors.Is(err, io.EOF)
}

func IsCancelError(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return true
	}

	if err.Error() == "context canceled" { // go-amqp is returning this when I cancel
		return true
	}

	return false
}

func IsDrainingError(err error) bool {
	// TODO: we should be able to identify these errors programatically
	return strings.Contains(err.Error(), "link is currently draining")
}

var amqpConditionsToRecoveryKind = map[amqp.ErrorCondition]recoveryKind{
	// no recovery needed, these are temporary errors.
	amqp.ErrorInternalError:                                     recoveryKindNone, // "amqp:internal-error"
	amqp.ErrorCondition("com.microsoft:server-busy"):            recoveryKindNone,
	amqp.ErrorCondition("com.microsoft:timeout"):                recoveryKindNone,
	amqp.ErrorCondition("com.microsoft:operation-cancelled"):    recoveryKindNone,
	amqp.ErrorCondition("client.sender:not-enough-link-credit"): recoveryKindNone,
	amqp.ErrorTransferLimitExceeded:                             recoveryKindNone, // "amqp:link:transfer-limit-exceeded"

	// Link recovery needed
	amqp.ErrorDetachForced: recoveryKindLink, // "amqp:link:detach-forced"

	// Connection recovery needed
	amqp.ErrorConnectionForced: recoveryKindConnection, // "amqp:connection:forced"

	// No recovery possible - this operation is non retriable.
	amqp.ErrorCondition("com.microsoft:session-cannot-be-locked"): recoveryKindNonRetriable,
	amqp.ErrorCondition("com.microsoft:message-lock-lost"):        recoveryKindNonRetriable,
	amqp.ErrorUnauthorizedAccess:                                  recoveryKindNonRetriable, // creds are bad
	amqp.ErrorNotFound:                                            recoveryKindNonRetriable,
	amqp.ErrorNotAllowed:                                          recoveryKindNonRetriable,
}

func getRecoveryKind(ctxForLogging context.Context, err error) recoveryKind {
	if IsCancelError(err) {
		return recoveryKindNonRetriable
	}

	if isPermanentNetError(err) || isEOF(err) {
		return recoveryKindConnection
	}

	// check the "special" AMQP errors that aren't condition-based.
	if errors.Is(err, amqp.ErrSessionClosed) ||
		errors.Is(err, amqp.ErrLinkClosed) ||
		isDetachError(err) {
		return recoveryKindLink
	}

	if errors.Is(err, amqp.ErrConnClosed) {
		return recoveryKindConnection
	}

	if IsDrainingError(err) {
		// temporary, operation should just be retryable since drain will
		// eventually complete.
		return recoveryKindNone
	}

	// then it's _probably_ an actual *amqp.Error, in which case we bucket it by
	// the 'condition'.
	var amqpError *amqp.Error

	if errors.As(err, &amqpError) {
		recoveryKind, ok := amqpConditionsToRecoveryKind[amqpError.Condition]

		if ok {
			return recoveryKind
		}
	}

	// this is some error type we've never seen.
	tab.For(ctxForLogging).Fatal(fmt.Sprintf("No recovery possible with error: %#v", err))
	return recoveryKindNonRetriable
}

const (
	amqpRetryDefaultTimes int           = 3
	amqpRetryDefaultDelay time.Duration = time.Second
)

type (
	// ErrMissingField indicates that an expected property was missing from an AMQP message. This should only be
	// encountered when there is an error with this library, or the server has altered its behavior unexpectedly.
	ErrMissingField string

	// ErrMalformedMessage indicates that a message was expected in the form of []byte was not a []byte. This is likely
	// a bug and should be reported.
	ErrMalformedMessage string

	// ErrIncorrectType indicates that type assertion failed. This should only be encountered when there is an error
	// with this library, or the server has altered its behavior unexpectedly.
	ErrIncorrectType struct {
		Key          string
		ExpectedType reflect.Type
		ActualValue  interface{}
	}

	// ErrAMQP indicates that the server communicated an AMQP error with a particular
	ErrAMQP rpc.Response

	// ErrNoMessages is returned when an operation returned no messages. It is not indicative that there will not be
	// more messages in the future.
	ErrNoMessages struct{}

	// ErrNotFound is returned when an entity is not found (404)
	ErrNotFound struct {
		EntityPath string
	}

	// ErrConnectionClosed indicates that the connection has been closed.
	ErrConnectionClosed string
)

func (e ErrMissingField) Error() string {
	return fmt.Sprintf("missing value %q", string(e))
}

func (e ErrMalformedMessage) Error() string {
	return "message was expected in the form of []byte was not a []byte"
}

// NewErrIncorrectType lets you skip using the `reflect` package. Just provide a variable of the desired type as
// 'expected'.
func NewErrIncorrectType(key string, expected, actual interface{}) ErrIncorrectType {
	return ErrIncorrectType{
		Key:          key,
		ExpectedType: reflect.TypeOf(expected),
		ActualValue:  actual,
	}
}

func (e ErrIncorrectType) Error() string {
	return fmt.Sprintf(
		"value at %q was expected to be of type %q but was actually of type %q",
		e.Key,
		e.ExpectedType,
		reflect.TypeOf(e.ActualValue))
}

func (e ErrAMQP) Error() string {
	return fmt.Sprintf("server says (%d) %s", e.Code, e.Description)
}

func (e ErrNoMessages) Error() string {
	return "no messages available"
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("entity at %s not found", e.EntityPath)
}

// IsErrNotFound returns true if the error argument is an ErrNotFound type
func IsErrNotFound(err error) bool {
	_, ok := err.(ErrNotFound)
	return ok
}

func (e ErrConnectionClosed) Error() string {
	return fmt.Sprintf("the connection has been closed: %s", string(e))
}
