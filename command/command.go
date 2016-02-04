package command

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Command interface {
	Metadata() command_metadata.CommandMetadata
	Run(scope scope.Scope, context *cli.Context)
}
