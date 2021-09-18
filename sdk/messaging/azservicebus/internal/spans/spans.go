// Package spans centralizes all the span names used in the Service Bus
// library.
// Span names that end in 'Fmt' are format strings, which will be filled in with
// an appropriate link type ('Sender', 'Receiver', 'Mgmt')
package spans

import (
	"context"

	"github.com/devigned/tab"
)

// Version is the semantic version number
const Version = "0.1.0"

const (
	// RecoverFmt happens when a link recovers (typically after a detach or network disconnect occurs)
	RecoverFmt  = "sb.%s.Recover"
	InitFmt     = "sb.%s.Init"
	SendMessage = "sb.Sender.SendMessage"
)

func ForMessage(ctx context.Context, operationName string, messageID string, sessionID *string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	ApplyComponentInfo(span)
	attrs := []tab.Attribute{tab.StringAttribute("amqp.message.id", messageID)}

	if sessionID != nil {
		attrs = append(attrs, tab.StringAttribute("amqp.session.id", *sessionID))
	}

	// if m.GroupSequence != nil {
	// 	attrs = append(attrs, tab.Int64Attribute("amqp.sequence_number", int64(*m.GroupSequence)))
	// }
	span.AddAttributes(attrs...)
	return ctx, span
}

func ForEntity(ctx context.Context, operationName string, managementPath string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	ApplyComponentInfo(span)
	span.AddAttributes(tab.StringAttribute("message_bus.destination", managementPath))
	return ctx, span
}

func ForSender(ctx context.Context, operationName string, entityPath string, host string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	ApplyComponentInfo(span)
	span.AddAttributes(
		tab.StringAttribute("span.kind", "producer"),
		tab.StringAttribute("message_bus.destination", entityPath),
		tab.StringAttribute("peer.address", host),
	)
	return ctx, span
}

func ForReceiver(ctx context.Context, operationName string, entityPath string, host string) (context.Context, tab.Spanner) {
	ctx, span := StartConsumerSpanFromContext(ctx, operationName)
	span.AddAttributes(
		tab.StringAttribute("message_bus.destination", entityPath),
		tab.StringAttribute("peer.address", host),
	)
	return ctx, span
}

func ApplyComponentInfo(span tab.Spanner) {
	span.AddAttributes(
		tab.StringAttribute("component", "github.com/Azure/azure-sdk-for-go"),
		tab.StringAttribute("version", Version),
	)
	// ApplyNetworkInfo(span)
}

func StartConsumerSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	ApplyComponentInfo(span)
	span.AddAttributes(tab.StringAttribute("span.kind", "consumer"))
	return ctx, span
}

// func ApplyNetworkInfo(span tab.Spanner) {
// 	hostname, err := os.Hostname()
// 	if err == nil {
// 		span.AddAttributes(
// 			tab.StringAttribute("peer.hostname", hostname),
// 		)
// 	}
// }
