package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_sensors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
)

type Sensor struct {
	network *net.Network
}

func NewSensor(network *net.Network) (cmd *Sensor) {
	cmd = new(Sensor)
	cmd.network = network
	return
}

func (cmd *Sensor) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensor",
		Description: "Show the value of a sensor for an application and entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] sensor SENSOR_NAME",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Sensor) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().Present() {
		cmd.show(scope.Application, scope.Entity, c.Args().First())
	} else {
		cmd.list(scope.Application, scope.Entity, )
	}
}

func (cmd *Sensor) show(application, entity, sensor string) {
	sensorValue := entity_sensors.SensorValue(cmd.network, application, entity, sensor)
	fmt.Println(sensorValue)
}


func (cmd *Sensor) list(application, entity string) {
	sensors := entity_sensors.SensorList(cmd.network, application, entity)
	table := terminal.NewTable([]string{"Name", "Description", "Value"})
	for _, sensor := range sensors {
		value := entity_sensors.SensorValue(cmd.network, application, entity, sensor.Name)
		table.Add(sensor.Name, sensor.Description, value)
	}
	table.Print()
}
