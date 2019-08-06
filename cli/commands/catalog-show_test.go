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
package commands

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/apache/brooklyn-client/cli/models"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func testCatalogEntitySummary() models.CatalogEntitySummary {
	summary := models.CatalogEntitySummary{
		CatalogItemSummary: models.CatalogItemSummary{
			IdentityDetails: models.IdentityDetails{
				Id:      "server:1.0.0-SNAPSHOT",
				Name:    "Server",
				Version: "1.0.0-SNAPSHOT",
			},
			JavaType: "org.apache.brooklyn.entity.software.base.EmptySoftwareProcess",
			Config: []models.ConfigSummary{
				{
					Name:   "testConfigKey",
					Pinned: true,
				},
			},
		},
		IconUrl: "/v1/catalog/icon/server/1.0.0-SNAPSHOT",
	}
	return summary
}

func assertItemSummaryFields(t *testing.T, expected models.CatalogItemSummary, actual models.CatalogItemSummary) {
	assert.Equal(t, expected.Id, actual.Id, "id")
	assert.Equal(t, expected.Name, actual.Name, "name")
	assert.Equal(t, expected.Version, actual.Version, "version")
	assert.Equal(t, expected.JavaType, actual.JavaType, "javaType")
	assert.Equal(t, expected.Config[0].Name, actual.Config[0].Name, "Config[0].Name")
	assert.Equal(t, expected.Config[0].Pinned, actual.Config[0].Pinned, "Config[0].Pinned")
}

func testDisplay(t *testing.T, fn func(c *cli.Context) error, testPath string) string {

	testApp := cli.NewApp()
	testApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "json, j",
			Usage: "Render value as json with json path selector as described at https://github.com/NodePrime/jsonpath. (Experimental, not supported on all commands at present) ",
		},
	}
	var out string
	var err error
	testApp.Action = func(c *cli.Context) {
		out, err = divertStdoutToString(func() error {
			return fn(c)
		})
	}

	err = testApp.Run([]string{"", "--json", testPath})
	if err != nil {
		t.Fatalf("Error from Run: %s\n", err)
	}

	return out
}

func TestCatalogItemSummaryDisplay(t *testing.T) {
	expected := testCatalogEntitySummary()
	displayOutput := testDisplay(t, func(c *cli.Context) error {
		return expected.Display(c)
	}, "$")
	var actual models.CatalogItemSummary
	actualBytes := []byte(displayOutput)
	unmarshalErr := json.Unmarshal(actualBytes, &actual)
	assert.Nil(t, unmarshalErr, "result is not JSON: %s", actualBytes)

	assertItemSummaryFields(t, expected.CatalogItemSummary, actual)
}

func TestCatalogEntitySummaryDisplay(t *testing.T) {

	expected := testCatalogEntitySummary()
	displayOutput := testDisplay(t, func(c *cli.Context) error {
		return expected.Display(c)
	}, "$")

	var actual models.CatalogEntitySummary
	actualBytes := []byte(displayOutput)
	unmarshalErr := json.Unmarshal(actualBytes, &actual)
	assert.Nil(t, unmarshalErr, "result is not JSON: %s", actualBytes)

	assertItemSummaryFields(t, expected.CatalogItemSummary, actual.CatalogItemSummary)
	assert.Equal(t, expected.IconUrl, actual.IconUrl, "iconUrl")
}

func TestPaths(t *testing.T) {
	type pathTest struct {
		testPath string
		expected string
	}

	testObject := testCatalogEntitySummary()

	tests := []pathTest{
		{"$.name+", testObject.Name},
		{"$.id+", testObject.Id},
	}

	for _, test := range tests {
		actual := testDisplay(t, func(c *cli.Context) error {
			return testObject.Display(c)
		}, test.testPath)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("path %s", test.testPath))
	}
}
