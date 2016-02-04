package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_effectors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type Effector struct {
	network *net.Network
}

func NewEffector(network *net.Network) (cmd *Effector) {
	cmd = new(Effector)
	cmd.network = network
	return
}

func (cmd *Effector) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "effector",
		Description: "Show the effectors for an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE effector [ NAME ]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Effector) Run(scope scope.Scope, c *cli.Context) {
    if err := net.VerifyLoginURL(cmd.network); err != nil {
        error_handler.ErrorExit(err)
    }
	effectors, err := entity_effectors.EffectorList(cmd.network, scope.Application, scope.Entity)
    if nil != err {
        error_handler.ErrorExit(err)
    }
	table := terminal.NewTable([]string{"Name", "Description", "Parameters"})
	for _, effector := range effectors {
		var parameters []string
		for _, parameter := range effector.Parameters {
			parameters = append(parameters, parameter.Name)
		}
        if !c.Args().Present() || c.Args().First() == effector.Name {
            table.Add(effector.Name, effector.Description, strings.Join(parameters, ","))
        }
	}
	table.Print()
}
