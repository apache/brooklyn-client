package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/locations"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/terminal"
	"github.com/brooklyncentral/brooklyn-cli/scope"
)

type Locations struct {
	network *net.Network
}

func NewLocations(network *net.Network) (cmd *Locations) {
	cmd = new(Locations)
	cmd.network = network
	return
}

func (cmd *Locations) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "locations",
		Description: "List the available locations",
		Usage:       "BROOKLYN_NAME [ SCOPE ] locations",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Locations) Run(scope scope.Scope, c *cli.Context) {
	locationList := locations.LocationList(cmd.network)
	table := terminal.NewTable([]string{"Id", "Name", "Spec"})
	for _, location := range locationList {
		table.Add(location.Id, location.Name, location.Spec)
	}
	table.Print()
}
