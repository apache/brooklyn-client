package command_factory

import(
	"errors"
	"github.com/robertgmoss/brooklyn-cli/command"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/commands"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type concreteFactory struct {
	cmdsByName map[string]command.Command
}

func NewFactory(network *net.Network) (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.cmdsByName["login"] = commands.NewLogin(network)
	factory.cmdsByName["tree"] = commands.NewTree(network)
	factory.cmdsByName["entities"] = commands.NewEntities(network)
	factory.cmdsByName["catalog"] = commands.NewCatalog(network)
	factory.cmdsByName["version"] = commands.NewVersion(network)
	factory.cmdsByName["create"] = commands.NewCreate(network)
	factory.cmdsByName["application"] = commands.NewApplication(network)
	factory.cmdsByName["sensors"] = commands.NewSensors(network)
	factory.cmdsByName["sensor"] = commands.NewSensor(network)
	factory.cmdsByName["effectors"] = commands.NewEffectors(network)
	factory.cmdsByName["policies"] = commands.NewPolicies(network)
	factory.cmdsByName["policy"] = commands.NewPolicy(network)
	factory.cmdsByName["config"] = commands.NewConfig(network)
	factory.cmdsByName["locations"] = commands.NewLocations(network)
	factory.cmdsByName["activities"] = commands.NewActivities(network)
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