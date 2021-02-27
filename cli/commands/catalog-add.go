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
	"github.com/apache/brooklyn-client/cli/api/catalog"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli/v2"
)

type CatalogAdd struct {
	network *net.Network
}

func NewCatalogAdd(network *net.Network) (cmd *CatalogAdd) {
	cmd = new(CatalogAdd)
	cmd.network = network
	return
}

func (cmd *CatalogAdd) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add",
		Description: "Add catalog items from the supplied BOM YAML or from a ZIP or local folder containing a catalog.bom and optionally other resources",
		Usage:       "BROOKLYN_NAME catalog add ( FILEPATH | URL )",
		Flags:       []cli.Flag{},
	}
}

func (cmd *CatalogAdd) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().First() == "" {
		error_handler.ErrorExit("A filename or URL must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	create, err := catalog.AddCatalog(cmd.network, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	if "" != create.Message {
		fmt.Println(create.Message)
		for id, _ := range create.Types {
			fmt.Printf("* %s\n", id)
		}
	} else {
		for id, _ := range create.Types {
			fmt.Println(id)
		}
	}
}
