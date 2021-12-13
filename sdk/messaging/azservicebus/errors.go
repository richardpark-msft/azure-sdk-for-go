// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"

var errReceiveAndDeleteReceiver = internal.ErrNonRetriable{Message: "messages that are received in `ReceiveModeReceiveAndDelete` mode are not settleable"}

func (e errClosedPermanently) Error() string { return "Link has been closed permanently" }
func (e errClosedPermanently) nonRetriable() {}

// 	return fmt.Sprintf("%s is closed and can no longer be used", ec.link)
