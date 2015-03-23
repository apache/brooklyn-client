package main

import (
	"os"
	//"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/app"
	"github.com/robertgmoss/brooklyn-cli/command_factory"
	"github.com/robertgmoss/brooklyn-cli/command_runner"
)

func main() {
	cmdFactory := command_factory.NewFactory()
	cmdRunner := command_runner.NewRunner(cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(cmdRunner, metaDatas...)
	theApp.Run(os.Args)
}
