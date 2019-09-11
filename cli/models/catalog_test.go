/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package models

import (
	"bytes"
	"encoding/json"
	"github.com/matryer/is"
	"testing"
)

func TestDisplayCatalogEntity(t *testing.T) {
	assert := is.New(t)

	var summary CatalogEntitySummary
	je := json.Unmarshal([]byte(testCaseJson), &summary)
	if je != nil {
		t.Fatal("failed to unmarshal test object", je)
	}

	var bb bytes.Buffer
	err := displayAsJson(&bb, summary, "$.name", false)
	if err != nil {
		t.Fatal("display error", err)
	}
	assert.Equal(bb.String(), `"Test Case"`)

	bb.Reset()
	err = displayAsJson(&bb, summary, "$.effectors[*].name", false)
	if err != nil {
		t.Fatal("display error", err)
	}
	assert.Equal(bb.String(), `"restart" "start" "stop"`)
}
