package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Children struct {
	network *net.Network
}

func NewChildren(network *net.Network) (cmd *Children) {
	cmd = new(Children)
	cmd.network = network
	return
}

func (cmd *Children) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity-children",
		Description: "Show the children of an entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] entity-children",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Children) Run(scope scope.Scope, c *cli.Context) {
	entityList := entities.Children(cmd.network, scope.Application, scope.Entity)
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entity := range entityList {
		table.Add(entity.Id, entity.Name, entity.Type)
	}
	table.Print()
}
