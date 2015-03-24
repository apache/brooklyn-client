package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/entity_config"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Config struct {
	network *net.Network
}

func NewConfig(network *net.Network) (cmd *Config){
	cmd = new(Config)
	cmd.network = network
	return
}

func (cmd *Config) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "config",
		Description: "Show the config for an application and entity",
		Usage:       "BROOKLYN_NAME config APPLICATION ENTITY",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Config) Run(c *cli.Context) {
	config := entity_config.ConfigCurrentState(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Key", "Value"})
	for key, value := range config {
		table.Add(key, fmt.Sprintf("%v", value))
	}
	table.Print()
}