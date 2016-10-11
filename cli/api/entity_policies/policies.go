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
package entity_policies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"net/url"
	"strconv"
)

// WIP
func AddPolicy(network *net.Network, application, entity, policy string, config map[string]string) (models.PolicySummary, error) {
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/policies", application, entity)
	data := url.Values{}
	data.Set("policyType", policy)
	//data.Add("config", config)
	req := network.NewPostRequest(path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	var policySummary models.PolicySummary
	body, err := network.SendRequest(req)
	if err != nil {
		return policySummary, err
	}
	err = json.Unmarshal(body, &policySummary)
	return policySummary, err
}

func PolicyList(network *net.Network, application, entity string) ([]models.PolicySummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies", application, entity)
	var policyList []models.PolicySummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return policyList, err
	}

	err = json.Unmarshal(body, &policyList)
	return policyList, err
}

func PolicyStatus(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s", application, entity, policy)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func CurrentState(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/current-state", application, entity)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func StartPolicy(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/start", application, entity, policy)
	body, err := network.SendEmptyPostRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func StopPolicy(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/stop", application, entity, policy)
	body, err := network.SendEmptyPostRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func DestroyPolicy(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/destroy", application, entity, policy)
	body, err := network.SendEmptyPostRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}
