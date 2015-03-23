package commands

import(
	"strings"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/application"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Application struct {
	network *net.Network
}

func NewApplication(network *net.Network) (cmd *Application){
	cmd = new(Application)
	cmd.network = network
	return
}

func (cmd *Application) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Description: "show the status and location of a running application",
		Usage:       "BROOKLYN_NAME application APPLICATION",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Application) Run(c *cli.Context) {
	application := application.Application(cmd.network, c.Args()[0])
	
	table := terminal.NewTable([]string{"Name", "Id", "Status", "Location"})
	table.Add(application.Spec.Name, application.Id, string(application.Status), strings.Join(application.Spec.Locations, ", "))
	table.Print()
}