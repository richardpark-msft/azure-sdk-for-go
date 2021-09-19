// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import "fmt"

// implements `internal/errorinfo/NonRetriable`
type errClosed struct {
	link string
}

func (ec errClosed) NonRetriable() {}
func (ec errClosed) Error() string {
	return fmt.Sprintf("%s is closed and can no longer be used", ec.link)
}

// MessageErrorCondition represents a well-known collection of AMQP errors
type MessageErrorCondition string

// Error Conditions
const (
	ErrorInternalError         MessageErrorCondition = "amqp:internal-error"
	ErrorNotFound              MessageErrorCondition = "amqp:not-found"
	ErrorUnauthorizedAccess    MessageErrorCondition = "amqp:unauthorized-access"
	ErrorDecodeError           MessageErrorCondition = "amqp:decode-error"
	ErrorResourceLimitExceeded MessageErrorCondition = "amqp:resource-limit-exceeded"
	ErrorNotAllowed            MessageErrorCondition = "amqp:not-allowed"
	ErrorInvalidField          MessageErrorCondition = "amqp:invalid-field"
	ErrorNotImplemented        MessageErrorCondition = "amqp:not-implemented"
	ErrorResourceLocked        MessageErrorCondition = "amqp:resource-locked"
	ErrorPreconditionFailed    MessageErrorCondition = "amqp:precondition-failed"
	ErrorResourceDeleted       MessageErrorCondition = "amqp:resource-deleted"
	ErrorIllegalState          MessageErrorCondition = "amqp:illegal-state"
)
