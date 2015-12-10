package commands

import (
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/api/entity_effectors"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/error_handler"
	"strings"
	"errors"
)

type Invoke struct {
	network *net.Network
}

func NewInvoke(network *net.Network) (cmd *Invoke) {
	cmd = new(Invoke)
	cmd.network = network
	return
}

func (cmd *Invoke) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "invoke",
		Description: "Invoke an effector of an application and entity",
		Usage:       "BROOKLYN_NAME EFF-SCOPE invoke [ parameter-options ]",
		Flags:       []cli.Flag {
			cli.StringSliceFlag{
				Name: "param, P",
				Usage: "Parameter and value separated by '=', e.g. -P x=y",
			},
		},
	}
}

func (cmd *Invoke) Run(scope scope.Scope, c *cli.Context) {

	parms := c.StringSlice("param")
	names, vals, err := extractParams(parms)
	result, err := entity_effectors.TriggerEffector(cmd.network, scope.Application, scope.Entity, scope.Effector, names, vals)
	if nil != err  {
		error_handler.ErrorExit(err)
	} else {
		if "" != result {
			fmt.Println(result)
		}
	}

}

func extractParams(parms []string) ([]string, []string, error) {
	names := make([]string, len(parms))
	vals := make([]string, len(parms))
	for i, parm := range parms {
		index := strings.Index(parm, "=")
		if index < 0 {
			return names, vals, errors.New("Parameter value not provided: " + parm)
		}
		names[i] = parm[0:index]
		vals[i] = parm[index+1:]
	}
	return names, vals, nil
}