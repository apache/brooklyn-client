package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
    "github.com/brooklyncentral/brooklyn-cli/command_metadata"
)

type ListApplicationSubCommand struct {
	network *net.Network
}

func NewListApplication(network *net.Network) (cmd *ListApplicationSubCommand) {
	cmd = new(ListApplicationSubCommand)
	cmd.network = network
	return
}

func (cmd *ListApplicationSubCommand) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Description: "Show the status and location of running applications",
		Usage:       "list application",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ListApplicationSubCommand) Run(c *cli.Context) {
	applications := application.Applications(cmd.network)

	table := terminal.NewTable([]string{"Id", "Name", "Status", "Location"})
	for _, app := range applications {
		table.Add(app.Id, app.Spec.Name, string(app.Status), strings.Join(app.Spec.Locations, ", "))
	}
	table.Print()
}
