package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
    "github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type ListEntitySubCommand struct {
	network *net.Network
}

func NewListEntity(network *net.Network) (cmd *ListEntitySubCommand) {
	cmd = new(ListEntitySubCommand)
	cmd.network = network
	return
}

func (cmd *ListEntitySubCommand) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entities",
		Description: "Show the entities for an application",
		Usage:       "BROOKLYN_NAME [ SCOPE ] entities",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ListEntitySubCommand) Run(scope scope.Scope, c *cli.Context) {
	entityList := entities.EntityList(cmd.network, scope.Application)
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entity := range entityList {
		table.Add(entity.Id, entity.Name, entity.Type)
	}
	table.Print()
}
