package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
    "github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type ListEntities struct {
	network *net.Network
}

func NewEntities(network *net.Network) (cmd *ListEntities) {
	cmd = new(ListEntities)
	cmd.network = network
	return
}

func (cmd *ListEntities) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entities",
		Description: "Show the entities for an application",
		Usage:       "BROOKLYN_NAME APP-SCOPE entities",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ListEntities) Run(scope scope.Scope, c *cli.Context) {
	entityList := entities.EntityList(cmd.network, scope.Application)
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entity := range entityList {
		table.Add(entity.Id, entity.Name, entity.Type)
	}
	table.Print()
}
