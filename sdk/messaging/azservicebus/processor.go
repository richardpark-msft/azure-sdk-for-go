// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/Azure/go-amqp"
)

type processorConfig struct {
	FullEntityPath string
	ReceiveMode    ReceiveMode

	Entity entity

	// determines if auto completion or abandonment of messages
	// happens based on the return value of user's processMessage handler.
	ShouldAutoComplete bool
	MaxConcurrentCalls int
}

// Processor is a push-based receiver for Service Bus.
type Processor struct {
	// configuration data that is read-only after the Processor has been created
	config processorConfig

	mu *sync.Mutex

	receiversCtx    context.Context
	cancelReceivers func()

	processorCtx    context.Context
	cancelProcessor func()

	activeReceiversWg *sync.WaitGroup

	// replaceable for unit tests
	subscribe func(
		ctx context.Context,
		links *links,
		shouldAutoComplete bool,
		handleMessage func(message *ReceivedMessage) error,
		notifyError func(err error)) bool

	links *links
}

// ProcessorOption represents an option on the Processor.
// Some examples:
// - `ProcessorWithReceiveMode` to configure the receive mode,
// - `ProcessorWithQueue` to target a queue.
type ProcessorOption func(processor *Processor) error

// ProcessorWithSubQueue allows you to open the sub queue (ie: dead letter queues, transfer dead letter queues)
// for a queue or subscription.
func ProcessorWithSubQueue(subQueue SubQueue) ProcessorOption {
	return func(receiver *Processor) error {
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

// ProcessorWithReceiveMode controls the receive mode for the processor.
func ProcessorWithReceiveMode(receiveMode ReceiveMode) ProcessorOption {
	return func(processor *Processor) error {
		if receiveMode != PeekLock && receiveMode != ReceiveAndDelete {
			return fmt.Errorf("invalid receive mode specified %d", receiveMode)
		}

		processor.config.ReceiveMode = receiveMode
		return nil
	}
}

// ProcessorWithQueue configures a processor to connect to a queue.
func ProcessorWithQueue(queue string) ProcessorOption {
	return func(processor *Processor) error {
		processor.config.Entity.Queue = queue
		return nil
	}
}

// ProcessorWithSubscription configures a processor to connect to a subscription
// associated with a topic.
func ProcessorWithSubscription(topic string, subscription string) ProcessorOption {
	return func(processor *Processor) error {
		processor.config.Entity.Topic = topic
		processor.config.Entity.Subscription = subscription
		return nil
	}
}

// ProcessorWithAutoComplete enables or disables auto-completion/abandon of messages
// When this option is enabled the result of the `processMessage` handler determines whether
// the message is abandoned (if an `error` is returned) or completed (if `nil` is returned).
// This option is enabled, by default.
func ProcessorWithAutoComplete(enableAutoCompleteMessages bool) ProcessorOption {
	return func(processor *Processor) error {
		processor.config.ShouldAutoComplete = enableAutoCompleteMessages
		return nil
	}
}

// ProcessorWithMaxConcurrentCalls controls the maximum number of message processing
// goroutines that are active at any time.
// Default is 1.
func ProcessorWithMaxConcurrentCalls(maxConcurrentCalls int) ProcessorOption {
	return func(processor *Processor) error {
		processor.config.MaxConcurrentCalls = maxConcurrentCalls
		return nil
	}
}

func newProcessor(ns *internal.Namespace, options ...ProcessorOption) (*Processor, error) {
	processor := &Processor{
		config: processorConfig{
			ReceiveMode:        PeekLock,
			ShouldAutoComplete: true,
			MaxConcurrentCalls: 1,
			RetryOptions: struct {
				Times int
				Delay time.Duration
			}{
				// TODO: allow these to be configured.
				Times: 10,
				Delay: time.Second * 5,
			},
		},

		mu:                &sync.Mutex{},
		activeReceiversWg: &sync.WaitGroup{},
		subscribe:         subscribe,
	}

	for _, opt := range options {
		if err := opt(processor); err != nil {
			return nil, err
		}
	}

	entityPath, err := processor.config.Entity.String()

	if err != nil {
		return nil, err
	}

	processor.config.FullEntityPath = entityPath
	processor.links = internal.NewLinks(ns, entityPath, processor.createReceiverLink)

	processor.processorCtx, processor.cancelProcessor = context.WithCancel(context.Background())
	processor.receiversCtx, processor.cancelReceivers = context.WithCancel(context.Background())

	return processor, nil
}

// compatible with `CreateLinkFunc`
func (p *Processor) createReceiverLink(ctx context.Context, session *amqp.Session) (*amqp.Sender, *amqp.Receiver, error) {
	options := createLinkOptions(p.config.ReceiveMode, p.config.FullEntityPath)
	return createReceiverLink(ctx, session, options)
}

// Start will start receiving messages from the queue or subscription.
//
//   if err := processor.Start(messageHandler, errorHandler); err != nil {
//     log.Fatalf("Processor failed to start: %s", err.Error())
//   }
//
//   <- processor.Done()
//
// Any errors that occur (such as network disconnects, failures in handleMessage) will be
// sent to your handleError function. The processor will retry and restart as needed -
// no user intervention is required.
func (p *Processor) Start(handleMessage func(message *ReceivedMessage) error, handleError func(err error)) error {
	select {
	case <-p.Done():
		return errClosed{link: "processor"}
	default:
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.activeReceiversWg.Add(1)

	go func(ctx context.Context) {
		defer p.activeReceiversWg.Done()

		for {
			backoff := internal.DefaultBackoffPolicy

			for i := 0; i < p.config.RetryOptions.Times; i++ {
				// we retry infinitely, but do it in the pattern they specify via their retryOptions for each "round" of retries.
				retry := p.subscribe(ctx, links, p.config.MaxConcurrentCalls, p.config.ShouldAutoComplete, handleMessage, handleError)

				if !retry {
					break
				}

				select {
				case <-time.After(backoff.Duration()):
					if err := p.links.Recover(ctx); err != nil {
						handleError(err)
					}
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}
	}(p.receiversCtx)

	return nil
}

// Done returns a channel that will be close()'d when the Processor
// has been closed.
func (p *Processor) Done() <-chan struct{} {
	return p.processorCtx.Done()
}

// Close will wait for any pending callbacks to complete.
func (p *Processor) Close(ctx context.Context) error {
	select {
	case <-p.Done():
		return nil
	default:
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.cancelReceivers()

	err := utils.WaitForGroupOrContext(ctx, p.activeReceiversWg)

	// now unlock anyone _external_ to the processor that's waiting for us to exit or close.
	p.cancelProcessor()

	return err
}

func subscribe(
	ctx context.Context,
	links *links,
	maxConcurrency int,
	shouldAutoComplete bool,
	handleMessage func(message *ReceivedMessage) error,
	notifyError func(err error)) bool {

	_, receiver, mgmt, linkRevision, err := links.Get(ctx)

	if err != nil {
		notifyError(err)
		return true
	}

	if err := receiver.IssueCredit(uint32(p.config.MaxConcurrentCalls)); err != nil {
		// something is very broken - definitely want to try to recreate the link.
		notifyError(err)
		return true
	}

	activeCallbacksWg := &sync.WaitGroup{}
	notifyErrorAsync := wrapNotifyError(notifyError, activeCallbacksWg)
	maxCh := make(chan bool, maxConcurrency)

	for {
		message, err := receiver.Receive(ctx)

		maxCh <- true

		go func() {
			defer func() { <-maxCh }()

			// this shouldn't happen since we do a `select` above that prevents it.
			// errors from their handler are sent to their error handler but do not terminate the
			// subscription.
			activeCallbacksWg.Add(1)
			defer activeCallbacksWg.Done()

			handleSingleMessage(handleMessage, notifyErrorAsync, shouldAutoComplete, receiver, message)

			// user callback completes and they get a new credit
			if err := receiver.IssueCredit(1); err != nil {
				notifyErrorAsync(err)
			}

			return nil
		}()
	}

	select {
	case <-ctx.Done():
		notifyErrorAsync(ctx.Err())
		activeCallbacksWg.Wait()

		// user's quitting, don't restart
		return false
	default:
		activeCallbacksWg.Wait()

		// we should retry since the listen handle can be closed if we did a .Recover() on the receiver.
		return true
	}
}

func handleSingleMessage(handleMessage func(message *ReceivedMessage) error, notifyErrorAsync func(err error), shouldAutoComplete bool, settler settler, message *amqp.Message) {
	err := handleMessage(convertToReceivedMessage(legacyMessage))

	if err != nil {
		notifyErrorAsync(err)
	}

	var settleErr error

	if shouldAutoComplete {
		// NOTE: we ignore the passed in context. Since we're settling behind the scenes
		// it's nice to wrap it up so users don't have to track it.
		if err != nil {
			settleErr = receiver.AbandonMessage(context.Background(), legacyMessage)
		} else {
			settleErr = receiver.CompleteMessage(context.Background(), legacyMessage)
		}

		if settleErr != nil {
			notifyErrorAsync(settleErr)
		}
	}
}

//
// settlement methods
// TODO: in other processor implementations this is implemented in the argument for the processMessage
// callback. You need some sort of association or else you have to track message <-> receiver mappings.
//

// CompleteMessage completes a message, deleting it from the queue or subscription.
func (p *Processor) CompleteMessage(ctx context.Context, message *ReceivedMessage) error {
	return message.legacyMessage.Complete(ctx)
}

// DeadLetterMessage settles a message by moving it to the dead letter queue for a
// queue or subscription.
func (p *Processor) DeadLetterMessage(ctx context.Context, message *ReceivedMessage) error {
	// TODO: expand to let them set the reason and description.
	return message.legacyMessage.DeadLetter(ctx, nil)
}

// AbandonMessage will cause a message to be returned to the queue or subscription.
// This will increment its delivery count, and potentially cause it to be dead lettered
// depending on your queue or subscription's configuration.
func (p *Processor) AbandonMessage(ctx context.Context, message *ReceivedMessage) error {
	return message.legacyMessage.Abandon(ctx)
}

// DeferMessage will cause a message to be deferred.
// Messages that are deferred by can be retrieved using `Receiver.ReceiveDeferredMessages()`.
func (p *Processor) DeferMessage(ctx context.Context, message *ReceivedMessage) error {
	return message.legacyMessage.Defer(ctx)
}

func wrapNotifyError(fn func(err error), wg *sync.WaitGroup) func(err error) {
	return func(err error) {
		if err == nil {
			return
		}

		wg.Add(1)
		go func() {
			fn(err)
			wg.Done()
		}()
	}
}
