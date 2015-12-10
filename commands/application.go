package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Application struct {
	network *net.Network
}

func NewApplication(network *net.Network) (cmd *Application) {
	cmd = new(Application)
	cmd.network = network
	return
}

func (cmd *Application) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Aliases:     []string{"app"},
		Description: "Show the status and location of a running application",
		Usage:       "BROOKLYN_NAME application APP",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Application) Run(scope scope.Scope, c *cli.Context) {
	application := application.Application(cmd.network, c.Args().First())

	table := terminal.NewTable([]string{"Name", "Id", "Status", "Location"})
	table.Add(application.Spec.Name, application.Id, string(application.Status), strings.Join(application.Spec.Locations, ", "))
	table.Print()
}
