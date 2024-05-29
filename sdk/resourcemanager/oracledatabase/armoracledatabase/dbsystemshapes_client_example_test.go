//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbSystemShapes_listByLocation.json
func ExampleDbSystemShapesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbSystemShapesClient().NewListByLocationPager("eastus", nil)
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
		// page.DbSystemShapeListResult = armoracledatabase.DbSystemShapeListResult{
		// 	Value: []*armoracledatabase.DbSystemShape{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/locations/dbSystemShapes"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dbSystemShapes/EXADATA.X9M"),
		// 			Properties: &armoracledatabase.DbSystemShapeProperties{
		// 				AvailableCoreCount: to.Ptr[int32](100),
		// 				AvailableCoreCountPerNode: to.Ptr[int32](1000),
		// 				AvailableDataStorageInTbs: to.Ptr[int32](10),
		// 				AvailableDataStoragePerServerInTbs: to.Ptr[float64](100),
		// 				AvailableDbNodePerNodeInGbs: to.Ptr[int32](10),
		// 				AvailableDbNodeStorageInGbs: to.Ptr[int32](10),
		// 				AvailableMemoryInGbs: to.Ptr[int32](10),
		// 				AvailableMemoryPerNodeInGbs: to.Ptr[int32](10),
		// 				CoreCountIncrement: to.Ptr[int32](1),
		// 				MaxStorageCount: to.Ptr[int32](100),
		// 				MaximumNodeCount: to.Ptr[int32](1000),
		// 				MinCoreCountPerNode: to.Ptr[int32](0),
		// 				MinDataStorageInTbs: to.Ptr[int32](0),
		// 				MinDbNodeStoragePerNodeInGbs: to.Ptr[int32](0),
		// 				MinMemoryPerNodeInGbs: to.Ptr[int32](0),
		// 				MinStorageCount: to.Ptr[int32](0),
		// 				MinimumCoreCount: to.Ptr[int32](1),
		// 				MinimumNodeCount: to.Ptr[int32](0),
		// 				RuntimeMinimumCoreCount: to.Ptr[int32](1),
		// 				ShapeFamily: to.Ptr("EXADATA"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbSystemShapes_get.json
func ExampleDbSystemShapesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDbSystemShapesClient().Get(ctx, "eastus", "EXADATA.X9M", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DbSystemShape = armoracledatabase.DbSystemShape{
	// 	Type: to.Ptr("Oracle.Database/locations/dbSystemShapes"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dbSystemShapes/EXADATA.X9M"),
	// 	Properties: &armoracledatabase.DbSystemShapeProperties{
	// 		AvailableCoreCount: to.Ptr[int32](100),
	// 		AvailableCoreCountPerNode: to.Ptr[int32](1000),
	// 		AvailableDataStorageInTbs: to.Ptr[int32](10),
	// 		AvailableDataStoragePerServerInTbs: to.Ptr[float64](100),
	// 		AvailableDbNodePerNodeInGbs: to.Ptr[int32](10),
	// 		AvailableDbNodeStorageInGbs: to.Ptr[int32](10),
	// 		AvailableMemoryInGbs: to.Ptr[int32](10),
	// 		AvailableMemoryPerNodeInGbs: to.Ptr[int32](10),
	// 		CoreCountIncrement: to.Ptr[int32](1),
	// 		MaxStorageCount: to.Ptr[int32](100),
	// 		MaximumNodeCount: to.Ptr[int32](1000),
	// 		MinCoreCountPerNode: to.Ptr[int32](0),
	// 		MinDataStorageInTbs: to.Ptr[int32](0),
	// 		MinDbNodeStoragePerNodeInGbs: to.Ptr[int32](0),
	// 		MinMemoryPerNodeInGbs: to.Ptr[int32](0),
	// 		MinStorageCount: to.Ptr[int32](0),
	// 		MinimumCoreCount: to.Ptr[int32](1),
	// 		MinimumNodeCount: to.Ptr[int32](0),
	// 		RuntimeMinimumCoreCount: to.Ptr[int32](1),
	// 		ShapeFamily: to.Ptr("EXADATA"),
	// 	},
	// }
}
