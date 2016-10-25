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
package locations

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func LocatedLocations(network *net.Network) (string, error) {
	url := "/v1/locations/usage/LocatedLocations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetLocation(network *net.Network, locationId string) (models.LocationSummary, error) {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	var locationDetail models.LocationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return locationDetail, err
	}
	err = json.Unmarshal(body, &locationDetail)
	return locationDetail, err
}

func DeleteLocation(network *net.Network, locationId string) (string, error) {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// WIP
func CreateLocation(network *net.Network, locationId string) (string, error) {
	url := fmt.Sprintf("/v1/locations", locationId)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func LocationList(network *net.Network) ([]models.LocationSummary, error) {
	url := "/v1/locations"
	var locationList []models.LocationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return locationList, err
	}

	err = json.Unmarshal(body, &locationList)
	return locationList, err
}
