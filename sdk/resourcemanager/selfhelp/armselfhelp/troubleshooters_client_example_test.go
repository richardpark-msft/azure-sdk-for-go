//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Create.json
func ExampleTroubleshootersClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTroubleshootersClient().Create(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", armselfhelp.TroubleshooterResource{
		Properties: &armselfhelp.TroubleshooterInstanceProperties{
			Parameters: map[string]*string{
				"ResourceURI": to.Ptr("subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp"),
			},
			SolutionID: to.Ptr("SampleTroubleshooterSolutionId"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TroubleshooterResource = armselfhelp.TroubleshooterResource{
	// 	Name: to.Ptr("abf168ed-1b54-454a-86f6-e4b62253d3b1"),
	// 	Type: to.Ptr("Microsoft.Help/troubleshooters"),
	// 	ID: to.Ptr("/subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp/providers/Microsoft.Help/troubleshooter/abf168ed-1b54-454a-86f6-e4b62253d3b1"),
	// 	Properties: &armselfhelp.TroubleshooterInstanceProperties{
	// 		Parameters: map[string]*string{
	// 			"ResourceURI": to.Ptr("subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp"),
	// 		},
	// 		ProvisioningState: to.Ptr(armselfhelp.TroubleshooterProvisioningStateSucceeded),
	// 		SolutionID: to.Ptr("SampleTroubleshooterSolutionId"),
	// 		Steps: []*armselfhelp.Step{
	// 			{
	// 				Type: to.Ptr(armselfhelp.TypeDecision),
	// 				Description: to.Ptr("step description"),
	// 				ExecutionStatus: to.Ptr(armselfhelp.ExecutionStatusSuccess),
	// 				ExecutionStatusDescription: to.Ptr("Step is success"),
	// 				Guidance: to.Ptr("IP address is used to check if the VM would be reachable from the given IP Address. We have prepopulated the IP address from your browser. If the field is left empty, the generic internet address space will be used."),
	// 				ID: to.Ptr("SampleStepId"),
	// 				Inputs: []*armselfhelp.StepInput{
	// 					{
	// 						QuestionContent: to.Ptr("Is VPN Connected?"),
	// 						QuestionContentType: to.Ptr(armselfhelp.QuestionContentTypeText),
	// 						QuestionID: to.Ptr("6cbe3f23-f3c2-41f3-886e-0cfbeee19b24"),
	// 						QuestionType: to.Ptr(armselfhelp.QuestionType("MultiChoice")),
	// 						ResponseOptions: []*armselfhelp.ResponseOption{
	// 							{
	// 								Key: to.Ptr("Yes"),
	// 								Value: to.Ptr("1"),
	// 							},
	// 							{
	// 								Key: to.Ptr("No"),
	// 								Value: to.Ptr("0"),
	// 						}},
	// 						ResponseValidationProperties: &armselfhelp.ResponseValidationProperties{
	// 							IsRequired: to.Ptr(true),
	// 							MaxLength: to.Ptr[int64](1),
	// 						},
	// 				}},
	// 				IsLastStep: to.Ptr(true),
	// 				Title: to.Ptr("Step title"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Get.json
func ExampleTroubleshootersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTroubleshootersClient().Get(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TroubleshooterResource = armselfhelp.TroubleshooterResource{
	// 	Name: to.Ptr("abf168ed-1b54-454a-86f6-e4b62253d3b1"),
	// 	Type: to.Ptr("Microsoft.Help/troubleshooters"),
	// 	ID: to.Ptr("/subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp/providers/Microsoft.Help/troubleshooter/abf168ed-1b54-454a-86f6-e4b62253d3b1"),
	// 	Properties: &armselfhelp.TroubleshooterInstanceProperties{
	// 		ProvisioningState: to.Ptr(armselfhelp.TroubleshooterProvisioningStateSucceeded),
	// 		SolutionID: to.Ptr("SampleTroubleshooterSolutionId"),
	// 		Steps: []*armselfhelp.Step{
	// 			{
	// 				Type: to.Ptr(armselfhelp.TypeDecision),
	// 				Description: to.Ptr("step description"),
	// 				ExecutionStatus: to.Ptr(armselfhelp.ExecutionStatusSuccess),
	// 				ExecutionStatusDescription: to.Ptr("Step is success"),
	// 				Guidance: to.Ptr("IP address is used to check if the VM would be reachable from the given IP Address. We have prepopulated the IP address from your browser. If the field is left empty, the generic internet address space will be used."),
	// 				ID: to.Ptr("SampleId"),
	// 				Inputs: []*armselfhelp.StepInput{
	// 					{
	// 						QuestionContent: to.Ptr("Is VPN Connected?"),
	// 						QuestionContentType: to.Ptr(armselfhelp.QuestionContentTypeText),
	// 						QuestionID: to.Ptr("SampleQuestionId"),
	// 						QuestionType: to.Ptr(armselfhelp.QuestionType("MultiChoice")),
	// 						ResponseOptions: []*armselfhelp.ResponseOption{
	// 							{
	// 								Key: to.Ptr("Yes"),
	// 								Value: to.Ptr("1"),
	// 							},
	// 							{
	// 								Key: to.Ptr("No"),
	// 								Value: to.Ptr("0"),
	// 						}},
	// 						ResponseValidationProperties: &armselfhelp.ResponseValidationProperties{
	// 							IsRequired: to.Ptr(true),
	// 							MaxLength: to.Ptr[int64](1),
	// 						},
	// 				}},
	// 				IsLastStep: to.Ptr(true),
	// 				Title: to.Ptr("Step title"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Continue.json
func ExampleTroubleshootersClient_Continue() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTroubleshootersClient().Continue(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", &armselfhelp.TroubleshootersClientContinueOptions{ContinueRequestBody: &armselfhelp.ContinueRequestBody{
		StepID: to.Ptr("SampleStepId"),
		Responses: []*armselfhelp.TroubleshooterResponse{
			{
				QuestionID:   to.Ptr("SampleQuestionId"),
				QuestionType: to.Ptr(armselfhelp.QuestionType("Text")),
				Response:     to.Ptr("Connection exception"),
			}},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_End.json
func ExampleTroubleshootersClient_End() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTroubleshootersClient().End(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c3cc9abe085093ba880ee3eeb792edb4fa789553/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Restart.json
func ExampleTroubleshootersClient_Restart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTroubleshootersClient().Restart(ctx, "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp", "abf168ed-1b54-454a-86f6-e4b62253d3b1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestartTroubleshooterResponse = armselfhelp.RestartTroubleshooterResponse{
	// 	TroubleshooterResourceName: to.Ptr("SampleTroubleshooterResourceName"),
	// }
}
