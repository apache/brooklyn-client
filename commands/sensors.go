package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_sensors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
)

type ListSensorsSubCommand struct {
	network *net.Network
}

func NewListSensor(network *net.Network) (cmd *ListSensorsSubCommand) {
	cmd = new(ListSensorsSubCommand)
	cmd.network = network
	return
}

func (cmd *ListSensorsSubCommand) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensor",
		Description: "Show the sensors for an application and entity",
		Usage:       "list sensor APPLICATION ENTITY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ListSensorsSubCommand) Run(c *cli.Context) {
	sensors := entity_sensors.SensorList(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Name", "Description", "Value"})
	for _, sensor := range sensors {
		value := entity_sensors.SensorValue(cmd.network, c.Args()[0], c.Args()[1], sensor.Name)
		table.Add(sensor.Name, sensor.Description, value)
	}
	table.Print()
}
