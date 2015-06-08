package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/io"
)

type Login struct {
	network *net.Network
	config *io.Config
}

func NewLogin(network *net.Network, config *io.Config) (cmd *Login) {
	cmd = new(Login)
	cmd.network = network
	cmd.config = config
	return
}

func (cmd *Login) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "login",
		Description: "Login to brooklyn",
		Usage:       "BROOKLYN_NAME login URL USER PASSWORD",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Login) Run(c *cli.Context) {
	cmd.network.BrooklynUrl = c.Args()[0]
	cmd.network.BrooklynUser = c.Args()[1]
	cmd.network.BrooklynPass = c.Args()[2]
	
	if cmd.config.Map == nil {
		cmd.config.Map = make(map[string]interface{})
	}
	// now persist these credentials to the yaml file
	auth, ok := cmd.config.Map["auth"].(map[string]interface{})
	if !ok {
		auth = make(map[string]interface{})
		cmd.config.Map["auth"] = auth
	}
	
	auth[cmd.network.BrooklynUrl] = map[string]string{
		"username": cmd.network.BrooklynUser,
		"password": cmd.network.BrooklynPass,
	}
	
	cmd.config.Map["target"] = cmd.network.BrooklynUrl
	cmd.config.Write()
}
