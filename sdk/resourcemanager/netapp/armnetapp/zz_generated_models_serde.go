//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", a.Etag)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountList.
func (a AccountList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountPatch.
func (a AccountPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AccountProperties.
func (a AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "activeDirectories", a.ActiveDirectories)
	populate(objectMap, "encryption", a.Encryption)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ActiveDirectory.
func (a ActiveDirectory) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "activeDirectoryId", a.ActiveDirectoryID)
	populate(objectMap, "adName", a.AdName)
	populate(objectMap, "administrators", a.Administrators)
	populate(objectMap, "aesEncryption", a.AesEncryption)
	populate(objectMap, "allowLocalNfsUsersWithLdap", a.AllowLocalNfsUsersWithLdap)
	populate(objectMap, "backupOperators", a.BackupOperators)
	populate(objectMap, "dns", a.DNS)
	populate(objectMap, "domain", a.Domain)
	populate(objectMap, "encryptDCConnections", a.EncryptDCConnections)
	populate(objectMap, "kdcIP", a.KdcIP)
	populate(objectMap, "ldapOverTLS", a.LdapOverTLS)
	populate(objectMap, "ldapSearchScope", a.LdapSearchScope)
	populate(objectMap, "ldapSigning", a.LdapSigning)
	populate(objectMap, "organizationalUnit", a.OrganizationalUnit)
	populate(objectMap, "password", a.Password)
	populate(objectMap, "securityOperators", a.SecurityOperators)
	populate(objectMap, "serverRootCACertificate", a.ServerRootCACertificate)
	populate(objectMap, "site", a.Site)
	populate(objectMap, "smbServerName", a.SmbServerName)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "statusDetails", a.StatusDetails)
	populate(objectMap, "username", a.Username)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPatch.
func (b BackupPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", b.Properties)
	populate(objectMap, "tags", b.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPoliciesList.
func (b BackupPoliciesList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", b.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPolicy.
func (b BackupPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", b.Etag)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "location", b.Location)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "properties", b.Properties)
	populate(objectMap, "systemData", b.SystemData)
	populate(objectMap, "tags", b.Tags)
	populate(objectMap, "type", b.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPolicyDetails.
func (b BackupPolicyDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", b.ID)
	populate(objectMap, "location", b.Location)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "properties", b.Properties)
	populate(objectMap, "tags", b.Tags)
	populate(objectMap, "type", b.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPolicyPatch.
func (b BackupPolicyPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", b.ID)
	populate(objectMap, "location", b.Location)
	populate(objectMap, "name", b.Name)
	populate(objectMap, "properties", b.Properties)
	populate(objectMap, "tags", b.Tags)
	populate(objectMap, "type", b.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupPolicyProperties.
func (b BackupPolicyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "backupPolicyId", b.BackupPolicyID)
	populate(objectMap, "dailyBackupsToKeep", b.DailyBackupsToKeep)
	populate(objectMap, "enabled", b.Enabled)
	populate(objectMap, "monthlyBackupsToKeep", b.MonthlyBackupsToKeep)
	populate(objectMap, "provisioningState", b.ProvisioningState)
	populate(objectMap, "volumeBackups", b.VolumeBackups)
	populate(objectMap, "volumesAssigned", b.VolumesAssigned)
	populate(objectMap, "weeklyBackupsToKeep", b.WeeklyBackupsToKeep)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type BackupProperties.
func (b BackupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "backupId", b.BackupID)
	populate(objectMap, "backupType", b.BackupType)
	populateTimeRFC3339(objectMap, "creationDate", b.CreationDate)
	populate(objectMap, "failureReason", b.FailureReason)
	populate(objectMap, "label", b.Label)
	populate(objectMap, "provisioningState", b.ProvisioningState)
	populate(objectMap, "size", b.Size)
	populate(objectMap, "useExistingSnapshot", b.UseExistingSnapshot)
	populate(objectMap, "volumeName", b.VolumeName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BackupProperties.
func (b *BackupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backupId":
			err = unpopulate(val, &b.BackupID)
			delete(rawMsg, key)
		case "backupType":
			err = unpopulate(val, &b.BackupType)
			delete(rawMsg, key)
		case "creationDate":
			err = unpopulateTimeRFC3339(val, &b.CreationDate)
			delete(rawMsg, key)
		case "failureReason":
			err = unpopulate(val, &b.FailureReason)
			delete(rawMsg, key)
		case "label":
			err = unpopulate(val, &b.Label)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &b.ProvisioningState)
			delete(rawMsg, key)
		case "size":
			err = unpopulate(val, &b.Size)
			delete(rawMsg, key)
		case "useExistingSnapshot":
			err = unpopulate(val, &b.UseExistingSnapshot)
			delete(rawMsg, key)
		case "volumeName":
			err = unpopulate(val, &b.VolumeName)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BackupsList.
func (b BackupsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", b.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CapacityPool.
func (c CapacityPool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", c.Etag)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "systemData", c.SystemData)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CapacityPoolList.
func (c CapacityPoolList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", c.NextLink)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CapacityPoolPatch.
func (c CapacityPoolPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecification.
func (m MetricSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "category", m.Category)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "enableRegionalMdmAccount", m.EnableRegionalMdmAccount)
	populate(objectMap, "fillGapWithZero", m.FillGapWithZero)
	populate(objectMap, "internalMetricName", m.InternalMetricName)
	populate(objectMap, "isInternal", m.IsInternal)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "resourceIdDimensionNameOverride", m.ResourceIDDimensionNameOverride)
	populate(objectMap, "sourceMdmAccount", m.SourceMdmAccount)
	populate(objectMap, "sourceMdmNamespace", m.SourceMdmNamespace)
	populate(objectMap, "supportedAggregationTypes", m.SupportedAggregationTypes)
	populate(objectMap, "supportedTimeGrainTypes", m.SupportedTimeGrainTypes)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MountTarget.
func (m MountTarget) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceSpecification.
func (s ServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "logSpecifications", s.LogSpecifications)
	populate(objectMap, "metricSpecifications", s.MetricSpecifications)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotPoliciesList.
func (s SnapshotPoliciesList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotPolicy.
func (s SnapshotPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotPolicyDetails.
func (s SnapshotPolicyDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotPolicyPatch.
func (s SnapshotPolicyPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotPolicyVolumeList.
func (s SnapshotPolicyVolumeList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotProperties.
func (s SnapshotProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "created", s.Created)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "snapshotId", s.SnapshotID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SnapshotProperties.
func (s *SnapshotProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeRFC3339(val, &s.Created)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "snapshotId":
			err = unpopulate(val, &s.SnapshotID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotRestoreFiles.
func (s SnapshotRestoreFiles) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "destinationPath", s.DestinationPath)
	populate(objectMap, "filePaths", s.FilePaths)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotsList.
func (s SnapshotsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionQuotaItemList.
func (s SubscriptionQuotaItemList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubvolumeModelProperties.
func (s SubvolumeModelProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "accessedTimeStamp", s.AccessedTimeStamp)
	populate(objectMap, "bytesUsed", s.BytesUsed)
	populateTimeRFC3339(objectMap, "changedTimeStamp", s.ChangedTimeStamp)
	populateTimeRFC3339(objectMap, "creationTimeStamp", s.CreationTimeStamp)
	populateTimeRFC3339(objectMap, "modifiedTimeStamp", s.ModifiedTimeStamp)
	populate(objectMap, "parentPath", s.ParentPath)
	populate(objectMap, "path", s.Path)
	populate(objectMap, "permissions", s.Permissions)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "size", s.Size)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubvolumeModelProperties.
func (s *SubvolumeModelProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessedTimeStamp":
			err = unpopulateTimeRFC3339(val, &s.AccessedTimeStamp)
			delete(rawMsg, key)
		case "bytesUsed":
			err = unpopulate(val, &s.BytesUsed)
			delete(rawMsg, key)
		case "changedTimeStamp":
			err = unpopulateTimeRFC3339(val, &s.ChangedTimeStamp)
			delete(rawMsg, key)
		case "creationTimeStamp":
			err = unpopulateTimeRFC3339(val, &s.CreationTimeStamp)
			delete(rawMsg, key)
		case "modifiedTimeStamp":
			err = unpopulateTimeRFC3339(val, &s.ModifiedTimeStamp)
			delete(rawMsg, key)
		case "parentPath":
			err = unpopulate(val, &s.ParentPath)
			delete(rawMsg, key)
		case "path":
			err = unpopulate(val, &s.Path)
			delete(rawMsg, key)
		case "permissions":
			err = unpopulate(val, &s.Permissions)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "size":
			err = unpopulate(val, &s.Size)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SubvolumePatchRequest.
func (s SubvolumePatchRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubvolumesList.
func (s SubvolumesList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VaultList.
func (v VaultList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Volume.
func (v Volume) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", v.Etag)
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "systemData", v.SystemData)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroup.
func (v VolumeGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupDetails.
func (v VolumeGroupDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupList.
func (v VolumeGroupList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupMetaData.
func (v VolumeGroupMetaData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationIdentifier", v.ApplicationIdentifier)
	populate(objectMap, "applicationType", v.ApplicationType)
	populate(objectMap, "deploymentSpecId", v.DeploymentSpecID)
	populate(objectMap, "globalPlacementRules", v.GlobalPlacementRules)
	populate(objectMap, "groupDescription", v.GroupDescription)
	populate(objectMap, "volumesCount", v.VolumesCount)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupProperties.
func (v VolumeGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupMetaData", v.GroupMetaData)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "volumes", v.Volumes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupVolumeProperties.
func (v VolumeGroupVolumeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeList.
func (v VolumeList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumePatch.
func (v VolumePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumePatchPropertiesExportPolicy.
func (v VolumePatchPropertiesExportPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "rules", v.Rules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumeProperties.
func (v VolumeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "avsDataStore", v.AvsDataStore)
	populate(objectMap, "backupId", v.BackupID)
	populate(objectMap, "baremetalTenantId", v.BaremetalTenantID)
	populate(objectMap, "capacityPoolResourceId", v.CapacityPoolResourceID)
	populate(objectMap, "cloneProgress", v.CloneProgress)
	populate(objectMap, "coolAccess", v.CoolAccess)
	populate(objectMap, "coolnessPeriod", v.CoolnessPeriod)
	populate(objectMap, "creationToken", v.CreationToken)
	populate(objectMap, "dataProtection", v.DataProtection)
	populate(objectMap, "defaultGroupQuotaInKiBs", v.DefaultGroupQuotaInKiBs)
	populate(objectMap, "defaultUserQuotaInKiBs", v.DefaultUserQuotaInKiBs)
	populate(objectMap, "enableSubvolumes", v.EnableSubvolumes)
	populate(objectMap, "encryptionKeySource", v.EncryptionKeySource)
	populate(objectMap, "exportPolicy", v.ExportPolicy)
	populate(objectMap, "fileSystemId", v.FileSystemID)
	populate(objectMap, "isDefaultQuotaEnabled", v.IsDefaultQuotaEnabled)
	populate(objectMap, "isRestoring", v.IsRestoring)
	populate(objectMap, "kerberosEnabled", v.KerberosEnabled)
	populate(objectMap, "ldapEnabled", v.LdapEnabled)
	populate(objectMap, "maximumNumberOfFiles", v.MaximumNumberOfFiles)
	populate(objectMap, "mountTargets", v.MountTargets)
	populate(objectMap, "networkFeatures", v.NetworkFeatures)
	populate(objectMap, "networkSiblingSetId", v.NetworkSiblingSetID)
	populate(objectMap, "placementRules", v.PlacementRules)
	populate(objectMap, "protocolTypes", v.ProtocolTypes)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	populate(objectMap, "proximityPlacementGroup", v.ProximityPlacementGroup)
	populate(objectMap, "securityStyle", v.SecurityStyle)
	populate(objectMap, "serviceLevel", v.ServiceLevel)
	populate(objectMap, "smbContinuouslyAvailable", v.SmbContinuouslyAvailable)
	populate(objectMap, "smbEncryption", v.SmbEncryption)
	populate(objectMap, "snapshotDirectoryVisible", v.SnapshotDirectoryVisible)
	populate(objectMap, "snapshotId", v.SnapshotID)
	populate(objectMap, "storageToNetworkProximity", v.StorageToNetworkProximity)
	populate(objectMap, "subnetId", v.SubnetID)
	populate(objectMap, "t2Network", v.T2Network)
	populate(objectMap, "throughputMibps", v.ThroughputMibps)
	populate(objectMap, "unixPermissions", v.UnixPermissions)
	populate(objectMap, "usageThreshold", v.UsageThreshold)
	populate(objectMap, "volumeGroupName", v.VolumeGroupName)
	populate(objectMap, "volumeSpecName", v.VolumeSpecName)
	populate(objectMap, "volumeType", v.VolumeType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type VolumePropertiesExportPolicy.
func (v VolumePropertiesExportPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "rules", v.Rules)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}