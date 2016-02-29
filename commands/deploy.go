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
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"strings"
)

type Deploy struct {
	network *net.Network
}

func NewDeploy(network *net.Network) (cmd *Deploy) {
	cmd = new(Deploy)
	cmd.network = network
	return
}

func (cmd *Deploy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "deploy",
		Description: "Deploy a new application from the given YAML (read from file or stdin)",
		Usage:       "BROOKLYN_NAME deploy ( <FILE> | - )",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Deploy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}

	var create models.TaskSummary
	var err error
	var blueprint []byte
	if c.Args().First() == "" {
		error_handler.ErrorExit("A filename or '-' must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}
	if c.Args().First() == "-" {
		blueprint, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			error_handler.ErrorExit(err)
		}
		create, err = application.CreateFromBytes(cmd.network, blueprint)
	} else {
		create, err = application.Create(cmd.network, c.Args().First())
	}
	if nil != err {
		if httpErr, ok := err.(net.HttpError); ok {
			error_handler.ErrorExit(strings.Join([]string{httpErr.Status, httpErr.Body}, "\n"), httpErr.Code)
		} else {
			error_handler.ErrorExit(err)
		}
	}
	table := terminal.NewTable([]string{"Id:", create.EntityId})
	table.Add("Name:", create.EntityDisplayName)
	table.Add("Status:", create.CurrentStatus)
	table.Print()
}
