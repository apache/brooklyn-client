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
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/urfave/cli"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"strings"
	"github.com/apache/brooklyn-client/cli/api/catalog"
)

type DeleteCatalogItem  struct {
	network *net.Network
}

func NewDeleteCatalogItem(network *net.Network) (cmd *DeleteCatalogItem) {
	cmd = new(DeleteCatalogItem)
	cmd.network = network
	return
}

func (cmd *DeleteCatalogItem) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "delete",
		Description: "delete the given catalog item",
		Usage:       "BROOKLYN_NAME catalog delete [ITEM_ID:VERSION]",
		Flags:       []cli.Flag{
			cli.BoolFlag{
				Name:  "applications, a",
				Usage: "delete application (default)",
			},
			cli.BoolFlag{
				Name:  "entities, e",
				Usage: "delete entity",
			},
			cli.BoolFlag{
				Name:  "locations, l",
				Usage: "delete location",
			},
			cli.BoolFlag{
				Name:  "policies, p",
				Usage: "delete policy",
			},
		},
	}
}

func (cmd *DeleteCatalogItem) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if len(c.Args()) != 1 {
		error_handler.ErrorExit("command requires single argument ITEM_ID:VERSION")
	}
	itemVersion := strings.Split(c.Args().First(), ":")
	if len(itemVersion) != 2 {
		error_handler.ErrorExit("command requires single argument ITEM_ID:VERSION")
	}
	itemId := itemVersion[0]
	version := itemVersion[1]
	err := cmd.deleteItem(c, itemId, version)
	if nil != err {
		error_handler.ErrorExit(err)
	}
}

func (cmd *DeleteCatalogItem) deleteItem(c *cli.Context, itemId string, version string) (error){
	if c.IsSet("entities") {
		return cmd.deleteEntity(c, itemId, version)
	} else if c.IsSet("locations") {
		return cmd.deleteLocation(c, itemId, version)
	} else if c.IsSet("policies") {
		return cmd.deletePolicy(c, itemId, version)
	}
	return cmd.deleteApplication(c, itemId, version)
}

func (cmd *DeleteCatalogItem) deleteApplication(c *cli.Context, itemId string, version string) (error){
	_, err := catalog.DeleteApplicationWithVersion(cmd.network, itemId, version)
	return err
}

func (cmd *DeleteCatalogItem) deleteEntity(c *cli.Context, itemId string, version string) (error){
	_, err := catalog.DeleteEntityWithVersion(cmd.network, itemId, version)
	return err
}

func (cmd *DeleteCatalogItem) deletePolicy(c *cli.Context, itemId string, version string) (error){
	_, err := catalog.DeletePolicyWithVersion(cmd.network, itemId, version)
	return err
}

func (cmd *DeleteCatalogItem) deleteLocation(c *cli.Context, itemId string, version string) (error){
	_, err := catalog.DeleteLocationWithVersion(cmd.network, itemId, version)
	return err
}