package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/version"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

type Version struct {
	
}

func NewVersion() (cmd *Version){
	cmd = new(Version)
	return
}

func (cmd *Version) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "version",
		Description: "",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Version) Run(c *cli.Context) {
	version := version.Version()
	fmt.Println(version)
}