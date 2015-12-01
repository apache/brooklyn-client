package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
    "github.com/brooklyncentral/brooklyn-cli/command_metadata"
)

type Entities struct {
	network *net.Network
}

func NewListEntity(network *net.Network) (cmd *Entities) {
	cmd = new(Entities)
	cmd.network = network
	return
}

func (cmd *Entities) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity",
		Description: "Show the entites for an application",
		Usage:       "list entity APPLICATION",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Entities) Run(c *cli.Context) {
	entityList := entities.EntityList(cmd.network, c.Args()[0])
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entity := range entityList {
		table.Add(entity.Id, entity.Name, entity.Type)
	}
	table.Print()
}
