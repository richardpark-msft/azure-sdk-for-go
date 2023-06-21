//go:build go1.18
// +build go1.18

//go:generate pwsh build-typespec.ps1
//go:generate autorest ./autorest.azure.md
//go:generate autorest ./autorest.openai.md
//go:generate go mod tidy
//go:generate goimports -w .

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azopenai
