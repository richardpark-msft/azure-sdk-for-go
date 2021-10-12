// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package sberrors

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"

	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// Error Conditions
const (
	// Service Bus Errors
	errorServerBusy         amqp.ErrorCondition = "com.microsoft:server-busy"
	errorTimeout            amqp.ErrorCondition = "com.microsoft:timeout"
	errorOperationCancelled amqp.ErrorCondition = "com.microsoft:operation-cancelled"
	errorContainerClose     amqp.ErrorCondition = "com.microsoft:container-close"
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

// Leveraging @serbrech's fine work from go-shuttle:
// https://github.com/Azure/go-shuttle/blob/ea882947109ade9b34d4d69642fdf7aec4570fee/common/errorhandling/recovery.go

var retryableAMQPConditions = map[string]bool{
	string(amqp.ErrorInternalError):         true,
	string(errorServerBusy):                 true, // "com.microsoft:server-busy"
	string(errorTimeout):                    true, // "com.microsoft:timeout"
	string(errorOperationCancelled):         true, // "com.microsoft:operation-cancelled"
	"client.sender:not-enough-link-credit":  true,
	string(amqp.ErrorUnauthorizedAccess):    true,
	string(amqp.ErrorDetachForced):          true,
	string(amqp.ErrorConnectionForced):      true,
	string(amqp.ErrorTransferLimitExceeded): true,
	"amqp: connection closed":               true,
	"unexpected frame":                      true,
	string(amqp.ErrorNotFound):              true,
}

func isRetryableAMQPError(ctxForLogging context.Context, err error) bool {
	var amqpErr *amqp.Error
	var isAMQPError = errors.As(err, &amqpErr)

	if isAMQPError {
		_, ok := retryableAMQPConditions[string(amqpErr.Condition)]
		return ok
	}

	// TODO: there is a bug somewhere that seems to be errorString'ing errors. Need to track that down.
	// In the meantime, try string matching instead
	for condition := range retryableAMQPConditions {
		if strings.Contains(err.Error(), condition) {
			tab.For(ctxForLogging).Error(fmt.Errorf("error needed to be matched by a string matcher, rather than by type: %w", err))
			return true
		}
	}

	return false
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

func shouldRecreateLink(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, amqp.ErrLinkDetached) ||
		// TODO: proper error types needs to happen
		strings.Contains(err.Error(), "detach frame link detached")
}

func shouldRecreateConnection(ctxForLogging context.Context, err error) bool {
	if err == nil {
		return false
	}

	shouldRecreate := isPermanentNetError(err) ||
		isRetryableAMQPError(ctxForLogging, err) ||
		isEOF(err) ||
		// these are distinct from a detach and probably indicate something
		// wrong with the connection itself, rather than just the link
		errors.Is(err, amqp.ErrSessionClosed) ||
		errors.Is(err, amqp.ErrLinkClosed)

	return shouldRecreate
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

// IsSessionLockedError checks to see if this is the "you tried to get a session that was already locked"
// error.
func IsSessionLockedError(err error) bool {
	var amqpError *amqp.Error

	if !errors.As(err, &amqpError) {
		return false
	}

	// Example:
	//{Condition: com.microsoft:session-cannot-be-locked, Description: The requested session 'session-1' cannot be accepted. It may be locked by another receiver.

	return amqpError.Condition == "com.microsoft:session-cannot-be-locked"
}

type FixBy int

const (
	// FixByRecoveringConnection being true means you need to recreate the connection before this
	// operation could succeed. This implies that NeedsNewLink is also true, and should be checked
	// first.
	FixByRecoveringConnection FixBy = 1

	// FixByRecoveringLink means you need to recreate the link before this operation could succeed.
	// (typically as a result of a detach)
	FixByRecoveringLink FixBy = 2

	// RecoverNotPossible means the operation can be retried without any form of recovery.
	FixByRetrying FixBy = 3

	// FixNotPossible means the operation should not be retried.
	FixNotPossible FixBy = 4
)

type ServiceBusError struct {
	Fix FixBy
	Err error
}

func NewServiceBusError(fix FixBy, err error) *ServiceBusError {
	if err == nil {
		panic("Nil error passed to ServiceBusError")
	}

	return &ServiceBusError{fix, err}
}

func (e *ServiceBusError) Error() string {
	return e.Err.Error()
}

func (e *ServiceBusError) Unwrap() error {
	return e.Err
}

func AsServiceBusError(ctxForLogging context.Context, err error) *ServiceBusError {
	asSBE, ok := err.(*ServiceBusError)

	if ok {
		return asSBE
	}

	if IsCancelError(err) {
		return NewServiceBusError(FixNotPossible, err)
	}

	/*
		The underlying call (doRPCWithRetry) from amqp-common-go stringizes the
		error that comes back in the response, so we can't access it programatically
		 https://github.com/Azure/azure-amqp-common-go/issues/59

		unhandled error link ef694da5-411b-4b3c-a586-4f060a564968: status code 410 and description: The lock supplied is invalid. Either the lock expired, or the message has already been removed from the queue. Reference:35361e4c-1403-487a-826f-69f21e1654b2, TrackingId:5598ae0d-6fab-4e40-9712-ac0fc411af12_B2, SystemTracker:riparkdev2:Queue:queue-16ac1f8180646f48, Timestamp:2021-10-08T17:50:01
	*/
	if strings.Contains(err.Error(), "status code 410 ") {
		return NewServiceBusError(
			FixNotPossible,
			err,
		)
	}

	// proper AMQP code, as received from go-amqp via the receiver.
	var amqpError *amqp.Error
	if errors.As(err, &amqpError) && amqpError.Condition == "com.microsoft:message-lock-lost" {
		return NewServiceBusError(
			FixNotPossible,
			err,
		)
	}

	if shouldRecreateConnection(ctxForLogging, err) {
		return NewServiceBusError(
			FixByRecoveringConnection,
			err,
		)
	}

	if shouldRecreateLink(err) {
		return NewServiceBusError(
			FixByRecoveringLink,
			err,
		)
	}

	// TODO: it'd be nice to get more crisp on this (for instance, if we get a throttling error)
	return NewServiceBusError(
		FixByRetrying,
		err,
	)
}

// returns true if the AMQP error is considered transient
func IsAMQPTransientError(ctx context.Context, err error) bool {
	// always retry on a detach error
	var amqpDetach *amqp.DetachError
	if errors.As(err, &amqpDetach) {
		return true
	}
	// for an AMQP error, only retry depending on the condition
	var amqpErr *amqp.Error
	if errors.As(err, &amqpErr) {
		switch amqpErr.Condition {
		case errorServerBusy, errorTimeout, errorOperationCancelled, errorContainerClose:
			return true
		default:
			tab.For(ctx).Debug(fmt.Sprintf("isAMQPTransientError: condition %s is not transient", amqpErr.Condition))
			return false
		}
	}
	tab.For(ctx).Debug(fmt.Sprintf("isAMQPTransientError: %T is not transient", err))
	return false
}
