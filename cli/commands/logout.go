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
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/io"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli/v2"
)

type Logout struct {
	network *net.Network
	config  *io.Config
}

func NewLogout(network *net.Network, config *io.Config) (cmd *Logout) {
	cmd = new(Logout)
	cmd.network = network
	cmd.config = config
	return
}

func (cmd *Logout) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "logout",
		Description: "Logout of brooklyn",
		Usage:       "BROOKLYN_NAME logout",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Logout) Run(scope scope.Scope, c *cli.Context) {
	config := io.GetConfig()
	err := config.Delete()
	if err != nil {
		error_handler.ErrorExit(err)
	}
}
