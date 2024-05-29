//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceAssessmentList_example.json
func ExampleRegulatoryComplianceAssessmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegulatoryComplianceAssessmentsClient().NewListPager("PCI-DSS-3.2", "1.1", &armsecurity.RegulatoryComplianceAssessmentsClientListOptions{Filter: nil})
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
		// page.RegulatoryComplianceAssessmentList = armsecurity.RegulatoryComplianceAssessmentList{
		// 	Value: []*armsecurity.RegulatoryComplianceAssessment{
		// 		{
		// 			Name: to.Ptr("968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceAssessment"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2/regulatoryComplianceControls/1.1/regulatoryComplianceAssessments/968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
		// 			Properties: &armsecurity.RegulatoryComplianceAssessmentProperties{
		// 				Description: to.Ptr("Troubleshoot missing scan data on your machines"),
		// 				AssessmentDetailsLink: to.Ptr("https://management.azure.com/subscriptions/a27e854a-8578-4395-8eaf-6fc7849f3050/providers/Microsoft.Security/securityStatuses/968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
		// 				AssessmentType: to.Ptr("Assessment"),
		// 				FailedResources: to.Ptr[int32](4),
		// 				PassedResources: to.Ptr[int32](7),
		// 				SkippedResources: to.Ptr[int32](0),
		// 				State: to.Ptr(armsecurity.StateFailed),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("3bcd234d-c9c7-c2a2-89e0-c01f419c1a8a"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceAssessment"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2/regulatoryComplianceControls/2/regulatoryComplianceAssessments/3bcd234d-c9c7-c2a2-89e0-c01f419c1a8a"),
		// 			Properties: &armsecurity.RegulatoryComplianceAssessmentProperties{
		// 				Description: to.Ptr("Resolve endpoint protection health issues on your machines"),
		// 				AssessmentDetailsLink: to.Ptr("https://management.azure.com/subscriptions/a27e854a-8578-4395-8eaf-6fc7849f3050/providers/Microsoft.Security/securityStatuses/3bcd234d-c9c7-c2a2-89e0-c01f419c1a8a"),
		// 				AssessmentType: to.Ptr("Assessment"),
		// 				FailedResources: to.Ptr[int32](0),
		// 				PassedResources: to.Ptr[int32](0),
		// 				SkippedResources: to.Ptr[int32](10),
		// 				State: to.Ptr(armsecurity.StateSkipped),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("d1db3318-01ff-16de-29eb-28b344515626"),
		// 			Type: to.Ptr("Microsoft.Security/regulatoryComplianceAssessment"),
		// 			ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2/regulatoryComplianceControls/2.1/regulatoryComplianceAssessments/d1db3318-01ff-16de-29eb-28b344515626"),
		// 			Properties: &armsecurity.RegulatoryComplianceAssessmentProperties{
		// 				Description: to.Ptr("Install monitoring agent on your machines"),
		// 				AssessmentDetailsLink: to.Ptr("https://management.azure.com/subscriptions/a27e854a-8578-4395-8eaf-6fc7849f3050/providers/Microsoft.Security/securityStatuses/d1db3318-01ff-16de-29eb-28b344515626"),
		// 				AssessmentType: to.Ptr("Assessment"),
		// 				FailedResources: to.Ptr[int32](0),
		// 				PassedResources: to.Ptr[int32](8),
		// 				SkippedResources: to.Ptr[int32](0),
		// 				State: to.Ptr(armsecurity.StatePassed),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ac34f238dd6b9071f486b57e9f9f1a0c43ec6f6/specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceAssessment_example.json
func ExampleRegulatoryComplianceAssessmentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegulatoryComplianceAssessmentsClient().Get(ctx, "PCI-DSS-3.2", "1.1", "968548cb-02b3-8cd2-11f8-0cf64ab1a347", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegulatoryComplianceAssessment = armsecurity.RegulatoryComplianceAssessment{
	// 	Name: to.Ptr("968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
	// 	Type: to.Ptr("Microsoft.Security/regulatoryComplianceAssessment"),
	// 	ID: to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/regulatoryComplianceStandards/PCI-DSS-3.2/regulatoryComplianceControls/1.1/regulatoryComplianceAssessments/968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
	// 	Properties: &armsecurity.RegulatoryComplianceAssessmentProperties{
	// 		Description: to.Ptr("Troubleshoot missing scan data on your machines"),
	// 		AssessmentDetailsLink: to.Ptr("https://management.azure.com/subscriptions/a27e854a-8578-4395-8eaf-6fc7849f3050/providers/Microsoft.Security/securityStatuses/968548cb-02b3-8cd2-11f8-0cf64ab1a347"),
	// 		AssessmentType: to.Ptr("Assessment"),
	// 		FailedResources: to.Ptr[int32](4),
	// 		PassedResources: to.Ptr[int32](7),
	// 		SkippedResources: to.Ptr[int32](0),
	// 		State: to.Ptr(armsecurity.StateFailed),
	// 	},
	// }
}
