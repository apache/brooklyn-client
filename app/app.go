package app

import (
	"fmt"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/command_runner"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/codegangsta/cli"
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
	Version:  "0.9.0",
}

func NewApp(baseName string, cmdRunner command_runner.Runner, metadatas ...command_metadata.CommandMetadata) (app *cli.App) {

	cli.AppHelpTemplate = appHelpTemplate()
	cli.CommandHelpTemplate = commandHelpTemplate()
	app = cli.NewApp()
	app.Name = appConfig.Name
	app.HelpName = appConfig.HelpName
	app.Usage = appConfig.Usage
	app.Version = appConfig.Version

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
				Usage:           operand.Usage,
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
			fmt.Fprintln(os.Stderr, err)
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
