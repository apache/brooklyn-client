package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_sensors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
"github.com/brooklyncentral/brooklyn-cli/models"
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
		cmd.list(scope.Application, scope.Entity, )
	}
}

func (cmd *Sensor) show(application, entity, sensor string) {
	sensorValue, err := entity_sensors.SensorValue(cmd.network, application, entity, sensor)
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(sensorValue)
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
		table.Add(sensor.Name, sensor.Description, value)
	}
	table.Print()
}
