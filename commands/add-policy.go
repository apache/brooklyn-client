package commands

import(
	"github.com/codegangsta/cli"
	//"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type AddPolicy struct {
	network *net.Network
}

func NewAddPolicy(network *net.Network) (cmd *AddPolicy){
	cmd = new(AddPolicy)
	cmd.network = network
	return
}

func (cmd *AddPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add-policy",
		Description: "Add a new policy",
		Usage:       "BROOKLYN_NAME add-policy APPLICATION ENTITY POLICY_TYPE",
		Flags: []cli.Flag{},
	}
}	

func (cmd *AddPolicy) Run(c *cli.Context) {
	// Todo
}