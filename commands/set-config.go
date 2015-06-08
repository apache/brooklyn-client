package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_config"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type SetConfig struct {
	network *net.Network
}

func NewSetConfig(network *net.Network) (cmd *SetConfig) {
	cmd = new(SetConfig)
	cmd.network = network
	return
}

func (cmd *SetConfig) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "set-config",
		Description: "Set config for an entity",
		Usage:       "BROOKLYN_NAME set-config APPLICATION ENTITY CONFIG_KEY VALUE",
		Flags:       []cli.Flag{},
	}
}

func (cmd *SetConfig) Run(c *cli.Context) {
	response := entity_config.SetConfig(cmd.network, c.Args()[0], c.Args()[1], c.Args()[2], c.Args()[3])
	fmt.Println(response)
}
