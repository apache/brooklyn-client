package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/api/entities"
	"github.com/apache/brooklyn-client/api/entity_sensors"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/apache/brooklyn-client/terminal"
	"github.com/codegangsta/cli"
	"os"
)

type Entity struct {
	network *net.Network
}

func NewEntity(network *net.Network) (cmd *Entity) {
	cmd = new(Entity)
	cmd.network = network
	return
}

func (cmd *Entity) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "entity",
		Aliases:     []string{"entities", "ent", "ents"},
		Description: "Show the entities of an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE entity [ENTITYID]",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "children, c",
				Usage: "List children of the entity",
			},
		},
	}
}

func (cmd *Entity) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.NumFlags() > 0 && c.FlagNames()[0] == "children" {
		cmd.listentity(scope.Application, c.StringSlice("children")[0])
	} else {
		if c.Args().Present() {
			cmd.show(scope.Application, c.Args().First())
		} else {
			if scope.Entity == scope.Application {
				cmd.listapp(scope.Application)
			} else {
				cmd.listentity(scope.Application, scope.Entity)
			}
		}
	}
}

const serviceStateSensor = "service.state"
const serviceIsUp = "service.isUp"

func (cmd *Entity) show(application, entity string) {
	summary, err := entities.GetEntity(cmd.network, application, entity)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	table := terminal.NewTable([]string{"Id:", summary.Id})
	table.Add("Name:", summary.Name)
	configState, err := entity_sensors.CurrentState(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	if serviceState, ok := configState[serviceStateSensor]; ok {
		table.Add("Status:", fmt.Sprintf("%v", serviceState))
	}
	if serviceIsUp, ok := configState[serviceIsUp]; ok {
		table.Add("ServiceUp:", fmt.Sprintf("%v", serviceIsUp))
	}
	table.Add("Type:", summary.Type)
	table.Add("CatalogItemId:", summary.CatalogItemId)
	table.Print()
}

func (cmd *Entity) listapp(application string) {
	entitiesList, err := entities.EntityList(cmd.network, application)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}

func (cmd *Entity) listentity(application string, entity string) {
	entitiesList, err := entities.Children(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}

	table := terminal.NewTable([]string{"Id", "Name", "Type"})
	for _, entityitem := range entitiesList {
		table.Add(entityitem.Id, entityitem.Name, entityitem.Type)
	}
	table.Print()
}
