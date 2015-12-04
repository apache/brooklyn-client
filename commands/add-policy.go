package commands

import (
	"github.com/codegangsta/cli"
	//"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type AddPolicy struct {
	network *net.Network
}

func NewAddPolicy(network *net.Network) (cmd *AddPolicy) {
	cmd = new(AddPolicy)
	cmd.network = network
	return
}

func (cmd *AddPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add-policy",
		Description: "Add a new policy",
		Usage:       "BROOKLYN_NAME [ SCOPE ] add-policy APPLICATION ENTITY POLICY_TYPE",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddPolicy) Run(scope scope.Scope, c *cli.Context) {
	// Todo
}
