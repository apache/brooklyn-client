package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_policies"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
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
		Usage:       "BROOKLYN_NAME stop-policy APPLICATION ENTITY POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StopPolicy) Run(c *cli.Context) {
	spec := entity_policies.StopPolicy(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(spec)
}
