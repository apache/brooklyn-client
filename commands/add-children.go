package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"time"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type AddChildren struct {
	network *net.Network
}

func NewAddChildren(network *net.Network) (cmd *AddChildren) {
	cmd = new(AddChildren)
	cmd.network = network
	return
}

func (cmd *AddChildren) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "add-children",
		Description: "Add a child or children to this entity from the supplied YAML",
		Usage:       "BROOKLYN_NAME [ SCOPE ] add-children FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddChildren) Run(scope scope.Scope, c *cli.Context) {
	activity, err := entities.AddChildren(cmd.network, scope.Application, scope.Entity, c.Args().First())
    if nil != err {
        error_handler.ErrorExit(err)
    }
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)

	table.Print()
}
