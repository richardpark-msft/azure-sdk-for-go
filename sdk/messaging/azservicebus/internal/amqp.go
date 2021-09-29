// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"time"

	"github.com/Azure/azure-amqp-common-go/v3/rpc"
	"github.com/Azure/go-amqp"
)

// AMQPReceiver is implemented by *amqp.Receiver
type AMQPReceiver interface {
	IssueCredit(credit uint32) error
	DrainCredit(ctx context.Context) error
	Receive(ctx context.Context) (*amqp.Message, error)
}

// AMQPReceiver is implemented by *amqp.Receiver
type AMQPReceiverCloser interface {
	AMQPReceiver
	Close(ctx context.Context) error
}

// AMQPSession is implemented by *amqp.Session
type AMQPSession interface {
	NewReceiver(opts ...amqp.LinkOption) (*amqp.Receiver, error)
	NewSender(opts ...amqp.LinkOption) (*amqp.Sender, error)
}

// AMQPSessionCloser is implemented by *amqp.Session
type AMQPSessionCloser interface {
	AMQPSession
	Close(ctx context.Context) error
}

// AMQPSender is implemented by *amqp.Sender
type AMQPSender interface {
	Send(ctx context.Context, msg *amqp.Message) error
	MaxMessageSize() uint64
}

// AMQPSenderCloser is implemented by *amqp.Sender
type AMQPSenderCloser interface {
	AMQPSender
	Close(ctx context.Context) error
}

// RPCLink is implemented by *rpc.Link
type RPCLink interface {
	Close(ctx context.Context) error
	RetryableRPC(ctx context.Context, times int, delay time.Duration, msg *amqp.Message) (*rpc.Response, error)
}

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

func (r *RecoverableAMQPReceiver) IssueCredit(credit uint32) error {
	ctx := context.Background()

	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return err
	}

	err = receiver.IssueCredit(credit)

	return r.Links.RecoverIfNeeded(ctx, err)
}

func (r *RecoverableAMQPReceiver) DrainCredit(ctx context.Context) error {
	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return err
	}

	err = receiver.DrainCredit(ctx)

	return r.Links.RecoverIfNeeded(ctx, err)
}

func (r *RecoverableAMQPReceiver) Receive(ctx context.Context) (*amqp.Message, error) {
	_, receiver, _, _, err := r.Links.Get(ctx)

	if err != nil {
		return nil, err
	}

	msg, err := receiver.Receive(ctx)
	return msg, r.Links.RecoverIfNeeded(ctx, err)
}

// recoverableAMQPSender

func (s *RecoverableAMQPSender) Send(ctx context.Context, msg *amqp.Message) error {
	sender, _, _, _, err := s.Links.Get(ctx)

	if err != nil {
		return err
	}

	err = sender.Send(ctx, msg)

	return s.Links.RecoverIfNeeded(ctx, err)
}

func (s *RecoverableAMQPSender) MaxMessageSize() uint64 {
	sender, _, _, _, err := s.Links.Get(context.Background())

	if err != nil {
		// TODO: what to do here.
		return 0
	}

	return sender.MaxMessageSize()
}
