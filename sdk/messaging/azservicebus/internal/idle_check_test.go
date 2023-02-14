package internal

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIdleCheck(t *testing.T) {
	start := time.Now()
	ic := NewIdleChecker(time.Second)

	newCtx, newCancel := ic.NewContext(context.Background())
	defer newCancel()

	deadline, hasDeadline := newCtx.Deadline()
	require.True(t, hasDeadline)
	require.GreaterOrEqual(t, deadline.Sub(start), time.Second)
}

func blockForOneMinute(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Minute):
		return nil
	}
}
