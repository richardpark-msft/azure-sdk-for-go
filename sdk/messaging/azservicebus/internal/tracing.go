// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/spans"
	"github.com/devigned/tab"
)

func (ns *Namespace) startSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	spans.ApplyComponentInfo(span)
	return ctx, span
}

func (em *entityManager) startSpanFromContext(ctx context.Context, operationName string) (context.Context, tab.Spanner) {
	ctx, span := tab.StartSpan(ctx, operationName)
	spans.ApplyComponentInfo(span)
	span.AddAttributes(tab.StringAttribute("span.kind", "client"))
	return ctx, span
}

func applyRequestInfo(span tab.Spanner, req *http.Request) {
	span.AddAttributes(
		tab.StringAttribute("http.url", req.URL.String()),
		tab.StringAttribute("http.method", req.Method),
	)
}

func applyResponseInfo(span tab.Spanner, res *http.Response) {
	if res != nil {
		span.AddAttributes(tab.Int64Attribute("http.status_code", int64(res.StatusCode)))
	}
}
