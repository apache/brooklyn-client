package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type DestroyPolicy struct {
	network *net.Network
}

func NewDestroyPolicy(network *net.Network) (cmd *DestroyPolicy){
	cmd = new(DestroyPolicy)
	cmd.network = network
	return
}

func (cmd *DestroyPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "destroy-policy",
		Description: "Destroy a policy",
		Usage:       "BROOKLYN_NAME destroy-policy APPLICATION ENTITY POLICY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *DestroyPolicy) Run(c *cli.Context) {
	spec := entity_policies.DestroyPolicy(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(spec)
}