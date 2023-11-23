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

// CertificateClient contains the methods for the Certificate group.
// Don't use this type directly, use NewCertificateClient() instead.
type CertificateClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCertificateClient creates a new instance of CertificateClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCertificateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CertificateClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CancelDeletion - If you try to delete a certificate that is being used by a pool or compute node, the status of the certificate
// changes to deleteFailed. If you decide that you want to continue using the certificate,
// you can use this operation to set the status of the certificate back to active. If you intend to delete the certificate,
// you do not need to run this operation after the deletion failed. You must make
// sure that the certificate is not being used by any resources, and then you can try again to delete the certificate.
// Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault Extension
// [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - certificateName - The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash,
//     and must match the certificate data in the request. For example SHA1-a3d1c5.
//   - options - CertificateClientCancelDeletionOptions contains the optional parameters for the CertificateClient.CancelDeletion
//     method.
func (client *CertificateClient) CancelDeletion(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientCancelDeletionOptions) (CertificateClientCancelDeletionResponse, error) {
	var err error
	const operationName = "CertificateClient.CancelDeletion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelDeletionCreateRequest(ctx, resourceGroupName, accountName, certificateName, options)
	if err != nil {
		return CertificateClientCancelDeletionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateClientCancelDeletionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateClientCancelDeletionResponse{}, err
	}
	resp, err := client.cancelDeletionHandleResponse(httpResp)
	return resp, err
}

// cancelDeletionCreateRequest creates the CancelDeletion request.
func (client *CertificateClient) cancelDeletionCreateRequest(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientCancelDeletionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}/cancelDelete"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelDeletionHandleResponse handles the CancelDeletion response.
func (client *CertificateClient) cancelDeletionHandleResponse(resp *http.Response) (CertificateClientCancelDeletionResponse, error) {
	result := CertificateClientCancelDeletionResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientCancelDeletionResponse{}, err
	}
	return result, nil
}

// Create - Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault
// Extension [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - certificateName - The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash,
//     and must match the certificate data in the request. For example SHA1-a3d1c5.
//   - parameters - Additional parameters for certificate creation.
//   - options - CertificateClientCreateOptions contains the optional parameters for the CertificateClient.Create method.
func (client *CertificateClient) Create(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientCreateOptions) (CertificateClientCreateResponse, error) {
	var err error
	const operationName = "CertificateClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, certificateName, parameters, options)
	if err != nil {
		return CertificateClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *CertificateClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *CertificateClient) createHandleResponse(resp *http.Response) (CertificateClientCreateResponse, error) {
	result := CertificateClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault
// Extension [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - certificateName - The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash,
//     and must match the certificate data in the request. For example SHA1-a3d1c5.
//   - options - CertificateClientBeginDeleteOptions contains the optional parameters for the CertificateClient.BeginDelete method.
func (client *CertificateClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientBeginDeleteOptions) (*runtime.Poller[CertificateClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, certificateName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CertificateClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CertificateClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault
// Extension [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
func (client *CertificateClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CertificateClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, certificateName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault Extension
// [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - certificateName - The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash,
//     and must match the certificate data in the request. For example SHA1-a3d1c5.
//   - options - CertificateClientGetOptions contains the optional parameters for the CertificateClient.Get method.
func (client *CertificateClient) Get(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientGetOptions) (CertificateClientGetResponse, error) {
	var err error
	const operationName = "CertificateClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, certificateName, options)
	if err != nil {
		return CertificateClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CertificateClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, certificateName string, options *CertificateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *CertificateClient) getHandleResponse(resp *http.Response) (CertificateClientGetResponse, error) {
	result := CertificateClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBatchAccountPager - Warning: This operation is deprecated and will be removed after February, 2024. Please use
// the Azure KeyVault Extension [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - options - CertificateClientListByBatchAccountOptions contains the optional parameters for the CertificateClient.NewListByBatchAccountPager
//     method.
func (client *CertificateClient) NewListByBatchAccountPager(resourceGroupName string, accountName string, options *CertificateClientListByBatchAccountOptions) *runtime.Pager[CertificateClientListByBatchAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateClientListByBatchAccountResponse]{
		More: func(page CertificateClientListByBatchAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateClientListByBatchAccountResponse) (CertificateClientListByBatchAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CertificateClient.NewListByBatchAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBatchAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return CertificateClientListByBatchAccountResponse{}, err
			}
			return client.listByBatchAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBatchAccountCreateRequest creates the ListByBatchAccount request.
func (client *CertificateClient) listByBatchAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *CertificateClientListByBatchAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBatchAccountHandleResponse handles the ListByBatchAccount response.
func (client *CertificateClient) listByBatchAccountHandleResponse(resp *http.Response) (CertificateClientListByBatchAccountResponse, error) {
	result := CertificateClientListByBatchAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListCertificatesResult); err != nil {
		return CertificateClientListByBatchAccountResponse{}, err
	}
	return result, nil
}

// Update - Warning: This operation is deprecated and will be removed after February, 2024. Please use the Azure KeyVault
// Extension [https://learn.microsoft.com/azure/batch/batch-certificate-migration-guide]
// instead.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - certificateName - The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash,
//     and must match the certificate data in the request. For example SHA1-a3d1c5.
//   - parameters - Certificate entity to update.
//   - options - CertificateClientUpdateOptions contains the optional parameters for the CertificateClient.Update method.
func (client *CertificateClient) Update(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientUpdateOptions) (CertificateClientUpdateResponse, error) {
	var err error
	const operationName = "CertificateClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, certificateName, parameters, options)
	if err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CertificateClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CertificateClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CertificateClient) updateHandleResponse(resp *http.Response) (CertificateClientUpdateResponse, error) {
	result := CertificateClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	return result, nil
}
