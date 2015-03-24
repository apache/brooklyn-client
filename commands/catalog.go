package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/catalog"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/terminal"
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
		Usage:       "BROOKLYN_NAME catalog",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Catalog) Run(c *cli.Context) {
	catalog := catalog.Catalog(cmd.network)
	table := terminal.NewTable([]string{"Id", "Name", "Description"})
	for _, app := range catalog {
		table.Add(app.Id, app.Name, app.Description)
	}
	table.Print()
}
