package commands

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/api/application"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/codegangsta/cli"
)

type Delete struct {
	network *net.Network
}

func NewDelete(network *net.Network) (cmd *Delete) {
	cmd = new(Delete)
	cmd.network = network
	return
}

func (cmd *Delete) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "delete",
		Description: "* Delete (expunge) a brooklyn application",
		Usage:       "BROOKLYN_NAME SCOPE delete",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Delete) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	del, err := application.Delete(cmd.network, scope.Application)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(del)
}
