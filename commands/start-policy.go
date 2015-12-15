package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type StartPolicy struct {
	network *net.Network
}

func NewStartPolicy(network *net.Network) (cmd *StartPolicy) {
	cmd = new(StartPolicy)
	cmd.network = network
	return
}

func (cmd *StartPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "start-policy",
		Description: "Start or resume a policy",
		Usage:       "BROOKLYN_NAME SCOPE start-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StartPolicy) Run(scope scope.Scope, c *cli.Context) {
	spec, err := entity_policies.StartPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(spec)
}
