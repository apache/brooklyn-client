package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
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
		Description: "Deploy a new brooklyn application from the supplied YAML",
		Usage:       "BROOKLYN_NAME deploy <FILEPATH|->",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Deploy) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().First() == "" {
		error_handler.ErrorExit("A filename or '-' must be provided as the first argument",error_handler.CLIUsageErrorExitCode)
	}
	if c.Args().First() == "-" {
		blueprint, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			error_handler.ErrorExit(err)
		}
		create, err := application.CreateFromBytes(cmd.network, blueprint)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		fmt.Println(create)
	} else {
		create, err := application.Create(cmd.network, c.Args().First())
		if nil != err {
			error_handler.ErrorExit(err)
		}
		fmt.Println(create)
	}
}
