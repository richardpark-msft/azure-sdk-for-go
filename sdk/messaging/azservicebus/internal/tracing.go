package internal

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/go-amqp"
)

func AddTracingLink(span tracing.Span, spanContext tracing.Span) {
	//panic("Not implemented")
}

// https://gist.github.com/lmolkova/e4215c0f44a49ef824983382762e6b92#:~:text=String%20traceparent%20%3D%20getDiagnosticId(msgSpan.getContext())%3B
func GetTraceParent(span tracing.Span) string {
	//	panic("Not implemented")
	//       msg.getProperties().putIfAbsent("Diagnostic-Id", traceparent);
	return ""
}

// TODO: it's supposed to be a SpanContext. That might literally just be a `context.Context`
// but not entirely sure.
func GetSpanContext(msg *amqp.Message) (tracing.Span, bool) {
	return tracing.Span{}, false
}
