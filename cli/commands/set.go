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
package commands

import (
	"fmt"
	"errors"
	"github.com/apache/brooklyn-client/cli/api/entity_sensors"
	"github.com/apache/brooklyn-client/cli/api/entity_config"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
)

type SetConfig struct {
	network *net.Network
}

func NewSetConfig(network *net.Network) (cmd *SetConfig) {
	cmd = new(SetConfig)
	cmd.network = network
	return
}

func (cmd *SetConfig) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "set",
		Description: "Set config for an entity",
		Usage:       "BROOKLYN_NAME CONFIG-SCOPE set VALUE",
		Flags:       []cli.Flag{},
	}
}

func (cmd *SetConfig) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	
	config := scope.Config
	sensor := scope.Sensor
	if config != "" && sensor != "" {
		error_handler.ErrorExit(errors.New("Both config and sensor supplied when calling set. Please only specify one."))
	}
	
	if config == "" && sensor == "" {
		error_handler.ErrorExit(errors.New("Please specify either config or scope when calling set."))
	}
	
	if config != "" {
		_, err := entity_config.SetConfig(cmd.network, scope.Application, scope.Entity, config, c.Args().First())
		if nil != err {
			error_handler.ErrorExit(err)
		}
		fmt.Println("Config set correctly")
	} else {
		_, err := entity_sensors.SetSensor(cmd.network, scope.Application, scope.Entity, sensor, c.Args().First())
		if nil != err {
			error_handler.ErrorExit(err)
		}
		fmt.Println("Sensor set correctly")
	}
}
