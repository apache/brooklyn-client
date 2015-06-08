package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type Create struct {
	network *net.Network
}

func NewCreate(network *net.Network) (cmd *Create) {
	cmd = new(Create)
	cmd.network = network
	return
}

func (cmd *Create) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "create",
		Description: "Create a new brooklyn application from the supplied YAML",
		Usage:       "BROOKLYN_NAME create FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Create) Run(c *cli.Context) {
	create := application.Create(cmd.network, c.Args()[0])
	fmt.Println(create)
}
