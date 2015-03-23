package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

type Policy struct {
	
}

func NewPolicy() (cmd *Policy){
	cmd = new(Policy)
	return
}

func (cmd *Policy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policy",
		Description: "show the status of a policy for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Policy) Run(c *cli.Context) {
	policy := entity_policies.PolicyStatus(c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(policy)
}