package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/version"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
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
    fmt.Println(version)
}
