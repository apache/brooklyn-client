package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
    "github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
    "github.com/brooklyncentral/brooklyn-cli/terminal"
    "io/ioutil"
	"os"
)

type Deploy struct {
	network *net.Network
}

func NewDeploy(network *net.Network) (cmd *Deploy) {
	cmd = new(Deploy)
	cmd.network = network
	return
}

func (cmd *Deploy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "deploy",
		Description: "Deploy a new application from the given YAML (read from file or stdin)",
		Usage:       "BROOKLYN_NAME deploy ( <FILE> | - )",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Deploy) Run(scope scope.Scope, c *cli.Context) {
    if err := net.VerifyLoginURL(cmd.network); err != nil {
        error_handler.ErrorExit(err)
    }
    
    var create models.TaskSummary
    var err error
	if c.Args().First() == "" {
		error_handler.ErrorExit("A filename or '-' must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}
	if c.Args().First() == "-" {
		blueprint, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			error_handler.ErrorExit(err)
		}
		create, err = application.CreateFromBytes(cmd.network, blueprint)
		if nil != err {
			error_handler.ErrorExit(err)
		}
	} else {
		create, err = application.Create(cmd.network, c.Args().First())
		if nil != err {
			error_handler.ErrorExit(err)
		}
	}
    table := terminal.NewTable([]string{"Id:",create.EntityId})
    table.Add("Name:", create.EntityDisplayName)
    table.Add("Status:", create.CurrentStatus)
    table.Print()
}
