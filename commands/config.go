package commands

import(
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
		Description: "show the config for an application and entity",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Config) Run(c *cli.Context) {
	config := entity_config.ConfigList(cmd.network, c.Args()[0], c.Args()[1])
	table := terminal.NewTable([]string{"Key", "Value"})
	for _, key := range config {
		value := entity_config.ConfigValue(cmd.network, c.Args()[0], c.Args()[1], key.Name)
		table.Add(key.Name, value)
	}
	table.Print()
}