//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnginx

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

// NewAPIKeysClient creates a new instance of APIKeysClient.
func (c *ClientFactory) NewAPIKeysClient() *APIKeysClient {
	return &APIKeysClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewCertificatesClient creates a new instance of CertificatesClient.
func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	return &CertificatesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient.
func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	return &ConfigurationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeploymentsClient creates a new instance of DeploymentsClient.
func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	return &DeploymentsClient{
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
