package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
    "github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "os"
    "fmt"
)

type Entity struct {
	network *net.Network
}

func NewEntity(network *net.Network) (cmd *Entity) {
	cmd = new(Entity)
	cmd.network = network
	return
}

func (cmd *Entity) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity",
		Aliases:     []string{"entities","ent","ents"},
		Description: "Show the entities of an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE entity",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Entity) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().Present() {
		cmd.show(scope.Application, c.Args().First())
	} else {
		if scope.Entity == scope.Application {
			cmd.listapp(scope.Application)
		} else {
			cmd.listentity(scope.Application, scope.Entity)
		}
	}
}

func (cmd *Entity) show(application, entity string) {
	summary, err := entities.GetEntity(cmd.network, application, entity)
    if nil != err {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        os.Exit(1)
    }
	table := terminal.NewTable([]string{"Id:", summary.Id})
	table.Add("Name:", summary.Name)
	table.Add("Type:", summary.Type)
	table.Add("CatalogItemId:", summary.CatalogItemId)
	table.Print()
}


func (cmd *Entity) listapp(application string) {
	entitiesList := entities.EntityList(cmd.network, application)
	
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}

func (cmd *Entity) listentity(application string, entity string) {
	entitiesList := entities.Children(cmd.network, application, entity)

	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}