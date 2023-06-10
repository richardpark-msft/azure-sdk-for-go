//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/require"
)

func TestClient_GetEmbeddings(t *testing.T) {
	dac, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	t.Run("azure", func(t *testing.T) {
		azureClient, err := NewClient(testVars.endpoint, dac, testVars.embeddingDeploymentID, nil)
		require.NoError(t, err)
		runEmbeddingsTest(t, azureClient)
	})

	t.Run("openAI", func(t *testing.T) {
		client := getOpenAIClient(t)
		runEmbeddingsTest(t, client)
	})
}

func runEmbeddingsTest(t *testing.T, client *Client) {
	resp, err := client.GetEmbeddings(context.Background(), EmbeddingsOptions{
		Input: "Your text string goes here",

		// NOTE: this parameter is ignored, in Azure OpenAI. The model is specified when
		// the client is opened as the deploymentID.
		Model: to.Ptr("text-similarity-curie-001"),
	}, nil)
	require.NoError(t, err)

	require.Equal(t, len(resp.Data[0].Embedding), 4096)
}

func getOpenAIClient(t *testing.T) *Client {
	if testVars.openAIKey == "" {
		fmt.Printf("WARNING: no OpenAI key, skipping OpenAI endpoint tests.\n")
		t.SkipNow()
	}

	openAIClient, err := NewClientForOpenAI(testVars.openAIEndpoint, KeyCredential{APIKey: testVars.openAIKey}, nil)
	require.NoError(t, err)

	return openAIClient
}
