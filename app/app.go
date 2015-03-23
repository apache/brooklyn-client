package app

import (
	"fmt"
	"strings"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/command_runner"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

func NewApp(baseName string, cmdRunner command_runner.Runner, metadatas ...command_metadata.CommandMetadata) (app *cli.App) {
	
	cli.AppHelpTemplate = appHelpTemplate()
	cli.CommandHelpTemplate = commandHelpTemplate()
	app = cli.NewApp()
	app.Commands = []cli.Command{}

	for _, metadata := range metadatas {
		app.Commands = append(app.Commands, getCommand(baseName, metadata, cmdRunner))
	}
	return
}

func getCommand(baseName string, metadata command_metadata.CommandMetadata, runner command_runner.Runner) cli.Command {
	return cli.Command{
		Name:        metadata.Name,
		ShortName:   metadata.ShortName,
		Description: metadata.Description,
		Usage:       strings.Replace(metadata.Usage, "BROOKLYN_NAME", baseName, -1),
		Action: func(context *cli.Context) {
			err := runner.RunCmdByName(metadata.Name, context)
			if err != nil {
				fmt.Println(err)
			}
		},
		Flags:           metadata.Flags,
		SkipFlagParsing: metadata.SkipFlagParsing,
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