package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strings"
)

type Applications struct {
	network *net.Network
}

func NewApplications(network *net.Network) (cmd *Applications) {
	cmd = new(Applications)
	cmd.network = network
	return
}

func (cmd *Applications) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "applications",
		Description: "Show the status and location of running applications",
		Usage:       "BROOKLYN_NAME applications",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Applications) Run(c *cli.Context) {
	applications := application.Applications(cmd.network)

	table := terminal.NewTable([]string{"Id", "Name", "Status", "Location"})
	for _, app := range applications {
		table.Add(app.Id, app.Spec.Name, string(app.Status), strings.Join(app.Spec.Locations, ", "))
	}
	table.Print()
}
