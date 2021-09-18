package azservicebus

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

type createLinkFunc func(ctx context.Context, session *amqp.Session) (*amqp.Sender, *amqp.Receiver, error)

// Links groups together:
// - An AMQP link (either *amqp.Sender or *amqp.Receiver)
// - A `$management` link
// - an AMQP Session
// And centralizes recovery (`Recover`) and creation (`Get`) of
// links.
type links struct {
	EntityPath     string
	ManagementPath string
	createLink     createLinkFunc

	mu sync.RWMutex

	// mgmt lets you interact with the $management link for your entity.
	mgmt *mgmtClient

	// the AMQP session for either the 'sender' or 'receiver' link
	session *amqp.Session

	// these are mutually exclusive - we don't expect `linkCreatorFn`
	// to populate both.
	sender   *amqp.Sender
	receiver *amqp.Receiver

	// the current 'revision' of our set of links.
	// starts at 1, increments each time you call Recover().
	revision uint64

	// whether this links set has been closed permanently (via Close)
	// Recover() does not affect this value.
	closedPermanently bool

	cancelAuthRefreshLink     func() <-chan struct{}
	cancelAuthRefreshMgmtLink func() <-chan struct{}

	ns *internal.Namespace
}

// NewLinkish creates a session, starts the claim refresher and creates an associated
// management link for a specific entity path.
func newLinks(ns *internal.Namespace, entityPath string, createLink createLinkFunc) *links {
	l := &links{
		EntityPath:        entityPath,
		ManagementPath:    fmt.Sprintf("%s/$management", entityPath),
		createLink:        createLink,
		closedPermanently: false,
		revision:          1,
		ns:                ns,
	}

	return l
}

// Recover will recycle all associated links (mgmt, receiver, sender and session)
// and recreate them using the link.linkCreator function.
func (l *links) Recover(ctx context.Context) error {
	l.mu.RLock()
	closedPermanently := l.closedPermanently
	l.mu.RUnlock()

	if closedPermanently {
		return amqp.ErrLinkClosed
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.revision++

	if err := l.closeWithoutLocking(ctx, false); err != nil {
		tab.For(ctx)
	}

	return l.initWithoutLocking(ctx)
}

// Get will initialize a session and call its link.linkCreator function.
// If this link has been closed via Close() it will return ErrLinkClosed.
func (l *links) Get(ctx context.Context) (*amqp.Sender, *amqp.Receiver, *mgmtClient, uint64, error) {
	l.mu.RLock()
	sender, receiver, mgmt, closedPermanently := l.sender, l.receiver, l.mgmt, l.closedPermanently
	l.mu.RUnlock()

	if closedPermanently {
		return nil, nil, nil, 0, amqp.ErrLinkClosed
	}

	if sender != nil || receiver != nil {
		return sender, receiver, mgmt, 0, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.initWithoutLocking(ctx); err != nil {
		return nil, nil, nil, 0, err
	}

	return l.sender, l.receiver, l.mgmt, 0, nil
}

// Close will close the the link permanently.
// Any further calls to Get()/Recover() to return ErrLinkClosed.
func (l *links) Close(ctx context.Context) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.closeWithoutLocking(ctx, true)
}

// initWithoutLocking will create a new link, unconditionally.
func (l *links) initWithoutLocking(ctx context.Context) error {
	client, err := l.ns.GetAMQPClient(ctx)

	if err != nil {
		return err
	}

	l.cancelAuthRefreshLink, err = l.ns.NegotiateClaim(ctx, l.EntityPath)

	if err != nil {
		l.closeWithoutLocking(ctx, false)
		return err
	}

	l.cancelAuthRefreshMgmtLink, err = l.ns.NegotiateClaim(ctx, l.ManagementPath)

	if err != nil {
		l.closeWithoutLocking(ctx, false)
		return err
	}

	l.session, err = client.NewSession()

	if err != nil {
		l.closeWithoutLocking(ctx, false)
		return err
	}

	l.sender, l.receiver, err = l.createLink(ctx, l.session)

	if err != nil {
		l.closeWithoutLocking(ctx, false)
		return err
	}

	return nil
}

// close closes the link.
// NOTE: No locking is done in this function, call `Close` if you require locking.
func (l *links) closeWithoutLocking(ctx context.Context, permanent bool) error {
	if l.closedPermanently {
		return amqp.ErrLinkClosed
	}

	defer func() {
		l.closedPermanently = true
	}()

	var messages []string

	l.cancelAuthRefreshLink()
	l.cancelAuthRefreshMgmtLink()

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
