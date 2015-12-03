package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/codegangsta/cli"
	"fmt"
	"strings"
	"github.com/brooklyncentral/brooklyn-cli/command"
)

type List struct {
	network *net.Network
    listCommands map[string]command.Command
}

func NewList(network *net.Network) (cmd *List) {
	cmd = new(List)
	cmd.network = network
	cmd.listCommands = map[string]command.Command{
		ListApplicationCommand: NewListApplication(cmd.network),
		ListEntityCommand: NewListEntity(cmd.network),
		ListSensorCommand: NewListSensor(cmd.network),
		ListEffectorCommand: NewListEffector(cmd.network),
	}
	return
}

const ListApplicationCommand = "application"
const ListEntityCommand = "entity"
const ListSensorCommand = "sensor"
const ListEffectorCommand = "effector"

var listCommands = []string {
	ListApplicationCommand,
	ListEntityCommand,
	ListSensorCommand,
	ListEffectorCommand,
}
var listCommandsUsage = strings.Join(listCommands, " | ")

func (cmd *List) SubCommand(name string) command.Command {
	return cmd.listCommands[name]
}

func (cmd *List) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "list",
		Description: "List details for a variety of operands",
		Usage:       "BROOKLYN_NAME list (" + listCommandsUsage + ")",
		Flags:       []cli.Flag{},
		Operands:    []command_metadata.CommandMetadata {
			cmd.SubCommand(ListApplicationCommand).Metadata(),
			cmd.SubCommand(ListEntityCommand).Metadata(),
			cmd.SubCommand(ListSensorCommand).Metadata(),
			cmd.SubCommand(ListEffectorCommand).Metadata(),
		},
	}
}

func (cmd *List) Run(c *cli.Context) {
	fmt.Printf( "Unrecognised item for list, please use one of (%s)\n", listCommandsUsage)
}