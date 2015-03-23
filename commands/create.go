package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/application"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

type Create struct {
	
}

func NewCreate() (cmd *Create){
	cmd = new(Create)
	return
}

func (cmd *Create) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "create",
		Description: "create a new brooklyn application from the supplied YAML",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Create) Run(c *cli.Context) {
	create := application.Create(c.Args()[0])
	fmt.Println(create)
}