package command_factory

import(
	"errors"
	"github.com/robertgmoss/brooklyn-cli/command"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/commands"
)

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type concreteFactory struct {
	cmdsByName map[string]command.Command
}

func NewFactory() (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.cmdsByName["login"] = commands.NewLogin()
	factory.cmdsByName["tree"] = commands.NewTree()
	factory.cmdsByName["catalog"] = commands.NewCatalog()
	factory.cmdsByName["version"] = commands.NewVersion()
	factory.cmdsByName["create"] = commands.NewCreate()
	factory.cmdsByName["application"] = commands.NewApplication()
	factory.cmdsByName["sensors"] = commands.NewSensors()
	factory.cmdsByName["sensor"] = commands.NewSensor()
	factory.cmdsByName["effectors"] = commands.NewEffectors()
	factory.cmdsByName["policies"] = commands.NewPolicies()
	factory.cmdsByName["policy"] = commands.NewPolicy()
	factory.cmdsByName["config"] = commands.NewConfig()
	factory.cmdsByName["locations"] = commands.NewLocations()
	return factory
}

func (f concreteFactory) GetByCmdName(cmdName string) (cmd command.Command, err error) {
	cmd, found := f.cmdsByName[cmdName]
	if !found {
		for _, c := range f.cmdsByName {
			if c.Metadata().ShortName == cmdName {
				return c, nil
			}
		}

		err = errors.New("Command not found")
	}
	return
}

func (factory concreteFactory) CommandMetadatas() (commands []command_metadata.CommandMetadata) {
	for _, command := range factory.cmdsByName {
		commands = append(commands, command.Metadata())
	}
	return
}