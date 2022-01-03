// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/errorinfo"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrMissingField_Error(t *testing.T) {
	const fieldName = "fieldName"
	var subject ErrMissingField = fieldName
	var cast error = subject

	got := cast.Error()
	const want = `missing value "` + fieldName + `"`

	if got != want {
		t.Logf("\n\tgot: \t%q\n\twant:\t%q", got, want)
		t.Fail()
	}
}

func TestErrIncorrectType_Error(t *testing.T) {
	var a int
	var b map[string]interface{}
	var c *float64

	types := map[reflect.Type]interface{}{
		reflect.TypeOf(a): 7.0,
		reflect.TypeOf(b): map[string]string{},
		reflect.TypeOf(c): int(2),
	}

	const key = "myFieldName"
	for expected, actual := range types {
		actualType := reflect.TypeOf(actual)
		t.Run(fmt.Sprintf("%s-%s", expected, actualType), func(t *testing.T) {
			expectedMessage := fmt.Sprintf(
				"value at %q was expected to be of type %q but was actually of type %q",
				key,
				expected.String(),
				actualType.String())

			subject := ErrIncorrectType{
				Key:          key,
				ActualValue:  actual,
				ExpectedType: expected,
			}

			var cast error = subject

			got := cast.Error()
			if got != expectedMessage {
				t.Logf("\n\tgot: \t%q\n\twant:\t%q", got, expectedMessage)
				t.Fail()
			}
		})
	}
}

func TestErrNotFound_Error(t *testing.T) {
	err := ErrNotFound{EntityPath: "/foo/bar"}
	assert.Equal(t, "entity at /foo/bar not found", err.Error())
	assert.True(t, IsErrNotFound(err))

	otherErr := errors.New("foo")
	assert.False(t, IsErrNotFound(otherErr))
}

func Test_isPermanentNetError(t *testing.T) {
	require.False(t, isPermanentNetError(&permanentNetError{
		temp: true,
	}))

	require.False(t, isPermanentNetError(&permanentNetError{
		timeout: true,
	}))

	require.False(t, isPermanentNetError(errors.New("not a net error")))

	require.True(t, isPermanentNetError(&permanentNetError{}))
}

func Test_isRetryableAMQPError(t *testing.T) {
	ctx := context.Background()

	retryableCodes := []string{
		string(amqp.ErrorInternalError),
		string(errorServerBusy),
		string(errorTimeout),
		string(errorOperationCancelled),
		"client.sender:not-enough-link-credit",
		string(amqp.ErrorUnauthorizedAccess),
		string(amqp.ErrorDetachForced),
		string(amqp.ErrorConnectionForced),
		string(amqp.ErrorTransferLimitExceeded),
		"amqp: connection closed",
		"unexpected frame",
		string(amqp.ErrorNotFound),
	}

	for _, code := range retryableCodes {
		require.True(t, isRetryableAMQPError(ctx, &amqp.Error{
			Condition: amqp.ErrorCondition(code),
		}))

		// it works equally well if the error is just in the String().
		// Need to narrow this down some more to see where the errors
		// might not be getting converted properly.
		require.True(t, isRetryableAMQPError(ctx, errors.New(code)))
	}

	require.False(t, isRetryableAMQPError(ctx, errors.New("some non-amqp related error")))
}

func assertNonRetriable(t *testing.T, expectedMsg string, err error) {
	var nonRetriable errorinfo.NonRetriable
	require.True(t, errors.As(err, &nonRetriable), "fatal errors are not retriable")
	require.EqualValues(t, expectedMsg, nonRetriable.Error())
}

func assertRetriable(t *testing.T, err error) {
	var nonRetriable errorinfo.NonRetriable
	require.False(t, errors.As(err, &nonRetriable), "should be retriable (ie, no marker method)")
}

func TestServiceBusError_Nil(t *testing.T) {
	sbe := ToSBE(context.Background(), nil)
	require.Nil(t, sbe)
}

func TestServiceBusError_RecoveryLink(t *testing.T) {
	// going to treat these as "connection troubles" and throw them into the
	// connection recovery scenario instead.
	linkRecoveryErrors := []error{
		amqp.ErrLinkDetached,
		amqp.ErrLinkClosed,
		amqp.ErrLinkDetached,
	}

	for _, origErr := range linkRecoveryErrors {
		t.Run(fmt.Sprintf("error: %s", origErr.Error()), func(t *testing.T) {
			sbe := ToSBE(context.Background(), origErr)
			require.EqualValues(t, recoveryKindLink, sbe.RecoveryKind)

			sbe = ToSBE(context.Background(), fmt.Errorf("this error is wrapped: %w", origErr))
			require.EqualValues(t, recoveryKindLink, sbe.RecoveryKind)
		})
	}
}

func TestServiceBusError_RecoveryConnection(t *testing.T) {
	connectionErrors := []error{
		&permanentNetError{},
		&amqp.Error{Condition: amqp.ErrorConnectionForced},
		amqp.ErrConnClosed,
	}

	for _, origErr := range connectionErrors {
		sbe := ToSBE(context.Background(), origErr)
		require.EqualValues(t, recoveryKindConnection, sbe.RecoveryKind)

		// should work the same, even if the error is wrapped
		sbe = ToSBE(context.Background(), fmt.Errorf("this error is wrapped: %w", origErr))
		require.EqualValues(t, recoveryKindConnection, sbe.RecoveryKind)
	}
}

func TestServiceBusError_Fatal(t *testing.T) {
	fatalErrors := []error{
		&amqp.Error{Condition: amqp.ErrorUnauthorizedAccess, Description: "description"},
		&amqp.Error{Condition: amqp.ErrorNotFound, Description: "description"},
		&amqp.Error{Condition: amqp.ErrorNotAllowed, Description: "description"},
		&amqp.Error{Condition: "com.microsoft:session-cannot-be-locked", Description: "description"},
		&amqp.Error{Condition: "com.microsoft:message-lock-lost", Description: "description"},
	}

	for _, origErr := range fatalErrors {
		sbe := ToSBE(context.Background(), origErr)
		require.EqualValues(t, recoveryKindNonRetriable, sbe.RecoveryKind)
		assertNonRetriable(t, origErr.Error(), sbe.AsError())

		// should work the same, even if the error is wrapped
		sbe = ToSBE(context.Background(), fmt.Errorf("this error is wrapped: %w", origErr))
		require.EqualValues(t, recoveryKindNonRetriable, sbe.RecoveryKind)
		assertNonRetriable(t, "this error is wrapped: "+origErr.Error(), sbe.AsError())
	}
}

func Test_IsCancelError(t *testing.T) {
	require.False(t, IsCancelError(nil))
	require.False(t, IsCancelError(errors.New("not a cancel error")))

	require.True(t, IsCancelError(errors.New("context canceled")))

	require.True(t, IsCancelError(context.Canceled))
	require.True(t, IsCancelError(context.DeadlineExceeded))
	require.True(t, IsCancelError(fmt.Errorf("wrapped: %w", context.Canceled)))
	require.True(t, IsCancelError(fmt.Errorf("wrapped: %w", context.DeadlineExceeded)))
}

func Test_IsDetachError(t *testing.T) {
	require.True(t, isDetachError(&amqp.DetachError{}))
	require.True(t, isDetachError(amqp.ErrLinkDetached))
	require.False(t, isDetachError(amqp.ErrConnClosed))
}
