package commands

import(
	"time"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entities"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Activities struct {
	network *net.Network
}

func NewActivities(network *net.Network) (cmd *Activities){
	cmd = new(Activities)
	cmd.network = network
	return
}

func (cmd *Activities) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "activities",
		Description: "show the activities for an entity",
		Usage:       "BROOKLYN_NAME activities APPLICATION ENTITY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Activities) Run(c *cli.Context) {
	activityList := entities.ActivityList(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Task", "Submitted", "Status"})
	for _, activity := range activityList {
		table.Add(activity.DisplayName, time.Unix(activity.SubmitTimeUtc / 1000, 0).Format(time.UnixDate), activity.CurrentStatus)
	}
	table.Print()
}