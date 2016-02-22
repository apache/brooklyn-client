package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/access_control"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Access struct {
	network *net.Network
}

func NewAccess(network *net.Network) (cmd *Access) {
	cmd = new(Access)
	cmd.network = network
	return
}

func (cmd *Access) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "access",
		Description: "Show access control",
		Usage:       "BROOKLYN_NAME access",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Access) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	access, err := access_control.Access(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println("Location Provisioning Allowed:", access.LocationProvisioningAllowed)
}
