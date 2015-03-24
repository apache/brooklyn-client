package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/locations"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/terminal"
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
		Usage:       "BROOKLYN_NAME locations",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Locations) Run(c *cli.Context) {
	locationList := locations.LocationList(cmd.network)
	table := terminal.NewTable([]string{"Id", "Name", "Spec"})
	for _, location := range locationList {
		table.Add(location.Id, location.Name, location.Spec)
	}
	table.Print()
}
