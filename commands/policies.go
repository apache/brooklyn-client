package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Policies struct {
	network *net.Network
}

func NewPolicies(network *net.Network) (cmd *Policies) {
	cmd = new(Policies)
	cmd.network = network
	return
}

func (cmd *Policies) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policies",
		Description: "Show the list of policies for an application and entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] policies APPLICATION ENTITY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Policies) Run(scope scope.Scope, c *cli.Context) {
	policies := entity_policies.PolicyList(cmd.network, scope.Application, scope.Entity)
	table := terminal.NewTable([]string{"Name", "State"})
	for _, policy := range policies {
		table.Add(policy.Name, string(policy.State))
	}
	table.Print()
}
