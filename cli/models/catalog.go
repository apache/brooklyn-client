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

type IdentityDetails struct {
	Id           string                 `json:"id"`
	Name         string                 `json:"name"`
	SymbolicName string                 `json:"symbolicName"`
	Version      string                 `json:"version"`
	Description  string                 `json:"description"`
}

type CatalogItemSummary struct {
	IdentityDetails
	JavaType     string                 `json:"javaType"`
	PlanYaml     string                 `json:"planYaml"`
	Deprecated   bool                   `json:"deprecated"`
	Links        map[string]interface{} `json:"links"`
	Type         string                 `json:"type"`
}

type CatalogPolicySummary struct {
	IdentityDetails
	javaType     string         `json:"javaType"`
	planYaml     string         `json:"planYaml"`
	iconUrl      string         `json:"iconUrl"`
	deprecated   bool           `json:"deprecated"`
	links        map[string]URI `json:"links"`
}


type CatalogLocationSummary struct {
	CatalogItemSummary
	IconUrl     string                 `json:"iconUrl"`
}

type CatalogEntitySummary struct {
	CatalogItemSummary
	Config       []ConfigSummary        `json:"config"`
	Effectors    []EffectorSummary      `json:"effectors"`
	Sensors      []SensorSummary        `json:"sensors"`
}
