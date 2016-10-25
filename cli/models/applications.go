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

type MemberSpec struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
}

type Tree struct {
	Id            string   `json:"id"`
	ParentId      string   `json:"parentId"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	CatalogItemId string   `json:"catalogItemId"`
	Children      []Tree   `json:"children"`
	GroupIds      []string `json:"groupIds"`
	Members       []MemberSpec `json:"members"`
}

type TaskSummary struct {
	SubmitTimeUtc     int64                              `json:"submitTimeUtc"`
	EndTimeUtc        int64                              `json:"endTimeUtc"`
	IsCancelled       bool                               `json:"isCancelled"`
	CurrentStatus     string                             `json:"currentStatus"`
	BlockingTask      LinkTaskWithMetadata               `json:"blockingTask"`
	DisplayName       string                             `json:"displayName"`
	Streams           map[string]LinkStreamsWithMetadata `json:"streams"`
	Description       string                             `json:"description"`
	EntityId          string                             `json:"entityId"`
	EntityDisplayName string                             `json:"entityDisplayName"`
	Error             bool                               `json:"error"`
	SubmittedByTask   LinkTaskWithMetadata               `json:"submittedByTask"`
	Result            interface{}                        `json:"result"`
	IsError           bool                               `json:"isError"`
	DetailedStatus    string                             `json:"detailedStatus"`
	Children          []LinkTaskWithMetadata             `json:"children"`
	BlockingDetails   string                             `json:"blockingDetails"`
	Cancelled         bool                               `json:"cancelled"`
	Links             map[string]URI                     `json:"links"`
	Id                string                             `json:"id"`
	StartTimeUtc      int64                              `json:"startTimeUtc"`
}

type ApplicationSummary struct {
	Links  map[string]URI  `json:"links"`
	Id     string          `json:"id"`
	Spec   ApplicationSpec `json:"spec"`
	Status Status          `json:"status"`
}

type ApplicationSpec struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Locations []string `json:"locations"`
}

type Status string

type LinkWithMetadata struct {
}

type LinkStreamsWithMetadata struct {
	Link     string             `json:"link"`
	Metadata LinkStreamMetadata `json:"metadata"`
}

type LinkStreamMetadata struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	SizeText string `json:"sizeText"`
}

type LinkTaskWithMetadata struct {
	Link     string           `json:"link"`
	Metadata LinkTaskMetadata `json:"metadata"`
}

type LinkTaskMetadata struct {
	Id                string `json:"id"`
	TaskName          string `json:"taskName"`
	EntityId          string `json:"entityId"`
	EntityDisplayName string `json:"entityDisplayName"`
}

type URI string
