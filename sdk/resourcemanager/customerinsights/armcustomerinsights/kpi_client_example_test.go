//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiCreateOrUpdate.json
func ExampleKpiClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKpiClient().BeginCreateOrUpdate(ctx, "TestHubRG", "sdkTestHub", "kpiTest45453647", armcustomerinsights.KpiResourceFormat{
		Properties: &armcustomerinsights.KpiDefinition{
			Description: map[string]*string{
				"en-us": to.Ptr("Kpi Description"),
			},
			Aliases: []*armcustomerinsights.KpiAlias{
				{
					AliasName:  to.Ptr("alias"),
					Expression: to.Ptr("Id+4"),
				}},
			CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
			DisplayName: map[string]*string{
				"en-us": to.Ptr("Kpi DisplayName"),
			},
			EntityType:     to.Ptr(armcustomerinsights.EntityTypesProfile),
			EntityTypeName: to.Ptr("testProfile2327128"),
			Expression:     to.Ptr("SavingAccountBalance"),
			Function:       to.Ptr(armcustomerinsights.KpiFunctionsSum),
			GroupBy: []*string{
				to.Ptr("SavingAccountBalance")},
			ThresHolds: &armcustomerinsights.KpiThresholds{
				IncreasingKpi: to.Ptr(true),
				LowerLimit:    to.Ptr[float64](5),
				UpperLimit:    to.Ptr[float64](50),
			},
			Unit: to.Ptr("unit"),
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
	// res.KpiResourceFormat = armcustomerinsights.KpiResourceFormat{
	// 	Name: to.Ptr("sdkTestHub/kpiTest45453647"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/kpi"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/kpi/kpiTest45453647"),
	// 	Properties: &armcustomerinsights.KpiDefinition{
	// 		Description: map[string]*string{
	// 			"en-us": to.Ptr("Kpi Description"),
	// 		},
	// 		Aliases: []*armcustomerinsights.KpiAlias{
	// 			{
	// 				AliasName: to.Ptr("alias"),
	// 				Expression: to.Ptr("Id+4"),
	// 		}},
	// 		CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
	// 		DisplayName: map[string]*string{
	// 			"en-us": to.Ptr("Kpi DisplayName"),
	// 		},
	// 		EntityType: to.Ptr(armcustomerinsights.EntityTypesProfile),
	// 		EntityTypeName: to.Ptr("testProfile2327128"),
	// 		Expression: to.Ptr("SavingAccountBalance"),
	// 		Function: to.Ptr(armcustomerinsights.KpiFunctionsSum),
	// 		GroupBy: []*string{
	// 			to.Ptr("SavingAccountBalance")},
	// 			KpiName: to.Ptr("kpiTest45453647"),
	// 			ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 			TenantID: to.Ptr("sdktesthub"),
	// 			ThresHolds: &armcustomerinsights.KpiThresholds{
	// 				IncreasingKpi: to.Ptr(true),
	// 				LowerLimit: to.Ptr[float64](5),
	// 				UpperLimit: to.Ptr[float64](50),
	// 			},
	// 			Unit: to.Ptr("unit"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiGet.json
func ExampleKpiClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKpiClient().Get(ctx, "TestHubRG", "sdkTestHub", "kpiTest45453647", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KpiResourceFormat = armcustomerinsights.KpiResourceFormat{
	// 	Name: to.Ptr("sdkTestHub/kpiTest45453647"),
	// 	Type: to.Ptr("Microsoft.CustomerInsights/hubs/kpi"),
	// 	ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/kpi/kpiTest45453647"),
	// 	Properties: &armcustomerinsights.KpiDefinition{
	// 		Description: map[string]*string{
	// 			"en-us": to.Ptr("Kpi Description"),
	// 		},
	// 		Aliases: []*armcustomerinsights.KpiAlias{
	// 			{
	// 				AliasName: to.Ptr("alias"),
	// 				Expression: to.Ptr("Id+4"),
	// 		}},
	// 		CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
	// 		DisplayName: map[string]*string{
	// 			"en-us": to.Ptr("Kpi DisplayName"),
	// 		},
	// 		EntityType: to.Ptr(armcustomerinsights.EntityTypesProfile),
	// 		EntityTypeName: to.Ptr("testProfile2327128"),
	// 		Expression: to.Ptr("SavingAccountBalance"),
	// 		Function: to.Ptr(armcustomerinsights.KpiFunctionsSum),
	// 		GroupBy: []*string{
	// 			to.Ptr("SavingAccountBalance")},
	// 			KpiName: to.Ptr("kpiTest45453647"),
	// 			ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
	// 			TenantID: to.Ptr("sdktesthub"),
	// 			ThresHolds: &armcustomerinsights.KpiThresholds{
	// 				IncreasingKpi: to.Ptr(true),
	// 				LowerLimit: to.Ptr[float64](5),
	// 				UpperLimit: to.Ptr[float64](50),
	// 			},
	// 			Unit: to.Ptr("unit"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiDelete.json
func ExampleKpiClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKpiClient().BeginDelete(ctx, "TestHubRG", "sdkTestHub", "kpiTest45453647", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiReprocess.json
func ExampleKpiClient_Reprocess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewKpiClient().Reprocess(ctx, "TestHubRG", "sdkTestHub", "kpiTest45453647", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiListByHub.json
func ExampleKpiClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKpiClient().NewListByHubPager("TestHubRG", "sdkTestHub", nil)
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
		// page.KpiListResult = armcustomerinsights.KpiListResult{
		// 	Value: []*armcustomerinsights.KpiResourceFormat{
		// 		{
		// 			Name: to.Ptr("sdkTestHub/kpiTest45453647"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/kpi"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/kpi/kpiTest45453647"),
		// 			Properties: &armcustomerinsights.KpiDefinition{
		// 				Description: map[string]*string{
		// 					"en-us": to.Ptr("Kpi Description"),
		// 				},
		// 				Aliases: []*armcustomerinsights.KpiAlias{
		// 					{
		// 						AliasName: to.Ptr("alias"),
		// 						Expression: to.Ptr("Id+4"),
		// 				}},
		// 				CalculationWindow: to.Ptr(armcustomerinsights.CalculationWindowTypesDay),
		// 				DisplayName: map[string]*string{
		// 					"en-us": to.Ptr("Kpi DisplayName"),
		// 				},
		// 				EntityType: to.Ptr(armcustomerinsights.EntityTypesProfile),
		// 				EntityTypeName: to.Ptr("testProfile2327128"),
		// 				Expression: to.Ptr("SavingAccountBalance"),
		// 				Function: to.Ptr(armcustomerinsights.KpiFunctionsSum),
		// 				GroupBy: []*string{
		// 					to.Ptr("SavingAccountBalance")},
		// 					KpiName: to.Ptr("kpiTest45453647"),
		// 					ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 					TenantID: to.Ptr("sdktesthub"),
		// 					ThresHolds: &armcustomerinsights.KpiThresholds{
		// 						IncreasingKpi: to.Ptr(true),
		// 						LowerLimit: to.Ptr[float64](5),
		// 						UpperLimit: to.Ptr[float64](50),
		// 					},
		// 					Unit: to.Ptr("unit"),
		// 				},
		// 		}},
		// 	}
	}
}