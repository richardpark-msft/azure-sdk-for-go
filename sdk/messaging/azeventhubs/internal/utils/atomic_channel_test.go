// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package utils_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestAtomicChannel(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		v := utils.AtomicChannel[string]{}
		require.Nil(t, v.Load())
	})

	t.Run("basic", func(t *testing.T) {
		v := utils.AtomicChannel[string]{}
		ch := make(chan string, 1)
		v.Store(ch)

		storedCh := v.Load()
		storedCh <- "hello world"

		helloWorld := <-ch
		require.Equal(t, "hello world", helloWorld)
	})

	t.Run("replace and close", func(t *testing.T) {
		v := utils.AtomicChannel[string]{}
		smallerCh := make(chan string, 1)
		v.Store(smallerCh)

		v.ReplaceAndClose(make(chan string, 2))

		// our old channel should be closed now so we'll
		// only get the zero value from it.
		val := <-smallerCh
		require.Empty(t, val)

		// and the channel stored in there should be the larger one we just created
		largerCh := v.Load()
		require.Equal(t, 2, cap(largerCh))
	})

	t.Run("replace and close (starts empty)", func(t *testing.T) {
		v := utils.AtomicChannel[string]{}

		v.ReplaceAndClose(make(chan string, 2))

		// and the channel stored in there should be the one we just created
		largerCh := v.Load()
		require.Equal(t, 2, cap(largerCh))
	})

	t.Run("nil", func(t *testing.T) {
		v := utils.AtomicChannel[string]{}
		v.ReplaceAndClose(nil)

		a := v.Load()
		require.Nil(t, a)
		require.True(t, a == nil)
		require.True(t, a == (chan string)(nil))

		v.Store(nil)

		a = v.Load()
		require.Nil(t, a)
		require.True(t, a == nil)
		require.True(t, a == (chan string)(nil))
	})
}
