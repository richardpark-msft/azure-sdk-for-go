// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/atom"
)

// AdminClient allows you to administer resources in a Service Bus Namespace.
// For example, you can create queues, enabling capabilities like partitioning, duplicate detection, etc..
// NOTE: For sending and receiving messages you'll need to use the `Client` type instead.
type AdminClient struct {
	em *atom.EntityManager
}

type AdminClientOptions struct {
	// for future expansion
}

// NewAdminClient creates an AdminClient authenticating using a connection string.
func NewAdminClientWithConnectionString(connectionString string, options *AdminClientOptions) (*AdminClient, error) {
	em, err := atom.NewEntityManagerWithConnectionString(connectionString, internal.Version)

	if err != nil {
		return nil, err
	}

	return &AdminClient{em: em}, nil
}

// NewAdminClient creates an AdminClient authenticating using a TokenCredential.
func NewAdminClient(fullyQualifiedNamespace string, tokenCredential azcore.TokenCredential, options *AdminClientOptions) (*AdminClient, error) {
	em, err := atom.NewEntityManager(fullyQualifiedNamespace, tokenCredential, internal.Version)

	if err != nil {
		return nil, err
	}

	return &AdminClient{em: em}, nil
}

func (ac *AdminClient) GetNamespaceProperties() {

}

func ifMatchMiddleware(creating bool) []atom.MiddlewareFunc {
	if creating {
		return nil
	}

	// an update requires the entity to already exist.
	return []atom.MiddlewareFunc{
		func(next atom.RestHandler) atom.RestHandler {
			return func(ctx context.Context, req *http.Request) (*http.Response, error) {
				req.Header.Set("If-Match", "*")
				return next(ctx, req)
			}
		}}
}
