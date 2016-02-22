package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/application"
	"github.com/apache/brooklyn-client/api/entities"
	"github.com/apache/brooklyn-client/api/entity_sensors"
	"github.com/apache/brooklyn-client/api/locations"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/codegangsta/cli"
	"strings"
)

type Application struct {
	network *net.Network
}

func NewApplication(network *net.Network) (cmd *Application) {
	cmd = new(Application)
	cmd.network = network
	return
}

func (cmd *Application) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "application",
		Aliases:     []string{"applications", "app", "apps"},
		Description: "Show the status and location of running applications",
		Usage:       "BROOKLYN_NAME application [APP]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Application) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.Args().Present() {
		cmd.show(c.Args().First())
	} else {
		cmd.list()
	}
}

const serviceIsUpStr = "service.isUp"

func (cmd *Application) show(appName string) {
	application, err := application.Application(cmd.network, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	entity, err := entities.GetEntity(cmd.network, appName, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	state, err := entity_sensors.CurrentState(cmd.network, appName, appName)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	location, err := locations.GetLocation(cmd.network, application.Spec.Locations[0])
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id:", application.Id})
	table.Add("Name:", application.Spec.Name)
	table.Add("Status:", string(application.Status))
	if serviceUp, ok := state[serviceIsUpStr]; ok {
		table.Add("ServiceUp:", fmt.Sprintf("%v", serviceUp))
	}
	table.Add("Type:", application.Spec.Type)
	table.Add("CatalogItemId:", entity.CatalogItemId)
	table.Add("LocationId:", strings.Join(application.Spec.Locations, ", "))
	table.Add("LocationName:", location.Name)
	table.Add("LocationSpec:", location.Spec)
	table.Add("LocationType:", location.Type)
	table.Print()
}

func (cmd *Application) list() {
	applications, err := application.Applications(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Status", "Location"})
	for _, app := range applications {
		table.Add(app.Id, app.Spec.Name, string(app.Status), strings.Join(app.Spec.Locations, ", "))
	}
	table.Print()
}
