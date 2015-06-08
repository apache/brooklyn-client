package command

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
)

type Command interface {
	Metadata() command_metadata.CommandMetadata
	Run(context *cli.Context)
}
