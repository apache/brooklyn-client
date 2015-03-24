package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entities"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Spec struct {
	network *net.Network
}

func NewSpec(network *net.Network) (cmd *Spec){
	cmd = new(Spec)
	cmd.network = network
	return
}

func (cmd *Spec) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "spec",
		Description: "Get the YAML spec used to create the entity, if available",
		Usage:       "BROOKLYN_NAME spec APPLICATION ENTITY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Spec) Run(c *cli.Context) {
	spec := entities.Spec(cmd.network, c.Args()[0], c.Args()[1])
	fmt.Println(spec)
}