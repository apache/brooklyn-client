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
	"github.com/apache/brooklyn-client/api/catalog"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/urfave/cli"
)

type Catalog struct {
	network *net.Network
}

func NewCatalog(network *net.Network) (cmd *Catalog) {
	cmd = new(Catalog)
	cmd.network = network
	return
}

func (cmd *Catalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "catalog",
		Description: "* List the available catalog applications",
		Usage:       "BROOKLYN_NAME catalog",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Catalog) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	catalog, err := catalog.Catalog(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Description"})
	for _, app := range catalog {
		table.Add(app.Id, app.Name, app.Description)
	}
	table.Print()
}
