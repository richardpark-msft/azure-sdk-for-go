//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatadog

// MarketplaceAgreementsClientCreateOrUpdateResponse contains the response from method MarketplaceAgreementsClient.CreateOrUpdate.
type MarketplaceAgreementsClientCreateOrUpdateResponse struct {
	AgreementResource
}

// MarketplaceAgreementsClientListResponse contains the response from method MarketplaceAgreementsClient.NewListPager.
type MarketplaceAgreementsClientListResponse struct {
	AgreementResourceListResponse
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.BeginCreate.
type MonitorsClientCreateResponse struct {
	MonitorResource
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.BeginDelete.
type MonitorsClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitorsClientGetDefaultKeyResponse contains the response from method MonitorsClient.GetDefaultKey.
type MonitorsClientGetDefaultKeyResponse struct {
	APIKey
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorResource
}

// MonitorsClientListAPIKeysResponse contains the response from method MonitorsClient.NewListAPIKeysPager.
type MonitorsClientListAPIKeysResponse struct {
	APIKeyListResponse
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.NewListByResourceGroupPager.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientListHostsResponse contains the response from method MonitorsClient.NewListHostsPager.
type MonitorsClientListHostsResponse struct {
	HostListResponse
}

// MonitorsClientListLinkedResourcesResponse contains the response from method MonitorsClient.NewListLinkedResourcesPager.
type MonitorsClientListLinkedResourcesResponse struct {
	LinkedResourceListResponse
}

// MonitorsClientListMonitoredResourcesResponse contains the response from method MonitorsClient.NewListMonitoredResourcesPager.
type MonitorsClientListMonitoredResourcesResponse struct {
	MonitoredResourceListResponse
}

// MonitorsClientListResponse contains the response from method MonitorsClient.NewListPager.
type MonitorsClientListResponse struct {
	MonitorResourceListResponse
}

// MonitorsClientRefreshSetPasswordLinkResponse contains the response from method MonitorsClient.RefreshSetPasswordLink.
type MonitorsClientRefreshSetPasswordLinkResponse struct {
	SetPasswordLink
}

// MonitorsClientSetDefaultKeyResponse contains the response from method MonitorsClient.SetDefaultKey.
type MonitorsClientSetDefaultKeyResponse struct {
	// placeholder for future response values
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.BeginUpdate.
type MonitorsClientUpdateResponse struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// SingleSignOnConfigurationsClientCreateOrUpdateResponse contains the response from method SingleSignOnConfigurationsClient.BeginCreateOrUpdate.
type SingleSignOnConfigurationsClientCreateOrUpdateResponse struct {
	SingleSignOnResource
}

// SingleSignOnConfigurationsClientGetResponse contains the response from method SingleSignOnConfigurationsClient.Get.
type SingleSignOnConfigurationsClientGetResponse struct {
	SingleSignOnResource
}

// SingleSignOnConfigurationsClientListResponse contains the response from method SingleSignOnConfigurationsClient.NewListPager.
type SingleSignOnConfigurationsClientListResponse struct {
	SingleSignOnResourceListResponse
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	MonitoringTagRules
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.NewListPager.
type TagRulesClientListResponse struct {
	MonitoringTagRulesListResponse
}