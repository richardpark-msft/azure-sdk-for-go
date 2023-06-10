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
	endpoint              string
	apiKey                string
	embeddingDeploymentID string
	chatDeploymentID      string
}

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: failed to load env file: %s\n", err)
	}

	testVars.endpoint = os.Getenv("AOAI_ENDPOINT")
	testVars.apiKey = os.Getenv("AOAI_API_KEY")
	testVars.embeddingDeploymentID = os.Getenv("AOAI_EMBEDDING_DEPLOYMENT")
	testVars.chatDeploymentID = os.Getenv("AOAI_CHAT_DEPLOYMENT")

	os.Exit(m.Run())
}
