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
package application

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

//WIP
func Fetch(network *net.Network) (string, error) {
	url := "/v1/applications/fetch"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	// TODO return model
	return string(body), nil
}

func Applications(network *net.Network) ([]models.ApplicationSummary, error) {
	url := fmt.Sprintf("/v1/applications")
	var appSummary []models.ApplicationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return appSummary, err
	}

	err = json.Unmarshal(body, &appSummary)
	return appSummary, err
}

func Create(network *net.Network, resource string) (models.TaskSummary, error) {
	url := "/v1/applications"
	var response models.TaskSummary
	body, err := network.SendPostResourceRequest(url, resource, "application/json")
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

func CreateFromBytes(network *net.Network, blueprint []byte) (models.TaskSummary, error) {
	url := "/v1/applications"
	var response models.TaskSummary
	body, err := network.SendPostRequest(url, blueprint)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// WIP
func Descendants(network *net.Network, app string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/descendants", app)

	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

// WIP
func DescendantsSensor(network *net.Network, app, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/descendants/sensor/%s", app, sensor)

	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func Tree(network *net.Network) ([]models.Tree, error) {
	url := "/v1/applications/fetch"
	var tree []models.Tree
	body, err := network.SendGetRequest(url)
	if err != nil {
		return tree, err
	}

	err = json.Unmarshal(body, &tree)
	return tree, err
}

func Application(network *net.Network, app string) (models.ApplicationSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s", app)
	var appSummary models.ApplicationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return appSummary, err
	}

	err = json.Unmarshal(body, &appSummary)
	return appSummary, err
}

func Delete(network *net.Network, application string) (models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s", application)
	var response models.TaskSummary
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// WIP
func CreateLegacy(network *net.Network) (string, error) {
	url := fmt.Sprintf("/v1/applications/createLegacy")
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	// TODO return model
	return string(body), nil
}
