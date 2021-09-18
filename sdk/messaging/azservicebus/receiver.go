// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// ReceiveMode represents the lock style to use for a reciever - either
// `PeekLock` or `ReceiveAndDelete`
type ReceiveMode int

const (
	// PeekLock will lock messages as they are received and can be settled
	// using the Receiver or Processor's (Complete|Abandon|DeadLetter|Defer)Message
	// functions.
	PeekLock ReceiveMode = 0
	// ReceiveAndDelete will delete messages as they are received.
	ReceiveAndDelete ReceiveMode = 1
)

// SubQueue allows you to target a subqueue of a queue or subscription.
// Ex: the dead letter queue (SubQueueDeadLetter).
type SubQueue string

const (
	// SubQueueDeadLetter targets the dead letter queue for a queue or subscription.
	SubQueueDeadLetter = "deadLetter"
	// SubQueueTransfer targets the transfer dead letter queue for a queue or subscription.
	SubQueueTransfer = "transferDeadLetter"
)

// Receiver receives messages using pull based functions (ReceiveMessages).
// For push-based receiving via callbacks look at the `Processor` type.
type Receiver struct {
	config struct {
		ReceiveMode    ReceiveMode
		FullEntityPath string
		Entity         entity

		RetryOptions struct {
			Times int
			Delay time.Duration
		}
	}

	mu    sync.Mutex
	links *links
	*settler
}

const defaultLinkRxBuffer = 2048

type amqpReceiver interface {
	Receive(ctx context.Context) (*amqp.Message, error)
	DrainCredit(ctx context.Context) error
	IssueCredit(credit uint32) error

	Close(ctx context.Context) error
}

// ReceiverOption represents an option for a receiver.
// Some examples:
// - `ReceiverWithReceiveMode` to configure the receive mode,
// - `ReceiverWithQueue` to target a queue.
type ReceiverOption func(receiver *Receiver) error

// ReceiverWithSubQueue allows you to open the sub queue (ie: dead letter queues, transfer dead letter queues)
// for a queue or subscription.
func ReceiverWithSubQueue(subQueue SubQueue) ReceiverOption {
	return func(receiver *Receiver) error {
		switch subQueue {
		case SubQueueDeadLetter:
		case SubQueueTransfer:
		case "":
			receiver.config.Entity.Subqueue = subQueue
		default:
			return fmt.Errorf("unknown SubQueue %s", subQueue)
		}

		return nil
	}
}

// ReceiverWithReceiveMode controls the receive mode for the receiver.
func ReceiverWithReceiveMode(receiveMode ReceiveMode) ReceiverOption {
	return func(receiver *Receiver) error {
		if receiveMode != PeekLock && receiveMode != ReceiveAndDelete {
			return fmt.Errorf("invalid receive mode specified %d", receiveMode)
		}

		receiver.config.ReceiveMode = receiveMode
		return nil
	}
}

// ReceiverWithQueue configures a receiver to connect to a queue.
func ReceiverWithQueue(queue string) ReceiverOption {
	return func(receiver *Receiver) error {
		receiver.config.Entity.Queue = queue
		return nil
	}
}

// ReceiverWithSubscription configures a receiver to connect to a subscription
// associated with a topic.
func ReceiverWithSubscription(topic string, subscription string) ReceiverOption {
	return func(receiver *Receiver) error {
		receiver.config.Entity.Topic = topic
		receiver.config.Entity.Subscription = subscription
		return nil
	}
}

func newReceiver(ns *internal.Namespace, options ...ReceiverOption) (*Receiver, error) {
	receiver := &Receiver{
		config: struct {
			ReceiveMode    ReceiveMode
			FullEntityPath string
			Entity         entity
			RetryOptions   struct {
				Times int
				Delay time.Duration
			}
		}{
			ReceiveMode: PeekLock,
		},
	}

	for _, opt := range options {
		if err := opt(receiver); err != nil {
			return nil, err
		}
	}

	entityPath, err := receiver.config.Entity.String()

	if err != nil {
		return nil, err
	}

	receiver.config.FullEntityPath = entityPath

	receiver.links = internal.NewLinks(ns, entityPath, receiver.linkCreator)

	// 'nil' settler handles returning an error message for receiveAndDelete links.
	if receiver.config.ReceiveMode == PeekLock {
		receiver.settler = &settler{links: links}
	}

	return receiver, nil
}

// ReceiveOptions are options for the ReceiveMessages function.
type ReceiveOptions struct {
	maxWaitTime                  time.Duration
	maxWaitTimeAfterFirstMessage time.Duration
}

// ReceiveOption represents an option for a `ReceiveMessages`.
// For example, `ReceiveWithMaxWaitTime` will let you configure the
// maxmimum amount of time to wait for messages to arrive.
type ReceiveOption func(options *ReceiveOptions) error

// ReceiveWithMaxWaitTime configures how long to wait for the first
// message in a set of messages to arrive.
// Default: 60 seconds
func ReceiveWithMaxWaitTime(max time.Duration) ReceiveOption {
	return func(options *ReceiveOptions) error {
		options.maxWaitTime = max
		return nil
	}
}

// ReceiveWithMaxTimeAfterFirstMessage confiures how long, after the first
// message arrives, to wait before returning.
// Default: 1 second
func ReceiveWithMaxTimeAfterFirstMessage(max time.Duration) ReceiveOption {
	return func(options *ReceiveOptions) error {
		options.maxWaitTimeAfterFirstMessage = max
		return nil
	}
}

// ReceiveMessages receives a fixed number of messages, up to numMessages.
// There are two timeouts involved in receiving messages:
// 1. An explicit timeout set with `ReceiveWithMaxWaitTime` (default: 60 seconds)
// 2. An implicit timeout (default: 1 second) that starts after the first
//    message has been received. This time can be adjusted with `ReceiveWithMaxTimeAfterFirstMessage`
func (r *Receiver) ReceiveMessages(ctx context.Context, maxMessages int, options ...ReceiveOption) ([]*ReceivedMessage, error) {
	if r.linkState.Closed() {
		return nil, r.linkState.Err()
	}

	ropts := &ReceiveOptions{
		maxWaitTime:                  time.Minute,
		maxWaitTimeAfterFirstMessage: time.Second,
	}

	for _, opt := range options {
		if err := opt(ropts); err != nil {
			return nil, err
		}
	}

	_, receiver, _, err := r.links.Get(ctx)

	if err != nil {
		return nil, err
	}

	if err := receiver.IssueCredit(uint32(maxMessages)); err != nil {
		return nil, err
	}

	var messages []*ReceivedMessage

	// There are two timers - an initial one that starts counting down
	// before receiving starts and a second, shorter timer, that's starts
	// after the first message is received
	//
	// The first timer is there just to make sure that if we _never_ receive
	// messages that we don't just wait forever. That one starts here.
	//
	// The second timer starts after we receive one message. That timer is there
	// to make sure we don't hold onto the messages for too long, and risk them
	// expiring.
	receiveCtx, startFirstMessageCountdown := contextWithTwoTimeouts(ctx, ropts.maxWaitTime, ropts.maxWaitTimeAfterFirstMessage)
	receivingComplete := make(chan error, 1)

	go func() {
		for {
			message, err := receiver.Receive(receiveCtx)

			if err != nil {
				receivingComplete <- err
				break
			}

			// TODO: temporary while we don't yet have the "final" messaging
			// structure from https://github.com/Azure/azure-sdk-for-go/issues/15094
			msg, err := internal.MessageFromAMQPMessage(message)

			if err != nil {
				// TODO: there is only one reason this happens (mapstructure fails).
				// We're only using to peel out a handful of attributes so I'd rather just
				// get rid of it (and not have the error condition at all)
				receivingComplete <- err
				break
			}

			receivedMessage := convertToReceivedMessage(msg)

			messages = append(messages, receivedMessage)

			if len(messages) == maxMessages {
				close(receivingComplete)
				break
			}

			if len(messages) == 1 {
				startFirstMessageCountdown()
			}
		}
	}()

	select {
	case err := <-receivingComplete:
		return messages, err
	case <-ctx.Done():
		err = ctx.Err()
	}

	// make sure we leave the link in a consistent state.
	if err := receiver.DrainCredit(ctx); err != nil {
		return nil, err
	}

	return messages, err
}

// CompleteMessage completes a message, deleting it from the queue or subscription.
// func (r *Receiver) CompleteMessage(ctx context.Context, message *ReceivedMessage) error {
// 	return r.settler.Complete(ctx, message)
// }

// DeadLetterMessage settles a message by moving it to the dead letter queue for a
// queue or subscription.
// func (r *Receiver) DeadLetterMessage(ctx context.Context, message *ReceivedMessage) error {
// 	// TODO: dead letter reasons.
// 	return r.settler.DeadLetter(ctx, message)
// }

// AbandonMessage will cause a message to be returned to the queue or subscription.
// This will increment its delivery count, and potentially cause it to be dead lettered
// depending on your queue or subscription's configuration.
// func (r *Receiver) AbandonMessage(ctx context.Context, message *ReceivedMessage) error {
// 	return r.settler.Abandon(ctx, message)
// }

// // DeferMessage will cause a message to be deferred.
// // Messages that are deferred by can be retrieved using `Receiver.ReceiveDeferredMessages()`.
// func (r *Receiver) DeferMessage(ctx context.Context, message *ReceivedMessage) error {
// 	return r.settler.Defer(ctx, message)
// }

// Close permanently closes the receiver.
func (r *Receiver) Close(ctx context.Context) error {
	return r.links.Close(ctx)
}

type entity struct {
	Subqueue     SubQueue
	Queue        string
	Topic        string
	Subscription string
}

func (e *entity) String() (string, error) {
	entityPath := ""

	if e.Queue != "" {
		entityPath = e.Queue
	} else if e.Topic != "" && e.Subscription != "" {
		entityPath = fmt.Sprintf("%s/Subscriptions/%s", e.Topic, e.Subscription)
	} else {
		return "", errors.New("a queue or subscription was not specified")
	}

	if e.Subqueue == SubQueueDeadLetter {
		entityPath += "/$DeadLetterQueue"
	} else if e.Subqueue == SubQueueTransfer {
		entityPath += "/$Transfer/$DeadLetterQueue"
	}

	return entityPath, nil
}

func contextWithTwoTimeouts(ctx context.Context, initial time.Duration, timeAfterFirstMessage time.Duration) (context.Context, func()) {
	ctx, cancel := context.WithTimeout(ctx, initial)

	afterFirstMessage := func() {
		go func() {
			select {
			case <-time.After(timeAfterFirstMessage):
			case <-ctx.Done():
				cancel()
			}
		}()
	}

	return ctx, afterFirstMessage
}

func createReceiverLink(ctx context.Context, session *amqp.Session, linkOptions []amqp.LinkOption) (*amqp.Sender, *amqp.Receiver, error) {
	opts := r.createLinkOptions()

	amqpReceiver, err := session.NewReceiver(opts...)

	if err != nil {
		tab.For(ctx).Error(err)
		return nil, nil, err
	}

	return nil, amqpReceiver, nil
}

func createLinkOptions(mode ReceiveMode, entityPath string) []amqp.LinkOption {
	receiveMode := amqp.ModeSecond

	if mode == ReceiveAndDelete {
		receiveMode = amqp.ModeFirst
	}

	opts := []amqp.LinkOption{
		amqp.LinkSourceAddress(r.config.FullEntityPath),
		amqp.LinkReceiverSettle(receiveMode),
		amqp.LinkWithManualCredits(),
		amqp.LinkCredit(defaultLinkRxBuffer),
	}

	if mode == ReceiveAndDelete {
		opts = append(opts, amqp.LinkSenderSettle(amqp.ModeSettled))
	}

	return opts
}
