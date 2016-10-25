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
package entity_policy_config

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func CurrentState(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/current-state", application, entity, policy)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func GetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

// WIP
func SetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendEmptyPostRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func GetAllConfigValues(network *net.Network, application, entity, policy string) ([]models.PolicyConfigList, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config", application, entity, policy)
	var policyConfigList []models.PolicyConfigList
	body, err := network.SendGetRequest(url)
	if nil != err {
		return policyConfigList, err
	}
	err = json.Unmarshal(body, &policyConfigList)
	return policyConfigList, err
}
