//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/extension/ExtensionMetadata_Get.json
func ExampleExtensionMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtensionMetadataClient().Get(ctx, "EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent", "1.9.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExtensionValue = armhybridcompute.ExtensionValue{
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.1"),
	// 	Properties: &armhybridcompute.ExtensionValueProperties{
	// 		ExtensionType: to.Ptr("azuremonitorlinuxagent"),
	// 		Publisher: to.Ptr("microsoft.azure.monitor"),
	// 		Version: to.Ptr("1.9.1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-03-31-preview/examples/extension/ExtensionMetadata_List.json
func ExampleExtensionMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExtensionMetadataClient().NewListPager("EastUS", "microsoft.azure.monitor", "azuremonitorlinuxagent", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExtensionValueListResult = armhybridcompute.ExtensionValueListResult{
		// 	Value: []*armhybridcompute.ExtensionValue{
		// 		{
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.1"),
		// 			Properties: &armhybridcompute.ExtensionValueProperties{
		// 				ExtensionType: to.Ptr("azuremonitorlinuxagent"),
		// 				Publisher: to.Ptr("microsoft.azure.monitor"),
		// 				Version: to.Ptr("1.9.1"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/Providers/Microsoft.HybridCompute/locations/eastus/publishers/microsoft.azure.monitor/extensionTypes/azuremonitorlinuxagent/versions/1.9.2"),
		// 			Properties: &armhybridcompute.ExtensionValueProperties{
		// 				ExtensionType: to.Ptr("azuremonitorlinuxagent"),
		// 				Publisher: to.Ptr("microsoft.azure.monitor"),
		// 				Version: to.Ptr("1.9.2"),
		// 			},
		// 	}},
		// }
	}
}
