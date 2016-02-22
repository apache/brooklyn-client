package command

import (
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Command interface {
	Metadata() command_metadata.CommandMetadata
	Run(scope scope.Scope, context *cli.Context)
}
