package azservicebus

import "testing"

func TestRenewer_MessageLock_Expired(t *testing.T) {
}

func TestRenewer_MessageLock_Unowned(t *testing.T) {

}

func TestRenewer_MessageLock_Success(t *testing.T) {

}

func TestRenewer_MessageLock_AvoidWastedRenews(t *testing.T) {
	// - lock is already expired
	// - lock is already locked for the length of the max duration they've asked us to renew for.

	t.Fail()
}
