package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/locations"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
)

type Locations struct {
	
}

func NewLocations() (cmd *Locations){
	cmd = new(Locations)
	return
}

func (cmd *Locations) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "locations",
		Description: "",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Locations) Run(c *cli.Context) {
	locationList := locations.LocationList()
	table := terminal.NewTable([]string{"Id", "Name", "Spec"})
	for _, location := range locationList {
		table.Add(location.Id, location.Name, location.Spec)
	}
	table.Print()
}