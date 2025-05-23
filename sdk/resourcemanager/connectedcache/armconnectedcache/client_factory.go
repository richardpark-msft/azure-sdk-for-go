// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armconnectedcache

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

// NewCacheNodesOperationsClient creates a new instance of CacheNodesOperationsClient.
func (c *ClientFactory) NewCacheNodesOperationsClient() *CacheNodesOperationsClient {
	return &CacheNodesOperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEnterpriseCustomerOperationsClient creates a new instance of EnterpriseCustomerOperationsClient.
func (c *ClientFactory) NewEnterpriseCustomerOperationsClient() *EnterpriseCustomerOperationsClient {
	return &EnterpriseCustomerOperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEnterpriseMccCacheNodesOperationsClient creates a new instance of EnterpriseMccCacheNodesOperationsClient.
func (c *ClientFactory) NewEnterpriseMccCacheNodesOperationsClient() *EnterpriseMccCacheNodesOperationsClient {
	return &EnterpriseMccCacheNodesOperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewEnterpriseMccCustomersClient creates a new instance of EnterpriseMccCustomersClient.
func (c *ClientFactory) NewEnterpriseMccCustomersClient() *EnterpriseMccCustomersClient {
	return &EnterpriseMccCustomersClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIspCacheNodesOperationsClient creates a new instance of IspCacheNodesOperationsClient.
func (c *ClientFactory) NewIspCacheNodesOperationsClient() *IspCacheNodesOperationsClient {
	return &IspCacheNodesOperationsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewIspCustomersClient creates a new instance of IspCustomersClient.
func (c *ClientFactory) NewIspCustomersClient() *IspCustomersClient {
	return &IspCustomersClient{
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
