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
package entity_config

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func ConfigValue(network *net.Network, application, entity, config string) (interface{}, error) {
	bytes, err := ConfigValueAsBytes(network, application, entity, config)
	if nil != err || 0 == len(bytes) {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal(bytes, &value)
	if nil != err {
		return nil, err
	}

	return value, nil
}

func ConfigValueAsBytes(network *net.Network, application, entity, config string) ([]byte, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	body, err := network.SendGetRequest(url)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func SetConfig(network *net.Network, application, entity, config, value string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func ConfigList(network *net.Network, application, entity string) ([]models.ConfigSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
	var configList []models.ConfigSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return configList, err
	}

	err = json.Unmarshal(body, &configList)
	return configList, err
}

func PostConfig(network *net.Network, application, entity, config, value string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func ConfigCurrentState(network *net.Network, application, entity string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/current-state", application, entity)
	var currentState map[string]interface{}
	body, err := network.SendGetRequest(url)
	if err != nil {
		return currentState, err
	}
	err = json.Unmarshal(body, &currentState)
	return currentState, err
}
