//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import "encoding/json"

func unmarshalAdditionalDataClassification(rawMsg json.RawMessage) (AdditionalDataClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AdditionalDataClassification
	switch m["assessedResourceType"] {
	case "ServerVulnerabilityAssessment":
		b = &ServerVulnerabilityProperties{}
	case string(AssessedResourceTypeContainerRegistryVulnerability):
		b = &ContainerRegistryVulnerabilityProperties{}
	case string(AssessedResourceTypeSQLServerVulnerability):
		b = &SQLServerVulnerabilityProperties{}
	default:
		b = &AdditionalData{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAlertSimulatorRequestPropertiesClassification(rawMsg json.RawMessage) (AlertSimulatorRequestPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AlertSimulatorRequestPropertiesClassification
	switch m["kind"] {
	case string(KindBundles):
		b = &AlertSimulatorBundlesRequestProperties{}
	default:
		b = &AlertSimulatorRequestProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAllowlistCustomAlertRuleClassification(rawMsg json.RawMessage) (AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AllowlistCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ConnectionFromIpNotAllowed":
		b = &ConnectionFromIPNotAllowed{}
	case "ConnectionToIpNotAllowed":
		b = &ConnectionToIPNotAllowed{}
	case "LocalUserNotAllowed":
		b = &LocalUserNotAllowed{}
	case "ProcessNotAllowed":
		b = &ProcessNotAllowed{}
	default:
		b = &AllowlistCustomAlertRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAllowlistCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AllowlistCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAllowlistCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAuthenticationDetailsPropertiesClassification(rawMsg json.RawMessage) (AuthenticationDetailsPropertiesClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthenticationDetailsPropertiesClassification
	switch m["authenticationType"] {
	case string(AuthenticationTypeAwsAssumeRole):
		b = &AwAssumeRoleAuthenticationDetailsProperties{}
	case string(AuthenticationTypeAwsCreds):
		b = &AwsCredsAuthenticationDetailsProperties{}
	case string(AuthenticationTypeGcpCredentials):
		b = &GcpCredentialsDetailsProperties{}
	default:
		b = &AuthenticationDetailsProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAutomationActionClassification(rawMsg json.RawMessage) (AutomationActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AutomationActionClassification
	switch m["actionType"] {
	case string(ActionTypeEventHub):
		b = &AutomationActionEventHub{}
	case string(ActionTypeLogicApp):
		b = &AutomationActionLogicApp{}
	case string(ActionTypeWorkspace):
		b = &AutomationActionWorkspace{}
	default:
		b = &AutomationAction{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalAutomationActionClassificationArray(rawMsg json.RawMessage) ([]AutomationActionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AutomationActionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAutomationActionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalAwsOrganizationalDataClassification(rawMsg json.RawMessage) (AwsOrganizationalDataClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AwsOrganizationalDataClassification
	switch m["organizationMembershipType"] {
	case string(OrganizationMembershipTypeMember):
		b = &AwsOrganizationalDataMember{}
	case string(OrganizationMembershipTypeOrganization):
		b = &AwsOrganizationalDataMaster{}
	default:
		b = &AwsOrganizationalData{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCloudOfferingClassification(rawMsg json.RawMessage) (CloudOfferingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b CloudOfferingClassification
	switch m["offeringType"] {
	case string(OfferingTypeCspmMonitorAws):
		b = &CspmMonitorAwsOffering{}
	case string(OfferingTypeCspmMonitorAzureDevOps):
		b = &CspmMonitorAzureDevOpsOffering{}
	case string(OfferingTypeCspmMonitorGcp):
		b = &CspmMonitorGcpOffering{}
	case string(OfferingTypeCspmMonitorGitLab):
		b = &CspmMonitorGitLabOffering{}
	case string(OfferingTypeCspmMonitorGithub):
		b = &CspmMonitorGithubOffering{}
	case string(OfferingTypeDefenderCspmAws):
		b = &DefenderCspmAwsOffering{}
	case string(OfferingTypeDefenderCspmGcp):
		b = &DefenderCspmGcpOffering{}
	case string(OfferingTypeDefenderForContainersAws):
		b = &DefenderForContainersAwsOffering{}
	case string(OfferingTypeDefenderForContainersGcp):
		b = &DefenderForContainersGcpOffering{}
	case string(OfferingTypeDefenderForDatabasesAws):
		b = &DefenderFoDatabasesAwsOffering{}
	case string(OfferingTypeDefenderForDatabasesGcp):
		b = &DefenderForDatabasesGcpOffering{}
	case string(OfferingTypeDefenderForServersAws):
		b = &DefenderForServersAwsOffering{}
	case string(OfferingTypeDefenderForServersGcp):
		b = &DefenderForServersGcpOffering{}
	default:
		b = &CloudOffering{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalCloudOfferingClassificationArray(rawMsg json.RawMessage) ([]CloudOfferingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CloudOfferingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCloudOfferingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalEnvironmentDataClassification(rawMsg json.RawMessage) (EnvironmentDataClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EnvironmentDataClassification
	switch m["environmentType"] {
	case string(EnvironmentTypeAwsAccount):
		b = &AwsEnvironmentData{}
	case string(EnvironmentTypeAzureDevOpsScope):
		b = &AzureDevOpsScopeEnvironmentData{}
	case string(EnvironmentTypeGcpProject):
		b = &GcpProjectEnvironmentData{}
	case string(EnvironmentTypeGithubScope):
		b = &GithubScopeEnvironmentData{}
	case string(EnvironmentTypeGitlabScope):
		b = &GitlabScopeEnvironmentData{}
	default:
		b = &EnvironmentData{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalExternalSecuritySolutionClassification(rawMsg json.RawMessage) (ExternalSecuritySolutionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ExternalSecuritySolutionClassification
	switch m["kind"] {
	case string(ExternalSecuritySolutionKindAAD):
		b = &AADExternalSecuritySolution{}
	case string(ExternalSecuritySolutionKindATA):
		b = &AtaExternalSecuritySolution{}
	case string(ExternalSecuritySolutionKindCEF):
		b = &CefExternalSecuritySolution{}
	default:
		b = &ExternalSecuritySolution{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalExternalSecuritySolutionClassificationArray(rawMsg json.RawMessage) ([]ExternalSecuritySolutionClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ExternalSecuritySolutionClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalExternalSecuritySolutionClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalGcpOrganizationalDataClassification(rawMsg json.RawMessage) (GcpOrganizationalDataClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GcpOrganizationalDataClassification
	switch m["organizationMembershipType"] {
	case string(OrganizationMembershipTypeMember):
		b = &GcpOrganizationalDataMember{}
	case string(OrganizationMembershipTypeOrganization):
		b = &GcpOrganizationalDataOrganization{}
	default:
		b = &GcpOrganizationalData{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNotificationsSourceClassification(rawMsg json.RawMessage) (NotificationsSourceClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NotificationsSourceClassification
	switch m["sourceType"] {
	case string(SourceTypeAlert):
		b = &NotificationsSourceAlert{}
	case string(SourceTypeAttackPath):
		b = &NotificationsSourceAttackPath{}
	default:
		b = &NotificationsSource{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNotificationsSourceClassificationArray(rawMsg json.RawMessage) ([]NotificationsSourceClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]NotificationsSourceClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalNotificationsSourceClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalResourceDetailsClassification(rawMsg json.RawMessage) (ResourceDetailsClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceDetailsClassification
	switch m["source"] {
	case string(SourceAzure):
		b = &AzureResourceDetails{}
	case string(SourceOnPremise):
		b = &OnPremiseResourceDetails{}
	case string(SourceOnPremiseSQL):
		b = &OnPremiseSQLResourceDetails{}
	default:
		b = &ResourceDetails{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalResourceIdentifierClassification(rawMsg json.RawMessage) (ResourceIdentifierClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourceIdentifierClassification
	switch m["type"] {
	case string(ResourceIdentifierTypeAzureResource):
		b = &AzureResourceIdentifier{}
	case string(ResourceIdentifierTypeLogAnalytics):
		b = &LogAnalyticsIdentifier{}
	default:
		b = &ResourceIdentifier{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalResourceIdentifierClassificationArray(rawMsg json.RawMessage) ([]ResourceIdentifierClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ResourceIdentifierClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalResourceIdentifierClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalServerVulnerabilityAssessmentsSettingClassification(rawMsg json.RawMessage) (ServerVulnerabilityAssessmentsSettingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServerVulnerabilityAssessmentsSettingClassification
	switch m["kind"] {
	case string(ServerVulnerabilityAssessmentsSettingKindAzureServersSetting):
		b = &AzureServersSetting{}
	default:
		b = &ServerVulnerabilityAssessmentsSetting{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalServerVulnerabilityAssessmentsSettingClassificationArray(rawMsg json.RawMessage) ([]ServerVulnerabilityAssessmentsSettingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ServerVulnerabilityAssessmentsSettingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalServerVulnerabilityAssessmentsSettingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalSettingClassification(rawMsg json.RawMessage) (SettingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b SettingClassification
	switch m["kind"] {
	case string(SettingKindAlertSyncSettings):
		b = &AlertSyncSettings{}
	case string(SettingKindDataExportSettings):
		b = &DataExportSettings{}
	default:
		b = &Setting{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalSettingClassificationArray(rawMsg json.RawMessage) ([]SettingClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]SettingClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalSettingClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalThresholdCustomAlertRuleClassification(rawMsg json.RawMessage) (ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ThresholdCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TimeWindowCustomAlertRule":
		b = &TimeWindowCustomAlertRule{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &ThresholdCustomAlertRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalThresholdCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ThresholdCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalThresholdCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTimeWindowCustomAlertRuleClassification(rawMsg json.RawMessage) (TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TimeWindowCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &TimeWindowCustomAlertRule{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTimeWindowCustomAlertRuleClassificationArray(rawMsg json.RawMessage) ([]TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TimeWindowCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTimeWindowCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
