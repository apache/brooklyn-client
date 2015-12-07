package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"time"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Activity struct {
	network *net.Network
}

func NewActivity(network *net.Network) (cmd *Activity) {
	cmd = new(Activity)
	cmd.network = network
	return
}

func (cmd *Activity) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "activity",
		Description: "Show the activity for an entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] activity ACTIVITYID",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Activity) Run(scope scope.Scope, c *cli.Context) {
	activity := activities.Activity(cmd.network, c.Args().First())
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)

	table.Print()
}
