// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

type (
	// SenderOption specifies an option that can configure a Sender.
	SenderOption func(sender *Sender) error

	// Sender is used to send messages as well as schedule them to be delivered at a later date.
	Sender struct {
		queueOrTopic string
		links        internal.AMQPLinks
		inner        internal.RecoverableSender
	}

	// SendableMessage are sendable using Sender.SendMessage.
	// Message, MessageBatch implement this interface.
	SendableMessage interface {
		toAMQPMessage() *amqp.Message
		messageType() string
	}
)

// tracing
const (
	spanNameSendMessageFmt string = "sb.sender.SendMessage.%s"
)

type messageBatchOptions struct {
	maxSizeInBytes *uint64
}

// MessageBatchOption is an option for configuring batch creation in
// `NewMessageBatch`.
type MessageBatchOption func(options *messageBatchOptions) error

// MessageBatchWithMaxSize overrides the max size (in bytes) for a batch.
// By default NewMessageBatch will use the max message size provided by the service.
func MessageBatchWithMaxSize(maxSizeInBytes uint64) func(options *messageBatchOptions) error {
	return func(options *messageBatchOptions) error {
		options.maxSizeInBytes = &maxSizeInBytes
		return nil
	}
}

// NewMessageBatch can be used to create a batch that contain multiple
// messages. Sending a batch of messages is more efficient than sending the
// messages one at a time.
func (s *Sender) NewMessageBatch(ctx context.Context, options ...MessageBatchOption) (*MessageBatch, error) {
	maxMessageSize, err := s.inner.MaxMessageSize(ctx)

	if err != nil {
		return nil, err
	}

	opts := &messageBatchOptions{
		maxSizeInBytes: &maxMessageSize,
	}

	for _, opt := range options {
		if err := opt(opts); err != nil {
			return nil, err
		}
	}

	return &MessageBatch{maxBytes: *opts.maxSizeInBytes}, nil
}

// SendMessage sends a message to a queue or topic.
// Message can be a MessageBatch (created using `Sender.CreateMessageBatch`) or
// a Message.
func (s *Sender) SendMessage(ctx context.Context, message SendableMessage) error {
	ctx, span := s.startProducerSpanFromContext(ctx, fmt.Sprintf(spanNameSendMessageFmt, message.messageType()))
	defer span.End()

	return s.inner.Send(ctx, message.toAMQPMessage())
}

// Close permanently closes the Sender.
func (s *Sender) Close(ctx context.Context) error {
	return s.links.Close(ctx, true)
}

func (sender *Sender) createSenderLink(ctx context.Context, session internal.AMQPSession) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
	amqpSender, err := session.NewSender(
		amqp.LinkSenderSettle(amqp.ModeMixed),
		amqp.LinkReceiverSettle(amqp.ModeFirst),
		amqp.LinkTargetAddress(sender.queueOrTopic))

	return amqpSender, nil, err
}

func newSender(ns *internal.Namespace, queueOrTopic string) (*Sender, error) {
	sender := &Sender{
		queueOrTopic: queueOrTopic,
	}

	links := ns.NewAMQPLinks(queueOrTopic, sender.createSenderLink)

	sender.inner = &internal.RecoverableAMQPSender{Links: links}
	sender.links = links
	return sender, nil
}

func (s *Sender) startProducerSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	internal.ApplyComponentInfo(span)
	span.AddAttributes(
		tab.StringAttribute("span.kind", "producer"),
		tab.StringAttribute("message_bus.destination", s.links.Audience()),
	)
	return ctx, span
}
