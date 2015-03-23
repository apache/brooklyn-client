package commands

import(
	"strings"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/application"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
)

type Application struct {
	
}

func NewApplication() (cmd *Application){
	cmd = new(Application)
	return
}

func (cmd *Application) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Description: "",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Application) Run(c *cli.Context) {
	application := application.Application(c.Args()[0])
	
	table := terminal.NewTable([]string{"Name", "Id", "Status", "Location"})
	table.Add(application.Spec.Name, application.Id, string(application.Status), strings.Join(application.Spec.Locations, ", "))
	table.Print()
}