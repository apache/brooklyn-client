package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/version"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Version struct {
	network *net.Network
}

func NewVersion(network *net.Network) (cmd *Version){
	cmd = new(Version)
	cmd.network = network
	return
}

func (cmd *Version) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "version",
		Description: "Display the version of the connected Brooklyn",
		Usage:       "BROOKLYN_NAME version",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Version) Run(c *cli.Context) {
	version := version.Version(cmd.network)
	fmt.Println(version)
}