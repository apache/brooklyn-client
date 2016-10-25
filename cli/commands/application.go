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
	"github.com/apache/brooklyn-client/cli/api/application"
	"github.com/apache/brooklyn-client/cli/api/entities"
	"github.com/apache/brooklyn-client/cli/api/entity_sensors"
	"github.com/apache/brooklyn-client/cli/api/locations"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"strings"
)

type Application struct {
	network *net.Network
}

func NewApplication(network *net.Network) (cmd *Application) {
	cmd = new(Application)
	cmd.network = network
	return
}

func (cmd *Application) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Aliases:     []string{"applications", "app", "apps"},
		Description: "Show the status and location of running applications",
		Usage:       "BROOKLYN_NAME application [APP]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Application) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.Args().Present() {
		cmd.show(c.Args().First())
	} else {
		cmd.list()
	}
}

const serviceIsUpStr = "service.isUp"

func (cmd *Application) show(appName string) {
	application, err := application.Application(cmd.network, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	entity, err := entities.GetEntity(cmd.network, appName, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	state, err := entity_sensors.CurrentState(cmd.network, appName, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id:", application.Id})
	table.Add("Name:", application.Spec.Name)
	table.Add("Status:", string(application.Status))
	if serviceUp, ok := state[serviceIsUpStr]; ok {
		table.Add("ServiceUp:", fmt.Sprintf("%v", serviceUp))
	}
	table.Add("Type:", application.Spec.Type)
	table.Add("CatalogItemId:", entity.CatalogItemId)
	if len(application.Spec.Locations) > 0 {
		table.Add("LocationId:", strings.Join(application.Spec.Locations, ", "))
		location, err := locations.GetLocation(cmd.network, application.Spec.Locations[0])
		if nil != err {
			error_handler.ErrorExit(err)
		}
		table.Add("LocationName:", location.Name)
		table.Add("LocationSpec:", location.Spec)
		table.Add("LocationType:", location.Type)
	}
	table.Print()
}

func (cmd *Application) list() {
	applications, err := application.Applications(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Status", "Location"})
	for _, app := range applications {
		table.Add(app.Id, app.Spec.Name, string(app.Status), strings.Join(app.Spec.Locations, ", "))
	}
	table.Print()
}
