package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/catalog"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type AddCatalog struct {
	network *net.Network
}

func NewAddCatalog(network *net.Network) (cmd *AddCatalog) {
	cmd = new(AddCatalog)
	cmd.network = network
	return
}

func (cmd *AddCatalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add-catalog",
		Description: "Add a new catalog item from the supplied YAML",
		Usage:       "BROOKLYN_NAME [ SCOPE ] add-catalog FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddCatalog) Run(scope scope.Scope, c *cli.Context) {
	create := catalog.AddCatalog(cmd.network, c.Args().First())
	fmt.Println(create)
}
