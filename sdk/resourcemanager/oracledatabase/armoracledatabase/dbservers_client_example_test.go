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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbServers_listByParent.json
func ExampleDbServersClient_NewListByCloudExadataInfrastructurePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDbServersClient().NewListByCloudExadataInfrastructurePager("rg000", "infra1", nil)
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
		// page.DbServerListResult = armoracledatabase.DbServerListResult{
		// 	Value: []*armoracledatabase.DbServer{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/cloudVmClusters/dbServers"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1/dbServers/ocid1"),
		// 			Properties: &armoracledatabase.DbServerProperties{
		// 				AutonomousVirtualMachineIDs: []*string{
		// 					to.Ptr("ocid1..aaaaa")},
		// 					AutonomousVMClusterIDs: []*string{
		// 						to.Ptr("ocid1..aaaaa")},
		// 						CompartmentID: to.Ptr("ocid1....aaaa"),
		// 						CPUCoreCount: to.Ptr[int32](100),
		// 						DbNodeIDs: []*string{
		// 							to.Ptr("ocid1..aaaaa")},
		// 							DbNodeStorageSizeInGbs: to.Ptr[int32](150),
		// 							DisplayName: to.Ptr("dbserver1"),
		// 							ExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
		// 							LifecycleState: to.Ptr(armoracledatabase.DbServerProvisioningStateAvailable),
		// 							MaxCPUCount: to.Ptr[int32](1000),
		// 							MaxMemoryInGbs: to.Ptr[int32](1000),
		// 							Ocid: to.Ptr("ocid1"),
		// 							VMClusterIDs: []*string{
		// 								to.Ptr("ocid1..aaaaa")},
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/520e274d7d95fc6d1002dd3c1fcaf8d55d27f63e/specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbServers_get.json
func ExampleDbServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDbServersClient().Get(ctx, "rg000", "infra1", "ocid1....aaaaaa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DbServer = armoracledatabase.DbServer{
	// 	Type: to.Ptr("Oracle.Database/cloudVmClusters/dbServers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1/dbServers/ocid1"),
	// 	Properties: &armoracledatabase.DbServerProperties{
	// 		AutonomousVirtualMachineIDs: []*string{
	// 			to.Ptr("ocid1..aaaaa")},
	// 			AutonomousVMClusterIDs: []*string{
	// 				to.Ptr("ocid1..aaaaa")},
	// 				CompartmentID: to.Ptr("ocid1....aaaa"),
	// 				CPUCoreCount: to.Ptr[int32](100),
	// 				DbNodeIDs: []*string{
	// 					to.Ptr("ocid1..aaaaa")},
	// 					DbNodeStorageSizeInGbs: to.Ptr[int32](150),
	// 					DisplayName: to.Ptr("dbserver1"),
	// 					ExadataInfrastructureID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 					LifecycleState: to.Ptr(armoracledatabase.DbServerProvisioningStateAvailable),
	// 					MaxCPUCount: to.Ptr[int32](1000),
	// 					MaxMemoryInGbs: to.Ptr[int32](1000),
	// 					Ocid: to.Ptr("ocid1"),
	// 					VMClusterIDs: []*string{
	// 						to.Ptr("ocid1..aaaaa")},
	// 					},
	// 				}
}
