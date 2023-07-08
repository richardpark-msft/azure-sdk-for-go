//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// azureCoreFoundationsError - The error object.
type azureCoreFoundationsError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []*azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// azureCoreFoundationsErrorInnererror - An object containing more specific information than the current object about the
// error.
type azureCoreFoundationsErrorInnererror struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// azureCoreFoundationsErrorResponse - A response containing error details.
type azureCoreFoundationsErrorResponse struct {
	// REQUIRED; The error object.
	Error *azureCoreFoundationsErrorResponseError
}

// azureCoreFoundationsErrorResponseError - The error object.
type azureCoreFoundationsErrorResponseError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []*azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// azureCoreFoundationsInnerError - An object containing more specific information about the error. As per Microsoft One API
// guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type azureCoreFoundationsInnerError struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// azureCoreFoundationsInnerErrorInnererror - Inner error.
type azureCoreFoundationsInnerErrorInnererror struct {
	// One of a server-defined set of error codes.
	Code *string

	// Inner error.
	Innererror *azureCoreFoundationsInnerErrorInnererror
}

// ChatChoice - The representation of a single prompt completion as part of an overall chat completions request. Generally,
// n choices are generated per provided prompt with a default value of 1. Token limits and
// other settings may limit the number of choices generated.
type ChatChoice struct {
	// REQUIRED; The reason that this chat completions choice completed its generated.
	FinishReason *CompletionsFinishReason

	// REQUIRED; The ordered index associated with this chat completions choice.
	Index *int32

	// The delta message content for a streaming response.
	Delta *ChatChoiceDelta

	// The chat message for a given chat completions prompt.
	Message *ChatChoiceMessage
}

// ChatChoiceDelta - The delta message content for a streaming response.
type ChatChoiceDelta struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatChoiceMessage - The chat message for a given chat completions prompt.
type ChatChoiceMessage struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatCompletions - Representation of the response data from a chat completions request. Completions support a wide variety
// of tasks and generate text that continues from or "completes" provided prompt data.
type ChatCompletions struct {
	// REQUIRED; The collection of completions choices associated with this completions response. Generally, n choices are generated
	// per provided prompt with a default value of 1. Token limits and other settings may
	// limit the number of choices generated.
	Choices []*ChatChoice

	// REQUIRED; The first timestamp associated with generation activity for this completions response, represented as seconds
	// since the beginning of the Unix epoch of 00:00 on 1 Jan 1970.
	Created *int32

	// REQUIRED; A unique identifier associated with this chat completions response.
	ID *string

	// REQUIRED; Usage information for tokens processed and generated as part of this completions operation.
	Usage *CompletionsUsage
}

// ChatCompletionsOptions - The configuration information for a chat completions request. Completions support a wide variety
// of tasks and generate text that continues from or "completes" provided prompt data.
type ChatCompletionsOptions struct {
	// REQUIRED; The collection of context messages associated with this chat completions request. Typical usage begins with a
	// chat message for the System role that provides instructions for the behavior of the
	// assistant, followed by alternating messages between the User and Assistant roles.
	Messages []*ChatMessage

	// A value that influences the probability of generated tokens appearing based on their cumulative frequency in generated
	// text. Positive values will make tokens less likely to appear as their frequency
	// increases and decrease the likelihood of the model repeating the same statements verbatim.
	FrequencyPenalty *float32

	// Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the
	// end-user. "auto" means the model can pick between an end-user or calling a
	// function. Specifying a particular function via {"name": "my_function"} forces the model to call that function. "none" is
	// the default when no functions are present. "auto" is the default if functions
	// are present.
	FunctionCall ChatCompletionsOptionsFunctionCall

	// A list of functions the model may generate JSON inputs for.
	Functions []*FunctionDefinition

	// A map between GPT token IDs and bias scores that influences the probability of specific tokens appearing in a completions
	// response. Token IDs are computed via external tokenizer tools, while bias
	// scores reside in the range of -100 to 100 with minimum and maximum values corresponding to a full ban or exclusive selection
	// of a token, respectively. The exact behavior of a given bias score varies
	// by model.
	LogitBias map[string]*int32

	// The maximum number of tokens to generate.
	MaxTokens *int32

	// The model name to provide as part of this completions request. Not applicable to Azure OpenAI, where deployment information
	// should be included in the Azure resource URI that's connected to.
	Model *string

	// The number of chat completions choices that should be generated for a chat completions response. Because this setting can
	// generate many completions, it may quickly consume your token quota. Use
	// carefully and ensure reasonable settings for max_tokens and stop.
	N *int32

	// A value that influences the probability of generated tokens appearing based on their existing presence in generated text.
	// Positive values will make tokens less likely to appear when they already exist
	// and increase the model's likelihood to output new topics.
	PresencePenalty *float32

	// A collection of textual sequences that will end completions generation.
	Stop []*string

	// The sampling temperature to use that controls the apparent creativity of generated completions. Higher values will make
	// output more random while lower values will make results more focused and
	// deterministic. It is not recommended to modify temperature and top_p for the same completions request as the interaction
	// of these two settings is difficult to predict.
	Temperature *float32

	// An alternative to sampling with temperature called nucleus sampling. This value causes the model to consider the results
	// of tokens with the provided probability mass. As an example, a value of 0.15
	// will cause only the tokens comprising the top 15% of probability mass to be considered. It is not recommended to modify
	// temperature and top_p for the same completions request as the interaction of
	// these two settings is difficult to predict.
	TopP *float32

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// ChatMessage - A single, role-attributed message within a chat completion interaction.
type ChatMessage struct {
	// REQUIRED; The role associated with this message payload.
	Role *ChatRole

	// The text associated with this message payload.
	Content *string

	// The name and arguments of a function that should be called, as generated by the model.
	FunctionCall *ChatMessageFunctionCall

	// The name of the author of this message. name is required if role is function, and it should be the name of the function
	// whose response is in the content. May contain a-z, A-Z, 0-9, and underscores,
	// with a maximum length of 64 characters.
	Name *string
}

// ChatMessageFunctionCall - The name and arguments of a function that should be called, as generated by the model.
type ChatMessageFunctionCall struct {
	// REQUIRED; The arguments to call the function with, as generated by the model in JSON format. Note that the model does not
	// always generate valid JSON, and may hallucinate parameters not defined by your function
	// schema. Validate the arguments in your code before calling your function.
	Arguments *string

	// REQUIRED; The name of the function to call.
	Name *string
}

// Choice - The representation of a single prompt completion as part of an overall completions request. Generally, n choices
// are generated per provided prompt with a default value of 1. Token limits and other
// settings may limit the number of choices generated.
type Choice struct {
	// REQUIRED; Reason for finishing
	FinishReason *CompletionsFinishReason

	// REQUIRED; The ordered index associated with this completions choice.
	Index *int32

	// REQUIRED; The log probabilities model for tokens associated with this completions choice.
	Logprobs *ChoiceLogprobs

	// REQUIRED; The generated text for a given completions prompt.
	Text *string
}

// ChoiceLogprobs - The log probabilities model for tokens associated with this completions choice.
type ChoiceLogprobs struct {
	// REQUIRED; The text offsets associated with tokens in this completions data.
	TextOffset []*int32

	// REQUIRED; A collection of log probability values for the tokens in this completions data.
	TokenLogprobs []*float32

	// REQUIRED; The textual forms of tokens evaluated in this probability model.
	Tokens []*string

	// REQUIRED; A mapping of tokens to maximum log probability values in this completions data.
	TopLogprobs []any
}

// Completions - Representation of the response data from a completions request. Completions support a wide variety of tasks
// and generate text that continues from or "completes" provided prompt data.
type Completions struct {
	// REQUIRED; The collection of completions choices associated with this completions response. Generally, n choices are generated
	// per provided prompt with a default value of 1. Token limits and other settings may
	// limit the number of choices generated.
	Choices []*Choice

	// REQUIRED; The first timestamp associated with generation activity for this completions response, represented as seconds
	// since the beginning of the Unix epoch of 00:00 on 1 Jan 1970.
	Created *int32

	// REQUIRED; A unique identifier associated with this completions response.
	ID *string

	// REQUIRED; Usage information for tokens processed and generated as part of this completions operation.
	Usage *CompletionsUsage
}

// CompletionsLogProbabilityModel - Representation of a log probabilities model for a completions generation.
type CompletionsLogProbabilityModel struct {
	// REQUIRED; The text offsets associated with tokens in this completions data.
	TextOffset []*int32

	// REQUIRED; A collection of log probability values for the tokens in this completions data.
	TokenLogprobs []*float32

	// REQUIRED; The textual forms of tokens evaluated in this probability model.
	Tokens []*string

	// REQUIRED; A mapping of tokens to maximum log probability values in this completions data.
	TopLogprobs []any
}

// CompletionsOptions - The configuration information for a completions request. Completions support a wide variety of tasks
// and generate text that continues from or "completes" provided prompt data.
type CompletionsOptions struct {
	// REQUIRED; The prompts to generate completions from.
	Prompt []*string

	// A value that controls how many completions will be internally generated prior to response formulation. When used together
	// with n, bestof controls the number of candidate completions and must be
	// greater than n. Because this setting can generate many completions, it may quickly consume your token quota. Use carefully
	// and ensure reasonable settings for maxtokens and stop.
	BestOf *int32

	// A value specifying whether completions responses should include input prompts as prefixes to their generated output.
	Echo *bool

	// A value that influences the probability of generated tokens appearing based on their cumulative frequency in generated
	// text. Positive values will make tokens less likely to appear as their frequency
	// increases and decrease the likelihood of the model repeating the same statements verbatim.
	FrequencyPenalty *float32

	// A map between GPT token IDs and bias scores that influences the probability of specific tokens appearing in a completions
	// response. Token IDs are computed via external tokenizer tools, while bias
	// scores reside in the range of -100 to 100 with minimum and maximum values corresponding to a full ban or exclusive selection
	// of a token, respectively. The exact behavior of a given bias score varies
	// by model.
	LogitBias map[string]*int32

	// A value that controls the emission of log probabilities for the provided number of most likely tokens within a completions
	// response.
	Logprobs *int32

	// The maximum number of tokens to generate.
	MaxTokens *int32

	// The model name to provide as part of this completions request. Not applicable to Azure OpenAI, where deployment information
	// should be included in the Azure resource URI that's connected to.
	Model *string

	// The number of completions choices that should be generated per provided prompt as part of an overall completions response.
	// Because this setting can generate many completions, it may quickly consume
	// your token quota. Use carefully and ensure reasonable settings for max_tokens and stop.
	N *int32

	// A value that influences the probability of generated tokens appearing based on their existing presence in generated text.
	// Positive values will make tokens less likely to appear when they already exist
	// and increase the model's likelihood to output new topics.
	PresencePenalty *float32

	// A collection of textual sequences that will end completions generation.
	Stop []*string

	// The sampling temperature to use that controls the apparent creativity of generated completions. Higher values will make
	// output more random while lower values will make results more focused and
	// deterministic. It is not recommended to modify temperature and top_p for the same completions request as the interaction
	// of these two settings is difficult to predict.
	Temperature *float32

	// An alternative to sampling with temperature called nucleus sampling. This value causes the model to consider the results
	// of tokens with the provided probability mass. As an example, a value of 0.15
	// will cause only the tokens comprising the top 15% of probability mass to be considered. It is not recommended to modify
	// temperature and top_p for the same completions request as the interaction of
	// these two settings is difficult to predict.
	TopP *float32

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// CompletionsUsage - Representation of the token counts processed for a completions request. Counts consider all tokens across
// prompts, choices, choice alternates, best_of generations, and other consumers.
type CompletionsUsage struct {
	// REQUIRED; The number of tokens generated across all completions emissions.
	CompletionTokens *int32

	// REQUIRED; The number of tokens in the provided prompts for the completions request.
	PromptTokens *int32

	// REQUIRED; The total number of tokens processed for the completions request and response.
	TotalTokens *int32
}

// Deployment - A specific deployment
type Deployment struct {
	// READ-ONLY; deployment id of the deployed model
	DeploymentID *string
}

// EmbeddingItem - Representation of a single embeddings relatedness comparison.
type EmbeddingItem struct {
	// REQUIRED; List of embeddings value for the input prompt. These represent a measurement of the vector-based relatedness
	// of the provided input.
	Embedding []*float32

	// REQUIRED; Index of the prompt to which the EmbeddingItem corresponds.
	Index *int32
}

// Embeddings - Representation of the response data from an embeddings request. Embeddings measure the relatedness of text
// strings and are commonly used for search, clustering, recommendations, and other similar
// scenarios.
type Embeddings struct {
	// REQUIRED; Embedding values for the prompts submitted in the request.
	Data []*EmbeddingItem

	// REQUIRED; Usage counts for tokens input using the embeddings API.
	Usage *EmbeddingsUsage
}

// EmbeddingsOptions - The configuration information for an embeddings request. Embeddings measure the relatedness of text
// strings and are commonly used for search, clustering, recommendations, and other similar scenarios.
type EmbeddingsOptions struct {
	// REQUIRED; Input texts to get embeddings for, encoded as a an array of strings. Each input must not exceed 2048 tokens in
	// length.
	// Unless you are embedding code, we suggest replacing newlines (\n) in your input with a single space, as we have observed
	// inferior results when newlines are present.
	Input []*string

	// The model name to provide as part of this embeddings request. Not applicable to Azure OpenAI, where deployment information
	// should be included in the Azure resource URI that's connected to.
	Model *string

	// An identifier for the caller or end user of the operation. This may be used for tracking or rate-limiting purposes.
	User *string
}

// EmbeddingsUsage - Usage counts for tokens input using the embeddings API.
type EmbeddingsUsage struct {
	// REQUIRED; Number of tokens sent in the original request.
	PromptTokens *int32

	// REQUIRED; Total number of tokens transacted in this request/response.
	TotalTokens *int32
}

// EmbeddingsUsageAutoGenerated - Measurement of the amount of tokens used in this request and response.
type EmbeddingsUsageAutoGenerated struct {
	// REQUIRED; Number of tokens sent in the original request.
	PromptTokens *int32

	// REQUIRED; Total number of tokens transacted in this request/response.
	TotalTokens *int32
}

// FunctionCall - The name and arguments of a function that should be called, as generated by the model.
type FunctionCall struct {
	// REQUIRED; The arguments to call the function with, as generated by the model in JSON format. Note that the model does not
	// always generate valid JSON, and may hallucinate parameters not defined by your function
	// schema. Validate the arguments in your code before calling your function.
	Arguments *string

	// REQUIRED; The name of the function to call.
	Name *string
}

// FunctionDefinition - The definition of a function that can be called by the model. See the guide [/docs/guides/gpt/function-calling]
// for examples.
type FunctionDefinition struct {
	// REQUIRED; The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum
	// length of 64.
	Name *string

	// The description of what the function does.
	Description *string

	// The parameters the functions accepts, described as a JSON Schema object. See the guide [/docs/guides/gpt/function-calling]
	// for examples, and the JSON Schema reference
	// [https://json-schema.org/understanding-json-schema/]for documentation about the format.
	Parameters any
}

// FunctionName - Specify the name of the only function to call.
type FunctionName struct {
	// REQUIRED; The name of the function to call.
	Name *string
}

// ImageGenerationOptions - Represents the request data used to generate images.
type ImageGenerationOptions struct {
	// REQUIRED; A description of the desired images.
	Prompt *string

	// The number of images to generate (defaults to 1).
	N *int32

	// The desired size of the generated images. Must be one of 256x256, 512x512, or 1024x1024 (defaults to 1024x1024).
	Size *ImageSize

	// A unique identifier representing your end-user, which can help to monitor and detect abuse.
	User *string
}

// ImageLocation - The image url if successful, and an error otherwise.
type ImageLocation struct {
	// The error if the operation failed.
	Error *ImageLocationError

	// The URL that provides temporary access to download the generated image.
	URL *string
}

// ImageLocationError - The error if the operation failed.
type ImageLocationError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []*azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// ImageOperationResponse - The result of the operation if the operation succeeded.
type ImageOperationResponse struct {
	// REQUIRED; A timestamp when this job or item was created (in unix epochs).
	Created *int64

	// REQUIRED; The ID of the operation.
	ID *string

	// REQUIRED; The status of the operation
	Status *State

	// The error if the operation failed.
	Error *ImageOperationResponseError

	// A timestamp when this operation and its associated images expire and will be deleted (in unix epochs).
	Expires *int64

	// The result of the operation if the operation succeeded.
	Result *ImageResponse
}

// ImageOperationResponseError - The error if the operation failed.
type ImageOperationResponseError struct {
	// REQUIRED; One of a server-defined set of error codes.
	Code *string

	// REQUIRED; A human-readable representation of the error.
	Message *string

	// An array of details about specific errors that led to this reported error.
	Details []*azureCoreFoundationsError

	// An object containing more specific information than the current object about the error.
	Innererror *azureCoreFoundationsErrorInnererror

	// The target of the error.
	Target *string
}

// ImageOperationStatus - Provides status details for long running operations.
type ImageOperationStatus struct {
	// REQUIRED; The unique ID of the operation.
	ID *string

	// REQUIRED; The status of the operation
	Status *State
}

// ImageResponse - The result of the operation if the operation succeeded.
type ImageResponse struct {
	// REQUIRED; A timestamp when this job or item was created (in unix epochs).
	Created *int64

	// REQUIRED; The images generated by the operator.
	Data []*ImageLocation
}
