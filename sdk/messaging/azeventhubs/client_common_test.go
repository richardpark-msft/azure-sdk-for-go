// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/internal/test"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"github.com/stretchr/testify/require"
)

func TestConsumerClientNormalizeEndpoint(t *testing.T) {
	testParams := test.GetConnectionParamsForTest(t)

	armClient, err := armeventhub.NewNamespacesClient(testParams.SubscriptionID, testParams.Cred, nil)
	require.NoError(t, err)

	hostnameNoSuffix := strings.Split(testParams.EventHubNamespace, ".")[0]

	resp, err := armClient.Get(context.Background(), testParams.ResourceGroup, hostnameNoSuffix, nil)
	require.NoError(t, err)

	// Even though it says 'ServiceBusEndpoint', we are using the Event Hubs ARM package. They have a shared heritage.
	consumerClient, err := azeventhubs.NewConsumerClient(*resp.Properties.ServiceBusEndpoint, testParams.EventHubName, azeventhubs.DefaultConsumerGroup, testParams.Cred, nil)
	require.NoError(t, err)

	_, err = consumerClient.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)

	producerClient, err := azeventhubs.NewProducerClient(*resp.Properties.ServiceBusEndpoint, testParams.EventHubName, testParams.Cred, nil)
	require.NoError(t, err)

	_, err = producerClient.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)
}
