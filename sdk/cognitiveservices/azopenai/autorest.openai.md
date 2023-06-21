``` yaml
input-file:
- https://raw.githubusercontent.com/openai/openai-openapi/master/openapi.yaml
output-folder: ../azopenai/internal/openai
clear-output-folder: true
# module: github.com/Azure/azure-sdk-for-go/sdk/cognitiveservices/azopenai/openai
license-header: MICROSOFT_MIT_NO_VERSION
openapi-type: data-plane
go: true
use: "@autorest/go@4.0.0-preview.50"
directive:
  - from:
      - api_client.go
      - client.go
      - models.go
      - response_types.go
    where: $
    transform: return $.replace(/OpenAIAPI(\w+)((?:Options|Response|Client))/g, "$1$2");
  - from:
      - api_client.go
    where: $
    transform: > 
      return $.replace(/"N": N,/g, '"N": options.N,')
        .replace(/"Size": Size,/g, '"Size": options.Size,')
        .replace(/"Mask": Mask,/g, '"Mask": options.Mask,')
        .replace(/"Prompt": Prompt,/g, '"Prompt": options.Prompt,')
        .replace(/"Temperature": Temperature,/g, '"Temperature": options.Temperature,')
        .replace(/"Language": Language,/g, '"Language": options.Language,')
        .replace(/"ResponseFormat": ResponseFormat,/g, '"ResponseFormat": options.ResponseFormat,')
        .replace(/"User": User,/g, '"User": options.User,')
```

