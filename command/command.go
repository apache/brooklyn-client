package command

import (
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/codegangsta/cli"
)

type Command interface {
	Metadata() command_metadata.CommandMetadata
	Run(context *cli.Context)
}