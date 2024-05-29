//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// AgentPoolsClient contains the methods for the AgentPools group.
// Don't use this type directly, use NewAgentPoolsClient() instead.
type AgentPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAgentPoolsClient creates a new instance of AgentPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentPoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AgentPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAbortLatestOperation - Aborts the currently running operation on the agent pool. The Agent Pool will be moved to a
// Canceling state and eventually to a Canceled state when cancellation finishes. If the operation completes
// before cancellation can take place, an error is returned.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginAbortLatestOperationOptions contains the optional parameters for the AgentPoolsClient.BeginAbortLatestOperation
//     method.
func (client *AgentPoolsClient) BeginAbortLatestOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginAbortLatestOperationOptions) (*runtime.Poller[AgentPoolsClientAbortLatestOperationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.abortLatestOperation(ctx, resourceGroupName, resourceName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientAbortLatestOperationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AgentPoolsClientAbortLatestOperationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// AbortLatestOperation - Aborts the currently running operation on the agent pool. The Agent Pool will be moved to a Canceling
// state and eventually to a Canceled state when cancellation finishes. If the operation completes
// before cancellation can take place, an error is returned.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *AgentPoolsClient) abortLatestOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginAbortLatestOperationOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginAbortLatestOperation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.abortLatestOperationCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// abortLatestOperationCreateRequest creates the AbortLatestOperation request.
func (client *AgentPoolsClient) abortLatestOperationCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginAbortLatestOperationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedclusters/{resourceName}/agentPools/{agentPoolName}/abort"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginCreateOrUpdate - Creates or updates an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - parameters - The agent pool to create or update.
//   - options - AgentPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginCreateOrUpdate
//     method.
func (client *AgentPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AgentPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, agentPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AgentPoolsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *AgentPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AgentPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
func (client *AgentPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*runtime.Poller[AgentPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AgentPoolsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *AgentPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AgentPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	if options != nil && options.IgnorePodDisruptionBudget != nil {
		reqQP.Set("ignore-pod-disruption-budget", strconv.FormatBool(*options.IgnorePodDisruptionBudget))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteMachines - Deletes specific machines in an agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - machines - A list of machines from the agent pool to be deleted.
//   - options - AgentPoolsClientBeginDeleteMachinesOptions contains the optional parameters for the AgentPoolsClient.BeginDeleteMachines
//     method.
func (client *AgentPoolsClient) BeginDeleteMachines(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, machines AgentPoolDeleteMachinesParameter, options *AgentPoolsClientBeginDeleteMachinesOptions) (*runtime.Poller[AgentPoolsClientDeleteMachinesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteMachines(ctx, resourceGroupName, resourceName, agentPoolName, machines, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientDeleteMachinesResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AgentPoolsClientDeleteMachinesResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DeleteMachines - Deletes specific machines in an agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *AgentPoolsClient) deleteMachines(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, machines AgentPoolDeleteMachinesParameter, options *AgentPoolsClientBeginDeleteMachinesOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginDeleteMachines"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteMachinesCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, machines, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteMachinesCreateRequest creates the DeleteMachines request.
func (client *AgentPoolsClient) deleteMachinesCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, machines AgentPoolDeleteMachinesParameter, options *AgentPoolsClientBeginDeleteMachinesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/deleteMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, machines); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets the specified managed cluster agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
func (client *AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetOptions) (AgentPoolsClientGetResponse, error) {
	var err error
	const operationName = "AgentPoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AgentPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AgentPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentPoolsClient) getHandleResponse(resp *http.Response) (AgentPoolsClientGetResponse, error) {
	result := AgentPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPool); err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	return result, nil
}

// GetAvailableAgentPoolVersions - See supported Kubernetes versions [https://docs.microsoft.com/azure/aks/supported-kubernetes-versions]
// for more details about the version lifecycle.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - AgentPoolsClientGetAvailableAgentPoolVersionsOptions contains the optional parameters for the AgentPoolsClient.GetAvailableAgentPoolVersions
//     method.
func (client *AgentPoolsClient) GetAvailableAgentPoolVersions(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientGetAvailableAgentPoolVersionsOptions) (AgentPoolsClientGetAvailableAgentPoolVersionsResponse, error) {
	var err error
	const operationName = "AgentPoolsClient.GetAvailableAgentPoolVersions"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAvailableAgentPoolVersionsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	resp, err := client.getAvailableAgentPoolVersionsHandleResponse(httpResp)
	return resp, err
}

// getAvailableAgentPoolVersionsCreateRequest creates the GetAvailableAgentPoolVersions request.
func (client *AgentPoolsClient) getAvailableAgentPoolVersionsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientGetAvailableAgentPoolVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableAgentPoolVersions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAvailableAgentPoolVersionsHandleResponse handles the GetAvailableAgentPoolVersions response.
func (client *AgentPoolsClient) getAvailableAgentPoolVersionsHandleResponse(resp *http.Response) (AgentPoolsClientGetAvailableAgentPoolVersionsResponse, error) {
	result := AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolAvailableVersions); err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	return result, nil
}

// GetUpgradeProfile - Gets the upgrade profile for an agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetUpgradeProfileOptions contains the optional parameters for the AgentPoolsClient.GetUpgradeProfile
//     method.
func (client *AgentPoolsClient) GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetUpgradeProfileOptions) (AgentPoolsClientGetUpgradeProfileResponse, error) {
	var err error
	const operationName = "AgentPoolsClient.GetUpgradeProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getUpgradeProfileCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	resp, err := client.getUpgradeProfileHandleResponse(httpResp)
	return resp, err
}

// getUpgradeProfileCreateRequest creates the GetUpgradeProfile request.
func (client *AgentPoolsClient) getUpgradeProfileCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetUpgradeProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUpgradeProfileHandleResponse handles the GetUpgradeProfile response.
func (client *AgentPoolsClient) getUpgradeProfileHandleResponse(resp *http.Response) (AgentPoolsClientGetUpgradeProfileResponse, error) {
	result := AgentPoolsClientGetUpgradeProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolUpgradeProfile); err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of agent pools in the specified managed cluster.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - AgentPoolsClientListOptions contains the optional parameters for the AgentPoolsClient.NewListPager method.
func (client *AgentPoolsClient) NewListPager(resourceGroupName string, resourceName string, options *AgentPoolsClientListOptions) *runtime.Pager[AgentPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgentPoolsClientListResponse]{
		More: func(page AgentPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgentPoolsClientListResponse) (AgentPoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AgentPoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return AgentPoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AgentPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentPoolsClient) listHandleResponse(resp *http.Response) (AgentPoolsClientListResponse, error) {
	result := AgentPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolListResult); err != nil {
		return AgentPoolsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpgradeNodeImageVersion - Upgrading the node image version of an agent pool applies the newest OS and runtime updates
// to the nodes. AKS provides one new image per week with the latest updates. For more details on node image
// versions, see: https://docs.microsoft.com/azure/aks/node-image-upgrade
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginUpgradeNodeImageVersionOptions contains the optional parameters for the AgentPoolsClient.BeginUpgradeNodeImageVersion
//     method.
func (client *AgentPoolsClient) BeginUpgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*runtime.Poller[AgentPoolsClientUpgradeNodeImageVersionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgradeNodeImageVersion(ctx, resourceGroupName, resourceName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AgentPoolsClientUpgradeNodeImageVersionResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AgentPoolsClientUpgradeNodeImageVersionResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpgradeNodeImageVersion - Upgrading the node image version of an agent pool applies the newest OS and runtime updates to
// the nodes. AKS provides one new image per week with the latest updates. For more details on node image
// versions, see: https://docs.microsoft.com/azure/aks/node-image-upgrade
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-02-preview
func (client *AgentPoolsClient) upgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginUpgradeNodeImageVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.upgradeNodeImageVersionCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// upgradeNodeImageVersionCreateRequest creates the UpgradeNodeImageVersion request.
func (client *AgentPoolsClient) upgradeNodeImageVersionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeNodeImageVersion"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
