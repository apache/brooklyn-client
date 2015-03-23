package app

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/command_runner"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
)

func NewApp(cmdRunner command_runner.Runner, metadatas ...command_metadata.CommandMetadata) (app *cli.App) {
	app = cli.NewApp()
	app.Commands = []cli.Command{}

	for _, metadata := range metadatas {
		app.Commands = append(app.Commands, getCommand(metadata, cmdRunner))
	}
	return
}

func getCommand(metadata command_metadata.CommandMetadata, runner command_runner.Runner) cli.Command {
	return cli.Command{
		Name:        metadata.Name,
		ShortName:   metadata.ShortName,
		Description: metadata.Description,
		Usage:       metadata.Usage,
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