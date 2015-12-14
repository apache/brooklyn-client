package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_config"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/brooklyncentral/brooklyn-cli/scope"
    "os"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

type Config struct {
	network *net.Network
}

func NewConfig(network *net.Network) (cmd *Config) {
	cmd = new(Config)
	cmd.network = network
	return
}

func (cmd *Config) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "config",
		Description: "Show the config for an application and entity",
		Usage:       "BROOKLYN_NAME [ SCOPE ] config",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Config) Run(scope scope.Scope, c *cli.Context) {

    if c.Args().Present() {
        config, err := entity_config.ConfigValue(cmd.network, scope.Application, scope.Entity, c.Args().First())
        if nil != err {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
        fmt.Println(config)

    } else {
        config, err := entity_config.ConfigCurrentState(cmd.network, scope.Application, scope.Entity)
        if nil != err {
            error_handler.ErrorExit(err)
        }
        table := terminal.NewTable([]string{"Key", "Value"})
        for key, value := range config {
            table.Add(key, fmt.Sprintf("%v", value))
        }
        table.Print()
    }
}
