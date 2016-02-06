package command_runner

import (
	"github.com/brooklyncentral/brooklyn-cli/command_factory"
	"github.com/brooklyncentral/brooklyn-cli/scope"
	"github.com/codegangsta/cli"
)

type Runner interface {
	RunCmdByName(cmdName string, c *cli.Context) (err error)
	RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) (err error)
}

type ConcreteRunner struct {
	cmdFactory command_factory.Factory
	scope      scope.Scope
}

func NewRunner(scope scope.Scope, cmdFactory command_factory.Factory) (runner ConcreteRunner) {
	runner.cmdFactory = cmdFactory
	runner.scope = scope
	return
}

func (runner ConcreteRunner) RunCmdByName(cmdName string, c *cli.Context) error {
	cmd, err := runner.cmdFactory.GetByCmdName(cmdName)
	if nil != err {
		return err
	}

	cmd.Run(runner.scope, c)
	return nil
}

func (runner ConcreteRunner) RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) error {
	cmd, err := runner.cmdFactory.GetBySubCmdName(cmdName, subCommand)
	if nil != err {
		return err
	}

	cmd.Run(runner.scope, c)
	return nil
}
