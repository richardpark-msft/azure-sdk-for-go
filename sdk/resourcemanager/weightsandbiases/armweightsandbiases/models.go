// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armweightsandbiases

import "time"

// InstanceProperties - Properties specific to Instance
type InstanceProperties struct {
	// REQUIRED; Marketplace details of the resource.
	Marketplace *MarketplaceDetails

	// REQUIRED; partner properties
	PartnerProperties *PartnerProperties

	// REQUIRED; Details of the user.
	User *UserDetails

	// Single sign-on properties
	SingleSignOnProperties *SingleSignOnPropertiesV2

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ResourceProvisioningState
}

// InstanceResource - Concrete tracked resource types can be created by aliasing this type using a specific property type.
type InstanceResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// READ-ONLY; Name of the Instance resource
	Name *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *InstanceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// InstanceResourceListResult - The response of a InstanceResource list operation.
type InstanceResourceListResult struct {
	// REQUIRED; The InstanceResource items on this page
	Value []*InstanceResource

	// The link to the next page of items
	NextLink *string
}

// InstanceResourceUpdate - The type used for update operations of the Instance Resource.
type InstanceResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// Resource tags.
	Tags map[string]*string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// MarketplaceDetails - Marketplace details for an organization
type MarketplaceDetails struct {
	// REQUIRED; Offer details for the marketplace that is selected by the user
	OfferDetails *OfferDetails

	// Azure subscription id for the the marketplace offer is purchased from
	SubscriptionID *string

	// READ-ONLY; Marketplace subscription status
	SubscriptionStatus *MarketplaceSubscriptionStatus
}

// OfferDetails - Offer details for the marketplace that is selected by the user
type OfferDetails struct {
	// REQUIRED; Offer Id for the marketplace offer
	OfferID *string

	// REQUIRED; Plan Id for the marketplace offer
	PlanID *string

	// REQUIRED; Publisher Id for the marketplace offer
	PublisherID *string

	// Plan Name for the marketplace offer
	PlanName *string

	// Plan Display Name for the marketplace offer
	TermID *string

	// Plan Display Name for the marketplace offer
	TermUnit *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// REQUIRED; The Operation items on this page
	Value []*Operation

	// The link to the next page of items
	NextLink *string
}

// PartnerProperties - Partner's specific Properties
type PartnerProperties struct {
	// REQUIRED; The region of the instance
	Region *Region

	// REQUIRED; The subdomain of the instance
	Subdomain *string
}

// SingleSignOnPropertiesV2 - Properties specific to Single Sign On Resource
type SingleSignOnPropertiesV2 struct {
	// REQUIRED; Type of Single Sign-On mechanism being used
	Type *SingleSignOnType

	// List of AAD domains fetched from Microsoft Graph for user.
	AADDomains []*string

	// AAD enterprise application Id used to setup SSO
	EnterpriseAppID *string

	// State of the Single Sign On for the resource
	State *SingleSignOnStates

	// URL for SSO to be used by the partner to redirect the user to their system
	URL *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// UserDetails - User details for an organization
type UserDetails struct {
	// Email address of the user
	EmailAddress *string

	// First name of the user
	FirstName *string

	// Last name of the user
	LastName *string

	// User's phone number
	PhoneNumber *string

	// User's principal name
	Upn *string
}
