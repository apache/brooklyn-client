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
	"github.com/apache/brooklyn-client/cli/api/entity_config"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
)

type Config struct {
	network *net.Network
}

func NewConfig(network *net.Network) (cmd *Config) {
	cmd = new(Config)
	cmd.network = network
	return
}

func (cmd *Config) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "config",
		Description: "Show the config for an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE config",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Config) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.Args().Present() {
		configValue, err := entity_config.ConfigValue(cmd.network, scope.Application, scope.Entity, c.Args().First())

		if nil != err {
			error_handler.ErrorExit(err)
		}
		displayValue, err := stringRepresentation(configValue)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		fmt.Println(displayValue)

	} else {
		config, err := entity_config.ConfigCurrentState(cmd.network, scope.Application, scope.Entity)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		table := terminal.NewTable([]string{"Key", "Value"})
		for key, value := range config {
			table.Add(key, fmt.Sprintf("%v", value))
		}
		table.Print()
	}
}
