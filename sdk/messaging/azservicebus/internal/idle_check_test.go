package internal

import (
	"context"
	"testing"
	"time"
)

func TestIdleCheck(t *testing.T) {
	ic := NewIdleChecker(time.Second)

	drainWithTimeout := func() error {}
	closeIfNeeded := func(err error) {}

	fr := func fakeReceiver(ctx context.Context, fn func(ctx context.Context) error) error {
		newCtx, newCancel := ic.NewContext(ctx)
		defer newCancel()

		err := fn(newCtx)
	
		if err != nil {
			if IsCancelError(err) {
				if ctx.Err() != nil {
					// user cancelled
					return err
				} else {
					// we cancelled - something's not working properly here.					
					err := drainWithTimeout()					
					closeIfNeeded(err)
					return nil	// let next call recover.
				}
			}
	
			return err
		}
	
		// we succeeded! we should debounce the timer.
		ic.Debounce()
		return nil
	}	
}

func blockForOneMinute(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Minute):
		return nil
	}
}
