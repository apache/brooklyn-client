package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_effectors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type ListEffectorSubCommand struct {
	network *net.Network
}

func NewListEffector(network *net.Network) (cmd *ListEffectorSubCommand) {
	cmd = new(ListEffectorSubCommand)
	cmd.network = network
	return
}

func (cmd *ListEffectorSubCommand) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "effectors",
		Description: "Show the list of effectors for an application and entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] effectors",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ListEffectorSubCommand) Run(scope scope.Scope, c *cli.Context) {
	effectors := entity_effectors.EffectorList(cmd.network, scope.Application, scope.Entity)
	table := terminal.NewTable([]string{"Name", "Description", "Parameters"})
	for _, effector := range effectors {
		var parameters []string
		for _, parameter := range effector.Parameters {
			parameters = append(parameters, parameter.Name)
		}
		table.Add(effector.Name, effector.Description, strings.Join(parameters, ","))
	}
	table.Print()
}
