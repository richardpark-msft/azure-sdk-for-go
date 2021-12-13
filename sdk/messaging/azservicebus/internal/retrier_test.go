// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"testing"

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
				return &amqp.DetachError{}
			}

			return nil
		}, nil, nil)

		require.Nil(t, err)
		require.EqualValues(t, 2, called)
	})

	t.Run("FatalFailure", func(t *testing.T) {
		ctx := context.Background()

		called := 0

		err := Retry(ctx, func(ctx context.Context, args RetryFnArgs) error {
			called++

			if called == 1 {
				return context.Canceled
			}

			return nil
		}, nil, nil)

		require.ErrorIs(t, err, context.Canceled)
		require.EqualValues(t, 2, called)
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

		require.NoError(t, err)
	})
}
