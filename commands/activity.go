package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"time"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
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
		Usage:       "BROOKLYN_NAME ENT-SCOPE activity [ ACTIVITYID]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Activity) Run(scope scope.Scope, c *cli.Context) {
	if c.Args().Present() {
		cmd.show(c.Args().First())
	} else {
		cmd.list(scope.Application, scope.Entity)
	}
}

func (cmd *Activity) show(activityId string) {
	activity := activities.Activity(cmd.network, activityId)
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	table.Add(activity.Id, activity.DisplayName,
		time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)

	table.Print()
}

func (cmd *Activity) list(application, entity string) {
	activityList := entities.GetActivities(cmd.network, application, entity)
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	for _, activity := range activityList {
		table.Add(activity.Id, truncate(activity.DisplayName),
			time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), truncate(activity.CurrentStatus))
	}
	table.Print()
}

const truncLimit = 40

func truncate(text string) string {
	if len(text) < truncLimit {
		return text
	}
	return text[0:(truncLimit-3)] + "..."
}