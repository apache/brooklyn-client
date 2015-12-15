package app

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/command_runner"
	"strings"
	"os"
    "github.com/brooklyncentral/brooklyn-cli/error_handler"
)

func NewApp(baseName string, cmdRunner command_runner.Runner, metadatas ...command_metadata.CommandMetadata) (app *cli.App) {

	cli.AppHelpTemplate = appHelpTemplate()
	cli.CommandHelpTemplate = commandHelpTemplate()
	app = cli.NewApp()
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
		command.Subcommands = make([]cli.Command,0)
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

COMMANDS:
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
