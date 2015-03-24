package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/application"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
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
		Description: "Delete a brooklyn application",
		Usage:       "BROOKLYN_NAME create FILEPATH",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Delete) Run(c *cli.Context) {
	del := application.Delete(cmd.network, c.Args()[0])
	fmt.Println(del)
}
