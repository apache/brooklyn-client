package commands

import(
	"strings"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_effectors"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
)

type Effectors struct {
	
}

func NewEffectors() (cmd *Effectors){
	cmd = new(Effectors)
	return
}

func (cmd *Effectors) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "effectors",
		Description: "show the list of effectors for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Effectors) Run(c *cli.Context) {
	effectors := entity_effectors.EffectorList(c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Name", "Description", "Parameters"})
	for _, effector := range effectors {
		var parameters []string
		for _, parameter := range effector.Parameters{
			parameters = append(parameters, parameter.Name)
		}
		table.Add(effector.Name, effector.Description, strings.Join(parameters, ","))
	}
	table.Print()
}