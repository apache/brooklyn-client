package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/catalog"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type Catalog struct {
	network *net.Network
}

func NewCatalog(network *net.Network) (cmd *Catalog) {
	cmd = new(Catalog)
	cmd.network = network
	return
}

func (cmd *Catalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "catalog",
		Description: "List the available catalog applications",
		Usage:       "BROOKLYN_NAME [ SCOPE ] catalog",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Catalog) Run(scope scope.Scope, c *cli.Context) {
	catalog, err := catalog.Catalog(cmd.network)
    if nil != err {
        error_handler.ErrorExit(err)
    }
	table := terminal.NewTable([]string{"Id", "Name", "Description"})
	for _, app := range catalog {
		table.Add(app.Id, app.Name, app.Description)
	}
	table.Print()
}
