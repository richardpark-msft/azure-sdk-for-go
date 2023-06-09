//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai

import (
	"context"
	"log"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

func TestClient_GetEmbeddings(t *testing.T) {
	type args struct {
		ctx          context.Context
		deploymentID string
		body         EmbeddingsOptions
		options      *ClientGetEmbeddingsOptions
	}
	deploymentID := "embedding"
	cred := KeyCredential{APIKey: apiKey}
	client, err := NewClientWithKeyCredential(endpoint, cred, deploymentID, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		want    ClientGetEmbeddingsResponse
		wantErr bool
	}{
		{
			name:   "Embeddings",
			client: client,
			args: args{
				ctx:          context.TODO(),
				deploymentID: "embedding",
				body: EmbeddingsOptions{
					Input: "Your text string goes here",
					Model: to.Ptr("text-similarity-curie-001"),
				},
				options: nil,
			},
			want: ClientGetEmbeddingsResponse{
				Embeddings{
					Data:  []*EmbeddingItem{},
					Usage: &EmbeddingsUsage{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.GetEmbeddings(tt.args.ctx, tt.args.body, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetEmbeddings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Embeddings.Data[0].Embedding) != 4096 {
				t.Errorf("Client.GetEmbeddings() len(Data) want 4096, got %d", len(got.Embeddings.Data))
				return
			}
		})
	}
}
