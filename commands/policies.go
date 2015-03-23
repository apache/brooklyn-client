package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_policies"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
)

type Policies struct {
	
}

func NewPolicies() (cmd *Policies){
	cmd = new(Policies)
	return
}

func (cmd *Policies) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policies",
		Description: "show the list of policies for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Policies) Run(c *cli.Context) {
	policies := entity_policies.PolicyList(c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Name", "State"})
	for _, policy := range policies {
		table.Add(policy.Name, string(policy.State))
	}
	table.Print()
}