package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Deploy struct {
	network *net.Network
}

func NewDeploy(network *net.Network) (cmd *Deploy) {
	cmd = new(Deploy)
	cmd.network = network
	return
}

func (cmd *Deploy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "deploy",
		Description: "Deploy a new brooklyn application from the supplied YAML",
		Usage:       "BROOKLYN_NAME [ SCOPE ] deploy FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Deploy) Run(scope scope.Scope, c *cli.Context) {
	create := application.Create(cmd.network, c.Args().First())
	fmt.Println(create)
}
