// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package disttrace

import (
	"context"

	aztracing "github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/constants"
)

// SpanAttributeMessageCount is used to add in the message count received or sent
// per operation.
const SpanAttributeMessageCount = "messaging.batch.message_count"
const SpanOperation = "messaging.operation"

// TODO: is this the right way to report a connection event?
const SpanNewConnection = "Microsoft.ServiceBus.NewConnection"
const SpanNewSession = "Microsoft.ServiceBus.NewSession"
const SpanNewLinks = "Microsoft.ServiceBus.NewLinks"

type Tracer struct {
	aztracing.Tracer
	baseAttrs []aztracing.Attribute
}

type TracerArgs struct {
	PeerName     string // PeerName is the host name of the remote service (ex: <name>.servicebus.windows.net)
	QueueOrTopic string // name of the queue or topic
}

func NewTracer(prov aztracing.Provider, args TracerArgs) Tracer {
	tracer := prov.NewTracer("Microsoft.ServiceBus", constants.Version)

	return Tracer{
		Tracer: tracer,
		baseAttrs: []aztracing.Attribute{
			{Key: "az.namespace", Value: "Microsoft.ServiceBus"},
			// TODO: I'd like to add in the link ID here.
			// {Key: "id", Value: args.ID},
			{Key: "messaging.destination.name", Value: args.QueueOrTopic},
			{Key: "messaging.system", Value: "servicebus"},
			{Key: "net.peer.name", Value: args.PeerName},
		},
	}
}

type SettleOperation string

const (
	SettleOperationComplete   SettleOperation = "ServiceBus.complete"
	SettleOperationAbandon    SettleOperation = "ServiceBus.abandon"
	SettleOperationDeadLetter SettleOperation = "ServiceBus.deadLetter"
	SettleOperationDefer      SettleOperation = "ServiceBus.defer"
)

func (sbt Tracer) StartSettleSpan(ctx context.Context, spanName SettleOperation) (context.Context, aztracing.Span) {
	// https://gist.github.com/lmolkova/e4215c0f44a49ef824983382762e6b92#other-api-calls-that-involve-communication-with-service

	return sbt.Start(ctx, string(spanName), &aztracing.SpanOptions{
		Kind: aztracing.SpanKindClient,
		Attributes: append(sbt.baseAttrs,
			aztracing.Attribute{Key: SpanOperation, Value: "settle"},
		),
	})
}

func (sbt Tracer) StartSBSpan(ctx context.Context, spanName string, options *aztracing.SpanOptions) (context.Context, aztracing.Span) {
	var opts aztracing.SpanOptions

	if options != nil {
		opts = *options
	}

	opts.Attributes = append(opts.Attributes, sbt.baseAttrs...)
	return sbt.Start(ctx, spanName, &opts)
}

func (sbt Tracer) StartReceivingSpan(ctx context.Context) (context.Context, aztracing.Span) {
	// https://gist.github.com/lmolkova/e4215c0f44a49ef824983382762e6b92#receiving-messages
	return sbt.Start(ctx, "ServiceBus.receive", &aztracing.SpanOptions{
		Kind: aztracing.SpanKindClient,
		Attributes: append(sbt.baseAttrs,
			aztracing.Attribute{Key: SpanOperation, Value: "receive"},
		),
	})
}

func (sbt Tracer) StartSendingSpan(ctx context.Context, count int) (context.Context, aztracing.Span) {
	// https://gist.github.com/lmolkova/e4215c0f44a49ef824983382762e6b92#sending-messages
	return sbt.Start(ctx, "ServiceBus.send", &aztracing.SpanOptions{
		Kind: aztracing.SpanKindClient,
		Attributes: append(sbt.baseAttrs,
			aztracing.Attribute{Key: SpanOperation, Value: "publish"},
			aztracing.Attribute{Key: SpanAttributeMessageCount, Value: count},
		),
	})
}

// TODO: we're not quite ready for this one - we need access to the SpanContext to create a Diagnostic-Id.
func (sbt Tracer) StartMessageSpan(ctx context.Context, count int) (context.Context, aztracing.Span) {
	return sbt.Start(ctx, "ServiceBus.message", &aztracing.SpanOptions{
		Kind:       aztracing.SpanKindProducer,
		Attributes: sbt.baseAttrs,
	})
}

// SetMessageCountAttr sets the message count attribute. Used by Receivers.
func SetMessageCountAttr(span aztracing.Span, count int) {
	span.SetAttributes(
		aztracing.Attribute{Key: SpanAttributeMessageCount, Value: count},
	)
}

// SetSpanStatus sets the span status to either SpanStatusError (if err is non-nil)
// or SpanStatusOK if err is nil.
// NOTE, like OpenTelemetry, the description is only included in a status when the code is for an error.
func SetSpanStatus(span aztracing.Span, err error, desc string) error {
	if err != nil {
		span.AddError(err)
		span.SetStatus(aztracing.SpanStatusError, desc)
		return err
	}

	span.SetStatus(aztracing.SpanStatusOK, desc)
	return nil
}
