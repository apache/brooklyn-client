package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/catalog"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
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
		Usage:       "BROOKLYN_NAME add-catalog FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddCatalog) Run(c *cli.Context) {
	create := catalog.AddCatalog(cmd.network, c.Args()[0])
	fmt.Println(create)
}
