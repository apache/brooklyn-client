package commands

import(
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/catalog"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/terminal"
)

type Catalog struct {
	
}

func NewCatalog() (cmd *Catalog){
	cmd = new(Catalog)
	return
}

func (cmd *Catalog) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "catalog",
		Description: "",
		Usage:       "",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Catalog) Run(c *cli.Context) {
	catalog := catalog.Catalog()
	table := terminal.NewTable([]string{"Id", "Name", "Description"})
	for _, app := range catalog {
		table.Add(app.Id, app.Name, app.Description)
	}
	table.Print()
}