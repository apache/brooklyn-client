package command_factory

import (
	"errors"
	"sort"
	"github.com/brooklyncentral/brooklyn-cli/command"
	"github.com/brooklyncentral/brooklyn-cli/command_metadata"
	"github.com/brooklyncentral/brooklyn-cli/commands"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/io"
	"strings"
)

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	GetBySubCmdName(cmdName string, subCmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type concreteFactory struct {
	cmdsByName map[string]command.Command
	subCommands map[string]map[string]command.Command
}


func NewFactory(network *net.Network, config *io.Config) (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.subCommands = make(map[string]map[string]command.Command)

	factory.command(commands.NewAccess(network))
	//factory.command(commands.NewActivities(network))
	factory.command(commands.NewActivity(network))
	factory.command(commands.NewActivityChildren(network))
	factory.command(commands.NewActivityStream(network))
	factory.command(commands.NewAddCatalog(network))
	factory.command(commands.NewAddChildren(network))
	factory.command(commands.NewApplication(network))
	factory.command(commands.NewCatalog(network))
	factory.command(commands.NewConfig(network))
	factory.command(commands.NewCreate(network))
	factory.command(commands.NewDelete(network))
	factory.command(commands.NewDestroyPolicy(network))
	factory.command(commands.NewChildren(network))
	listCommand := commands.NewList(network)
	factory.command(listCommand);
    factory.subCommand(listCommand, commands.ListApplicationCommand)
    factory.subCommand(listCommand, commands.ListEffectorCommand)
    factory.subCommand(listCommand, commands.ListEntityCommand)
    factory.subCommand(listCommand, commands.ListSensorCommand)
	factory.command(commands.NewLocations(network))
	factory.command(commands.NewLogin(network, config))
	factory.command(commands.NewPolicies(network))
	factory.command(commands.NewPolicy(network))
	factory.command(commands.NewRename(network))
	factory.command(commands.NewSensor(network))
	factory.command(commands.NewSetConfig(network))
	factory.command(commands.NewSpec(network))
	factory.command(commands.NewStartPolicy(network))
	factory.command(commands.NewStopPolicy(network))
	factory.command(commands.NewTree(network))
	factory.command(commands.NewVersion(network))

	return factory
}


func (factory *concreteFactory) command(command command.Command) {
	factory.cmdsByName[command.Metadata().Name] = command
}

// TODO make this more generic - instead of List use a generic Command type
func (factory concreteFactory) subCommand(listCommand *commands.List, subCommandName string)  {
	if nil == factory.subCommands[listCommand.Metadata().Name] {
		factory.subCommands[listCommand.Metadata().Name] = make(map[string]command.Command)
	}
	factory.subCommands[listCommand.Metadata().Name][subCommandName] = listCommand.SubCommand(subCommandName)
}

func (f concreteFactory) GetByCmdName(cmdName string) (cmd command.Command, err error) {
	cmd, found := f.cmdsByName[cmdName]
	if !found {
		for _, c := range f.cmdsByName {
			if c.Metadata().ShortName == cmdName {
				return c, nil
			}
		}

		err = errors.New(strings.Join([]string{"Command not found:", cmdName}, " "))
	}
	return
}

func (f concreteFactory) GetBySubCmdName(cmdName string, subCmdName string) (cmd command.Command, err error) {

	_, hasPrimary := f.subCommands[cmdName]
	if hasPrimary {
		cmd, found := f.subCommands[cmdName][subCmdName]
		if found {
			return cmd, nil
		}
	}
	return cmd, errors.New(strings.Join([]string{"Command not found:", cmdName, subCmdName}, " "))
}

func (factory concreteFactory) CommandMetadatas() (commands []command_metadata.CommandMetadata) {
	keys := make([]string, 0, len(factory.cmdsByName))
    for key := range factory.cmdsByName {
        keys = append(keys, key)
    }
    sort.Strings(keys)
	
	for _, key := range keys {
		command := factory.cmdsByName[key]
		commands = append(commands, command.Metadata())
	}
	return
}
