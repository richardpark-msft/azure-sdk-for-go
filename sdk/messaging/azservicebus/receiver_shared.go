// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/devigned/tab"
)

type ReceiveOperation struct {
	mu       sync.Mutex
	lastErr  error
	cancel   context.CancelFunc
	messages chan *ReceivedMessage
}

func (r *ReceiveOperation) Messages() <-chan *ReceivedMessage {
	return r.messages
}

func (r *ReceiveOperation) Err() error {
	return r.lastErr
}

func (r *ReceiveOperation) Stop() {
	r.cancel()
}

func receiveMessages(receiver internal.AMQPReceiver, initialCredit uint32, reissueCredit bool) *ReceiveOperation {
	ch := make(chan *ReceivedMessage, initialCredit)
	ctx, cancel := context.WithCancel(context.Background())

	rcv := &ReceiveOperation{
		cancel:   cancel,
		messages: ch,
	}

	go func() {
		defer close(ch)

		if err := getMessages(ctx, receiver, initialCredit, reissueCredit, ch); err != nil {
			rcv.lastErr = err
			return
		}

		if err := drainLink(ctx, receiver, ch); err != nil {
			rcv.lastErr = err
			return
		}
	}()

	return rcv
}

func getMessages(ctx context.Context, receiver internal.AMQPReceiver, initialCredit uint32, reissueCredit bool, messagesCh chan *ReceivedMessage) error {
	err := receiver.IssueCredit(initialCredit)

	if err != nil {
		return err
	}

	for {
		amqpMessage, err := receiver.Receive(ctx)

		if err != nil {
			if internal.IsCancelError(err) {
				// user's stopped the receiver operation.
				return nil
			}

			// fatal error that we can't handle
			return err
		}

		// TODO: there is some odd error happening where we occasionally get no error but also
		// get a nil AMQP message.
		if amqpMessage != nil {
			messagesCh <- newReceivedMessage(ctx, amqpMessage)
			if reissueCredit {
				if err := receiver.IssueCredit(1); err != nil {
					return err
				}
			}
		}
	}
}

// drainLink initiates a drainLink on the link. Service Bus will send whatever messages it might have still had and
// set our link credit to 0.
func drainLink(ctx context.Context, receiver internal.AMQPReceiver, messagesCh chan *ReceivedMessage) error {
	receiveCtx, cancelReceive := context.WithCancel(context.Background())

	// start the drain asynchronously. Note that we ignore the user's context at this point
	// since draining makes sure we don't get messages when nobody is receiving.
	go func() {
		if err := receiver.DrainCredit(context.Background()); err != nil {
			tab.For(receiveCtx).Debug(fmt.Sprintf("Draining of credit failed. link will be closed and will re-open on next receive: %s", err.Error()))
		}
		cancelReceive()
	}()

	// Receive until the drain completes, at which point it'll cancel
	// our context.
	// NOTE: That's a gap here where we need to be able to drain _only_ the internally cached messages
	// in the receiver. Filed as https://github.com/Azure/go-amqp/issues/71
	for {
		am, err := receiver.Receive(receiveCtx)

		if err != nil {
			if internal.IsCancelError(err) {
				return nil
			}

			return err
		}

		// TODO: there is some odd error happening where we occasionally get no error but also
		// get a nil AMQP message.
		if am != nil {
			// keep sending the messages to the channel
			messagesCh <- newReceivedMessage(ctx, am)
		}
	}
}
