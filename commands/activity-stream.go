package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type ActivityStream struct {
	network *net.Network
}

func NewActivityStream(network *net.Network) (cmd *ActivityStream) {
	cmd = new(ActivityStream)
	cmd.network = network
	return
}

func (cmd *ActivityStream) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "activity-stream",
		Description: "Show the stream for a given activity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] activity-stream STREAM_ID",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStream) Run(scope scope.Scope, c *cli.Context) {
	activityStream := activities.ActivityStream(cmd.network, scope.Activity, c.Args().First())
	fmt.Println(activityStream)
}
