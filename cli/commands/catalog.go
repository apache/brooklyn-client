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
	"github.com/apache/brooklyn-client/cli/command"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
	"strings"
	"fmt"
	"errors"
)

type Catalog struct {
	network *net.Network
	catalogCommands map[string]command.Command
}

func NewCatalog(network *net.Network) (cmd *Catalog) {
	cmd = new(Catalog)
	cmd.network = network
	cmd.catalogCommands = map[string]command.Command {
		ShowCatalogCommand: NewCatalogShow(cmd.network),
		ListCatalogCommand: NewCatalogList(cmd.network),
		AddCatalogCommand: NewCatalogAdd(cmd.network),
		DeleteCatalogCommand: NewDeleteCatalogItem(cmd.network),
	}
	return
}

const ShowCatalogCommand = "show"
const ListCatalogCommand = "list"
const AddCatalogCommand = "add"
const DeleteCatalogCommand = "delete"

var catalogCommands = []string{
	ShowCatalogCommand,
	ListCatalogCommand,
	AddCatalogCommand,
	DeleteCatalogCommand,
}
var catalogCommandsUsage = "list TYPE | add FILE/URL | delete TYPE ID:VERSION | show ITEM"

type CatalogItemType int
const  (
	Unknown = iota
	ApplicationsItemType
	EntitiesItemType
	LocationsItemType
	PoliciesItemType
)
const catalogItemTypesUsage = " ( application | entity | location | policy )"

func GetCatalogType(c *cli.Context) (CatalogItemType, error) {
	commandType := c.Args().First()
	if strings.HasPrefix("entity", commandType) {
		return EntitiesItemType, nil
	} else if strings.HasPrefix("location", commandType) {
		return LocationsItemType, nil
	} else if strings.HasPrefix("policy", commandType) {
		return PoliciesItemType, nil
	} else if strings.HasPrefix("application", commandType) {
		return ApplicationsItemType, nil
	}
	return Unknown, errors.New("Unknown type: " + commandType)
}

func (cmd *Catalog) SubCommandNames() []string {
	return catalogCommands
}

func (cmd *Catalog) SubCommand(name string) command.Command {
	return cmd.catalogCommands[name]
}

func (cmd *Catalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "catalog",
		Description: "Catalog operations",
		Usage:       "BROOKLYN_NAME catalog (" + catalogCommandsUsage + ")",
		Flags:       []cli.Flag{},
		Operands:    []command_metadata.CommandMetadata{
			cmd.SubCommand(ShowCatalogCommand).Metadata(),
			cmd.SubCommand(ListCatalogCommand).Metadata(),
			cmd.SubCommand(AddCatalogCommand).Metadata(),
			cmd.SubCommand(DeleteCatalogCommand).Metadata(),
		},
	}
}

func (cmd *Catalog) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
 	fmt.Printf("'catalog' requires one of (%s)\n", catalogCommandsUsage)
}
