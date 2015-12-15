package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"strconv"
	"time"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"fmt"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
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
		Aliases:     []string{"activities","act","acts"},
		Description: "Show the activity for an application / entity",
		Usage:       "BROOKLYN_NAME SCOPE activity [ ACTIVITYID]",
		Flags:       []cli.Flag{
			cli.StringSliceFlag{
				Name: "children, c",
				Usage: "List children of the activity",
			},
		},
	}
}

func (cmd *Activity) Run(scope scope.Scope, c *cli.Context) {
    if err := net.VerifyLoginURL(cmd.network); err != nil {
        error_handler.ErrorExit(err)
    }
	if c.NumFlags() > 0 && c.FlagNames()[0] == "children" {
		cmd.listchildren(c.StringSlice("children")[0])
	} else {
		if c.Args().Present() {
			cmd.show(c.Args().First())
		} else {
			if scope.Activity == "" {
				cmd.list(scope.Application, scope.Entity)
			} else {
				cmd.listchildren(scope.Activity)
			}
		}
	}
}

func (cmd *Activity) show(activityId string) {
	activity, err := activities.Activity(cmd.network, activityId)
    if nil != err {
        error_handler.ErrorExit(err)
    }
	
	table := terminal.NewTable([]string{"Id:", activity.Id})
	table.Add("DisplayName:", activity.DisplayName)
	table.Add("Description:", activity.Description)
	table.Add("EntityId:", activity.EntityId)
	table.Add("EntityDisplayName:", activity.EntityDisplayName)
	table.Add("Submitted:", time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate))
	table.Add("Started:", time.Unix(activity.StartTimeUtc/1000, 0).Format(time.UnixDate))
	table.Add("Ended:", time.Unix(activity.EndTimeUtc/1000, 0).Format(time.UnixDate))
	table.Add("CurrentStatus:", activity.CurrentStatus)
	table.Add("IsError:", strconv.FormatBool(activity.IsError))
	table.Add("IsCancelled:", strconv.FormatBool(activity.IsCancelled))
	table.Add("SubmittedByTask:", activity.SubmittedByTask.Metadata.Id)
	if activity.Streams["stdin"].Metadata.Size > 0 || 
	   activity.Streams["stdout"].Metadata.Size > 0 ||
	   activity.Streams["stderr"].Metadata.Size > 0  ||
	   activity.Streams["env"].Metadata.Size > 0 {
		  table.Add("Streams:", fmt.Sprintf("stdin: %d, stdout: %d, stderr: %d, env %d", 
			 activity.Streams["stdin"].Metadata.Size, 
			 activity.Streams["stdout"].Metadata.Size, 
			 activity.Streams["stderr"].Metadata.Size, 
			 activity.Streams["env"].Metadata.Size))
	} else {
		table.Add("Streams:","")
	}
	table.Add("DetailedStatus:", fmt.Sprintf("\"%s\"",activity.DetailedStatus))
	table.Print()
}

func (cmd *Activity) list(application, entity string) {
	activityList, err := entities.GetActivities(cmd.network, application, entity)
    if nil != err {
        error_handler.ErrorExit(err)
    }
	table := terminal.NewTable([]string{"Id", "Task", "Submitted", "Status"})
	for _, activity := range activityList {
		table.Add(activity.Id, truncate(activity.DisplayName),
			time.Unix(activity.SubmitTimeUtc/1000, 0).Format(time.UnixDate), truncate(activity.CurrentStatus))
	}
	table.Print()
}

func (cmd *Activity) listchildren(activity string) {
	activityList, err := activities.ActivityChildren(cmd.network, activity)
    if nil != err {
        error_handler.ErrorExit(err)
    }
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