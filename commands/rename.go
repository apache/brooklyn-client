package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/entities"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Rename struct {
	network *net.Network
}

func NewRename(network *net.Network) (cmd *Rename) {
	cmd = new(Rename)
	cmd.network = network
	return
}

func (cmd *Rename) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "rename",
		Description: "Rename an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE rename NEW_NAME",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Rename) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	rename, err := entities.Rename(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(rename)
}
