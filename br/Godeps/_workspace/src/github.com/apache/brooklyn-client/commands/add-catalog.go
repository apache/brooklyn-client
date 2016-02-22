package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/catalog"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
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
		Description: "* Add a new catalog item from the supplied YAML",
		Usage:       "BROOKLYN_NAME add-catalog FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddCatalog) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	create, err := catalog.AddCatalog(cmd.network, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(create)
}
