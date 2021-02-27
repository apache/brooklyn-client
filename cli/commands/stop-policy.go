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
	"github.com/apache/brooklyn-client/cli/api/entity_policies"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli/v2"
)

type StopPolicy struct {
	network *net.Network
}

func NewStopPolicy(network *net.Network) (cmd *StopPolicy) {
	cmd = new(StopPolicy)
	cmd.network = network
	return
}

func (cmd *StopPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stop-policy",
		Description: "Suspends a policy",
		Usage:       "BROOKLYN_NAME SCOPE stop-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StopPolicy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	spec, err := entity_policies.StopPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(spec)
}
