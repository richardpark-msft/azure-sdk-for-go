package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessMode enumerates the values for access mode.
type AccessMode string

const (
	// AccessModeOpen ...
	AccessModeOpen AccessMode = "Open"
	// AccessModePrivateOnly ...
	AccessModePrivateOnly AccessMode = "PrivateOnly"
)

// PossibleAccessModeValues returns an array of possible values for the AccessMode const type.
func PossibleAccessModeValues() []AccessMode {
	return []AccessMode{AccessModeOpen, AccessModePrivateOnly}
}

// CategoryType enumerates the values for category type.
type CategoryType string

const (
	// CategoryTypeLogs ...
	CategoryTypeLogs CategoryType = "Logs"
	// CategoryTypeMetrics ...
	CategoryTypeMetrics CategoryType = "Metrics"
)

// PossibleCategoryTypeValues returns an array of possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{CategoryTypeLogs, CategoryTypeMetrics}
}

// ComparisonOperationType enumerates the values for comparison operation type.
type ComparisonOperationType string

const (
	// ComparisonOperationTypeEquals ...
	ComparisonOperationTypeEquals ComparisonOperationType = "Equals"
	// ComparisonOperationTypeGreaterThan ...
	ComparisonOperationTypeGreaterThan ComparisonOperationType = "GreaterThan"
	// ComparisonOperationTypeGreaterThanOrEqual ...
	ComparisonOperationTypeGreaterThanOrEqual ComparisonOperationType = "GreaterThanOrEqual"
	// ComparisonOperationTypeLessThan ...
	ComparisonOperationTypeLessThan ComparisonOperationType = "LessThan"
	// ComparisonOperationTypeLessThanOrEqual ...
	ComparisonOperationTypeLessThanOrEqual ComparisonOperationType = "LessThanOrEqual"
	// ComparisonOperationTypeNotEquals ...
	ComparisonOperationTypeNotEquals ComparisonOperationType = "NotEquals"
)

// PossibleComparisonOperationTypeValues returns an array of possible values for the ComparisonOperationType const type.
func PossibleComparisonOperationTypeValues() []ComparisonOperationType {
	return []ComparisonOperationType{ComparisonOperationTypeEquals, ComparisonOperationTypeGreaterThan, ComparisonOperationTypeGreaterThanOrEqual, ComparisonOperationTypeLessThan, ComparisonOperationTypeLessThanOrEqual, ComparisonOperationTypeNotEquals}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// MetricStatisticType enumerates the values for metric statistic type.
type MetricStatisticType string

const (
	// MetricStatisticTypeAverage ...
	MetricStatisticTypeAverage MetricStatisticType = "Average"
	// MetricStatisticTypeCount ...
	MetricStatisticTypeCount MetricStatisticType = "Count"
	// MetricStatisticTypeMax ...
	MetricStatisticTypeMax MetricStatisticType = "Max"
	// MetricStatisticTypeMin ...
	MetricStatisticTypeMin MetricStatisticType = "Min"
	// MetricStatisticTypeSum ...
	MetricStatisticTypeSum MetricStatisticType = "Sum"
)

// PossibleMetricStatisticTypeValues returns an array of possible values for the MetricStatisticType const type.
func PossibleMetricStatisticTypeValues() []MetricStatisticType {
	return []MetricStatisticType{MetricStatisticTypeAverage, MetricStatisticTypeCount, MetricStatisticTypeMax, MetricStatisticTypeMin, MetricStatisticTypeSum}
}

// PredictiveAutoscalePolicyScaleMode enumerates the values for predictive autoscale policy scale mode.
type PredictiveAutoscalePolicyScaleMode string

const (
	// PredictiveAutoscalePolicyScaleModeDisabled ...
	PredictiveAutoscalePolicyScaleModeDisabled PredictiveAutoscalePolicyScaleMode = "Disabled"
	// PredictiveAutoscalePolicyScaleModeEnabled ...
	PredictiveAutoscalePolicyScaleModeEnabled PredictiveAutoscalePolicyScaleMode = "Enabled"
	// PredictiveAutoscalePolicyScaleModeForecastOnly ...
	PredictiveAutoscalePolicyScaleModeForecastOnly PredictiveAutoscalePolicyScaleMode = "ForecastOnly"
)

// PossiblePredictiveAutoscalePolicyScaleModeValues returns an array of possible values for the PredictiveAutoscalePolicyScaleMode const type.
func PossiblePredictiveAutoscalePolicyScaleModeValues() []PredictiveAutoscalePolicyScaleMode {
	return []PredictiveAutoscalePolicyScaleMode{PredictiveAutoscalePolicyScaleModeDisabled, PredictiveAutoscalePolicyScaleModeEnabled, PredictiveAutoscalePolicyScaleModeForecastOnly}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// PrivateEndpointServiceConnectionStatusApproved ...
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	// PrivateEndpointServiceConnectionStatusPending ...
	PrivateEndpointServiceConnectionStatusPending PrivateEndpointServiceConnectionStatus = "Pending"
	// PrivateEndpointServiceConnectionStatusRejected ...
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{PrivateEndpointServiceConnectionStatusApproved, PrivateEndpointServiceConnectionStatusPending, PrivateEndpointServiceConnectionStatusRejected}
}

// ReceiverStatus enumerates the values for receiver status.
type ReceiverStatus string

const (
	// ReceiverStatusDisabled ...
	ReceiverStatusDisabled ReceiverStatus = "Disabled"
	// ReceiverStatusEnabled ...
	ReceiverStatusEnabled ReceiverStatus = "Enabled"
	// ReceiverStatusNotSpecified ...
	ReceiverStatusNotSpecified ReceiverStatus = "NotSpecified"
)

// PossibleReceiverStatusValues returns an array of possible values for the ReceiverStatus const type.
func PossibleReceiverStatusValues() []ReceiverStatus {
	return []ReceiverStatus{ReceiverStatusDisabled, ReceiverStatusEnabled, ReceiverStatusNotSpecified}
}

// RecurrenceFrequency enumerates the values for recurrence frequency.
type RecurrenceFrequency string

const (
	// RecurrenceFrequencyDay ...
	RecurrenceFrequencyDay RecurrenceFrequency = "Day"
	// RecurrenceFrequencyHour ...
	RecurrenceFrequencyHour RecurrenceFrequency = "Hour"
	// RecurrenceFrequencyMinute ...
	RecurrenceFrequencyMinute RecurrenceFrequency = "Minute"
	// RecurrenceFrequencyMonth ...
	RecurrenceFrequencyMonth RecurrenceFrequency = "Month"
	// RecurrenceFrequencyNone ...
	RecurrenceFrequencyNone RecurrenceFrequency = "None"
	// RecurrenceFrequencySecond ...
	RecurrenceFrequencySecond RecurrenceFrequency = "Second"
	// RecurrenceFrequencyWeek ...
	RecurrenceFrequencyWeek RecurrenceFrequency = "Week"
	// RecurrenceFrequencyYear ...
	RecurrenceFrequencyYear RecurrenceFrequency = "Year"
)

// PossibleRecurrenceFrequencyValues returns an array of possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{RecurrenceFrequencyDay, RecurrenceFrequencyHour, RecurrenceFrequencyMinute, RecurrenceFrequencyMonth, RecurrenceFrequencyNone, RecurrenceFrequencySecond, RecurrenceFrequencyWeek, RecurrenceFrequencyYear}
}

// ScaleDirection enumerates the values for scale direction.
type ScaleDirection string

const (
	// ScaleDirectionDecrease ...
	ScaleDirectionDecrease ScaleDirection = "Decrease"
	// ScaleDirectionIncrease ...
	ScaleDirectionIncrease ScaleDirection = "Increase"
	// ScaleDirectionNone ...
	ScaleDirectionNone ScaleDirection = "None"
)

// PossibleScaleDirectionValues returns an array of possible values for the ScaleDirection const type.
func PossibleScaleDirectionValues() []ScaleDirection {
	return []ScaleDirection{ScaleDirectionDecrease, ScaleDirectionIncrease, ScaleDirectionNone}
}

// ScaleRuleMetricDimensionOperationType enumerates the values for scale rule metric dimension operation type.
type ScaleRuleMetricDimensionOperationType string

const (
	// ScaleRuleMetricDimensionOperationTypeEquals ...
	ScaleRuleMetricDimensionOperationTypeEquals ScaleRuleMetricDimensionOperationType = "Equals"
	// ScaleRuleMetricDimensionOperationTypeNotEquals ...
	ScaleRuleMetricDimensionOperationTypeNotEquals ScaleRuleMetricDimensionOperationType = "NotEquals"
)

// PossibleScaleRuleMetricDimensionOperationTypeValues returns an array of possible values for the ScaleRuleMetricDimensionOperationType const type.
func PossibleScaleRuleMetricDimensionOperationTypeValues() []ScaleRuleMetricDimensionOperationType {
	return []ScaleRuleMetricDimensionOperationType{ScaleRuleMetricDimensionOperationTypeEquals, ScaleRuleMetricDimensionOperationTypeNotEquals}
}

// ScaleType enumerates the values for scale type.
type ScaleType string

const (
	// ScaleTypeChangeCount ...
	ScaleTypeChangeCount ScaleType = "ChangeCount"
	// ScaleTypeExactCount ...
	ScaleTypeExactCount ScaleType = "ExactCount"
	// ScaleTypePercentChangeCount ...
	ScaleTypePercentChangeCount ScaleType = "PercentChangeCount"
	// ScaleTypeServiceAllowedNextValue ...
	ScaleTypeServiceAllowedNextValue ScaleType = "ServiceAllowedNextValue"
)

// PossibleScaleTypeValues returns an array of possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{ScaleTypeChangeCount, ScaleTypeExactCount, ScaleTypePercentChangeCount, ScaleTypeServiceAllowedNextValue}
}

// TimeAggregationType enumerates the values for time aggregation type.
type TimeAggregationType string

const (
	// TimeAggregationTypeAverage ...
	TimeAggregationTypeAverage TimeAggregationType = "Average"
	// TimeAggregationTypeCount ...
	TimeAggregationTypeCount TimeAggregationType = "Count"
	// TimeAggregationTypeLast ...
	TimeAggregationTypeLast TimeAggregationType = "Last"
	// TimeAggregationTypeMaximum ...
	TimeAggregationTypeMaximum TimeAggregationType = "Maximum"
	// TimeAggregationTypeMinimum ...
	TimeAggregationTypeMinimum TimeAggregationType = "Minimum"
	// TimeAggregationTypeTotal ...
	TimeAggregationTypeTotal TimeAggregationType = "Total"
)

// PossibleTimeAggregationTypeValues returns an array of possible values for the TimeAggregationType const type.
func PossibleTimeAggregationTypeValues() []TimeAggregationType {
	return []TimeAggregationType{TimeAggregationTypeAverage, TimeAggregationTypeCount, TimeAggregationTypeLast, TimeAggregationTypeMaximum, TimeAggregationTypeMinimum, TimeAggregationTypeTotal}
}