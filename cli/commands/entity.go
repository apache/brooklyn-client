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
	"github.com/apache/brooklyn-client/cli/api/entities"
	"github.com/apache/brooklyn-client/cli/api/entity_sensors"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"os"
)

type Entity struct {
	network *net.Network
}

func NewEntity(network *net.Network) (cmd *Entity) {
	cmd = new(Entity)
	cmd.network = network
	return
}

func (cmd *Entity) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity",
		Aliases:     []string{"entities", "ent", "ents"},
		Description: "Show the entities of an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE entity [ENTITYID]",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "children, c",
				Usage: "List children of the entity",
			},
		},
	}
}

func (cmd *Entity) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.NumFlags() > 0 && c.FlagNames()[0] == "children" {
		cmd.listentity(scope.Application, c.StringSlice("children")[0])
	} else {
		if c.Args().Present() {
			cmd.show(scope.Application, c.Args().First())
		} else {
			if scope.Entity == scope.Application {
				cmd.listapp(scope.Application)
			} else {
				cmd.listentity(scope.Application, scope.Entity)
			}
		}
	}
}

const serviceStateSensor = "service.state"
const serviceIsUp = "service.isUp"

func (cmd *Entity) show(application, entity string) {
	summary, err := entities.GetEntity(cmd.network, application, entity)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	table := terminal.NewTable([]string{"Id:", summary.Id})
	table.Add("Name:", summary.Name)
	configState, err := entity_sensors.CurrentState(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	if serviceState, ok := configState[serviceStateSensor]; ok {
		table.Add("Status:", fmt.Sprintf("%v", serviceState))
	}
	if serviceIsUp, ok := configState[serviceIsUp]; ok {
		table.Add("ServiceUp:", fmt.Sprintf("%v", serviceIsUp))
	}
	table.Add("Type:", summary.Type)
	table.Add("CatalogItemId:", summary.CatalogItemId)
	table.Print()
}

func (cmd *Entity) listapp(application string) {
	entitiesList, err := entities.EntityList(cmd.network, application)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}

func (cmd *Entity) listentity(application string, entity string) {
	entitiesList, err := entities.Children(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}

	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}
