package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type StopPolicy struct {
	network *net.Network
}

func NewStopPolicy(network *net.Network) (cmd *StopPolicy) {
	cmd = new(StopPolicy)
	cmd.network = network
	return
}

func (cmd *StopPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stop-policy",
		Description: "Suspends a policy",
		Usage:       "BROOKLYN_NAME [ SCOPE ] stop-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StopPolicy) Run(scope scope.Scope, c *cli.Context) {
	spec := entity_policies.StopPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
	fmt.Println(spec)
}
