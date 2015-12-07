package commands

import (
	"github.com/codegangsta/cli"
	//"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	//"github.com/brooklyncentral/brooklyn-cli/terminal"
	//"time"
)

type Activities struct {
	network *net.Network
}

func NewActivities(network *net.Network) (cmd *Activities) {
	cmd = new(Activities)
	cmd.network = network
	return
}

func (cmd *Activities) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "activities",
		Description: "Show the activities for an entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] activities",
		Flags:       []cli.Flag{},
	}
}
//WIP
//func (cmd *Activities) Run(scope scope.Scope, c *cli.Context) {
//	activityList := activities.ActivityList(cmd.network, c.Args().First(), c.Args()[1])
//	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
//	for _, activity := range activityList {
//		table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)
//	}
//	table.Print()
//}