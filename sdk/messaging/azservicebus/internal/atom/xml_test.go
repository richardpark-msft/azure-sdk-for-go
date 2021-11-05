// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package atom

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeserializeCorrelationFilter(t *testing.T) {
	bytes, err := ioutil.ReadFile("testdata/correlation_rule.xml")
	require.NoError(t, err)

	var entry *RuleEnvelope
	err = xml.Unmarshal(bytes, &entry)
	require.NoError(t, err)
}
