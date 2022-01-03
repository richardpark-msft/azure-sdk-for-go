// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/tracing"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

type (
	// Sender is used to send messages as well as schedule them to be delivered at a later date.
	Sender struct {
		queueOrTopic   string
		cleanupOnClose func()
		links          internal.AMQPLinks
		baseRetrier    internal.Retrier
	}
)

// tracing
const (
	spanNameSend string = "sb.sender.Send"
)

// MessageBatchOptions contains options for the `Sender.NewMessageBatch` function.
type MessageBatchOptions struct {
	// MaxBytes overrides the max size (in bytes) for a batch.
	// By default NewMessageBatch will use the max message size provided by the service.
	MaxBytes uint64
}

// NewMessageBatch can be used to create a batch that contain multiple
// messages. Sending a batch of messages is more efficient than sending the
// messages one at a time.
func (s *Sender) NewMessageBatch(ctx context.Context, options *MessageBatchOptions) (*MessageBatch, error) {
	sender, _, _, _, err := s.links.Get(ctx)

	if err != nil {
		return nil, err
	}

	maxBytes := sender.MaxMessageSize()

	if options != nil && options.MaxBytes != 0 {
		maxBytes = options.MaxBytes
	}

	return newMessageBatch(maxBytes), nil
}

// SendMessage sends a Message to a queue or topic.
func (s *Sender) SendMessage(ctx context.Context, message *Message) error {
	return s.sendWithRetries(ctx, message.toAMQPMessage())
}

// SendMessageBatch sends a MessageBatch to a queue or topic.
// Message batches can be created using `Sender.NewMessageBatch`.
func (s *Sender) SendMessageBatch(ctx context.Context, batch *MessageBatch) error {
	return s.sendWithRetries(ctx, batch.toAMQPMessage())
}

// ScheduleMessages schedules a slice of Messages to appear on Service Bus Queue/Subscription at a later time.
// Returns the sequence numbers of the messages that were scheduled.  Messages that haven't been
// delivered can be cancelled using `Receiver.CancelScheduleMessage(s)`
func (s *Sender) ScheduleMessages(ctx context.Context, messages []*Message, scheduledEnqueueTime time.Time) ([]int64, error) {
	var amqpMessages []*amqp.Message

	for _, m := range messages {
		amqpMessages = append(amqpMessages, m.toAMQPMessage())
	}

	return s.scheduleAMQPMessages(ctx, amqpMessages, scheduledEnqueueTime)
}

// MessageBatch changes

// CancelScheduledMessages cancels multiple messages that were scheduled.
func (s *Sender) CancelScheduledMessages(ctx context.Context, sequenceNumber []int64) error {
	_, _, mgmt, _, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	return mgmt.CancelScheduled(ctx, sequenceNumber...)
}

// Close permanently closes the Sender.
func (s *Sender) Close(ctx context.Context) error {
	s.cleanupOnClose()
	return s.links.Close(ctx, true)
}

func (s *Sender) scheduleAMQPMessages(ctx context.Context, messages []*amqp.Message, scheduledEnqueueTime time.Time) ([]int64, error) {
	_, _, mgmt, _, err := s.links.Get(ctx)

	if err != nil {
		return nil, err
	}

	return mgmt.ScheduleMessages(ctx, scheduledEnqueueTime, messages...)
}

func (sender *Sender) createSenderLink(ctx context.Context, session internal.AMQPSession) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
	amqpSender, err := session.NewSender(
		amqp.LinkSenderSettle(amqp.ModeMixed),
		amqp.LinkReceiverSettle(amqp.ModeFirst),
		amqp.LinkTargetAddress(sender.queueOrTopic))

	if err != nil {
		tab.For(ctx).Error(err)
		return nil, nil, err
	}

	return amqpSender, nil, nil
}

func newSender(ns internal.NamespaceWithNewAMQPLinks, queueOrTopic string, cleanupOnClose func()) (*Sender, error) {
	sender := &Sender{
		queueOrTopic:   queueOrTopic,
		cleanupOnClose: cleanupOnClose,
		baseRetrier: internal.NewBackoffRetrier(internal.BackoffRetrierParams{
			Factor:     2,
			Jitter:     true,
			Min:        10 * time.Millisecond,
			Max:        1024 * time.Millisecond,
			MaxRetries: 10,
		}),
	}

	sender.links = ns.NewAMQPLinks(queueOrTopic, sender.createSenderLink)
	return sender, nil
}

func (s *Sender) startProducerSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	tracing.ApplyComponentInfo(span, internal.Version)
	span.AddAttributes(
		tab.StringAttribute("span.kind", "producer"),
		tab.StringAttribute("message_bus.destination", s.links.Audience()),
	)
	return ctx, span
}

func (s *Sender) sendWithRetries(ctx context.Context, message *amqp.Message) error {
	retrier := s.baseRetrier.Copy()

	var err error

	for retrier.Try(ctx) {
		ctx, span := s.startProducerSpanFromContext(ctx, spanNameSend)
		span.AddAttributes(tab.Int64Attribute("attempt", int64(retrier.CurrentTry())))

		var sender internal.AMQPSender
		var linksRevision uint64
		sender, _, _, linksRevision, err = s.links.Get(ctx)

		if err != nil {
			if internal.IsNonRetriable(err) {
				break
			}

			span.AddAttributes(tab.StringAttribute("Error(links)", err.Error()))
			span.End()
			continue
		}

		err = sender.Send(ctx, message)

		if err != nil {
			asSBE := s.links.RecoverIfNeeded(ctx, linksRevision, err)
			span.AddAttributes(tab.StringAttribute("Error(send)", asSBE.Error()))

			if internal.IsNonRetriable(asSBE) {
				break
			}

			span.End()
			continue
		}

		span.End()
		break
	}

	return err
}
