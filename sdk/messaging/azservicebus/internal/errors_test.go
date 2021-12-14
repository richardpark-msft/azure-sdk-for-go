// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

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

func Test_recoveryKind(t *testing.T) {
	ctx := context.Background()

	t.Run("link", func(t *testing.T) {
		linkErrorCodes := []string{
			string(amqp.ErrorDetachForced),
		}

		for _, code := range linkErrorCodes {
			t.Run(code, func(t *testing.T) {
				sbe := ToSBE(ctx, &amqp.Error{Condition: amqp.ErrorCondition(code)})
				require.EqualValues(t, RecoveryKindLink, sbe.RecoveryKind, fmt.Sprintf("requires link recovery: %s", code))
			})
		}

		t.Run("sentintel errors", func(t *testing.T) {
			sbe := ToSBE(ctx, amqp.ErrLinkClosed)
			require.EqualValues(t, RecoveryKindLink, sbe.RecoveryKind)

			sbe = ToSBE(ctx, amqp.ErrSessionClosed)
			require.EqualValues(t, RecoveryKindLink, sbe.RecoveryKind)
		})
	})

	t.Run("connection", func(t *testing.T) {
		codes := []string{
			string(amqp.ErrorConnectionForced),
		}

		for _, code := range codes {
			t.Run(code, func(t *testing.T) {
				sbe := ToSBE(ctx, &amqp.Error{Condition: amqp.ErrorCondition(code)})
				require.EqualValues(t, RecoveryKindConn, sbe.RecoveryKind, fmt.Sprintf("requires connection recovery: %s", code))
			})
		}

		t.Run("sentinel errors", func(t *testing.T) {
			sbe := ToSBE(ctx, amqp.ErrConnClosed)
			require.EqualValues(t, RecoveryKindConn, sbe.RecoveryKind)
		})
	})

	t.Run("nonretriable", func(t *testing.T) {
		codes := []string{
			string(amqp.ErrorTransferLimitExceeded),
			string(amqp.ErrorInternalError),
			string(amqp.ErrorUnauthorizedAccess),
			string(amqp.ErrorNotFound),
			string(amqp.ErrorMessageSizeExceeded),
		}

		for _, code := range codes {
			t.Run(code, func(t *testing.T) {
				sbe := ToSBE(ctx, &amqp.Error{Condition: amqp.ErrorCondition(code)})
				require.EqualValues(t, RecoveryKindNonRetriable, sbe.RecoveryKind, fmt.Sprintf("cannot be recovered: %s", code))
			})
		}
	})

	t.Run("none", func(t *testing.T) {
		codes := []string{
			string(errorOperationCancelled),
			string(errorServerBusy),
			string(errorTimeout),
		}

		for _, code := range codes {
			t.Run(code, func(t *testing.T) {
				sbe := ToSBE(ctx, &amqp.Error{Condition: amqp.ErrorCondition(code)})
				require.EqualValues(t, RecoveryKindNone, sbe.RecoveryKind, fmt.Sprintf("no recovery needed: %s", code))
			})
		}
	})
}

func Test_recoveryKinds(t *testing.T) {
	ctx := context.Background()

	sbe := ToSBE(ctx, nil)
	require.Nil(t, sbe)

	// link recoveries
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, amqp.ErrLinkClosed).RecoveryKind)
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, amqp.ErrSessionClosed).RecoveryKind)
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, &amqp.DetachError{}).RecoveryKind)
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, fmt.Errorf("wrapped: %w", amqp.ErrLinkClosed)).RecoveryKind)
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, fmt.Errorf("wrapped: %w", amqp.ErrSessionClosed)).RecoveryKind)
	require.EqualValues(t, RecoveryKindLink, ToSBE(ctx, fmt.Errorf("wrapped: %w", &amqp.DetachError{})).RecoveryKind)

	// going to treat these as "connection troubles" and throw them into the
	// connection recovery scenario instead.
	require.EqualValues(t, RecoveryKindConn, ToSBE(ctx, &permanentNetError{}).RecoveryKind)
	require.EqualValues(t, RecoveryKindConn, ToSBE(ctx, fmt.Errorf("%w", &permanentNetError{})).RecoveryKind)
	require.EqualValues(t, RecoveryKindConn, ToSBE(ctx, amqp.ErrConnClosed).RecoveryKind)
	require.EqualValues(t, RecoveryKindConn, ToSBE(ctx, fmt.Errorf("%w", amqp.ErrConnClosed)).RecoveryKind)
}

func Test_IsNonRetriable(t *testing.T) {
	ctx := context.Background()

	errs := []error{
		context.Canceled,
		context.DeadlineExceeded,
		ErrNonRetriable{Message: "any message"},
		fmt.Errorf("wrapped: %w", context.Canceled),
		fmt.Errorf("wrapped: %w", context.DeadlineExceeded),
		fmt.Errorf("wrapped: %w", ErrNonRetriable{Message: "any message"}),
	}

	for _, err := range errs {
		require.EqualValues(t, RecoveryKindNonRetriable, ToSBE(ctx, err).RecoveryKind)
	}
}
