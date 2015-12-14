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

type DestroyPolicy struct {
	network *net.Network
}

func NewDestroyPolicy(network *net.Network) (cmd *DestroyPolicy) {
	cmd = new(DestroyPolicy)
	cmd.network = network
	return
}

func (cmd *DestroyPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "destroy-policy",
		Description: "Destroy a policy",
		Usage:       "BROOKLYN_NAME [ SCOPE ] destroy-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *DestroyPolicy) Run(scope scope.Scope, c *cli.Context) {
	spec, err := entity_policies.DestroyPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(spec)
}
