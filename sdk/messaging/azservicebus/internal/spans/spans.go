// Package spans centralizes all the span names used in the Service Bus
// library.
// Span names that end in 'Fmt' are format strings, which will be filled in with
// an appropriate link type ('Sender', 'Receiver', 'Mgmt')
package spans

import (
	"context"

	"github.com/devigned/tab"
)

const (
	// RecoverFmt happens when a link recovers (typically after a detach or network disconnect occurs)
	RecoverFmt  = "sb.%s.Recover"
	InitFmt     = "sb.%s.Init"
	SendMessage = "sb.Sender.SendMessage"
)

func forMessage(ctx context.Context, operationName string, messageID string, sessionID *string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	applyComponentInfo(span)
	attrs := []tab.Attribute{tab.StringAttribute("amqp.message.id", messageID}
	if m.SessionID != nil {
		attrs = append(attrs, tab.StringAttribute("amqp.session.id", sessionID))
	}

	// if m.GroupSequence != nil {
	// 	attrs = append(attrs, tab.Int64Attribute("amqp.sequence_number", int64(*m.GroupSequence)))
	// }
	span.AddAttributes(attrs...)
	return ctx, span
}

func forEntity(ctx context.Context, operationName string, managementPath string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	applyComponentInfo(span)
	span.AddAttributes(tab.StringAttribute("message_bus.destination", managementPath))
	return ctx, span
}

func forSender(ctx context.Context, operationName string, fqdn string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	applyComponentInfo(span)
	span.AddAttributes(
		tab.StringAttribute("span.kind", "producer"),
		// TODO: need to check this against lmolkova's spec.
		tab.StringAttribute("message_bus.destination", fqdn),
	)
	return ctx, span
}

func forReceiver(ctx context.Context, operationName string, fqdn string) (context.Context, tab.Spanner) {
	ctx, span := startConsumerSpanFromContext(ctx, operationName)
	// TODO: need to check this against lmolkova's spec.
	// oddly enough this was previously just the entityPath. Not sure if that was 
	// enough info.
	span.AddAttributes(tab.StringAttribute("message_bus.destination", fqdn))
	return ctx, span
}
