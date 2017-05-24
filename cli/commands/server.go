/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Server 2.0 (the
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
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
	"fmt"
)

type Server struct {
	network *net.Network
}

func NewServer(network *net.Network) (cmd *Server) {
	cmd = new(Server)
	cmd.network = network
	return
}

func (cmd *Server) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "server",
		Description: "Display the URL of the connected Brooklyn",
		Usage:       "BROOKLYN_NAME server",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Server) Run(scope scope.Scope, c *cli.Context) {
	fmt.Println(cmd.network.BrooklynUrl)
}
