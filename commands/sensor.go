package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_sensors"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

type Sensor struct {
	
}

func NewSensor() (cmd *Sensor){
	cmd = new(Sensor)
	return
}

func (cmd *Sensor) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "sensor",
		Description: "show the value of a sensor for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Sensor) Run(c *cli.Context) {
	sensor := entity_sensors.SensorValue(c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(sensor)
}