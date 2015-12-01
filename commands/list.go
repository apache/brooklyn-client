package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/codegangsta/cli"
	"fmt"
	"strings"
)

type List struct {
	network *net.Network
}

func NewList(network *net.Network) (cmd *List) {
	cmd = new(List)
	cmd.network = network
	return
}

const applicationCommand = "application"
const entityCommand = "entity"
const sensorCommand = "sensor"
const effectorCommand = "effector"

var listCommands = []string{
	applicationCommand,
    entityCommand,
    sensorCommand,
    effectorCommand,
}
var listCommandsUsage = strings.Join(listCommands, " | ")

func (cmd *List) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "list",
		Description: "List details for a variety of operands",
		Usage:       "BROOKLYN_NAME list (" + listCommandsUsage + ")",
		Flags:       []cli.Flag{},
		Operands:    []*command_metadata.CommandMetadata {
			&command_metadata.CommandMetadata{
				Name:        "application",
				Description: "List details of applications",
				Usage:       "list application",
				Flags:       []cli.Flag{},
			},
			&command_metadata.CommandMetadata{
				Name:        "entity",
				Description: "List details of entities of an application",
				Usage:       "list entity APPID ENTITYID",
				Flags:       []cli.Flag{},
			},
			&command_metadata.CommandMetadata{
				Name:        "sensor",
				Description: "List details of sensors of an entity",
				Usage:       "list sensor APPID ENTITYID",
				Flags:       []cli.Flag{},
			},
			&command_metadata.CommandMetadata{
				Name:        "effector",
				Description: "List details of effectors of an entity",
				Usage:       "list entity APPID ENTITYID",
				Flags:       []cli.Flag{},
			},
		},
	}
}

func (cmd *List) Run(c *cli.Context) {
	fmt.Printf( "Unrecognised item for list, please use one of (%s)\n", listCommandsUsage)
}