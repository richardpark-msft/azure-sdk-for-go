// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"

	"github.com/Azure/azure-amqp-common-go/v3/uuid"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpsb"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/spans"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/models"
	"github.com/Azure/go-amqp"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/devigned/tab"
)

type (
	// SenderOption specifies an option that can configure a Sender.
	SenderOption func(sender *Sender) error

	// Sender is used to send messages as well as schedule them to be delivered at a later date.
	Sender struct {
		queueOrTopic string
		links        *amqpsb.Links
	}

	// SendableMessage are sendable using Sender.SendMessage.
	// Message, MessageBatch implement this interface.
	SendableMessage interface {
		ToAMQPMessage() (*amqp.Message, error)
		MessageType() string
	}
)

// tracing
const (
	spanNameSendMessageFmt string = "sb.sender.SendMessage.%s"
)

type messageBatchOptions struct {
	maxSizeInBytes *int
}

// MessageBatchOption is an option for configuring batch creation in
// `NewMessageBatch`.
type MessageBatchOption func(options *messageBatchOptions) error

// MessageBatchWithMaxSize overrides the max size (in bytes) for a batch.
// By default NewMessageBatch will use the max message size provided by the service.
func MessageBatchWithMaxSize(maxSizeInBytes int) func(options *messageBatchOptions) error {
	return func(options *messageBatchOptions) error {
		options.maxSizeInBytes = &maxSizeInBytes
		return nil
	}
}

// NewMessageBatch can be used to create a batch that contain multiple
// messages. Sending a batch of messages is more efficient than sending the
// messages one at a time.
func (s *Sender) NewMessageBatch(ctx context.Context, options ...MessageBatchOption) (*models.MessageBatch, error) {
	sender, _, _, _, err := s.links.Get(ctx)

	if err != nil {
		return nil, err
	}

	opts := &messageBatchOptions{
		maxSizeInBytes: to.IntPtr(int(sender.MaxMessageSize())),
	}

	for _, opt := range options {
		if err := opt(opts); err != nil {
			return nil, err
		}
	}

	return &models.MessageBatch{MaxBytes: *opts.maxSizeInBytes}, nil
}

// SendMessage sends a message to a queue or topic.
// Message can be a MessageBatch (created using `Sender.CreateMessageBatch`) or
// a Message.
func (s *Sender) SendMessage(ctx context.Context, message SendableMessage) error {
	ctx, span := s.startProducerSpanFromContext(ctx, spans.SendMessage)
	defer span.End()

	sender, _, _, err := s.links.Get(ctx)

	if err != nil {
		return err
	}

	amqpMessage, err := message.toAMQPMessage()

	if err != nil {
		return err
	}

	if amqpMess.ID == "" {
		id, err := uuid.NewV4()
		if err != nil {
			tab.For(ctx).Error(err)
			return err
		}
		msg.ID = id.String()
	}

	sender.Send(ctx, amqpMessage)
}

// Close permanently closes the Sender.
func (s *Sender) Close(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer s.linkState.Close()

	var err error

	if s.legacySender != nil {
		err = s.legacySender.Close(ctx)
		s.legacySender = nil
	}

	return err
}

func (sender *Sender) createSenderLink(ctx context.Context, session *amqp.Session) (*amqp.Sender, *amqp.Receiver, error) {
	amqpSender, err := session.NewSender(
		amqp.LinkSenderSettle(amqp.ModeMixed),
		amqp.LinkReceiverSettle(amqp.ModeFirst),
		amqp.LinkTargetAddress(s.getAddress()))

	return amqpSender, nil, err
}

func newSender(ns *internal.Namespace, queueOrTopic string) (*Sender, error) {
	sender := &Sender{
		queueOrTopic: queueOrTopic,
	}

	sender.links = newAMQPLinks(ns, queueOrTopic, sender.createSenderLink)
	return sender, nil
}

// handleAMQPError is called internally when an event has failed to send so we
// can parse the error to determine whether we should attempt to retry sending the event again.
// func (s *Sender) handleAMQPError(ctx context.Context, err error) error {
// 	var amqpError *amqp.Error
// 	if errors.As(err, &amqpError) {
// 		switch amqpError.Condition {
// 		case errorServerBusy:
// 			return s.retryRetryableAmqpError(ctx, amqpRetryDefaultTimes, amqpRetryBusyServerDelay)
// 		case errorTimeout:
// 			return s.retryRetryableAmqpError(ctx, amqpRetryDefaultTimes, amqpRetryDefaultDelay)
// 		case errorOperationCancelled:
// 			return s.retryRetryableAmqpError(ctx, amqpRetryDefaultTimes, amqpRetryDefaultDelay)
// 		case errorContainerClose:
// 			return s.retryRetryableAmqpError(ctx, amqpRetryDefaultTimes, amqpRetryDefaultDelay)
// 		default:
// 			return err
// 		}
// 	}
// 	return s.retryRetryableAmqpError(ctx, amqpRetryDefaultTimes, amqpRetryDefaultDelay)
// }
