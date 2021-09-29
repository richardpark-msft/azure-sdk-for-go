package internal

import (
	"context"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/go-amqp"
)

// RecoverableSender is implemented by *RecoverableAMQPSender
// NOTE: It's almost an AMQPSender but MaxMessageSize requires an
// initialized link (and could fail if it can't obtain one)
type RecoverableSender interface {
	Send(ctx context.Context, msg *amqp.Message) error
	MaxMessageSize(ctx context.Context) (uint64, error)
}

// RecoverableReceiver is implemented by *RecoverableAMQPReceiver
// NOTE: It's almost an AMQPReceiver but IssueCredit (which normally)
// doesn't do any I/O) needs to take a context since it can do a links.Get())
type RecoverableReceiver interface {
	IssueCredit(ctx context.Context, credit uint32) error
	DrainCredit(ctx context.Context) error
	Receive(ctx context.Context) (*amqp.Message, error)
}

// recoverableRPCLink is a little different than the RecoverableSender/Receiver
// because it handles it's own recovery (it comes from the azure-amqp-common-go)
// library.
type recoverableRPCLink struct {
	inner RPCLink
	links AMQPLinks
}

type RecoverableAMQPReceiver struct {
	Links AMQPLinks
}

type RecoverableAMQPSender struct {
	Links AMQPLinks
}

// recoverableRPCLink

func (r *recoverableRPCLink) RetryableRPC(ctx context.Context, times int, delay time.Duration, msg *amqp.Message) (*rpc.Response, error) {
	resp, err := r.inner.RetryableRPC(ctx, times, delay, msg)
	return resp, r.links.RecoverIfNeeded(ctx, err)
}

func (r *recoverableRPCLink) Close(ctx context.Context) error {
	return r.inner.Close(ctx)
}

// recoverableAMQPReceiver

func (r *RecoverableAMQPReceiver) IssueCredit(ctx context.Context, credit uint32) error {
	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return r.Links.RecoverIfNeeded(ctx, err)
	}

	err = receiver.IssueCredit(credit)

	return r.Links.RecoverIfNeeded(ctx, err)
}

func (r *RecoverableAMQPReceiver) DrainCredit(ctx context.Context) error {
	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return r.Links.RecoverIfNeeded(ctx, err)
	}

	err = receiver.DrainCredit(ctx)

	return r.Links.RecoverIfNeeded(ctx, err)
}

func (r *RecoverableAMQPReceiver) Receive(ctx context.Context) (*amqp.Message, error) {
	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return nil, r.Links.RecoverIfNeeded(ctx, err)
	}

	msg, err := receiver.Receive(ctx)
	return msg, r.Links.RecoverIfNeeded(ctx, err)
}

// recoverableAMQPSender

func (s *RecoverableAMQPSender) Send(ctx context.Context, msg *amqp.Message) error {
	sender, _, _, _, err := s.Links.Get(ctx)

	if err != nil {
		return s.Links.RecoverIfNeeded(ctx, err)
	}

	err = sender.Send(ctx, msg)

	return s.Links.RecoverIfNeeded(ctx, err)
}

func (s *RecoverableAMQPSender) MaxMessageSize(ctx context.Context) (uint64, error) {
	sender, _, _, _, err := s.Links.Get(ctx)

	if err != nil {
		return 0, s.Links.RecoverIfNeeded(ctx, err)
	}

	return sender.MaxMessageSize(), nil
}
