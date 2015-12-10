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
		Aliases:     []string{"applications","app","apps"},
		Description: "Show the status and location of running applications",
		Usage:       "BROOKLYN_NAME application [APP]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Application) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().Present() {
		cmd.show(c.Args().First())
	} else {
		cmd.list()
	}	
}

func (cmd *Application) show(appName string) {
	application := application.Application(cmd.network, appName)
	table := terminal.NewTable([]string{"Id:", application.Id})
	table.Add("Name:", application.Spec.Name)
	table.Add("Status:", string(application.Status))
//	table.Add("Service Up:")
	table.Add("Type:", application.Spec.Type)
	table.Add("LocationId:", strings.Join(application.Spec.Locations, ", "))
//	table.Add("Location:")
//	table.Add("LocationType:")
	table.Print()
}

func (cmd *Application) list() {
	applications := application.Applications(cmd.network)
	table := terminal.NewTable([]string{"Id", "Name", "Status", "Location"})
	for _, app := range applications {
		table.Add(app.Id, app.Spec.Name, string(app.Status), strings.Join(app.Spec.Locations, ", "))
	}
	table.Print()
}