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
package entity_effectors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"net/url"
	"strconv"
	"strings"
)

func EffectorList(network *net.Network, application, entity string) ([]models.EffectorSummary, error) {
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors", application, entity)
	var effectorList []models.EffectorSummary
	body, err := network.SendGetRequest(path)
	if err != nil {
		return effectorList, err
	}

	err = json.Unmarshal(body, &effectorList)
	return effectorList, err
}

func TriggerEffector(network *net.Network, application, entity, effector string, params []string, args []string) (string, error) {
	if len(params) != len(args) {
		return "", errors.New(strings.Join([]string{"Parameters not supplied:", strings.Join(params, ", ")}, " "))
	}
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors/%s", application, entity, effector)
	data := url.Values{}
	for i := range params {
		data.Set(params[i], args[i])
	}
	req := network.NewPostRequest(path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	body, err := network.SendRequest(req)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
