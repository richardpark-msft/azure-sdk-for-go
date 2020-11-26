package migrate

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VMwareOperationsStatusClient is the discover your workloads for Azure.
type VMwareOperationsStatusClient struct {
	BaseClient
}

// NewVMwareOperationsStatusClient creates an instance of the VMwareOperationsStatusClient client.
func NewVMwareOperationsStatusClient() VMwareOperationsStatusClient {
	return NewVMwareOperationsStatusClientWithBaseURI(DefaultBaseURI)
}

// NewVMwareOperationsStatusClientWithBaseURI creates an instance of the VMwareOperationsStatusClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVMwareOperationsStatusClientWithBaseURI(baseURI string) VMwareOperationsStatusClient {
	return VMwareOperationsStatusClient{NewWithBaseURI(baseURI)}
}

// GetOperationStatus sends the get operation status request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// operationStatusName - operation status ARM name.
// APIVersion - the API version to use for this operation.
func (client VMwareOperationsStatusClient) GetOperationStatus(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, operationStatusName string, APIVersion string) (result OperationStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMwareOperationsStatusClient.GetOperationStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOperationStatusPreparer(ctx, subscriptionID, resourceGroupName, siteName, operationStatusName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VMwareOperationsStatusClient", "GetOperationStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOperationStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.VMwareOperationsStatusClient", "GetOperationStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOperationStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.VMwareOperationsStatusClient", "GetOperationStatus", resp, "Failure responding to request")
	}

	return
}

// GetOperationStatusPreparer prepares the GetOperationStatus request.
func (client VMwareOperationsStatusClient) GetOperationStatusPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, operationStatusName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationStatusName": autorest.Encode("path", operationStatusName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"siteName":            autorest.Encode("path", siteName),
		"subscriptionId":      autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/operationsStatus/{operationStatusName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOperationStatusSender sends the GetOperationStatus request. The method will close the
// http.Response Body if it receives an error.
func (client VMwareOperationsStatusClient) GetOperationStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetOperationStatusResponder handles the response to the GetOperationStatus request. The method always
// closes the http.Response Body.
func (client VMwareOperationsStatusClient) GetOperationStatusResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
