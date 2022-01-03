// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/tracing"
	"github.com/devigned/tab"
)

type AMQPLinks interface {
	EntityPath() string
	ManagementPath() string

	Audience() string

	// Get will initialize a session and call its link.linkCreator function.
	// If this link has been closed via Close() it will return an non retriable error.
	Get(ctx context.Context) (AMQPSender, AMQPReceiver, MgmtClient, uint64, *ServiceBusError)

	// RecoverIfNeeded will check if an error requires recovery, and will recover
	// the link or, possibly, the connection.
	// If the error is fatal it will conform to the errorinfo.NonRetriable interface.
	RecoverIfNeeded(ctx context.Context, linksRevision uint64, err *ServiceBusError) *ServiceBusError

	// Close will close the the link.
	// If permanent is true the link will not be auto-recreated if Get/Recover
	// are called. All functions will return `ErrLinksClosed`
	Close(ctx context.Context, permanent bool) error

	// ClosedPermanently is true if AMQPLinks.Close(ctx, true) has been called.
	ClosedPermanently() bool
}

// amqpLinks manages the set of AMQP links (and detritus) typically needed to work
//  within Service Bus:
// - An *goamqp.Sender or *goamqp.Receiver AMQP link (could also be 'both' if needed)
// - A `$management` link
// - an *goamqp.Session
//
// State management can be done through Recover (close and reopen), Close (close permanently, return failures)
// and Get() (retrieve latest version of all amqpLinks, or create if needed).
type amqpLinks struct {
	entityPath     string
	managementPath string
	audience       string
	createLink     CreateLinkFunc
	baseRetrier    Retrier

	mu sync.RWMutex

	// mgmt lets you interact with the $management link for your entity.
	mgmt MgmtClient

	// the AMQP session for either the 'sender' or 'receiver' link
	session AMQPSessionCloser

	// these are populated by your `createLinkFunc` when you construct
	// the amqpLinks
	sender   AMQPSenderCloser
	receiver AMQPReceiverCloser

	// last connection revision seen by this links instance.
	clientRevision uint64

	// the current 'revision' of our set of links.
	// starts at 1, increments each time you call Recover().
	revision uint64

	// whether this links set has been closed permanently (via Close)
	// Recover() does not affect this value.
	closedPermanently bool

	cancelAuthRefreshLink     func() <-chan struct{}
	cancelAuthRefreshMgmtLink func() <-chan struct{}

	ns NamespaceForAMQPLinks
}

// CreateLinkFunc creates the links, using the given session. Typically you'll only create either an
// *amqp.Sender or a *amqp.Receiver. AMQPLinks handles it either way.
type CreateLinkFunc func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, *ServiceBusError)

// NewAMQPLinks creates a session, starts the claim refresher and creates an associated
// management link for a specific entity path.
func newAMQPLinks(ns NamespaceForAMQPLinks, entityPath string, baseRetrier Retrier, createLink CreateLinkFunc) AMQPLinks {
	l := &amqpLinks{
		entityPath:        entityPath,
		managementPath:    fmt.Sprintf("%s/$management", entityPath),
		audience:          ns.GetEntityAudience(entityPath),
		createLink:        createLink,
		baseRetrier:       baseRetrier,
		closedPermanently: false,
		revision:          1,
		ns:                ns,
	}

	return l
}

// ManagementPath is the management path for the associated entity.
func (links *amqpLinks) ManagementPath() string {
	return links.managementPath
}

// recoverLink will recycle all associated links (mgmt, receiver, sender and session)
// and recreate them using the link.linkCreator function.
func (links *amqpLinks) recoverLink(ctx context.Context, theirLinkRevision *uint64) *ServiceBusError {
	ctx, span := tab.StartSpan(ctx, tracing.SpanRecoverLink)
	defer span.End()

	links.mu.RLock()
	closedPermanently := links.closedPermanently
	ourLinkRevision := links.revision
	links.mu.RUnlock()

	if closedPermanently {
		span.AddAttributes(tab.StringAttribute("outcome", "was_closed_permanently"))
		return sberrClosedPermanently
	}

	if theirLinkRevision != nil && ourLinkRevision > *theirLinkRevision {
		// we've already recovered past their failure.
		span.AddAttributes(
			tab.StringAttribute("outcome", "already_recovered"),
			tab.StringAttribute("lock", "readlock"),
			tab.StringAttribute("revisions", fmt.Sprintf("ours(%d), theirs(%d)", ourLinkRevision, *theirLinkRevision)),
		)
		return nil
	}

	links.mu.Lock()
	defer links.mu.Unlock()

	if theirLinkRevision != nil && ourLinkRevision > *theirLinkRevision {
		// we've already recovered past their failure.
		span.AddAttributes(
			tab.StringAttribute("outcome", "already_recovered"),
			tab.StringAttribute("lock", "writelock"),
			tab.StringAttribute("revisions", fmt.Sprintf("ours(%d), theirs(%d)", ourLinkRevision, *theirLinkRevision)),
		)
		return nil
	}

	if err := links.closeWithoutLocking(ctx, false); err != nil {
		span.Logger().Error(err)
	}

	sbe := links.initWithoutLocking(ctx)

	if sbe != nil {
		span.AddAttributes(tab.StringAttribute("init_error", sbe.String()))
		return sbe
	}

	links.revision++

	span.AddAttributes(
		tab.StringAttribute("outcome", "recovered"),
		tab.StringAttribute("revision_new", fmt.Sprintf("%d", links.revision)),
	)
	return nil
}

// Recover will recover the links or the connection, depending
// on the severity of the error. This function uses the `baseRetrier`
// defined in the links struct.
func (links *amqpLinks) RecoverIfNeeded(ctx context.Context, linksRevision uint64, origErr *ServiceBusError) *ServiceBusError {
	var sbe *ServiceBusError

	if sbe = links.recoverImpl(ctx, 0, linksRevision, origErr); sbe != nil {
		if sbe.RecoveryKind == recoveryKindNonRetriable {
			return sbe
		}

		return sbe
	}

	return nil
}

func (links *amqpLinks) recoverImpl(ctx context.Context, try int, linksRevision uint64, sbe *ServiceBusError) *ServiceBusError {
	if sbe == nil {
		return nil
	}

	if sbe.RecoveryKind == recoveryKindNone {
		return sbe
	}

	_, span := tab.StartSpan(ctx, tracing.SpanRecover)
	defer span.End()

	span.AddAttributes(
		tab.StringAttribute("recovery_kind", string(sbe.RecoveryKind)),
		tab.Int64Attribute("attempt", int64(try)),
	)

	switch sbe.RecoveryKind {
	case recoveryKindNonRetriable:
		span.AddAttributes(
			tab.StringAttribute("error", sbe.String()),
			tab.StringAttribute("error_verbose", fmt.Sprintf("%#v", sbe.AsError())))
		return sbe
	case recoveryKindLink:
		span.AddAttributes(
			tab.StringAttribute("error", sbe.String()),
			tab.StringAttribute("error_verbose", fmt.Sprintf("%#v", sbe.inner)))

		if sbe := links.recoverLink(ctx, &linksRevision); sbe != nil {
			span.AddAttributes(tab.StringAttribute("recoveryFailure", sbe.String()))
			return sbe
		}

		return sbe
	case recoveryKindConnection:
		span.AddAttributes(
			tab.StringAttribute("error", sbe.String()),
			tab.StringAttribute("error_verbose", fmt.Sprintf("%#v", sbe.inner)))

		if sbe := links.recoverConnection(ctx); sbe != nil {
			span.AddAttributes(tab.StringAttribute("recoveryFailure", sbe.String()))
			return sbe
		}

		// unconditionally recover the link if the connection died.
		if sbe := links.recoverLink(ctx, nil); sbe != nil {
			span.AddAttributes(tab.StringAttribute("recoveryFailure", sbe.String()))
			return sbe
		}

		return sbe
	default:
		span.AddAttributes(
			tab.StringAttribute("error", sbe.String()),
			tab.StringAttribute("errorType", fmt.Sprintf("%T", sbe.inner)))

		return sbe
	}
}

func (links *amqpLinks) recoverConnection(ctx context.Context) *ServiceBusError {
	tab.For(ctx).Info("Connection is dead, recovering")

	links.mu.RLock()
	clientRevision := links.clientRevision
	links.mu.RUnlock()

	sbe := links.ns.Recover(ctx, clientRevision)

	if sbe != nil {
		tab.For(ctx).Error(fmt.Errorf("Recover connection failure: %w", sbe))
		return sbe
	}

	return nil
}

// Get will initialize a session and call its link.linkCreator function.
// If this link has been closed via Close() it will return an non retriable error.
// Returns a *serviceBusError
func (l *amqpLinks) Get(ctx context.Context) (AMQPSender, AMQPReceiver, MgmtClient, uint64, *ServiceBusError) {
	l.mu.RLock()
	sender, receiver, mgmt, revision, closedPermanently := l.sender, l.receiver, l.mgmt, l.revision, l.closedPermanently
	l.mu.RUnlock()

	if closedPermanently {
		return nil, nil, nil, 0, sberrClosedPermanently
	}

	if sender != nil || receiver != nil {
		return sender, receiver, mgmt, revision, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if sbe := l.initWithoutLocking(ctx); sbe != nil {
		return nil, nil, nil, 0, sbe
	}

	return l.sender, l.receiver, l.mgmt, l.revision, nil
}

// EntityPath is the full entity path for the queue/topic/subscription.
func (l *amqpLinks) EntityPath() string {
	return l.entityPath
}

// EntityPath is the audience for the queue/topic/subscription.
func (l *amqpLinks) Audience() string {
	return l.audience
}

// ClosedPermanently is true if AMQPLinks.Close(ctx, true) has been called.
func (l *amqpLinks) ClosedPermanently() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.closedPermanently
}

// Close will close the the link permanently.
// Any further calls to Get()/Recover() to return ErrLinksClosed.
func (l *amqpLinks) Close(ctx context.Context, permanent bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.closeWithoutLocking(ctx, permanent)
}

// initWithoutLocking will create a new link, unconditionally.
func (l *amqpLinks) initWithoutLocking(ctx context.Context) *ServiceBusError {
	var sbe *ServiceBusError
	l.cancelAuthRefreshLink, sbe = l.ns.NegotiateClaim(ctx, l.entityPath)

	if sbe != nil {
		if err := l.closeWithoutLocking(ctx, false); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Failure during link cleanup after negotiateClaim: %s", err.Error()))
		}

		return sbe
	}

	l.cancelAuthRefreshMgmtLink, sbe = l.ns.NegotiateClaim(ctx, l.managementPath)

	if sbe != nil {
		if err := l.closeWithoutLocking(ctx, false); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Failure during link cleanup after negotiate claim for mgmt link: %s", err.Error()))
		}
		return sbe
	}

	l.session, l.clientRevision, sbe = l.ns.NewAMQPSession(ctx)

	if sbe != nil {
		if err := l.closeWithoutLocking(ctx, false); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Failure during link cleanup after creating AMQP session: %s", err.Error()))
		}
		return sbe
	}

	l.sender, l.receiver, sbe = l.createLink(ctx, l.session)

	if sbe != nil {
		if err := l.closeWithoutLocking(ctx, false); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Failure during link cleanup after creating link: %s", err.Error()))
		}
		return sbe
	}

	l.mgmt, sbe = l.ns.NewMgmtClient(ctx, l)

	if sbe != nil {
		if err := l.closeWithoutLocking(ctx, false); err != nil {
			tab.For(ctx).Debug(fmt.Sprintf("Failure during link cleanup after creating mgmt client: %s", err.Error()))
		}
		return sbe
	}

	return nil
}

// close closes the link.
// NOTE: No locking is done in this function, call `Close` if you require locking.
func (l *amqpLinks) closeWithoutLocking(ctx context.Context, permanent bool) error {
	if l.closedPermanently {
		return nil
	}

	defer func() {
		if permanent {
			l.closedPermanently = true
		}
	}()

	var messages []string

	if l.cancelAuthRefreshLink != nil {
		l.cancelAuthRefreshLink()
	}

	if l.cancelAuthRefreshMgmtLink != nil {
		l.cancelAuthRefreshMgmtLink()
	}

	if l.sender != nil {
		if err := l.sender.Close(ctx); err != nil {
			messages = append(messages, fmt.Sprintf("amqp sender close error: %s", err.Error()))
		}
		l.sender = nil
	}

	if l.receiver != nil {
		if err := l.receiver.Close(ctx); err != nil {
			messages = append(messages, fmt.Sprintf("amqp receiver close error: %s", err.Error()))
		}
		l.receiver = nil
	}

	if l.session != nil {
		if err := l.session.Close(ctx); err != nil {
			messages = append(messages, fmt.Sprintf("amqp session close error: %s", err.Error()))
		}
		l.session = nil
	}

	if l.mgmt != nil {
		if err := l.mgmt.Close(ctx); err != nil {
			messages = append(messages, fmt.Sprintf("$management link close error: %s", err.Error()))
		}
		l.mgmt = nil
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n"))
	}

	return nil
}
