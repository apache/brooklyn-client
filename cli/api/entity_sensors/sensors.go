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
package entity_sensors

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func SensorValue(network *net.Network, application, entity, sensor string) (interface{}, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	if nil != err || 0 == len(body) {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal(body, &value)
	if nil != err {
		return nil, err
	}

	return value, nil
}

// WIP
func DeleteSensor(network *net.Network, application, entity, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendDeleteRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}

// WIP
//func SetSensor(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}

//	return string(body)
//}

// WIP
//func SetSensors(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}

//	return string(body)
//}

func SensorList(network *net.Network, application, entity string) ([]models.SensorSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity)
	body, err := network.SendGetRequest(url)
	var sensorList []models.SensorSummary
	if err != nil {
		return sensorList, err
	}

	err = json.Unmarshal(body, &sensorList)
	return sensorList, err
}

func CurrentState(network *net.Network, application, entity string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/current-state", application, entity)
	var currentState map[string]interface{}
	body, err := network.SendGetRequest(url)
	if err != nil {
		return currentState, err
	}

	err = json.Unmarshal(body, &currentState)
	return currentState, err
}
