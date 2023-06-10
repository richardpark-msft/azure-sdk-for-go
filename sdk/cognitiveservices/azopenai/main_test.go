//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var testVars struct {
	endpoint                    string
	apiKey                      string
	embeddingDeploymentID       string
	completionsDeploymentID     string
	chatCompletionsDeploymentID string

	openAIKey      string
	openAIEndpoint string
}

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: failed to load env file: %s\n", err)
	}

	testVars.endpoint = os.Getenv("AOAI_ENDPOINT")
	testVars.apiKey = os.Getenv("AOAI_API_KEY")
	testVars.embeddingDeploymentID = os.Getenv("AOAI_EMBEDDING_DEPLOYMENT")
	testVars.completionsDeploymentID = os.Getenv("AOAI_COMPLETIONS_DEPLOYMENT")
	testVars.chatCompletionsDeploymentID = os.Getenv("AOAI_CHAT_COMPLETIONS_DEPLOYMENT")

	// OpenAI parameters - optional. Tests that use them will be skipped if you don't provide a key.
	testVars.openAIKey = os.Getenv("OPENAI_API_KEY")
	testVars.openAIEndpoint = os.Getenv("OPENAI_ENDPOINT")

	os.Exit(m.Run())
}
