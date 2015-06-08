package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entities"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
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
		Name:        "rename-entity",
		Description: "Rename an entity",
		Usage:       "BROOKLYN_NAME rename-entity APPLICATION ENTITY",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Rename) Run(c *cli.Context) {
	rename := entities.Rename(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2])
	fmt.Println(rename)
}
