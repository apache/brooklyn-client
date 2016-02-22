package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/entity_policies"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type DestroyPolicy struct {
	network *net.Network
}

func NewDestroyPolicy(network *net.Network) (cmd *DestroyPolicy) {
	cmd = new(DestroyPolicy)
	cmd.network = network
	return
}

func (cmd *DestroyPolicy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "destroy-policy",
		Description: "Destroy a policy",
		Usage:       "BROOKLYN_NAME SCOPE destroy-policy POLICY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *DestroyPolicy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	spec, err := entity_policies.DestroyPolicy(cmd.network, scope.Application, scope.Entity, c.Args().First())
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(spec)
}
