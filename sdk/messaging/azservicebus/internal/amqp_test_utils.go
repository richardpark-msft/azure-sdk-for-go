// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/go-amqp"
)

type fakeNS struct {
	claimNegotiated int
	recovered       uint64
	clientRevisions []uint64
	MgmtClient      MgmtClient
	Session         AMQPSessionCloser
}

type fakeAMQPReceiver struct {
	issueCreditParams []uint32
	issueCreditReturn error

	drainCreditCalls  int
	drainCreditReturn error

	receiveReturns []struct {
		Message *amqp.Message
		Err     error
	}
}

type fakeAMQPSession struct {
	AMQPSessionCloser
	closed int
}

type fakeMgmtClient struct {
	MgmtClient
	closed int
}

type fakeRPCLink struct {
	closed    int
	closedErr error

	rpcParams  []*amqp.Message
	rpcReturns []struct {
		Resp *rpc.Response
		Err  error
	}
}

type FakeAMQPLinks struct {
	AMQPLinks

	recoverParams []error

	// values to be returned for each `Get` call
	Revision uint64
	Receiver *fakeAMQPReceiver
	Sender   *fakeAMQPSender
	Mgmt     MgmtClient
	Err      error
}

// links

func (l *FakeAMQPLinks) Get(ctx context.Context) (AMQPSender, AMQPReceiver, MgmtClient, uint64, error) {
	return l.Sender, l.Receiver, l.Mgmt, l.Revision, l.Err
}

func (l *FakeAMQPLinks) RecoverIfNeeded(ctxForLogging context.Context, err error) error {
	l.recoverParams = append(l.recoverParams, err)
	return err
}

func (l *FakeAMQPLinks) ResetMock() {
	*l = FakeAMQPLinks{}
}

// sender

type fakeAMQPSender struct {
	closed int

	sendParams  []*amqp.Message
	sendReturns []error

	maxMessageSizeReturn []uint64
}

func (s *fakeAMQPSender) Send(ctx context.Context, msg *amqp.Message) error {
	s.sendParams = append(s.sendParams, msg)

	var resp error
	resp, s.sendReturns = s.sendReturns[0], s.sendReturns[1:]

	return resp
}

func (s *fakeAMQPSender) MaxMessageSize() uint64 {
	var resp uint64
	resp, s.maxMessageSizeReturn = s.maxMessageSizeReturn[0], s.maxMessageSizeReturn[1:]

	return resp
}

func (s *fakeAMQPSender) Close(ctx context.Context) error {
	s.closed++
	return nil
}

// receiver

func (r *fakeAMQPReceiver) IssueCredit(credit uint32) error {
	r.issueCreditParams = append(r.issueCreditParams, credit)
	return r.issueCreditReturn
}

func (r *fakeAMQPReceiver) DrainCredit(ctx context.Context) error {
	r.drainCreditCalls++
	return r.drainCreditReturn
}

func (r *fakeAMQPReceiver) Receive(ctx context.Context) (*amqp.Message, error) {
	ret := r.receiveReturns[0]
	r.receiveReturns = r.receiveReturns[1:]
	return ret.Message, ret.Err
}

// session

func (s *fakeAMQPSession) Close(ctx context.Context) error {
	s.closed++
	return nil
}

// mgmt

func (m *fakeMgmtClient) Close(ctx context.Context) error {
	m.closed++
	return nil
}

// rpclink

func (r *fakeRPCLink) Close(ctx context.Context) error {
	r.closed++
	return r.closedErr
}

func (r *fakeRPCLink) RetryableRPC(ctx context.Context, times int, delay time.Duration, msg *amqp.Message) (*rpc.Response, error) {
	r.rpcParams = append(r.rpcParams, msg)
	ret := r.rpcReturns[0]
	r.rpcReturns = r.rpcReturns[1:]

	return ret.Resp, ret.Err
}

// namespace

func (ns *fakeNS) NegotiateClaim(ctx context.Context, entityPath string) (func() <-chan struct{}, error) {
	ch := make(chan struct{})
	close(ch)

	ns.claimNegotiated++

	return func() <-chan struct{} {
		return ch
	}, nil
}

func (ns *fakeNS) GetEntityAudience(entityPath string) string {
	return fmt.Sprintf("audience: %s", entityPath)
}

func (ns *fakeNS) NewAMQPSession(ctx context.Context) (AMQPSessionCloser, uint64, error) {
	return ns.Session, ns.recovered + 100, nil
}

func (ns *fakeNS) NewMgmtClient(ctx context.Context, links AMQPLinks) (MgmtClient, error) {
	return ns.MgmtClient, nil
}

func (ns *fakeNS) Recover(ctx context.Context, clientRevision uint64) error {
	ns.clientRevisions = append(ns.clientRevisions, clientRevision)
	ns.recovered++
	return nil
}

type createLinkResponse struct {
	sender   AMQPSenderCloser
	receiver AMQPReceiverCloser
	err      error
}
