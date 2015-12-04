package commands

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/io"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/codegangsta/cli"
	"os"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Login struct {
	network *net.Network
	config  *io.Config
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
		Usage:       "BROOKLYN_NAME [ SCOPE ] login URL [USER PASSWORD]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Login) Run(scope scope.Scope, c *cli.Context) {
	defer func() {
		if str := recover(); str != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", str)
			os.Exit(1)
		}
	}()
	if !c.Args().Present() {
		panic("A URL must be provided as the first argument")
	}

	// If an argument was not supplied, it is set to empty string
	cmd.network.BrooklynUrl = c.Args().Get(0)
	cmd.network.BrooklynUser = c.Args().Get(1)
	cmd.network.BrooklynPass = c.Args().Get(2)
	if cmd.network.BrooklynUser != "" && cmd.network.BrooklynPass == "" {
		panic("If a username is provided, a password must also be provided")
	}

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
