package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/version"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Version struct {
	network *net.Network
}

func NewVersion(network *net.Network) (cmd *Version) {
	cmd = new(Version)
	cmd.network = network
	return
}

func (cmd *Version) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "version",
		Description: "Display the version of the connected Brooklyn",
		Usage:       "BROOKLYN_NAME version",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Version) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	version, err := version.Version(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(version.Version)
}
