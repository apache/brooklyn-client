package command_runner

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/command_factory"
)

type Runner interface {
	RunCmdByName(cmdName string, c *cli.Context) (err error)
}

type ConcreteRunner struct {
	cmdFactory command_factory.Factory
}

func NewRunner(cmdFactory command_factory.Factory) (runner ConcreteRunner) {
	runner.cmdFactory = cmdFactory
	return
}

func (runner ConcreteRunner) RunCmdByName(cmdName string, c *cli.Context) error {
	cmd, err := runner.cmdFactory.GetByCmdName(cmdName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	cmd.Run(c)
	return nil
}