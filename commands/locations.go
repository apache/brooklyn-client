package commands

import (
	"github.com/apache/brooklyn-client/api/locations"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/codegangsta/cli"
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
		Description: "* List the available locations",
		Usage:       "BROOKLYN_NAME locations",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Locations) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	locationList, err := locations.LocationList(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Spec"})
	for _, location := range locationList {
		table.Add(location.Id, location.Name, location.Spec)
	}
	table.Print()
}
