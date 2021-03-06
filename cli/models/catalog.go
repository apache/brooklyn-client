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
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli/v2"
	"k8s.io/client-go/util/jsonpath"
)

type IdentityDetails struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	SymbolicName string `json:"symbolicName"`
	Version      string `json:"version"`
	Description  string `json:"description"`
}

type CatalogItemSummary struct {
	IdentityDetails
	JavaType   string                 `json:"javaType"`
	PlanYaml   string                 `json:"planYaml"`
	Deprecated bool                   `json:"deprecated"`
	Config     []ConfigSummary        `json:"config"`
	Tags       []interface{}          `json:"tags"`
	Links      map[string]interface{} `json:"links"`
	Type       string                 `json:"type"`
}

type CatalogEntitySummary struct {
	CatalogItemSummary
	IconUrl   string            `json:"iconUrl"`
	Effectors []EffectorSummary `json:"effectors"`
	Sensors   []SensorSummary   `json:"sensors"`
}

type CatalogBundleAddResult struct {
	Message string                        `json:"message"`
	Bundle  string                        `json:"bundle"`
	Code    string                        `json:"code"`
	Types   map[string]CatalogItemSummary `json:"types"`
}

func createTableWithIdentityDetails(item IdentityDetails) terminal.Table {
	table := terminal.NewTable([]string{"Id:", item.Id})
	table.Add("Version:", item.Version)
	table.Add("Name:", item.Name)
	table.Add("Symbolic Name:", item.SymbolicName)
	table.Add("Description:", item.Description)
	return table
}

func (summary *CatalogItemSummary) Display(c *cli.Context) error {

	if jsonFlag := c.String("json"); jsonFlag != "" {
		raw := c.Bool("raw-output")
		err := displayAsJson(os.Stdout, summary, jsonFlag, raw)
		if err != nil {
			return fmt.Errorf("display error: %s", err)
		}
	} else {
		summary.displayAsTable()
	}
	return nil
}

func (summary *CatalogEntitySummary) Display(c *cli.Context) error {

	if jsonFlag := c.String("json"); jsonFlag != "" {
		raw := c.Bool("raw-output")
		err := displayAsJson(os.Stdout, summary, jsonFlag, raw)
		if err != nil {
			return fmt.Errorf("display error: %s\n", err)
		}
	} else {
		summary.displayAsTable()
	}
	return nil
}

func (summary *CatalogItemSummary) displayAsTable() {

	table := createTableWithIdentityDetails(summary.IdentityDetails)
	if summary.Deprecated {
		table.Add("Deprecated:", "true")
	}
	table.Add("Java Type:", summary.JavaType)
	table.Print()
}

func (summary *CatalogEntitySummary) displayAsTable() {

	table := createTableWithIdentityDetails(summary.IdentityDetails)
	if summary.Deprecated {
		table.Add("Deprecated:", "true")
	}
	table.Add("Java Type:", summary.JavaType)
	table.Add("Icon URL:", summary.IconUrl)

	for c, conf := range summary.Config {
		if c == 0 {
			table.Add("", "") // helps distinguish entries from one another
			table.Add("Config:", "")
		}
		table.Add("Name:", conf.Name)
		table.Add("Type:", conf.Type)
		table.Add("Description:", conf.Description)
		table.Add("Default Value:", fmt.Sprintf("%v", conf.DefaultValue))
		table.Add("Reconfigurable:", strconv.FormatBool(conf.Reconfigurable))
		table.Add("Label:", conf.Label)
		table.Add("Priority:", strconv.FormatFloat(conf.Priority, 'f', -1, 64))
		table.Add("Pinned:", strconv.FormatBool(conf.Pinned))

		if len(conf.PossibleValues) > 0 {
			var values bytes.Buffer
			for i, pv := range conf.PossibleValues {
				if i > 0 {
					values.WriteString(", ")
				}
				values.WriteString(pv["value"])
				if pv["value"] != pv["description"] {
					values.WriteString(" (" + pv["description"] + ")")
				}
			}
			table.Add("Possible Values:", values.String())
		}
		table.Add("", "") // helps distinguish entries from one another
	}

	for t, tag := range summary.Tags {
		if t == 0 {
			table.Add("", "") // helps distinguish entries from one another
			table.Add("Tags:", "")
		}
		if asJson, erj := json.Marshal(tag); nil == erj {
			table.Add("tag:", string(asJson))
		} else {
			table.Add("tag:", fmt.Sprintf("%v", tag))
		}
	}

	table.Print()
}

func resultsBackToJson(wr io.Writer, values []reflect.Value, raw bool) error {
	for i, r := range values {
		object := r.Interface()

		jsonText, err := json.Marshal(object)
		if err != nil {
			return fmt.Errorf("error converting object to JSON: %s", err)
		}

		if r.Kind() == reflect.String && raw {
			stringWithQuotes := string(jsonText)
			trimLeft := strings.TrimPrefix(stringWithQuotes, `"`)
			withoutQuotes := strings.TrimSuffix(trimLeft, `"`)
			jsonText = []byte(withoutQuotes)
		}

		if i != len(values)-1 {
			jsonText = append(jsonText, ' ')
		}
		if _, err := wr.Write(jsonText); err != nil {
			return err
		}
	}
	return nil
}

func displayAsJson(w io.Writer, v interface{}, displayPath string, raw bool) error {
	j := jsonpath.New("displayer")
	j.AllowMissingKeys(true)

	// wrap the path with k8s.io's conventional {} braces for convenience
	err := j.Parse(fmt.Sprintf("{%s}", displayPath))
	if err != nil {
		return fmt.Errorf("could not parse JSONPath expression (%s)", err)
	}

	allResults, err := j.FindResults(v)
	if err != nil {
		return fmt.Errorf("error evaluating JSONPath expression: %s", err)
	}
	for ix := range allResults {
		if err = resultsBackToJson(w, allResults[ix], raw); err != nil {
			return fmt.Errorf("display error: %s", err)
		}
	}
	return nil
}
