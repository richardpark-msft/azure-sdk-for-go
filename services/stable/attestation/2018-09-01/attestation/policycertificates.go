package attestation

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

// PolicyCertificatesClient is the describes the interface for the per-tenant enclave service.
type PolicyCertificatesClient struct {
	BaseClient
}

// NewPolicyCertificatesClient creates an instance of the PolicyCertificatesClient client.
func NewPolicyCertificatesClient() PolicyCertificatesClient {
	return PolicyCertificatesClient{New()}
}

// Add sends the add request.
// Parameters:
// tenantBaseURL - the tenant name, for example https://mytenant.attest.azure.net.
// policyCertificateToAdd - an RFC7519 JSON Web Token containing a claim named "maa-policyCertificate" whose
// value is an RFC7517 JSON Web Key which specifies a new key to add. The RFC7519 JWT must be signed with one
// of the existing signing certificates
func (client PolicyCertificatesClient) Add(ctx context.Context, tenantBaseURL string, policyCertificateToAdd string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyCertificatesClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx, tenantBaseURL, policyCertificateToAdd)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client PolicyCertificatesClient) AddPreparer(ctx context.Context, tenantBaseURL string, policyCertificateToAdd string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"tenantBaseUrl": tenantBaseURL,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{tenantBaseUrl}", urlParameters),
		autorest.WithPath("/operations/policy/certificates"),
		autorest.WithJSON(policyCertificateToAdd),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyCertificatesClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client PolicyCertificatesClient) AddResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// tenantBaseURL - the tenant name, for example https://mytenant.attest.azure.net.
func (client PolicyCertificatesClient) Get(ctx context.Context, tenantBaseURL string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyCertificatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, tenantBaseURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PolicyCertificatesClient) GetPreparer(ctx context.Context, tenantBaseURL string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"tenantBaseUrl": tenantBaseURL,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{tenantBaseUrl}", urlParameters),
		autorest.WithPath("/operations/policy/certificates"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyCertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PolicyCertificatesClient) GetResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove sends the remove request.
// Parameters:
// tenantBaseURL - the tenant name, for example https://mytenant.attest.azure.net.
// policyCertificateToRemove - an RFC7519 JSON Web Token containing a claim named "maa-policyCertificate" whose
// value is an RFC7517 JSON Web Key which specifies a new key to update. The RFC7519 JWT must be signed with
// one of the existing signing certificates
func (client PolicyCertificatesClient) Remove(ctx context.Context, tenantBaseURL string, policyCertificateToRemove string) (result SetObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolicyCertificatesClient.Remove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemovePreparer(ctx, tenantBaseURL, policyCertificateToRemove)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "attestation.PolicyCertificatesClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client PolicyCertificatesClient) RemovePreparer(ctx context.Context, tenantBaseURL string, policyCertificateToRemove string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"tenantBaseUrl": tenantBaseURL,
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{tenantBaseUrl}", urlParameters),
		autorest.WithPath("/operations/policy/certificates"),
		autorest.WithJSON(policyCertificateToRemove),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyCertificatesClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client PolicyCertificatesClient) RemoveResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusBadRequest, http.StatusUnauthorized),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
