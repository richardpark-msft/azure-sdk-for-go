//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai_test

// func TestImageGeneration_AzureOpenAI(t *testing.T) {
// 	if recording.GetRecordMode() == recording.PlaybackMode {
// 		t.Skipf("Ignoring poller-based test")
// 	}

// 	cred, err := azopenai.NewKeyCredential(apiKey)
// 	require.NoError(t, err)

// 	client, err := azopenai.NewClientWithKeyCredential(endpoint, cred, "", newClientOptionsForTest(t))
// 	require.NoError(t, err)

// 	testImageGeneration(t, client, azopenai.ImageGenerationResponseFormatURL)
// }

// func TestImageGeneration_OpenAI(t *testing.T) {
// 	client := newOpenAIClientForTest(t)
// 	testImageGeneration(t, client, azopenai.ImageGenerationResponseFormatURL)
// }

// func TestImageGeneration_OpenAI_Base64(t *testing.T) {
// 	client := newOpenAIClientForTest(t)
// 	testImageGeneration(t, client, azopenai.ImageGenerationResponseFormatB64JSON)
// }

// func testImageGeneration(t *testing.T, client *azopenai.Client, responseFormat azopenai.ImageGenerationResponseFormat) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()

// 	resp, err := client.CreateImage(ctx, azopenai.ImageGenerationOptions{
// 		Prompt:         to.Ptr("a cat"),
// 		Size:           to.Ptr(azopenai.ImageSize256x256),
// 		ResponseFormat: &responseFormat,
// 	}, nil)
// 	require.NoError(t, err)

// 	if recording.GetRecordMode() == recording.LiveMode {
// 		switch responseFormat {
// 		case azopenai.ImageGenerationResponseFormatURL:
// 			imageLocation := resp.Data[0].Result.(azopenai.ImageLocation)
// 			headResp, err := http.DefaultClient.Head(*imageLocation.URL)
// 			require.NoError(t, err)
// 			require.Equal(t, http.StatusOK, headResp.StatusCode)
// 		case azopenai.ImageGenerationResponseFormatB64JSON:
// 			imagePayload := resp.Data[0].Result.(azopenai.ImagePayload)
// 			bytes, err := base64.StdEncoding.DecodeString(*imagePayload.B64JSON)
// 			require.NoError(t, err)
// 			require.NotEmpty(t, bytes)
// 		}
// 	}
// }
