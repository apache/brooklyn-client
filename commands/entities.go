package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entities"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Entities struct {
	network *net.Network
}

func NewEntities(network *net.Network) (cmd *Entities){
	cmd = new(Entities)
	cmd.network = network
	return
}

func (cmd *Entities) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entities",
		Description: "show the entites for an application",
		Usage:       "BROOKLYN_NAME entities APPLICATION",
		Flags: []cli.Flag{},
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