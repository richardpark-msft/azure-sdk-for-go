// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package test

import (
	"context"
	"fmt"
	"sort"
	"sync/atomic"
	"testing"

	aztracing "github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/require"
)

type TestSpan struct {
	Status     *aztracing.SpanStatus
	StatusDesc *string
	Attrs      map[string]any

	ID       int64
	ParentID int64
	Ended    bool
}

type TestProvider struct {
	t          *testing.T
	nextSpanID int64

	AllSpans []TestSpan
}

func (tp *TestProvider) NewTracer(name, version string) aztracing.Tracer {
	type spanIDKey int

	newSpanFn := func(ctx context.Context, spanName string, options *aztracing.SpanOptions) (context.Context, aztracing.Span) {
		if options == nil {
			panic("Every span created in this library should set options")
		}

		id := atomic.AddInt64(&tp.nextSpanID, 1)

		span := &TestSpan{
			Attrs: map[string]any{},
			ID:    id,
		}

		for _, attr := range options.Attributes {
			span.Attrs[attr.Key] = attr.Value
		}

		// grab the parent span ID
		parentID, ok := ctx.Value(spanIDKey(0)).(int64)

		if ok {
			// we have one!
			span.ParentID = parentID
		} else {
			//panic("No parent span!")
		}

		// we're the parent now.
		ctx = context.WithValue(ctx, spanIDKey(0), span.ID)

		return ctx, aztracing.NewSpan(aztracing.SpanImpl{
			SetAttributes: func(a ...aztracing.Attribute) {
				require.False(tp.t, span.Ended)

				for _, attr := range a {
					span.Attrs[attr.Key] = attr.Value
				}
			},
			SetStatus: func(ss aztracing.SpanStatus, s string) {
				require.False(tp.t, span.Ended)
				require.Nil(tp.t, span.Status)

				span.Status = &ss
				span.StatusDesc = &s
			},
			End: func() {
				require.False(tp.t, span.Ended)

				span.Ended = true
				text := ""

				if options != nil {
					attrs := sortAttributes(span.Attrs)

					for _, attr := range attrs {
						text += fmt.Sprintf("\n  %s=%v", attr.Key, attr.Value)
					}
				}

				statusDesc := ""

				if span.StatusDesc != nil {
					statusDesc = *span.StatusDesc
				}

				fmt.Printf("\n===> TRACE spanName(id: %d, parent: %d): %q, spanKind: %d, status: %d, statusDesc: %q, attrs: %s\n===> END TRACE\n", span.ID, span.ParentID, spanName, options.Kind, *span.Status, statusDesc, text)
			},
		})
	}

	return aztracing.NewTracer(newSpanFn, &aztracing.TracerOptions{})
}

func sortAttributes(m map[string]any) []aztracing.Attribute {
	var attrs []aztracing.Attribute

	for k, v := range m {
		attrs = append(attrs, aztracing.Attribute{Key: k, Value: v})
	}

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Key < attrs[j].Key
	})

	return attrs
}

func NewProviderForTests(t *testing.T) aztracing.Provider {
	return aztracing.NewProvider((&TestProvider{t: t, nextSpanID: 100}).NewTracer, nil)
}
