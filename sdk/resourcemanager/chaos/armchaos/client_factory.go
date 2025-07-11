// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armchaos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewCapabilitiesClient creates a new instance of CapabilitiesClient.
func (c *ClientFactory) NewCapabilitiesClient() *CapabilitiesClient {
	return &CapabilitiesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCapabilityTypesClient creates a new instance of CapabilityTypesClient.
func (c *ClientFactory) NewCapabilityTypesClient() *CapabilityTypesClient {
	return &CapabilityTypesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExperimentsClient creates a new instance of ExperimentsClient.
func (c *ClientFactory) NewExperimentsClient() *ExperimentsClient {
	return &ExperimentsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusesClient creates a new instance of OperationStatusesClient.
func (c *ClientFactory) NewOperationStatusesClient() *OperationStatusesClient {
	return &OperationStatusesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewTargetTypesClient creates a new instance of TargetTypesClient.
func (c *ClientFactory) NewTargetTypesClient() *TargetTypesClient {
	return &TargetTypesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewTargetsClient creates a new instance of TargetsClient.
func (c *ClientFactory) NewTargetsClient() *TargetsClient {
	return &TargetsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
