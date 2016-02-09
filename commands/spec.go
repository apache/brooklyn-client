package commands

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/codegangsta/cli"
)

type Spec struct {
	network *net.Network
}

func NewSpec(network *net.Network) (cmd *Spec) {
	cmd = new(Spec)
	cmd.network = network
	return
}

func (cmd *Spec) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "spec",
		Description: "Get the YAML spec used to create the entity, if available",
		Usage:       "BROOKLYN_NAME SCOPE spec",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Spec) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	spec, err := entities.Spec(cmd.network, scope.Application, scope.Entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(spec)
}
