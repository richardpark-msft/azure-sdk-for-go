// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

// Operations
const (
	lockRenewalOperationName   = "com.microsoft:renew-lock"
	peekMessageOperationID     = "com.microsoft:peek-message"
	scheduleMessageOperationID = "com.microsoft:schedule-message"
	cancelScheduledOperationID = "com.microsoft:cancel-scheduled-message"
)

// Field Descriptions
const (
	OperationFieldName     = "operation"
	LockTokensFieldName    = "lock-tokens"
	ServerTimeoutFieldName = "com.microsoft:server-timeout"
	AssociatedLinkName     = "associated-link-name"
)
