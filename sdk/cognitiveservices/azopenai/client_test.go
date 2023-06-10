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

// From the docs on the difference between ChatCompletions and Completions:
//
// > The difference between these APIs derives mainly from the underlying GPT models that are available in each. The
// > chat completions API is the interface to our most capable model (gpt-4), and our most cost effective model
// >  (gpt-3.5-turbo). For reference, gpt-3.5-turbo performs at a similar capability level to text-davinci-003 but
// > at 10% the price per token! See pricing details here.

func TestClient_GetChatCompletions(t *testing.T) {
	cred := KeyCredential{APIKey: testVars.apiKey}
	chatClient, err := NewClientWithKeyCredential(testVars.endpoint, cred, testVars.chatCompletionsDeploymentID, nil)
	require.NoError(t, err)

	resp, err := chatClient.GetChatCompletions(context.Background(), ChatCompletionsOptions{
		Messages: []*ChatMessage{
			{
				Role:    to.Ptr(ChatRoleUser),
				Content: to.Ptr("Count to 10, with a comma between each number and no newlines. E.g., 1, 2, 3, ..."),
			},
		},
		MaxTokens:   to.Ptr(int32(1024)),
		Temperature: to.Ptr(float32(0.0)),
	}, nil)
	require.NoError(t, err)

	expected := ClientGetChatCompletionsResponse{
		ChatCompletions: ChatCompletions{
			Choices: []*ChatChoice{
				{
					Message: &ChatChoiceMessage{
						Role:    to.Ptr(ChatRoleAssistant),
						Content: to.Ptr("1, 2, 3, 4, 5, 6, 7, 8, 9, 10."),
					},
					Index:        to.Ptr(int32(0)),
					FinishReason: to.Ptr(CompletionsFinishReason("stop")),
				},
			},
			Usage: &CompletionsUsage{
				CompletionTokens: to.Ptr(int32(29)),
				PromptTokens:     to.Ptr(int32(37)),
				TotalTokens:      to.Ptr(int32(66)),
			},
		},
	}

	// nix the fields that are non-deterministic.
	resp.Created = nil
	resp.ID = nil

	require.Equal(t, expected, resp)
}

func TestClient_GetCompletions(t *testing.T) {
	cred := KeyCredential{APIKey: testVars.apiKey}
	client, err := NewClientWithKeyCredential(testVars.endpoint, cred, testVars.completionsDeploymentID, nil)
	require.NoError(t, err)

	resp, err := client.GetCompletions(context.Background(), CompletionsOptions{
		Prompt:      []*string{to.Ptr("What is Azure OpenAI?")},
		MaxTokens:   to.Ptr(int32(2048 - 127)),
		Temperature: to.Ptr(float32(0.0)),
	}, nil)
	require.NoError(t, err)

	expected := ClientGetCompletionsResponse{
		Completions: Completions{
			Choices: []*Choice{
				{
					Text:         to.Ptr("\n\nAzure OpenAI is a platform from Microsoft that provides access to OpenAI's artificial intelligence (AI) technologies. It enables developers to build, train, and deploy AI models in the cloud. Azure OpenAI provides access to OpenAI's powerful AI technologies, such as GPT-3, which can be used to create natural language processing (NLP) applications, computer vision models, and reinforcement learning models."),
					Index:        to.Ptr(int32(0)),
					FinishReason: to.Ptr(CompletionsFinishReason("stop")),
					Logprobs:     nil,
				},
			},
			Usage: &CompletionsUsage{
				CompletionTokens: to.Ptr(int32(85)),
				PromptTokens:     to.Ptr(int32(6)),
				TotalTokens:      to.Ptr(int32(91)),
			},
		},
	}

	resp.ID = nil
	resp.Created = nil

	require.Equal(t, expected, resp)
}
