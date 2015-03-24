package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
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
		Usage:       "BROOKLYN_NAME start-policy APPLICATION ENTITY POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *StartPolicy) Run(c *cli.Context) {
	spec := entity_policies.StartPolicy(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(spec)
}
