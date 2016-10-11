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
package activities

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func Activity(network *net.Network, activity string) (models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/activities/%s", activity)
	var task models.TaskSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(body, &task)
	return task, err
}

func ActivityChildren(network *net.Network, activity string) ([]models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/activities/%s/children", activity)
	var tasks []models.TaskSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(body, &tasks)
	return tasks, err
}

func ActivityStream(network *net.Network, activity, streamId string) (string, error) {
	url := fmt.Sprintf("/v1/activities/%s/stream/%s", activity, streamId)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}
