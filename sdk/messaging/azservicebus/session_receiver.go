package azservicebus

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/go-amqp"
)

type SessionReceiver struct {
	*Receiver
}

const sessionFilterName = "com.microsoft:session-filter"

func newSessionReceiver(receiver *Receiver, sessionID *string, options ...ReceiverOption) *SessionReceiver {
	return &SessionReceiver{
		Receiver: receiver,
	}
}

func (sr *SessionReceiver) updateReceiver(ctx context.Context, linkOptions []amqp.LinkOption) (amqpReceiver, error) {
	receiver, err := sr.Receiver.updateReceiver(ctx, getSessionFilter(s.sessionID)...)

	rawsid := sr.Receiver.receiver.LinkSourceFilterValue(sessionFilterName)

	if rawsid == nil && r.sessionID == nil {
		return errors.New("failed to create a receiver.  no unlocked sessions available")
	} else if rawsid != nil && r.sessionID != nil && rawsid != *r.sessionID {
		return fmt.Errorf("failed to create a receiver for session %s, it may be locked by another receiver", rawsid)
	} else if r.sessionID == nil {
		sid := rawsid.(string)
		r.sessionID = &sid
	}
}

func getSessionFilter(sessionID *string) amqp.LinkOption {
	const code = uint64(0x00000137000000C)

	if sessionID == nil {
		return amqp.LinkSourceFilter(sessionFilterName, code, nil), true
	}

	return amqp.LinkSourceFilter(sessionFilterName, code, sessionID), true
}
