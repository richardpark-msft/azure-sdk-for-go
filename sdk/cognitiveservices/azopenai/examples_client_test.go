// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai"
)

func ExampleNewClientForOpenAI() {
	// NOTE: this constructor creates a client that connects to the public OpenAI endpoint.
	// To connect to an Azure OpenAI endpoint, use azopenai.NewClient() or azopenai.NewClientWithyKeyCredential.
	keyCredential, err := azopenai.NewKeyCredential("<OpenAI-APIKey>")

	if err != nil {
		panic(err)
	}

	client, err := azopenai.NewClientForOpenAI("https://api.openai.com/v1", keyCredential, nil)

	if err != nil {
		panic(err)
	}

	_ = client
}

func ExampleNewClient() {
	// NOTE: this constructor creates a client that connects to an Azure OpenAI endpoint.
	// To connect to the public OpenAI endpoint, use azopenai.NewClientForOpenAI
	dac, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		panic(err)
	}

	modelDeploymentID := "model deployment ID"
	client, err := azopenai.NewClient("https://<your-azure-openai-host>.openai.azure.com", dac, modelDeploymentID, nil)

	if err != nil {
		panic(err)
	}

	_ = client
}

func ExampleNewClientWithKeyCredential() {
	// NOTE: this constructor creates a client that connects to an Azure OpenAI endpoint.
	// To connect to the public OpenAI endpoint, use azopenai.NewClientForOpenAI
	keyCredential, err := azopenai.NewKeyCredential("<Azure-OpenAI-APIKey>")

	if err != nil {
		panic(err)
	}

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	modelDeploymentID := "model deployment ID"
	client, err := azopenai.NewClientWithKeyCredential("https://<your-azure-openai-host>.openai.azure.com", keyCredential, modelDeploymentID, nil)

	if err != nil {
		panic(err)
	}

	_ = client
}

func ExampleClient_GetCompletionsStream() {
	azureOpenAIKey := os.Getenv("AOAI_API_KEY")
	modelDeploymentID := os.Getenv("AOAI_COMPLETIONS_MODEL_DEPLOYMENT")

	// Ex: "https://<your-azure-openai-host>.openai.azure.com"
	azureOpenAIEndpoint := os.Getenv("AOAI_ENDPOINT")

	if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
		return
	}

	keyCredential, err := azopenai.NewKeyCredential(azureOpenAIKey)

	if err != nil {
		panic(err)
	}

	// In Azure OpenAI you must deploy a model before you can use it in your client. For more information
	// see here: https://learn.microsoft.com/azure/cognitive-services/openai/how-to/create-resource
	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, modelDeploymentID, nil)

	if err != nil {
		panic(err)
	}

	resp, err := client.GetCompletionsStream(context.TODO(), azopenai.CompletionsOptions{
		Prompt:      []*string{to.Ptr("What is Azure OpenAI?")},
		MaxTokens:   to.Ptr(int32(2048)),
		Temperature: to.Ptr(float32(0.0)),
	}, nil)

	if err != nil {
		panic(err)
	}

	for {
		entry, err := resp.CompletionsStream.Read()

		if errors.Is(err, io.EOF) {
			fmt.Printf("\n *** No more completions ***\n")
			break
		}

		if err != nil {
			panic(err)
		}

		for _, choice := range entry.Choices {
			fmt.Printf("%s", *choice.Text)
		}
	}
}

func ExampleClient_CreateImage() {
	// azureOpenAIKey := os.Getenv("AOAI_API_KEY")

	// // Ex: "https://<your-azure-openai-host>.openai.azure.com"
	// azureOpenAIEndpoint := os.Getenv("AOAI_ENDPOINT")

	// if azureOpenAIKey == "" || azureOpenAIEndpoint == "" {
	// 	fmt.Printf("Skipping example, environment variables missing\n")
	// 	return
	// }

	// keyCredential, err := azopenai.NewKeyCredential(azureOpenAIKey)

	// if err != nil {
	// 	panic(err)
	// }

	// client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, "", nil)

	// if err != nil {
	// 	panic(err)
	// }

	// resp, err := client.CreateImage(context.TODO(), azopenai.ImageGenerationOptions{
	// 	Prompt:         to.Ptr("a cat"),
	// 	ResponseFormat: to.Ptr(azopenai.ImageGenerationResponseFormatURL),
	// }, nil)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, generatedImage := range resp.Data {
	// 	// the underlying type for the generatedImage is dictated by the value of
	// 	// ImageGenerationOptions.ResponseFormat. In this example we used `azopenai.ImageGenerationResponseFormatURL`,
	// 	// so the underlying type will be ImageLocation.
	// 	imageLocation := generatedImage.Result.(azopenai.ImageLocation)
	// 	fmt.Printf("Image generated at URL %q\n", *imageLocation.URL)
	// }
}

func ExamplesClient_GetChatCompletions_usingFunctions() {
	type ParamProperty struct {
		Type        string   `json:"type"`
		Description string   `json:"description,omitempty"`
		Enum        []string `json:"enum,omitempty"`
	}

	// (some JSON schema compatible types)
	type Params struct {
		Type       string                   `json:"type"`
		Properties map[string]ParamProperty `json:"properties"`
		Required   []string                 `json:"required,omitempty"`
	}

	// For examples on creating a client see:
	// (for Azure OpenAI)
	// - ExampleNewClient
	// - ExampleNewClientWithKeyCredential
	// (for OpenAI)
	// - ExampleNewClientForOpenAI
	resp, err := exampleClient.GetChatCompletions(context.Background(), azopenai.ChatCompletionsOptions{
		Model: to.Ptr("gpt-3.5-turbo-0613"),
		Messages: []*azopenai.ChatMessage{
			{
				Role:    to.Ptr(azopenai.ChatRoleUser),
				Content: to.Ptr("What's the weather like in Boston, in celsius?"),
			},
		},
		FunctionCall: azopenai.ChatCompletionsOptionsFunctionCall{
			Value: to.Ptr("auto"),
		},
		Functions: []*azopenai.FunctionDefinition{
			{
				Name:        to.Ptr("get_current_weather"),
				Description: to.Ptr("Get the current weather in a given location"),
				Parameters: Params{
					Required: []string{"location"},
					Type:     "object",
					Properties: map[string]ParamProperty{
						"location": {
							Type:        "string",
							Description: "The city and state, e.g. San Francisco, CA",
						},
						"unit": {
							Type: "string",
							Enum: []string{"celsius", "fahrenheit"},
						},
					},
				},
			},
		},
	}, nil)

	if err != nil {
		panic(err)
	}

	// now we can extract out the function name and the arguments, which will match the
	// JSON schema we specified in the FunctionDefinition.Parameters.Properties, above.
	funcCall := resp.ChatCompletions.Choices[0].Message.FunctionCall

	// NOTE: this location type just matches the JSON schema above.
	type location struct {
		Location string `json:"location"`
		Unit     string `json:"unit"`
	}

	// deserialize the function arguments structure
	var funcParams *location
	if err := json.Unmarshal([]byte(*funcCall.Arguments), &funcParams); err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Args: Location: %s, Unit: %s\n", *funcCall.Name, funcParams.Location, funcParams.Unit)
}

var exampleClient *azopenai.Client
