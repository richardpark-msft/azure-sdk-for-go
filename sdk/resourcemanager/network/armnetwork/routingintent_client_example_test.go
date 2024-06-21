//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/RoutingIntentPut.json
func ExampleRoutingIntentClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutingIntentClient().BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "Intent1", armnetwork.RoutingIntent{
		Properties: &armnetwork.RoutingIntentProperties{
			RoutingPolicies: []*armnetwork.RoutingPolicy{
				{
					Name: to.Ptr("InternetTraffic"),
					Destinations: []*string{
						to.Ptr("Internet")},
					NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
				},
				{
					Name: to.Ptr("PrivateTrafficPolicy"),
					Destinations: []*string{
						to.Ptr("PrivateTraffic")},
					NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoutingIntent = armnetwork.RoutingIntent{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routingIntent/Intent1"),
	// 	Name: to.Ptr("Intent1"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs/routingIntent"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RoutingIntentProperties{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingPolicies: []*armnetwork.RoutingPolicy{
	// 			{
	// 				Name: to.Ptr("InternetTraffic"),
	// 				Destinations: []*string{
	// 					to.Ptr("Internet")},
	// 					NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("PrivateTrafficPolicy"),
	// 					Destinations: []*string{
	// 						to.Ptr("PrivateTraffic")},
	// 						NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				}},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/RoutingIntentGet.json
func ExampleRoutingIntentClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoutingIntentClient().Get(ctx, "rg1", "virtualHub1", "Intent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RoutingIntent = armnetwork.RoutingIntent{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routingIntent/Intent1"),
	// 	Name: to.Ptr("Intent1"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs/routingIntent"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RoutingIntentProperties{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingPolicies: []*armnetwork.RoutingPolicy{
	// 			{
	// 				Name: to.Ptr("InternetTraffic"),
	// 				Destinations: []*string{
	// 					to.Ptr("Internet")},
	// 					NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				},
	// 				{
	// 					Name: to.Ptr("PrivateTrafficPolicy"),
	// 					Destinations: []*string{
	// 						to.Ptr("PrivateTraffic")},
	// 						NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
	// 				}},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/RoutingIntentDelete.json
func ExampleRoutingIntentClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRoutingIntentClient().BeginDelete(ctx, "rg1", "virtualHub1", "Intent1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/RoutingIntentList.json
func ExampleRoutingIntentClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoutingIntentClient().NewListPager("rg1", "virtualHub1", nil)
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
		// page.ListRoutingIntentResult = armnetwork.ListRoutingIntentResult{
		// 	Value: []*armnetwork.RoutingIntent{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routingIntent/Intent1"),
		// 			Name: to.Ptr("Intent1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualHubs/routingIntent"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.RoutingIntentProperties{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RoutingPolicies: []*armnetwork.RoutingPolicy{
		// 					{
		// 						Name: to.Ptr("InternetTraffic"),
		// 						Destinations: []*string{
		// 							to.Ptr("Internet")},
		// 							NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
		// 						},
		// 						{
		// 							Name: to.Ptr("PrivateTrafficPolicy"),
		// 							Destinations: []*string{
		// 								to.Ptr("PrivateTraffic")},
		// 								NextHop: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azfw1"),
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
