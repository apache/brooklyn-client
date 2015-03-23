package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

type Login struct {
	
}

func NewLogin() (cmd *Login){
	cmd = new(Login)
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