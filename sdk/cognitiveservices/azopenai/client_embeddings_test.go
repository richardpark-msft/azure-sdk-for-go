//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestClient_GetEmbeddings(t *testing.T) {
	cred := KeyCredential{APIKey: testVars.apiKey}
	client, err := NewClientWithKeyCredential(testVars.endpoint, cred, testVars.embeddingDeploymentID, nil)
	require.NoError(t, err)

	resp, err := client.GetEmbeddings(context.Background(), EmbeddingsOptions{
		Input: "Your text string goes here",
		Model: to.Ptr("text-similarity-curie-001"),
	}, nil)
	require.NoError(t, err)

	// require.Equal(t, ClientGetEmbeddingsResponse{
	// 	Embeddings{
	// 		Data:  []*EmbeddingItem{},
	// 		Usage: &EmbeddingsUsage{},
	// 	},
	// }, resp)

	require.Equal(t, len(resp.Data[0].Embedding), 4096)
}
