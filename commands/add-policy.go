package commands

import (
	"github.com/codegangsta/cli"
	//"github.com/apache/brooklyn-client/api/entity_policies"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
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
