//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_Get.json
func ExampleNamespaceTopicEventSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().Get(ctx, "examplerg", "examplenamespace2", "examplenamespacetopic2", "examplenamespacetopicEventSub1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subscription = armeventgrid.Subscription{
	// 	Name: to.Ptr("examplenamespacetopicEventSub2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub2"),
	// 	Properties: &armeventgrid.SubscriptionProperties{
	// 		DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
	// 			DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
	// 			Queue: &armeventgrid.QueueInfo{
	// 				EventTimeToLive: to.Ptr("P1D"),
	// 				MaxDeliveryCount: to.Ptr[int32](4),
	// 				ReceiveLockDurationInSeconds: to.Ptr[int32](60),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
	// 		ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_CreateOrUpdate.json
func ExampleNamespaceTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "examplerg", "examplenamespace2", "examplenamespacetopic2", "examplenamespacetopicEventSub2", armeventgrid.Subscription{
		Properties: &armeventgrid.SubscriptionProperties{
			DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
				DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
				Queue: &armeventgrid.QueueInfo{
					EventTimeToLive:              to.Ptr("P1D"),
					MaxDeliveryCount:             to.Ptr[int32](4),
					ReceiveLockDurationInSeconds: to.Ptr[int32](60),
				},
			},
			EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
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
	// res.Subscription = armeventgrid.Subscription{
	// 	Name: to.Ptr("examplenamespacetopicEventSub2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub2"),
	// 	Properties: &armeventgrid.SubscriptionProperties{
	// 		DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
	// 			DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
	// 			Queue: &armeventgrid.QueueInfo{
	// 				EventTimeToLive: to.Ptr("P1D"),
	// 				MaxDeliveryCount: to.Ptr[int32](4),
	// 				ReceiveLockDurationInSeconds: to.Ptr[int32](60),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
	// 		ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_Delete.json
func ExampleNamespaceTopicEventSubscriptionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().BeginDelete(ctx, "examplerg", "examplenamespace2", "examplenamespacetopic2", "examplenamespacetopicEventSub2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_Update.json
func ExampleNamespaceTopicEventSubscriptionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().BeginUpdate(ctx, "examplerg", "exampleNamespaceName1", "exampleNamespaceTopicName1", "exampleNamespaceTopicEventSubscriptionName1", armeventgrid.SubscriptionUpdateParameters{
		Properties: &armeventgrid.SubscriptionUpdateParametersProperties{
			DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
				DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
				Queue: &armeventgrid.QueueInfo{
					EventTimeToLive:              to.Ptr("P1D"),
					MaxDeliveryCount:             to.Ptr[int32](3),
					ReceiveLockDurationInSeconds: to.Ptr[int32](60),
				},
			},
			EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
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
	// res.Subscription = armeventgrid.Subscription{
	// 	Name: to.Ptr("exampleNamespaceTopicEventSubscriptionName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/exampleNamespaceTopicName1/eventSubscriptions/exampleNamespaceTopicEventSubscriptionName1"),
	// 	Properties: &armeventgrid.SubscriptionProperties{
	// 		DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
	// 			DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
	// 			Queue: &armeventgrid.QueueInfo{
	// 				EventTimeToLive: to.Ptr("P1D"),
	// 				MaxDeliveryCount: to.Ptr[int32](3),
	// 				ReceiveLockDurationInSeconds: to.Ptr[int32](60),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
	// 		ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_ListByNamespaceTopic.json
func ExampleNamespaceTopicEventSubscriptionsClient_NewListByNamespaceTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespaceTopicEventSubscriptionsClient().NewListByNamespaceTopicPager("examplerg", "examplenamespace2", "examplenamespacetopic2", &armeventgrid.NamespaceTopicEventSubscriptionsClientListByNamespaceTopicOptions{Filter: nil,
		Top: nil,
	})
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
		// page.SubscriptionsListResult = armeventgrid.SubscriptionsListResult{
		// 	Value: []*armeventgrid.Subscription{
		// 		{
		// 			Name: to.Ptr("examplenamespacetopicEventSub1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub1"),
		// 			Properties: &armeventgrid.SubscriptionProperties{
		// 				DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
		// 					DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
		// 					Queue: &armeventgrid.QueueInfo{
		// 						EventTimeToLive: to.Ptr("P1D"),
		// 						MaxDeliveryCount: to.Ptr[int32](5),
		// 						ReceiveLockDurationInSeconds: to.Ptr[int32](50),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
		// 				ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplenamespacetopicEventSub2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub2"),
		// 			Properties: &armeventgrid.SubscriptionProperties{
		// 				DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
		// 					DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
		// 					Queue: &armeventgrid.QueueInfo{
		// 						EventTimeToLive: to.Ptr("P1D"),
		// 						MaxDeliveryCount: to.Ptr[int32](4),
		// 						ReceiveLockDurationInSeconds: to.Ptr[int32](60),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
		// 				ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/NamespaceTopicEventSubscriptions_GetDeliveryAttributes.json
func ExampleNamespaceTopicEventSubscriptionsClient_GetDeliveryAttributes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().GetDeliveryAttributes(ctx, "examplerg", "exampleNamespace", "exampleNamespaceTopic", "exampleEventSubscriptionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeliveryAttributeListResult = armeventgrid.DeliveryAttributeListResult{
	// 	Value: []armeventgrid.DeliveryAttributeMappingClassification{
	// 		&armeventgrid.StaticDeliveryAttributeMapping{
	// 			Name: to.Ptr("header1"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeStatic),
	// 			Properties: &armeventgrid.StaticDeliveryAttributeMappingProperties{
	// 				IsSecret: to.Ptr(false),
	// 				Value: to.Ptr("NormalValue"),
	// 			},
	// 		},
	// 		&armeventgrid.DynamicDeliveryAttributeMapping{
	// 			Name: to.Ptr("header2"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeDynamic),
	// 			Properties: &armeventgrid.DynamicDeliveryAttributeMappingProperties{
	// 				SourceField: to.Ptr("data.foo"),
	// 			},
	// 		},
	// 		&armeventgrid.StaticDeliveryAttributeMapping{
	// 			Name: to.Ptr("header3"),
	// 			Type: to.Ptr(armeventgrid.DeliveryAttributeMappingTypeStatic),
	// 			Properties: &armeventgrid.StaticDeliveryAttributeMappingProperties{
	// 				IsSecret: to.Ptr(true),
	// 				Value: to.Ptr("mySecretValue"),
	// 			},
	// 	}},
	// }
}