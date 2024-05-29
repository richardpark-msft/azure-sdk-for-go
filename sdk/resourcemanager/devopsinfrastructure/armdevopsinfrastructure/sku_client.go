//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevopsinfrastructure

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SKUClient contains the methods for the SKU group.
// Don't use this type directly, use NewSKUClient() instead.
type SKUClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSKUClient creates a new instance of SKUClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSKUClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SKUClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByLocationPager - List ResourceSku resources by subscription ID
//
// Generated from API version 2024-04-04-preview
//   - locationName - Name of the location.
//   - options - SKUClientListByLocationOptions contains the optional parameters for the SKUClient.NewListByLocationPager method.
func (client *SKUClient) NewListByLocationPager(locationName string, options *SKUClientListByLocationOptions) *runtime.Pager[SKUClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[SKUClientListByLocationResponse]{
		More: func(page SKUClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SKUClientListByLocationResponse) (SKUClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SKUClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return SKUClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *SKUClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *SKUClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevOpsInfrastructure/locations/{locationName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *SKUClient) listByLocationHandleResponse(resp *http.Response) (SKUClientListByLocationResponse, error) {
	result := SKUClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUListResult); err != nil {
		return SKUClientListByLocationResponse{}, err
	}
	return result, nil
}
