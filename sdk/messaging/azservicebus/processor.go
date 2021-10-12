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
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sberrors"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// ProcessorOptions contains options for the `Client.NewProcessorForQueue` or
// `Client.NewProcessorForSubscription` functions.
type ProcessorOptions struct {
	// ReceiveMode controls when a message is deleted from Service Bus.
	//
	// `azservicebus.PeekLock` is the default. The message is locked, preventing multiple
	// receivers from processing the message at once. You control the lock state of the message
	// using one of the message settlement functions like processor.CompleteMessage(), which removes
	// it from Service Bus, or processor.AbandonMessage(), which makes it available again.
	//
	// `azservicebus.ReceiveAndDelete` causes Service Bus to remove the message as soon
	// as it's received.
	//
	// More information about receive modes:
	// https://docs.microsoft.com/azure/service-bus-messaging/message-transfers-locks-settlement#settling-receive-operations
	ReceiveMode ReceiveMode

	// SubQueue should be set to connect to the sub queue (ex: dead letter queue)
	// of the queue or subscription.
	SubQueue SubQueue

	// AutoComplete controls whether processor will automatically settle messages
	// using the return value of your `handleMessage` function.
	// If your handleMessage function returns a non-nil error the message will be abandoned.
	// If your handleMessage function returns nil, the message will be completed.
	//
	// This option is enabled, by default.
	AutoComplete *bool

	// AutoLockRenewal controls whether the processor will automatically renew your message
	// locks in the background, and for how long.
	// To disable, set to 0.
	//
	// Defaults to 5 minutes
	MaxAutoLockRenewal *time.Duration

	// MaxConcurrentCalls controls the maximum number of message processing
	// goroutines that are active at any time.
	// Default is 1.
	MaxConcurrentCalls int
}

// Processor is a push-based receiver for Service Bus.
type Processor struct {
	receiveMode        ReceiveMode
	autoComplete       bool
	maxConcurrentCalls int
	maxAutoLockRenewal time.Duration

	settler   settler
	amqpLinks internal.AMQPLinks

	mu *sync.Mutex

	userMessageHandler func(message *ReceivedMessage) error
	userErrorHandler   func(err error)

	receiversCtx    context.Context
	cancelReceivers func()

	wg sync.WaitGroup

	finiteRetrier  internal.Retrier
	cleanupOnClose func()

	messageLockMu sync.Mutex
	messageLocks  map[[16]byte]context.CancelFunc
}

func applyProcessorOptions(processor *Processor, entity *entity, options *ProcessorOptions) error {
	processor.maxConcurrentCalls = 1
	processor.receiveMode = PeekLock
	processor.autoComplete = true
	processor.maxAutoLockRenewal = 5 * time.Minute

	if options == nil {
		return nil
	}

	if options.AutoComplete != nil {
		processor.autoComplete = *options.AutoComplete
	}

	if options.MaxAutoLockRenewal != nil {
		processor.maxAutoLockRenewal = *options.MaxAutoLockRenewal
	}

	if err := checkReceiverMode(options.ReceiveMode); err != nil {
		return err
	}

	processor.receiveMode = options.ReceiveMode

	if err := entity.SetSubQueue(options.SubQueue); err != nil {
		return err
	}

	if options.MaxConcurrentCalls > 0 {
		processor.maxConcurrentCalls = options.MaxConcurrentCalls
	}

	return nil
}

func newProcessor(ns internal.NamespaceWithNewAMQPLinks, entity *entity, cleanupOnClose func(), options *ProcessorOptions) (*Processor, error) {
	processor := &Processor{
		// TODO: make this configurable
		finiteRetrier: internal.NewBackoffRetrier(internal.BackoffRetrierParams{
			Factor:     1.5,
			Min:        time.Second,
			Max:        time.Minute,
			MaxRetries: 10,
		}),
		cleanupOnClose: cleanupOnClose,
		mu:             &sync.Mutex{},
		messageLocks:   map[[16]byte]context.CancelFunc{},
	}

	if err := applyProcessorOptions(processor, entity, options); err != nil {
		return nil, err
	}

	entityPath, err := entity.String()

	if err != nil {
		return nil, err
	}

	processor.amqpLinks = ns.NewAMQPLinks(entityPath, func(ctx context.Context, session internal.AMQPSession) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
		linkOptions := createLinkOptions(processor.receiveMode, entityPath)
		_, receiver, err := createReceiverLink(ctx, session, linkOptions)

		if err != nil {
			return nil, nil, err
		}

		if err := receiver.IssueCredit(uint32(processor.maxConcurrentCalls)); err != nil {
			_ = receiver.Close(ctx)
			return nil, nil, err
		}

		return nil, receiver, nil
	})

	processor.settler = newMessageSettler(processor.amqpLinks, processor.finiteRetrier)
	processor.receiversCtx, processor.cancelReceivers = context.WithCancel(context.Background())

	return processor, nil
}

// Start will start receiving messages from the queue or subscription.
//
//   if err := processor.Start(context.TODO(), messageHandler, errorHandler); err != nil {
//     log.Fatalf("Processor failed to start: %s", err.Error())
//   }
//
// Any errors that occur (such as network disconnects, failures in handleMessage) will be
// sent to your handleError function. The processor will retry and restart as needed -
// no user intervention is required.
func (p *Processor) Start(ctx context.Context, handleMessage func(message *ReceivedMessage) error, handleError func(err error)) error {
	ctx, span := tab.StartSpan(ctx, tracing.SpanProcessorLoop)
	defer span.End()

	err := func() error {
		p.mu.Lock()
		defer p.mu.Unlock()

		if p.userMessageHandler != nil {
			return errors.New("processor already started")
		}

		p.userMessageHandler = handleMessage
		p.userErrorHandler = func(err error) {
			if sberrors.IsCancelError(err) {
				return
			}

			handleError(err)
		}

		p.receiversCtx, p.cancelReceivers = context.WithCancel(ctx)

		return nil
	}()

	if err != nil {
		return err
	}

	for {
		retrier := p.finiteRetrier.Copy()

		for retrier.Try(p.receiversCtx) {
			if err := p.subscribe(); err != nil {
				if sberrors.IsCancelError(err) {
					break
				}
			}
		}

		select {
		case <-p.receiversCtx.Done():
			// check, did they cancel or did we cancel?
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return nil
			}
		default:
		}
	}

}

// Close will wait for any pending callbacks to complete.
// NOTE: Close() cannot be called synchronously in a message
// or error handler. You must run it asynchronously using
// `go processor.Close(ctx)` or similar.
func (p *Processor) Close(ctx context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.amqpLinks.ClosedPermanently() {
		return nil
	}

	ctx, span := tab.StartSpan(ctx, tracing.SpanProcessorClose)
	defer span.End()

	defer func() {
		if err := p.amqpLinks.Close(ctx, true); err != nil {
			span.Logger().Debug(fmt.Sprintf("Error closing amqpLinks on processor.Close(): %s", err.Error()))
		}
	}()

	p.cleanupOnClose()

	_, receiver, _, _, err := p.amqpLinks.Get(ctx)

	if err != nil {
		span.Logger().Error(err)
		return err
	}

	if err := receiver.DrainCredit(ctx); err != nil {
		span.Logger().Error(err)
		// fall through for now and just let whatever is going on finish
		// otherwise they might not be able to actually close.
	}

	p.cancelReceivers()
	return utils.WaitForGroupOrContext(ctx, &p.wg)
}

// CompleteMessage completes a message, deleting it from the queue or subscription.
func (p *Processor) CompleteMessage(ctx context.Context, message *ReceivedMessage) error {
	if err := p.settler.CompleteMessage(ctx, message); err != nil {
		return err
	}

	p.cancelRenewal(message.LockToken)
	return nil
}

// AbandonMessage will cause a message to be returned to the queue or subscription.
// This will increment its delivery count, and potentially cause it to be dead lettered
// depending on your queue or subscription's configuration.
func (p *Processor) AbandonMessage(ctx context.Context, message *ReceivedMessage) error {
	if err := p.settler.AbandonMessage(ctx, message); err != nil {
		return err
	}

	p.cancelRenewal(message.LockToken)
	return nil
}

// DeferMessage will cause a message to be deferred. Deferred messages
// can be received using `Receiver.ReceiveDeferredMessages`.
func (p *Processor) DeferMessage(ctx context.Context, message *ReceivedMessage) error {
	if err := p.settler.DeferMessage(ctx, message); err != nil {
		return err
	}

	p.cancelRenewal(message.LockToken)
	return nil
}

// DeadLetterMessage settles a message by moving it to the dead letter queue for a
// queue or subscription. To receive these messages create a processor with `Client.NewProcessorForQueue()`
// or `Client.NewProcessorForSubscription()` using the `ProcessorOptions.SubQueue` option.
func (p *Processor) DeadLetterMessage(ctx context.Context, message *ReceivedMessage, options *DeadLetterOptions) error {
	if err := p.settler.DeadLetterMessage(ctx, message, options); err != nil {
		return err
	}

	p.cancelRenewal(message.LockToken)
	return nil
}

func (p *Processor) cancelRenewal(lockToken [16]byte) {
	p.messageLockMu.Lock()
	cancel := p.messageLocks[lockToken]
	delete(p.messageLocks, lockToken)
	p.messageLockMu.Unlock()

	cancel()
}

func (p *Processor) startMessageLockRenewal(receivedMessage *ReceivedMessage) {
	if p.maxAutoLockRenewal == 0 {
		return
	}

	ctx, cancel := context.WithDeadline(p.receiversCtx, time.Now().Add(p.maxAutoLockRenewal))
	go autoRenewMessageLock(ctx, internal.DefaultInfiniteRetrier, p.amqpLinks, receivedMessage)

	p.messageLockMu.Lock()
	p.messageLocks[receivedMessage.LockToken] = cancel
	p.messageLockMu.Unlock()
}

// subscribe continually receives messages from Service Bus, stopping
// if a fatal link/connection error occurs.
func (p *Processor) subscribe() error {
	p.wg.Add(1)
	defer p.wg.Done()

	for {
		_, receiver, _, linkRevision, err := p.amqpLinks.Get(p.receiversCtx)

		if err != nil {
			if err := p.amqpLinks.RecoverIfNeeded(p.receiversCtx, linkRevision, err); err != nil {
				p.userErrorHandler(err)
				return err
			}
		}

		amqpMessage, err := receiver.Receive(p.receiversCtx)

		if err != nil {
			if err := p.amqpLinks.RecoverIfNeeded(p.receiversCtx, linkRevision, err); err != nil {
				p.userErrorHandler(err)
			}

			return nil
		}

		if amqpMessage == nil {
			// amqpMessage shouldn't be nil here, but somehow it is.
			// need to track this down in the AMQP library.
			continue
		}

		p.wg.Add(1)

		go func() {
			defer p.wg.Done()

			// purposefully avoiding using `ctx`. We always let processing complete
			// for message threads to avoid potential message loss.
			_ = p.processMessage(context.Background(), receiver, amqpMessage)
		}()
	}
}

func (p *Processor) RenewMessageLock(ctx context.Context, msg *ReceivedMessage) (time.Time, error) {
	retrier := p.finiteRetrier.Copy()

	for retrier.Try(ctx) {
		var mgmt internal.MgmtClient
		_, _, mgmt, _, err := p.amqpLinks.Get(ctx)

		if err == nil {
			var newExpirationTimes []time.Time
			newExpirationTimes, err = mgmt.RenewLocks(ctx, msg.rawAMQPMessage.LinkName(), []amqp.UUID{msg.LockToken})

			if err == nil {
				return newExpirationTimes[0], nil
			}
		}

		sbe := sberrors.AsServiceBusError(ctx, err)

		if sbe.Fix != sberrors.FixByRetrying {
			return time.Time{}, sbe
		}
	}

	return time.Time{}, ctx.Err()
}

func (p *Processor) processMessage(ctx context.Context, receiver internal.AMQPReceiver, amqpMessage *amqp.Message) error {
	ctx, span := tab.StartSpan(ctx, tracing.SpanProcessorMessage)
	defer span.End()

	receivedMessage := newReceivedMessage(ctx, amqpMessage)
	p.startMessageLockRenewal(receivedMessage)

	messageHandlerErr := p.userMessageHandler(receivedMessage)

	if messageHandlerErr != nil {
		p.userErrorHandler(messageHandlerErr)
	}

	if p.autoComplete {
		var settleErr error

		if messageHandlerErr != nil {
			settleErr = p.AbandonMessage(ctx, receivedMessage)
		} else {
			settleErr = p.CompleteMessage(ctx, receivedMessage)
		}

		if settleErr != nil {
			p.userErrorHandler(fmt.Errorf("failed to settle message with ID '%s': %w", receivedMessage.ID, settleErr))
			return settleErr
		}
	}

	select {
	case <-p.receiversCtx.Done():
		return nil
	default:
	}

	if err := receiver.IssueCredit(1); err != nil {
		if !sberrors.IsDrainingError(err) {
			p.userErrorHandler(err)
			return fmt.Errorf("failed issuing additional credit, processor will be restarted: %w", err)
		}
	}

	return nil
}

func checkReceiverMode(receiveMode ReceiveMode) error {
	if receiveMode == PeekLock || receiveMode == ReceiveAndDelete {
		return nil
	} else {
		return fmt.Errorf("Invalid receive mode %d, must be either azservicebus.PeekLock or azservicebus.ReceiveAndDelete", receiveMode)
	}
}
