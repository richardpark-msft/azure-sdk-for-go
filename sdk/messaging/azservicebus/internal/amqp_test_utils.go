// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/golang/mock/gomock"
)

type FakeNS struct {
	claimNegotiated int
	recovered       uint64
	clientRevisions []uint64
	RPCLink         amqpwrap.RPCLink
	Session         amqpwrap.AMQPSession
	AMQPLinks       *FakeAMQPLinks

	CloseCalled int
}

type FakeAMQPSender struct {
	Closed int
	amqpwrap.AMQPSender
}

type FakeAMQPSession struct {
	amqpwrap.AMQPSession

	NewReceiverFn func(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error)

	closed int
}

type FakeAMQPLinks struct {
	AMQPLinks

	Closed              int
	CloseIfNeededCalled int

	// values to be returned for each `Get` call
	Revision LinkID
	Receiver amqpwrap.AMQPReceiver
	Sender   amqpwrap.AMQPSender
	RPC      amqpwrap.RPCLink

	// Err is the error returned as part of Get()
	Err error

	permanently bool
}

type FakeAMQPReceiver struct {
	amqpwrap.AMQPReceiver
	Closed  int
	CloseFn func(ctx context.Context) error

	CreditsCalled int
	CreditsImpl   func() uint32

	IssueCreditErr   error
	RequestedCredits uint32

	PrefetchedCalled int

	ReceiveCalled int
	ReceiveFn     func(ctx context.Context) (*amqp.Message, error)

	ReleaseMessageCalled int
	ReleaseMessageFn     func(ctx context.Context, msg *amqp.Message) error

	ReceiveResults []struct {
		M *amqp.Message
		E error
	}

	PrefetchedResults []*amqp.Message
}

type FakeRPCLink struct {
	Resp  *amqpwrap.RPCResponse
	Error error
}

func (r *FakeRPCLink) Close(ctx context.Context) error {
	return nil
}

func (r *FakeRPCLink) RPC(ctx context.Context, msg *amqp.Message) (*amqpwrap.RPCResponse, error) {
	return r.Resp, r.Error
}

func (r *FakeAMQPReceiver) LinkName() string {
	return "fakelink"
}

func (r *FakeAMQPReceiver) IssueCredit(credit uint32) error {
	r.RequestedCredits += credit

	if r.IssueCreditErr != nil {
		return r.IssueCreditErr
	}

	return nil
}

func (r *FakeAMQPReceiver) Credits() uint32 {
	r.CreditsCalled++

	if r.CreditsImpl != nil {
		return r.CreditsImpl()
	}

	return r.RequestedCredits
}

func (r *FakeAMQPReceiver) Prefetched() *amqp.Message {
	r.PrefetchedCalled++

	if len(r.PrefetchedResults) == 0 {
		return nil
	}

	res := r.PrefetchedResults[0]
	r.PrefetchedResults = r.PrefetchedResults[1:]

	return res
}

// Receive returns the next result from ReceiveResults or, if the ReceiveResults
// is empty, will block on ctx.Done().
func (r *FakeAMQPReceiver) Receive(ctx context.Context) (*amqp.Message, error) {
	r.ReceiveCalled++

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if r.ReceiveFn != nil {
			return r.ReceiveFn(ctx)
		}

		if len(r.ReceiveResults) == 0 {
			<-ctx.Done()
			return nil, ctx.Err()
		}

		res := r.ReceiveResults[0]
		r.ReceiveResults = r.ReceiveResults[1:]

		return res.M, res.E
	}
}

func (r *FakeAMQPReceiver) ReleaseMessage(ctx context.Context, msg *amqp.Message) error {
	r.ReleaseMessageCalled++

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if r.ReleaseMessageFn != nil {
			return r.ReleaseMessageFn(ctx, msg)
		}

		return nil
	}
}

func (r *FakeAMQPReceiver) Close(ctx context.Context) error {
	r.Closed++

	if r.CloseFn != nil {
		return r.CloseFn(ctx)
	}

	return nil
}

func (l *FakeAMQPLinks) Get(ctx context.Context) (*LinksWithID, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return &LinksWithID{
			Sender:   l.Sender,
			Receiver: l.Receiver,
			RPC:      l.RPC,
			ID:       l.Revision,
		}, l.Err
	}
}

func (l *FakeAMQPLinks) Retry(ctx context.Context, eventName log.Event, operation string, fn RetryWithLinksFn, o exported.RetryOptions) error {
	lwr, err := l.Get(ctx)

	if err != nil {
		return err
	}

	return fn(ctx, lwr, &utils.RetryFnArgs{})
}

func (l *FakeAMQPLinks) Close(ctx context.Context, permanently bool) error {
	if permanently {
		l.permanently = true
	}

	l.Closed++
	return nil
}

func (l *FakeAMQPLinks) CloseIfNeeded(ctx context.Context, err error) RecoveryKind {
	l.CloseIfNeededCalled++
	return GetRecoveryKind(err)
}

func (l *FakeAMQPLinks) ClosedPermanently() bool {
	return l.permanently
}

func (s *FakeAMQPSender) Close(ctx context.Context) error {
	s.Closed++
	return nil
}

func (s *FakeAMQPSession) NewReceiver(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error) {
	return s.NewReceiverFn(ctx, source, opts)
}

func (s *FakeAMQPSession) Close(ctx context.Context) error {
	s.closed++
	return nil
}

func (ns *FakeNS) NegotiateClaim(ctx context.Context, entityPath string) (context.CancelFunc, <-chan struct{}, error) {
	ch := make(chan struct{})
	close(ch)

	ns.claimNegotiated++

	return func() {}, ch, nil
}

func (ns *FakeNS) GetEntityAudience(entityPath string) string {
	return fmt.Sprintf("audience: %s", entityPath)
}

func (ns *FakeNS) NewAMQPSession(ctx context.Context) (amqpwrap.AMQPSession, uint64, error) {
	return ns.Session, ns.recovered + 100, nil
}

func (ns *FakeNS) NewRPCLink(ctx context.Context, managementPath string) (amqpwrap.RPCLink, error) {
	return ns.RPCLink, nil
}

func (ns *FakeNS) Recover(ctx context.Context, clientRevision uint64) (bool, error) {
	ns.clientRevisions = append(ns.clientRevisions, clientRevision)
	ns.recovered++
	return true, nil
}

func (ns *FakeNS) Close(ctx context.Context, permanently bool) error {
	ns.CloseCalled++
	return nil
}

func (ns *FakeNS) Check() error {
	return nil
}

type MockReceiverArgs struct {
	Err        error
	Msgs       []*amqp.Message
	Prefetched []*amqp.Message
}

func NewMockReceiver(ctrl *gomock.Controller, args MockReceiverArgs) *mock.MockAMQPReceiverCloser {
	receiver := mock.NewMockAMQPReceiverCloser(ctrl)

	var credits uint32
	linkName := fmt.Sprintf("link-%d", time.Now().UnixNano())

	receiver.EXPECT().LinkName().Return(linkName).AnyTimes()
	receiver.EXPECT().Credits().DoAndReturn(func() uint32 { return credits })
	receiver.EXPECT().IssueCredit(gomock.Any()).DoAndReturn(func(credit uint32) error {
		credits += credit
		return nil
	})

	ch := make(chan *amqp.Message, len(args.Msgs))

	for _, m := range args.Msgs {
		ch <- m
	}

	receiver.EXPECT().Receive(gomock.Any()).DoAndReturn(func(ctx context.Context) (*amqp.Message, error) {
		if len(ch) == 0 {
			if args.Err != nil {
				return nil, args.Err
			}

			<-ctx.Done()
			return nil, ctx.Err()
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case msg := <-ch:
			credits--
			return msg, nil
		}
	}).AnyTimes()

	receiver.EXPECT().Prefetched().DoAndReturn(func() *amqp.Message {
		if len(args.Prefetched) == 0 {
			return nil
		}

		m := args.Prefetched[0]
		args.Prefetched = args.Prefetched[1:]
		return m
	}).AnyTimes()

	return receiver
}
