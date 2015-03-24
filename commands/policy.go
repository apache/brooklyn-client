package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Policy struct {
	network *net.Network
}

func NewPolicy(network *net.Network) (cmd *Policy){
	cmd = new(Policy)
	cmd.network = network
	return
}

func (cmd *Policy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policy",
		Description: "Show the status of a policy for an application and entity",
		Usage:       "BROOKLYN_NAME policy APPLICATION ENITITY POLICY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Policy) Run(c *cli.Context) {
	policy := entity_policies.PolicyStatus(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(policy)
}