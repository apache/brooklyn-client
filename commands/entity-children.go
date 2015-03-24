package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entities"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Children struct {
	network *net.Network
}

func NewChildren(network *net.Network) (cmd *Children){
	cmd = new(Children)
	cmd.network = network
	return
}

func (cmd *Children) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity-children",
		Description: "Show the children of an application's entity",
		Usage:       "BROOKLYN_NAME children APPLICATION ENTITY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Children) Run(c *cli.Context) {
	entityList := entities.Children(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entity := range entityList {
		table.Add(entity.Id, entity.Name, entity.Type)
	}
	table.Print()
}