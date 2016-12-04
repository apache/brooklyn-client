/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package app

import (
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/command_runner"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/urfave/cli"
	"os"
	"strings"
)

type configDefaults struct {
	Name     string
	HelpName string
	Usage    string
	Version  string
}

var appConfig = configDefaults{
	Name:     os.Args[0],
	HelpName: os.Args[0],
	Usage:    "A Brooklyn command line client application",
	Version:  "0.11.0-SNAPSHOT",  // BROOKLYN_VERSION
}

func NewApp(baseName string, cmdRunner command_runner.Runner, metadatas ...command_metadata.CommandMetadata) (app *cli.App) {

	cli.AppHelpTemplate = appHelpTemplate()
	cli.CommandHelpTemplate = commandHelpTemplate()
	app = cli.NewApp()
	app.Name = appConfig.Name
	app.HelpName = appConfig.HelpName
	app.Usage = appConfig.Usage
	app.Version = appConfig.Version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "skipSslChecks",
			Usage: "Skip verification of server's certificate chain and hostname (for use with self-signed certs)",
		},
		cli.StringFlag{
			Name:  "json, j",
			Usage: "Render value as json with json path selector as described at https://github.com/NodePrime/jsonpath. (Experimental, not supported on all commands at present) ",
		},
	}

	app.Commands = []cli.Command{}

	for _, metadata := range metadatas {
		primaryCommand := getCommand(baseName, metadata, cmdRunner)
		app.Commands = append(app.Commands, primaryCommand)
	}
	return
}

func getCommand(baseName string, metadata command_metadata.CommandMetadata, runner command_runner.Runner) cli.Command {
	command := cli.Command{
		Name:        metadata.Name,
		Aliases:     metadata.Aliases,
		ShortName:   metadata.ShortName,
		Description: metadata.Description,
		Usage:       strings.Replace(metadata.Usage, "BROOKLYN_NAME", baseName, -1),
		Action: func(context *cli.Context) {
			err := runner.RunCmdByName(metadata.Name, context)
			if err != nil {
				error_handler.ErrorExit(err)
			}
		},
		Flags:           metadata.Flags,
		SkipFlagParsing: metadata.SkipFlagParsing,
	}

	if nil != metadata.Operands {
		command.Subcommands = make([]cli.Command, 0)
		for _, operand := range metadata.Operands {
			command.Subcommands = append(command.Subcommands, cli.Command{
				Name:            operand.Name,
				Aliases:         operand.Aliases,
				ShortName:       operand.ShortName,
				Description:     operand.Description,
				Usage:           strings.Replace(operand.Usage, "BROOKLYN_NAME", baseName, -1),
				Flags:           operand.Flags,
				SkipFlagParsing: operand.SkipFlagParsing,
				Action:          subCommandAction(command.Name, operand.Name, runner),
			})
			command.Usage = strings.Join([]string{
				command.Usage, "\n... ", operand.Usage, "\t", operand.Description,
			}, "")
		}
	}

	return command
}

func subCommandAction(command string, operand string, runner command_runner.Runner) func(context *cli.Context) {
	return func(context *cli.Context) {
		err := runner.RunSubCmdByName(command, operand, context)
		if err != nil {
			error_handler.ErrorExit(err)
		}
	}
}

func appHelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}{{if or .Author .Email}}

AUTHOR:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}


SCOPES:
   Many commands require a "scope" expression to indicate the target on which they operate. The scope expressions are
   as follows (values in brackets are aliases for the scope):
   - application APP-ID   (app, a) Selects and application, e.g. "br app myapp"
   - entity      ENT-ID   (ent, e) Selects an entity within an application scope, e.g. "br app myapp ent myserver"
   - effector    EFF-ID   (eff, f) Selects an effector of an entity or application, e.g. "br a myapp e myserver eff xyz"
   - config      CONF-KEY (conf, con, c) Selects a configuration key of an entity e.g. "br a myapp e myserver config jmx.agent.mode"
   - activity    ACT-ID   (act, v) Selects an activity of an entity e.g. "br a myapp e myserver act iHG7sq1"


COMMANDS:

   Commands whose description begins with a "*" character are particularly experimental and likely to change in upcoming
   releases.  If not otherwise specified, "SCOPE" below means application or entity scope.  If an entity scope is not
   specified, the application entity is used as a default.

   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Description}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
}

func commandHelpTemplate() string {
	return `NAME:
   {{.Name}} - {{.Description}}
{{with .ShortName}}
ALIAS:
   {{.}}
{{end}}
USAGE:
   {{.Usage}}{{with .Flags}}
OPTIONS:
{{range .}}   {{.}}
{{end}}{{else}}
{{end}}`
}
