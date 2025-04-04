// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armimpactreporting

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
	moduleVersion = "v0.1.0"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// ConfidenceLevel - Degree of confidence on the impact being a platform issue.
type ConfidenceLevel string

const (
	// ConfidenceLevelHigh - High confidence on azure being the source of impact
	ConfidenceLevelHigh ConfidenceLevel = "High"
	// ConfidenceLevelLow - Low confidence on azure being the source of impact
	ConfidenceLevelLow ConfidenceLevel = "Low"
	// ConfidenceLevelMedium - Medium confidence on azure being the source of impact
	ConfidenceLevelMedium ConfidenceLevel = "Medium"
)

// PossibleConfidenceLevelValues returns the possible values for the ConfidenceLevel const type.
func PossibleConfidenceLevelValues() []ConfidenceLevel {
	return []ConfidenceLevel{
		ConfidenceLevelHigh,
		ConfidenceLevelLow,
		ConfidenceLevelMedium,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// IncidentSource - List of incident interfaces.
type IncidentSource string

const (
	// IncidentSourceAzureDevops - When source of Incident is AzureDevops
	IncidentSourceAzureDevops IncidentSource = "AzureDevops"
	// IncidentSourceICM - When source of Incident is Microsoft ICM
	IncidentSourceICM IncidentSource = "ICM"
	// IncidentSourceJira - When source of Incident is Jira
	IncidentSourceJira IncidentSource = "Jira"
	// IncidentSourceOther - When source of Incident is Other
	IncidentSourceOther IncidentSource = "Other"
	// IncidentSourceServiceNow - When source of Incident is ServiceNow
	IncidentSourceServiceNow IncidentSource = "ServiceNow"
)

// PossibleIncidentSourceValues returns the possible values for the IncidentSource const type.
func PossibleIncidentSourceValues() []IncidentSource {
	return []IncidentSource{
		IncidentSourceAzureDevops,
		IncidentSourceICM,
		IncidentSourceJira,
		IncidentSourceOther,
		IncidentSourceServiceNow,
	}
}

// MetricUnit - List of unit of the metric.
type MetricUnit string

const (
	// MetricUnitByteSeconds - When measurement is in ByteSeconds
	MetricUnitByteSeconds MetricUnit = "ByteSeconds"
	// MetricUnitBytes - When measurement is in Bytes
	MetricUnitBytes MetricUnit = "Bytes"
	// MetricUnitBytesPerSecond - When measurement is in BytesPerSecond
	MetricUnitBytesPerSecond MetricUnit = "BytesPerSecond"
	// MetricUnitCores - When measurement is in Cores
	MetricUnitCores MetricUnit = "Cores"
	// MetricUnitCount - When measurement is in Count
	MetricUnitCount MetricUnit = "Count"
	// MetricUnitCountPerSecond - When measurement is in CountPerSecond
	MetricUnitCountPerSecond MetricUnit = "CountPerSecond"
	// MetricUnitMilliCores - When measurement is in MilliCores
	MetricUnitMilliCores MetricUnit = "MilliCores"
	// MetricUnitMilliSeconds - When measurement is in MilliSeconds
	MetricUnitMilliSeconds MetricUnit = "MilliSeconds"
	// MetricUnitNanoCores - When measurement is in NanoCores
	MetricUnitNanoCores MetricUnit = "NanoCores"
	// MetricUnitOther - When measurement is in Other than listed
	MetricUnitOther MetricUnit = "Other"
	// MetricUnitPercent - When measurement is in Percent
	MetricUnitPercent MetricUnit = "Percent"
	// MetricUnitSeconds - When measurement is in Seconds
	MetricUnitSeconds MetricUnit = "Seconds"
)

// PossibleMetricUnitValues returns the possible values for the MetricUnit const type.
func PossibleMetricUnitValues() []MetricUnit {
	return []MetricUnit{
		MetricUnitByteSeconds,
		MetricUnitBytes,
		MetricUnitBytesPerSecond,
		MetricUnitCores,
		MetricUnitCount,
		MetricUnitCountPerSecond,
		MetricUnitMilliCores,
		MetricUnitMilliSeconds,
		MetricUnitNanoCores,
		MetricUnitOther,
		MetricUnitPercent,
		MetricUnitSeconds,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// Platform - Enum for connector types
type Platform string

const (
	// PlatformAzureMonitor - Type of Azure Monitor
	PlatformAzureMonitor Platform = "AzureMonitor"
)

// PossiblePlatformValues returns the possible values for the Platform const type.
func PossiblePlatformValues() []Platform {
	return []Platform{
		PlatformAzureMonitor,
	}
}

// Protocol - List of protocols
type Protocol string

const (
	// ProtocolFTP - When communication protocol is FTP
	ProtocolFTP Protocol = "FTP"
	// ProtocolHTTP - When communication protocol is HTTP
	ProtocolHTTP Protocol = "HTTP"
	// ProtocolHTTPS - When communication protocol is HTTPS
	ProtocolHTTPS Protocol = "HTTPS"
	// ProtocolOther - When communication protocol is Other
	ProtocolOther Protocol = "Other"
	// ProtocolRDP - When communication protocol is RDP
	ProtocolRDP Protocol = "RDP"
	// ProtocolSSH - When communication protocol is SSH
	ProtocolSSH Protocol = "SSH"
	// ProtocolTCP - When communication protocol is TCP
	ProtocolTCP Protocol = "TCP"
	// ProtocolUDP - When communication protocol is UDP
	ProtocolUDP Protocol = "UDP"
)

// PossibleProtocolValues returns the possible values for the Protocol const type.
func PossibleProtocolValues() []Protocol {
	return []Protocol{
		ProtocolFTP,
		ProtocolHTTP,
		ProtocolHTTPS,
		ProtocolOther,
		ProtocolRDP,
		ProtocolSSH,
		ProtocolTCP,
		ProtocolUDP,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Provisioning Canceled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateFailed - Provisioning Failed
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Provisioning Succeeded
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// Toolset - List of azure interfaces.
type Toolset string

const (
	// ToolsetARM - If communication toolset is ARM
	ToolsetARM Toolset = "ARM"
	// ToolsetAnsible - If communication toolset is Ansible
	ToolsetAnsible Toolset = "Ansible"
	// ToolsetChef - If communication toolset is Chef
	ToolsetChef Toolset = "Chef"
	// ToolsetOther - If communication toolset is Other
	ToolsetOther Toolset = "Other"
	// ToolsetPortal - If communication toolset is Portal
	ToolsetPortal Toolset = "Portal"
	// ToolsetPuppet - If communication toolset is Puppet
	ToolsetPuppet Toolset = "Puppet"
	// ToolsetSDK - If communication toolset is SDK
	ToolsetSDK Toolset = "SDK"
	// ToolsetShell - If communication toolset is Shell
	ToolsetShell Toolset = "Shell"
	// ToolsetTerraform - If communication toolset is Terraform
	ToolsetTerraform Toolset = "Terraform"
)

// PossibleToolsetValues returns the possible values for the Toolset const type.
func PossibleToolsetValues() []Toolset {
	return []Toolset{
		ToolsetARM,
		ToolsetAnsible,
		ToolsetChef,
		ToolsetOther,
		ToolsetPortal,
		ToolsetPuppet,
		ToolsetSDK,
		ToolsetShell,
		ToolsetTerraform,
	}
}
