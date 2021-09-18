package azservicebus

import (
	"context"

	"github.com/Azure/go-amqp"
)

type settler struct {
	links *links
}

func (s *settler) Complete(ctx context.Context, message *ReceivedMessage) error {
	_, receiver, mgmt, linkRevision, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	// complete
	if message.linkRevision != linkRevision {
		return mgmt.SendDisposition(ctx, message.rawAMQPMessage.LockToken, disposition{Status: completedDisposition})
	}

	return message.rawAMQPMessage.Accept(ctx)
}

func (s *settler) Abandon(ctx context.Context, links *links, message *ReceivedMessage) error {
	_, receiver, mgmt, linkRevision, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	if linkRevision != message.linkRevision {
		// abandon
		d := disposition{
			Status: abandonedDisposition,
		}
		return sendMgmtDisposition(ctx, m, d)
	}

	message.rawAMQPMessage.Modify(ctx, false, false, nil)

}

func (s *settler) Defer(ctx context.Context, links *links, message *ReceivedMessage) error {
	_, receiver, _, linkRevision, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	return message.rawAMQPMessage.Modify(ctx, true, true, nil)
}

func (s *settler) DeadLetter(ctx context.Context, message *ReceivedMessage) {
	_, receiver, mgmt, linkRevision, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	if linkRevision != message.linkRevision {
		d := disposition{
			Status:                suspendedDisposition,
			DeadLetterDescription: ptrString(err.Error()),
			DeadLetterReason:      ptrString("amqp:error"),
		}
		return sendMgmtDisposition(ctx, m, d)
	}

	var info map[string]interface{}
	if additionalData != nil {
		info = make(map[string]interface{}, len(additionalData))
		for key, val := range additionalData {
			info[key] = val
		}
	}

	amqpErr := amqp.Error{
		Condition:   amqp.ErrorCondition(ErrorInternalError),
		Description: err.Error(),
		Info:        info,
	}
	return message.rawAMQPMessage.Reject(ctx, &amqpErr)
}
