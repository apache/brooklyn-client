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
package access_control

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
)

func Access(network *net.Network) (models.AccessSummary, error) {
	url := fmt.Sprintf("/v1/access")
	var access models.AccessSummary

	body, err := network.SendGetRequest(url)
	if err != nil {
		return access, err
	}

	err = json.Unmarshal(body, &access)
	return access, err
}

// WIP
//func LocationProvisioningAllowed(network *net.Network, allowed bool) {
//	url := fmt.Sprintf("/v1/access/locationProvisioningAllowed")
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}
//}
