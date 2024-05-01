// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package utils

import "sync/atomic"

type AtomicChannel[T any] struct {
	v atomic.Value
}

func (ac *AtomicChannel[T]) Load() chan T {
	v := ac.v.Load()

	if v == nil {
		return nil
	}

	return v.(chan T)
}

func (ac *AtomicChannel[T]) Store(v chan T) {
	ac.v.Store(v)
}

func (ac *AtomicChannel[T]) ReplaceAndClose(newV chan T) bool {
	v := ac.v.Swap(newV)

	if v != nil {
		close(v.(chan T))
		return true
	}

	return false
}
