package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Login struct {
	network *net.Network
}

func NewLogin(network *net.Network) (cmd *Login){
	cmd = new(Login)
	cmd.network = network
	return
}

func (cmd *Login) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "login",
		Description: "",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Login) Run(c *cli.Context) {
	// do login stuff here
	fmt.Println("logging in...")
}