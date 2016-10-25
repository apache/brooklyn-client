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
	"github.com/apache/brooklyn-client/cli/api/entity_effectors"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"strings"
)

type Effector struct {
	network *net.Network
}

func NewEffector(network *net.Network) (cmd *Effector) {
	cmd = new(Effector)
	cmd.network = network
	return
}

func (cmd *Effector) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "effector",
		Description: "Show the effectors for an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE effector [ NAME ]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Effector) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	effectors, err := entity_effectors.EffectorList(cmd.network, scope.Application, scope.Entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Name", "Description", "Parameters"})
	for _, effector := range effectors {
		var parameters []string
		for _, parameter := range effector.Parameters {
			parameters = append(parameters, parameter.Name)
		}
		if !c.Args().Present() || c.Args().First() == effector.Name {
			table.Add(effector.Name, effector.Description, strings.Join(parameters, ","))
		}
	}
	table.Print()
}
