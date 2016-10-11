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
	"github.com/apache/brooklyn-client/cli/api/entity_sensors"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"sort"
)

type Sensor struct {
	network *net.Network
}

type sensorList []models.SensorSummary

// Len is the number of elements in the collection.
func (sensors sensorList) Len() int {
	return len(sensors)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (sensors sensorList) Less(i, j int) bool {
	return sensors[i].Name < sensors[j].Name
}

// Swap swaps the elements with indexes i and j.
func (sensors sensorList) Swap(i, j int) {
	temp := sensors[i]
	sensors[i] = sensors[j]
	sensors[j] = temp
}

func NewSensor(network *net.Network) (cmd *Sensor) {
	cmd = new(Sensor)
	cmd.network = network
	return
}

func (cmd *Sensor) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensor",
		Description: "Show values of all sensors or named sensor for an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE sensor [ SENSOR_NAME ]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Sensor) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.Args().Present() {
		cmd.show(scope.Application, scope.Entity, c.Args().First())
	} else {
		cmd.list(scope.Application, scope.Entity)
	}
}

func (cmd *Sensor) show(application, entity, sensor string) {
	sensorValue, err := entity_sensors.SensorValue(cmd.network, application, entity, sensor)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	displayValue, err := stringRepresentation(sensorValue)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(displayValue)
}

func (cmd *Sensor) list(application, entity string) {
	sensors, err := entity_sensors.SensorList(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	var theSensors sensorList = sensors
	table := terminal.NewTable([]string{"Name", "Description", "Value"})

	sort.Sort(theSensors)

	for _, sensor := range theSensors {
		value, err := entity_sensors.SensorValue(cmd.network, application, entity, sensor.Name)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		displayValue, err := stringRepresentation(value)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		table.Add(sensor.Name, sensor.Description, displayValue)
	}
	table.Print()
}
