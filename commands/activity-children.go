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

type ActivityChildren struct {
	network *net.Network
}

func NewActivityChildren(network *net.Network) (cmd *ActivityChildren) {
	cmd = new(ActivityChildren)
	cmd.network = network
	return
}

func (cmd *ActivityChildren) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "activity-children",
		Description: "Show the child activities for an entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] activity-children",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityChildren) Run(scope scope.Scope, c *cli.Context) {
	activityList := activities.ActivityChildren(cmd.network, scope.Activity)
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	for _, activity := range activityList {
		table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)
	}
	table.Print()
}
