//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armguestconfiguration

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

// HCRPAssignmentsClient contains the methods for the GuestConfigurationHCRPAssignments group.
// Don't use this type directly, use NewHCRPAssignmentsClient() instead.
type HCRPAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHCRPAssignmentsClient creates a new instance of HCRPAssignmentsClient with the specified values.
//   - subscriptionID - Subscription ID which uniquely identify Microsoft Azure subscription. The subscription ID forms part of
//     the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHCRPAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HCRPAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName+".HCRPAssignmentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HCRPAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates an association between a ARC machine and guest configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - guestConfigurationAssignmentName - Name of the guest configuration assignment.
//   - resourceGroupName - The resource group name.
//   - machineName - The name of the ARC machine.
//   - parameters - Parameters supplied to the create or update guest configuration assignment.
//   - options - HCRPAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the HCRPAssignmentsClient.CreateOrUpdate
//     method.
func (client *HCRPAssignmentsClient) CreateOrUpdate(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, machineName string, parameters Assignment, options *HCRPAssignmentsClientCreateOrUpdateOptions) (HCRPAssignmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, guestConfigurationAssignmentName, resourceGroupName, machineName, parameters, options)
	if err != nil {
		return HCRPAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HCRPAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HCRPAssignmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HCRPAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, machineName string, parameters Assignment, options *HCRPAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HCRPAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (HCRPAssignmentsClientCreateOrUpdateResponse, error) {
	result := HCRPAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return HCRPAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a guest configuration assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - guestConfigurationAssignmentName - Name of the guest configuration assignment
//   - machineName - The name of the ARC machine.
//   - options - HCRPAssignmentsClientDeleteOptions contains the optional parameters for the HCRPAssignmentsClient.Delete method.
func (client *HCRPAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *HCRPAssignmentsClientDeleteOptions) (HCRPAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, machineName, options)
	if err != nil {
		return HCRPAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HCRPAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HCRPAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HCRPAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HCRPAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *HCRPAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get information about a guest configuration assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - guestConfigurationAssignmentName - The guest configuration assignment name.
//   - machineName - The name of the ARC machine.
//   - options - HCRPAssignmentsClientGetOptions contains the optional parameters for the HCRPAssignmentsClient.Get method.
func (client *HCRPAssignmentsClient) Get(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *HCRPAssignmentsClientGetOptions) (HCRPAssignmentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, guestConfigurationAssignmentName, machineName, options)
	if err != nil {
		return HCRPAssignmentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HCRPAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HCRPAssignmentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HCRPAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *HCRPAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if guestConfigurationAssignmentName == "" {
		return nil, errors.New("parameter guestConfigurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{guestConfigurationAssignmentName}", url.PathEscape(guestConfigurationAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HCRPAssignmentsClient) getHandleResponse(resp *http.Response) (HCRPAssignmentsClientGetResponse, error) {
	result := HCRPAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return HCRPAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all guest configuration assignments for an ARC machine.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - machineName - The name of the ARC machine.
//   - options - HCRPAssignmentsClientListOptions contains the optional parameters for the HCRPAssignmentsClient.NewListPager
//     method.
func (client *HCRPAssignmentsClient) NewListPager(resourceGroupName string, machineName string, options *HCRPAssignmentsClientListOptions) *runtime.Pager[HCRPAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HCRPAssignmentsClientListResponse]{
		More: func(page HCRPAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *HCRPAssignmentsClientListResponse) (HCRPAssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, machineName, options)
			if err != nil {
				return HCRPAssignmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HCRPAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HCRPAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HCRPAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *HCRPAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HCRPAssignmentsClient) listHandleResponse(resp *http.Response) (HCRPAssignmentsClientListResponse, error) {
	result := HCRPAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return HCRPAssignmentsClientListResponse{}, err
	}
	return result, nil
}