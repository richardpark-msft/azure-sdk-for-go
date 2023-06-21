pwsh ../../../eng/common/scripts/TypeSpec-Project-Sync.ps1 .
Push-Location ./TempTypeSpecFiles/OpenAI.Inference

$text = @"
parameters:
  "service-dir":
    default: "sdk/openai"
  "dependencies":
    default: ""
emit:
 - "@azure-tools/typespec-autorest"
options:
  "@azure-tools/typespec-autorest":
    emitter-output-dir: "{project-root}/../"
    output-file: "./generated.json"
    azure-resource-provider-folder: "data-plane"
    examples-directory: examples  
"@

Out-File -FilePath ./tspconfig.yaml -InputObject $text

$packageJson = @"
{
  "name": "OpenAI.Inference",
  "version": "0.1.0",
  "type": "module",
  "dependencies": {
    "@azure-tools/typespec-autorest": "^0.31.0",
    "@typespec/compiler": "latest"
  },
  "private": true
}
"@

Out-File -FilePath ./package.json -InputObject $packageJson
tsp install
tsp compile .
Pop-Location

