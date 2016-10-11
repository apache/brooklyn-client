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
package entities

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"net/url"
)

//WIP
func GetTask(network *net.Network, application, entity, task string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities/%s", application, entity, task)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

//WIP
func GetIcon(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/icon", application, entity)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func Children(network *net.Network, application, entity string) ([]models.EntitySummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	var entityList []models.EntitySummary
	body, err := network.SendGetRequest(urlStr)
	if err != nil {
		return entityList, err
	}

	err = json.Unmarshal(body, &entityList)
	return entityList, err
}

func AddChildren(network *net.Network, application, entity, resource string) (models.TaskSummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	var response models.TaskSummary
	body, err := network.SendPostResourceRequest(urlStr, resource, "application/json")
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	return response, err
}

//WIP
func GetLocations(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/locations", application, entity)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func Spec(network *net.Network, application, entity string) (string, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/spec", application, entity)
	body, err := network.SendGetRequest(urlStr)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

//WIP
func GetDescendants(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants", application, entity)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

//WIP
func GetDescendantsSensor(network *net.Network, application, entity, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants/sensor/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func GetActivities(network *net.Network, application, entity string) ([]models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities", application, entity)
	var activityList []models.TaskSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return activityList, err
	}

	err = json.Unmarshal(body, &activityList)
	return activityList, err
}

//WIP
func GetTags(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/tags", application, entity)
	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

//WIP
func Expunge(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/expunge", application, entity)
	body, err := network.SendEmptyPostRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

//WIP
func GetEntity(network *net.Network, application, entity string) (models.EntitySummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s", application, entity)
	summary := models.EntitySummary{}
	body, err := network.SendGetRequest(url)
	if err != nil {
		return summary, err
	}

	err = json.Unmarshal(body, &summary)
	return summary, err
}

func EntityList(network *net.Network, application string) ([]models.EntitySummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities", application)
	var entityList []models.EntitySummary
	body, err := network.SendGetRequest(urlStr)
	if err != nil {
		return entityList, err
	}

	err = json.Unmarshal(body, &entityList)
	return entityList, err
}

func Rename(network *net.Network, application, entity, newName string) (string, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/name?name=%s", application, entity, url.QueryEscape(newName))
	body, err := network.SendEmptyPostRequest(urlStr)
	if nil != err {
		return "", err
	}
	return string(body), nil
}
