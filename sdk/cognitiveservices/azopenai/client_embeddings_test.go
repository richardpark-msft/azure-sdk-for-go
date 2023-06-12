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
	t.Run("azure", func(t *testing.T) {
		dac, err := azidentity.NewDefaultAzureCredential(nil)
		require.NoError(t, err)

		azureClient, err := NewClient(testVars.endpoint, dac, testVars.embeddingDeploymentID, nil)
		require.NoError(t, err)
		runEmbeddingsTest(t, azureClient)
	})

	t.Run("openAI", func(t *testing.T) {
		client := getOpenAIClient(t)
		runEmbeddingsTest(t, client)
	})
}

func TestClient_ReplicatePyNoteBookTest(t *testing.T) {
	// embeddingClient := getTokenCredentialClient(t, testVars.embeddingDeploymentID)

	// mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(testVars.mongoCS))
	// require.NoError(t, err)

	// db := mongoClient.Database("test")

	// db.RunCommand(context.Background(), struct {} {
	// 	createIndexes: 'exampleCollection',
	// 	indexes: [
	// 	  {
	// 		name: 'vectorSearchIndex',
	// 		key: {
	// 		  "vectorContent": "cosmosSearch"
	// 		},
	// 		cosmosSearchOptions: {
	// 		  kind: 'vector-ivf',
	// 		  numLists: 100,
	// 		  similarity: 'COS',
	// 		  dimensions: 3
	// 		}
	// 	  }
	// 	]
	//   })

	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

	// redis.Comm

	// statusCmd := client.Ping(context.Background())
	// _, err := statusCmd.Result()
	// require.NoError(t, err)

	// TODO: this is really required, not optional, so it should be renamed.
	// chatClient := getTokenCredentialClient(t, testVars.completionsDeploymentID)
	// opts := ChatCompletionsOptions{
	// 	Messages: []*ChatMessage{
	// 		// TODO: Role is a bit annoying here. Maybe we can just export them as ptr-string, since
	// 		// that's how they'll be used.
	// 		{Role: to.Ptr(ChatRoleUser), Content: to.Ptr("Is Sam Bankman-Fried's company, FTX, considered a well-managed company?")},
	// 	},
	// 	Temperature: to.Ptr[float32](1.0),
	// }

	// https://platform.openai.com/docs/api-reference/embeddings
	// resp, err := embeddingClient.GetEmbeddings(context.Background(), EmbeddingsOptions{
	// 	// `Input` is an any field (can be string, or array of string)
	// 	Input: []string{
	// 		"context",
	// 	},
	// 	Model: &testVars.embeddingDeploymentID,
	// 	User:  to.Ptr(t.Name()),
	// }, nil)
	// require.NoError(t, err)
	// require.NotNil(t, resp)

	// So for the content that we want to use we calculate embeddings.
	// Then for the thing we want to check against we calculate embeddings.
	// Then we feed that into "something" that understands how to do vector comparisons in a useful way
	// and then use that.

	// in their example they use Redis, so let's just do that.
}

func runEmbeddingsTest(t *testing.T, client *Client) {
	resp, err := client.GetEmbeddings(context.Background(), EmbeddingsOptions{
		Input: "Your text string goes here",

		// NOTE: this parameter is ignored, in Azure OpenAI. The model is specified when
		// the client is opened as the deploymentID.
		Model: to.Ptr("text-similarity-curie-001"),
	}, nil)
	require.NoError(t, err)

	require.Equal(t, 4096, len(resp.Data[0].Embedding))
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

func getTokenCredentialClient(t *testing.T, deploymentID string) *Client {
	dac, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	client, err := NewClient(testVars.endpoint, dac, deploymentID, nil)
	require.NoError(t, err)

	return client
}
