package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/activities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type ActivityStreamEnv struct {
	network *net.Network
}

type ActivityStreamStderr struct {
	network *net.Network
}

type ActivityStreamStdin struct {
	network *net.Network
}

type ActivityStreamStdout struct {
	network *net.Network
}

func NewActivityStreamEnv(network *net.Network) (cmd *ActivityStreamEnv) {
	cmd = new(ActivityStreamEnv)
	cmd.network = network
	return
}

func NewActivityStreamStderr(network *net.Network) (cmd *ActivityStreamStderr) {
	cmd = new(ActivityStreamStderr)
	cmd.network = network
	return
}

func NewActivityStreamStdin(network *net.Network) (cmd *ActivityStreamStdin) {
	cmd = new(ActivityStreamStdin)
	cmd.network = network
	return
}

func NewActivityStreamStdout(network *net.Network) (cmd *ActivityStreamStdout) {
	cmd = new(ActivityStreamStdout)
	cmd.network = network
	return
}

func (cmd *ActivityStreamEnv) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "env",
		Description: "Show the ENV stream for a given activity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] env",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStderr) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stderr",
		Description: "Show the STDERR stream for a given activity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] stderr",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStdin) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stdin",
		Description: "Show the STDIN stream for a given activity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] stdin",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStdout) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stdout",
		Description: "Show the STDOUT stream for a given activity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] stdout",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamEnv) Run(scope scope.Scope, c *cli.Context) {
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "env")
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStderr) Run(scope scope.Scope, c *cli.Context) {
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stderr")
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStdin) Run(scope scope.Scope, c *cli.Context) {
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stdin")
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStdout) Run(scope scope.Scope, c *cli.Context) {
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stdout")
    if nil != err {
        error_handler.ErrorExit(err)
    }
	fmt.Println(activityStream)
}
