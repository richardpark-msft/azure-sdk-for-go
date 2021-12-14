// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

func TestRetrier(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		ctx := context.Background()

		called := 0

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			called++
			return nil
		}, nil, nil)

		require.Nil(t, err)
		require.EqualValues(t, 1, called)
	})

	t.Run("FailsThenSucceeds", func(t *testing.T) {
		ctx := context.Background()

		called := 0

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			called++

			if called == 1 {
				// first round always has a nil error
				// (mostly nobody will care and just no-op on recovery
				// if the error is nil)
				require.Nil(t, args.LastErr)
			} else {
				// subsequent calls should pass in the error from the
				// last failure (this makes it simple to do recovery at the
				// start of your function)
				var amqpErr *amqp.DetachError
				require.True(t, errors.As(args.LastErr, &amqpErr))
				require.EqualValues(t, fmt.Sprintf("Error from previous iteration %d", called-1), amqpErr.RemoteError.Description)
			}

			return &amqp.DetachError{
				RemoteError: &amqp.Error{
					// should be passeed into the callback on the next iteration.
					Description: fmt.Sprintf("Error from previous iteration %d", called)},
			}
		}, nil, fastRetryOptions)

		// if all the retries fail then we get the
		var amqpErr *amqp.DetachError
		require.True(t, errors.As(err, &amqpErr))
		require.EqualValues(t, "Error from previous iteration 3", amqpErr.RemoteError.Description)
		require.EqualValues(t, 3, called)
	})

	t.Run("FatalFailure", func(t *testing.T) {
		ctx := context.Background()

		called := 0

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			called++
			return context.Canceled
		}, nil, nil)

		require.ErrorIs(t, err, context.Canceled)
		require.EqualValues(t, 1, called)
	})

	t.Run("NonFatalFailures", func(t *testing.T) {
		ctx := context.Background()

		called := 0

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			called++
			if called == 1 {
				return &amqp.Error{
					Condition: "com.microsoft:message-lock-lost",
				}
			}

			panic("won't be called")
		}, nil, nil)

		var amqpErr *amqp.Error
		require.ErrorAs(t, err, &amqpErr)
		require.EqualValues(t, "com.microsoft:message-lock-lost", amqpErr.Condition)
		require.EqualValues(t, 1, called)
	})

	t.Run("Cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			// we propagate the context so this one should also be cancelled.
			select {
			case <-ctx.Done():
			default:
				require.Fail(t, "Context should have been cancelled")
			}

			return context.Canceled
		}, nil, nil)

		require.ErrorIs(t, context.Canceled, err)
	})

	t.Run("CustomFatalCheck", func(t *testing.T) {
		errCall := 0

		err := Retry(context.Background(), func(ctx context.Context, args RetryFnArgs) error {
			return fmt.Errorf("hello: %d", args.I)
		}, func(err error) bool {
			defer func() { errCall++ }()
			require.EqualValues(t, fmt.Sprintf("hello: %d", errCall), err.Error())

			// let one iteration go, then everything after that is fatal.
			return errCall != 0
		}, fastRetryOptions)

		require.EqualValues(t, "hello: 1", err.Error())
	})
}

func Test_calcDelay(t *testing.T) {
	t.Run("can't exceed max retry delay", func(t *testing.T) {
		duration := calcDelay(RetryOptions{
			RetryDelay:    time.Hour,
			MaxRetryDelay: time.Minute,
		}, 1)

		require.EqualValues(t, time.Minute, duration)
	})

	t.Run("increases with jitter", func(t *testing.T) {
		duration := calcDelay(RetryOptions{
			RetryDelay:    time.Minute,
			MaxRetryDelay: time.Hour,
		}, 1)

		require.GreaterOrEqual(t, duration, time.Duration((2-1)*time.Minute.Seconds()*0.8*float64(time.Second)))
		require.LessOrEqual(t, duration, time.Duration((2-1)*time.Minute.Seconds()*1.3*float64(time.Second)))

		duration = calcDelay(RetryOptions{
			RetryDelay:    time.Minute,
			MaxRetryDelay: time.Hour,
		}, 2)

		require.GreaterOrEqual(t, duration, time.Duration((2*2-1)*time.Minute.Seconds()*0.8*float64(time.Second)))
		require.LessOrEqual(t, duration, time.Duration((2*2-1)*time.Minute.Seconds()*1.3*float64(time.Second)))

		duration = calcDelay(RetryOptions{
			RetryDelay:    time.Minute,
			MaxRetryDelay: time.Hour,
		}, 3)

		require.GreaterOrEqual(t, duration, time.Duration((2*2*2-1)*time.Minute.Seconds()*0.8*float64(time.Second)))
		require.LessOrEqual(t, duration, time.Duration((2*2*2-1)*time.Minute.Seconds()*1.3*float64(time.Second)))
	})
}

var fastRetryOptions *RetryOptions = &RetryOptions{
	RetryDelay: time.Millisecond,
}
