package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_sensors"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Sensor struct {
	network *net.Network
}

func NewSensor(network *net.Network) (cmd *Sensor){
	cmd = new(Sensor)
	cmd.network = network
	return
}

func (cmd *Sensor) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensor",
		Description: "show the value of a sensor for an application and entity",
		Usage:       "BROOKLYN_NAME sensor APPLICATION ENTITY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Sensor) Run(c *cli.Context) {
	sensor := entity_sensors.SensorValue(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(sensor)
}