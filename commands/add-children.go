package commands

import (
	"github.com/apache/brooklyn-client/api/entities"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/codegangsta/cli"
	"time"
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
		Description: "* Add a child or children to this entity from the supplied YAML",
		Usage:       "BROOKLYN_NAME SCOPE add-children FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *AddChildren) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	activity, err := entities.AddChildren(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)

	table.Print()
}
