package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_sensors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
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
	sensor := entity_sensors.SensorValue(cmd.network, scope.Application, scope.Entity, c.Args().First())
	fmt.Println(sensor)
}
