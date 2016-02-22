package commands

import (
	"github.com/apache/brooklyn-client/api/entity_effectors"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/codegangsta/cli"
	"strings"
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
