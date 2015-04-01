package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/access_control"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
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

func (cmd *Access) Run(c *cli.Context) {
	access := access_control.Access(cmd.network)
	fmt.Println("Location Provisioning Allowed:", access.LocationProvisioningAllowed)
}
