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
	"github.com/apache/brooklyn-client/api/catalog"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type AddCatalog struct {
	network *net.Network
}

func NewAddCatalog(network *net.Network) (cmd *AddCatalog) {
	cmd = new(AddCatalog)
	cmd.network = network
	return
}

func (cmd *AddCatalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add-catalog",
		Description: "* Add a new catalog item from the supplied YAML",
		Usage:       "BROOKLYN_NAME add-catalog FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddCatalog) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	create, err := catalog.AddCatalog(cmd.network, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Fprintln(c.App.Writer, create)
}
