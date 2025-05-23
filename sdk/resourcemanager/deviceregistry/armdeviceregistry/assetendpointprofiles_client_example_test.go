// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdeviceregistry_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceregistry/armdeviceregistry"
	"log"
)

// Generated from example definition: 2024-11-01/Create_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginCreateOrReplace_createAssetEndpointProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginCreateOrReplace(ctx, "myResourceGroup", "my-assetendpointprofile", armdeviceregistry.AssetEndpointProfile{
		Location: to.Ptr("West Europe"),
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Type: to.Ptr("CustomLocation"),
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		},
		Tags: map[string]*string{
			"site": to.Ptr("building-1"),
		},
		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
			TargetAddress:       to.Ptr("https://www.example.com/myTargetAddress"),
			EndpointProfileType: to.Ptr("myEndpointProfileType"),
			Authentication: &armdeviceregistry.Authentication{
				Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
			},
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
	// res = armdeviceregistry.AssetEndpointProfilesClientCreateOrReplaceResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/Create_AssetEndpointProfile_With_DiscoveredAepRef.json
func ExampleAssetEndpointProfilesClient_BeginCreateOrReplace_createAssetEndpointProfileWithDiscoveredAepRef() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginCreateOrReplace(ctx, "myResourceGroup", "my-assetendpointprofile", armdeviceregistry.AssetEndpointProfile{
		Location: to.Ptr("West Europe"),
		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
			Type: to.Ptr("CustomLocation"),
			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		},
		Tags: map[string]*string{
			"site": to.Ptr("building-1"),
		},
		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
			TargetAddress:                     to.Ptr("https://www.example.com/myTargetAddress"),
			EndpointProfileType:               to.Ptr("myEndpointProfileType"),
			DiscoveredAssetEndpointProfileRef: to.Ptr("discoveredAssetEndpointProfile1"),
			Authentication: &armdeviceregistry.Authentication{
				Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
			},
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
	// res = armdeviceregistry.AssetEndpointProfilesClientCreateOrReplaceResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			DiscoveredAssetEndpointProfileRef: to.Ptr("discoveredAssetEndpointProfile1"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/Delete_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginDelete(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: 2024-11-01/Get_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_Get_getAssetEndpointProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetEndpointProfilesClient().Get(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.AssetEndpointProfilesClientGetResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodUsernamePassword),
	// 				UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
	// 					UsernameSecretName: to.Ptr("usernameRef"),
	// 					PasswordSecretName: to.Ptr("passwordRef"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/Get_AssetEndpointProfile_With_SyncStatus.json
func ExampleAssetEndpointProfilesClient_Get_getAssetEndpointProfileWithSyncStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetEndpointProfilesClient().Get(ctx, "myResourceGroup", "my-assetendpointprofile", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeviceregistry.AssetEndpointProfilesClientGetResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodUsernamePassword),
	// 				UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
	// 					UsernameSecretName: to.Ptr("usernameRef"),
	// 					PasswordSecretName: to.Ptr("passwordRef"),
	// 				},
	// 			},
	// 			Status: &armdeviceregistry.AssetEndpointProfileStatus{
	// 				Errors: []*armdeviceregistry.AssetEndpointProfileStatusError{
	// 					{
	// 						Code: to.Ptr[int32](500),
	// 						Message: to.Ptr("Internal Server Error"),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

// Generated from example definition: 2024-11-01/List_AssetEndpointProfiles_ResourceGroup.json
func ExampleAssetEndpointProfilesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssetEndpointProfilesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armdeviceregistry.AssetEndpointProfilesClientListByResourceGroupResponse{
		// 	AssetEndpointProfileListResult: armdeviceregistry.AssetEndpointProfileListResult{
		// 		Value: []*armdeviceregistry.AssetEndpointProfile{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
		// 				Name: to.Ptr("my-assetendpointprofile"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-1"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile1"),
		// 				Name: to.Ptr("my-assetendpointprofile1"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-2"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("7824a74f-21e1-4458-ae06-604d3a241d2c"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress1"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
		// 						X509Credentials: &armdeviceregistry.X509Credentials{
		// 							CertificateSecretName: to.Ptr("certificatRef"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile2"),
		// 				Name: to.Ptr("my-assetendpointprofile2"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-2"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress2"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodUsernamePassword),
		// 						UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
		// 							UsernameSecretName: to.Ptr("usernameRef"),
		// 							PasswordSecretName: to.Ptr("passwordRef"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-11-01/List_AssetEndpointProfiles_Subscription.json
func ExampleAssetEndpointProfilesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssetEndpointProfilesClient().NewListBySubscriptionPager(nil)
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
		// page = armdeviceregistry.AssetEndpointProfilesClientListBySubscriptionResponse{
		// 	AssetEndpointProfileListResult: armdeviceregistry.AssetEndpointProfileListResult{
		// 		Value: []*armdeviceregistry.AssetEndpointProfile{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
		// 				Name: to.Ptr("my-assetendpointprofile"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-1"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile1"),
		// 				Name: to.Ptr("my-assetendpointprofile1"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-2"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("7824a74f-21e1-4458-ae06-604d3a241d2c"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress1"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodCertificate),
		// 						X509Credentials: &armdeviceregistry.X509Credentials{
		// 							CertificateSecretName: to.Ptr("certificatRef"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile2"),
		// 				Name: to.Ptr("my-assetendpointprofile2"),
		// 				Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
		// 				Location: to.Ptr("West Europe"),
		// 				ExtendedLocation: &armdeviceregistry.ExtendedLocation{
		// 					Type: to.Ptr("CustomLocation"),
		// 					Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"site": to.Ptr("building-2"),
		// 				},
		// 				SystemData: &armdeviceregistry.SystemData{
		// 					CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
		// 					LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
		// 				},
		// 				Properties: &armdeviceregistry.AssetEndpointProfileProperties{
		// 					UUID: to.Ptr("1824a74f-21e1-4458-ae07-604d3a241d2e"),
		// 					TargetAddress: to.Ptr("https://www.example.com/myTargetAddress2"),
		// 					EndpointProfileType: to.Ptr("myEndpointProfileType"),
		// 					Authentication: &armdeviceregistry.Authentication{
		// 						Method: to.Ptr(armdeviceregistry.AuthenticationMethodUsernamePassword),
		// 						UsernamePasswordCredentials: &armdeviceregistry.UsernamePasswordCredentials{
		// 							UsernameSecretName: to.Ptr("usernameRef"),
		// 							PasswordSecretName: to.Ptr("passwordRef"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

// Generated from example definition: 2024-11-01/Update_AssetEndpointProfile.json
func ExampleAssetEndpointProfilesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssetEndpointProfilesClient().BeginUpdate(ctx, "myResourceGroup", "my-assetendpointprofile", armdeviceregistry.AssetEndpointProfileUpdate{
		Properties: &armdeviceregistry.AssetEndpointProfileUpdateProperties{
			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
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
	// res = armdeviceregistry.AssetEndpointProfilesClientUpdateResponse{
	// 	AssetEndpointProfile: &armdeviceregistry.AssetEndpointProfile{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.DeviceRegistry/assetEndpointProfiles/my-assetendpointprofile"),
	// 		Name: to.Ptr("my-assetendpointprofile"),
	// 		Type: to.Ptr("Microsoft.DeviceRegistry/assetEndpointProfiles"),
	// 		Location: to.Ptr("West Europe"),
	// 		ExtendedLocation: &armdeviceregistry.ExtendedLocation{
	// 			Type: to.Ptr("CustomLocation"),
	// 			Name: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"site": to.Ptr("building-1"),
	// 		},
	// 		SystemData: &armdeviceregistry.SystemData{
	// 			CreatedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			CreatedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T00:36:43.2516899Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ta23112-4596-44ff-b773-19405922bfc1"),
	// 			LastModifiedByType: to.Ptr(armdeviceregistry.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-16T01:37:16.0922793Z"); return t}()),
	// 		},
	// 		Properties: &armdeviceregistry.AssetEndpointProfileProperties{
	// 			UUID: to.Ptr("0796f7c1-f2c8-44d7-9f5b-9a6f9522a85d"),
	// 			TargetAddress: to.Ptr("https://www.example.com/myTargetAddress"),
	// 			EndpointProfileType: to.Ptr("myEndpointProfileType"),
	// 			Authentication: &armdeviceregistry.Authentication{
	// 				Method: to.Ptr(armdeviceregistry.AuthenticationMethodAnonymous),
	// 			},
	// 			ProvisioningState: to.Ptr(armdeviceregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
