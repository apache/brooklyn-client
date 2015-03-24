package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/activities"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"time"
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
		Usage:       "BROOKLYN_NAME activity ACTIVITY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Activity) Run(c *cli.Context) {
	activity := activities.Activity(cmd.network, c.Args()[0])
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	table.Add(activity.Id, activity.DisplayName, time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), activity.CurrentStatus)

	table.Print()
}
