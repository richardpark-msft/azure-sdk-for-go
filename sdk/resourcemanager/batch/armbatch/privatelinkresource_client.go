//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PrivateLinkResourceClient contains the methods for the PrivateLinkResource group.
// Don't use this type directly, use NewPrivateLinkResourceClient() instead.
type PrivateLinkResourceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourceClient creates a new instance of PrivateLinkResourceClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets information about the specified private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - privateLinkResourceName - The private link resource name. This must be unique within the account.
//   - options - PrivateLinkResourceClientGetOptions contains the optional parameters for the PrivateLinkResourceClient.Get method.
func (client *PrivateLinkResourceClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (PrivateLinkResourceClientGetResponse, error) {
	var err error
	const operationName = "PrivateLinkResourceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourceClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string, options *PrivateLinkResourceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources/{privateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourceClient) getHandleResponse(resp *http.Response) (PrivateLinkResourceClientGetResponse, error) {
	result := PrivateLinkResourceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourceClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBatchAccountPager - Lists all of the private link resources in the specified account.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - options - PrivateLinkResourceClientListByBatchAccountOptions contains the optional parameters for the PrivateLinkResourceClient.NewListByBatchAccountPager
//     method.
func (client *PrivateLinkResourceClient) NewListByBatchAccountPager(resourceGroupName string, accountName string, options *PrivateLinkResourceClientListByBatchAccountOptions) *runtime.Pager[PrivateLinkResourceClientListByBatchAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourceClientListByBatchAccountResponse]{
		More: func(page PrivateLinkResourceClientListByBatchAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourceClientListByBatchAccountResponse) (PrivateLinkResourceClientListByBatchAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkResourceClient.NewListByBatchAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBatchAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return PrivateLinkResourceClientListByBatchAccountResponse{}, err
			}
			return client.listByBatchAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBatchAccountCreateRequest creates the ListByBatchAccount request.
func (client *PrivateLinkResourceClient) listByBatchAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateLinkResourceClientListByBatchAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBatchAccountHandleResponse handles the ListByBatchAccount response.
func (client *PrivateLinkResourceClient) listByBatchAccountHandleResponse(resp *http.Response) (PrivateLinkResourceClientListByBatchAccountResponse, error) {
	result := PrivateLinkResourceClientListByBatchAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListPrivateLinkResourcesResult); err != nil {
		return PrivateLinkResourceClientListByBatchAccountResponse{}, err
	}
	return result, nil
}
