//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedDatabaseAdvancedThreatProtectionSettingsListByDatabase.json
func ExampleManagedDatabaseAdvancedThreatProtectionSettingsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedDatabaseAdvancedThreatProtectionSettingsClient().NewListByDatabasePager("threatprotection-6852", "threatprotection-2080", "testdb", nil)
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
		// page.ManagedDatabaseAdvancedThreatProtectionListResult = armsql.ManagedDatabaseAdvancedThreatProtectionListResult{
		// 	Value: []*armsql.ManagedDatabaseAdvancedThreatProtection{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/databases/advancedThreatProtectionSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-6852/providers/Microsoft.Sql/managedInstances/threatprotection-2080/databases/testdb/advancedThreatProtectionSettings"),
		// 			Properties: &armsql.AdvancedThreatProtectionProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
		// 			},
		// 			SystemData: &armsql.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedDatabaseAdvancedThreatProtectionSettingsGet.json
func ExampleManagedDatabaseAdvancedThreatProtectionSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseAdvancedThreatProtectionSettingsClient().Get(ctx, "threatprotection-6852", "threatprotection-2080", "testdb", armsql.AdvancedThreatProtectionNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedDatabaseAdvancedThreatProtection = armsql.ManagedDatabaseAdvancedThreatProtection{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/advancedThreatProtectionSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.Sql/managedInstances/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 	Properties: &armsql.AdvancedThreatProtectionProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedDatabaseAdvancedThreatProtectionSettingsCreateMax.json
func ExampleManagedDatabaseAdvancedThreatProtectionSettingsClient_CreateOrUpdate_updateAManagedDatabasesAdvancedThreatProtectionSettingsWithAllParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseAdvancedThreatProtectionSettingsClient().CreateOrUpdate(ctx, "threatprotection-4799", "threatprotection-6440", "testdb", armsql.AdvancedThreatProtectionNameDefault, armsql.ManagedDatabaseAdvancedThreatProtection{
		Properties: &armsql.AdvancedThreatProtectionProperties{
			State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedDatabaseAdvancedThreatProtection = armsql.ManagedDatabaseAdvancedThreatProtection{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/advancedThreatProtectionSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.Sql/managedInstances/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 	Properties: &armsql.AdvancedThreatProtectionProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		State: to.Ptr(armsql.AdvancedThreatProtectionStateEnabled),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedDatabaseAdvancedThreatProtectionSettingsCreateMin.json
func ExampleManagedDatabaseAdvancedThreatProtectionSettingsClient_CreateOrUpdate_updateAManagedDatabasesAdvancedThreatProtectionSettingsWithMinimalParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedDatabaseAdvancedThreatProtectionSettingsClient().CreateOrUpdate(ctx, "threatprotection-4799", "threatprotection-6440", "testdb", armsql.AdvancedThreatProtectionNameDefault, armsql.ManagedDatabaseAdvancedThreatProtection{
		Properties: &armsql.AdvancedThreatProtectionProperties{
			State: to.Ptr(armsql.AdvancedThreatProtectionStateDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedDatabaseAdvancedThreatProtection = armsql.ManagedDatabaseAdvancedThreatProtection{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/databases/advancedThreatProtectionSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/threatprotection-4799/providers/Microsoft.Sql/managedInstances/threatprotection-6440/advancedThreatProtectionSettings/Default"),
	// 	Properties: &armsql.AdvancedThreatProtectionProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		State: to.Ptr(armsql.AdvancedThreatProtectionStateDisabled),
	// 	},
	// 	SystemData: &armsql.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-03T04:41:33.937Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armsql.CreatedByTypeUser),
	// 	},
	// }
}
