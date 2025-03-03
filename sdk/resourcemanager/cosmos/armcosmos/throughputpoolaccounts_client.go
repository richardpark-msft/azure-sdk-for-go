//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// ThroughputPoolAccountsClient contains the methods for the ThroughputPoolAccounts group.
// Don't use this type directly, use NewThroughputPoolAccountsClient() instead.
type ThroughputPoolAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewThroughputPoolAccountsClient creates a new instance of ThroughputPoolAccountsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewThroughputPoolAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThroughputPoolAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ThroughputPoolAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all the Azure Cosmos DB accounts available under the subscription.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - throughputPoolName - Cosmos DB Throughput Pool name.
//   - options - ThroughputPoolAccountsClientListOptions contains the optional parameters for the ThroughputPoolAccountsClient.NewListPager
//     method.
func (client *ThroughputPoolAccountsClient) NewListPager(resourceGroupName string, throughputPoolName string, options *ThroughputPoolAccountsClientListOptions) *runtime.Pager[ThroughputPoolAccountsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThroughputPoolAccountsClientListResponse]{
		More: func(page ThroughputPoolAccountsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThroughputPoolAccountsClientListResponse) (ThroughputPoolAccountsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ThroughputPoolAccountsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, throughputPoolName, options)
			}, nil)
			if err != nil {
				return ThroughputPoolAccountsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ThroughputPoolAccountsClient) listCreateRequest(ctx context.Context, resourceGroupName string, throughputPoolName string, options *ThroughputPoolAccountsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/throughputPools/{throughputPoolName}/throughputPoolAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if throughputPoolName == "" {
		return nil, errors.New("parameter throughputPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{throughputPoolName}", url.PathEscape(throughputPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ThroughputPoolAccountsClient) listHandleResponse(resp *http.Response) (ThroughputPoolAccountsClientListResponse, error) {
	result := ThroughputPoolAccountsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputPoolAccountsListResult); err != nil {
		return ThroughputPoolAccountsClientListResponse{}, err
	}
	return result, nil
}
