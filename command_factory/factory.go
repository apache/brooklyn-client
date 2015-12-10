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

	factory.simpleCommand(commands.NewAccess(network))
	//factory.command(commands.NewActivities(network))
	factory.simpleCommand(commands.NewActivity(network))
	factory.simpleCommand(commands.NewActivityChildren(network))
	factory.simpleCommand(commands.NewActivityStream(network))
	factory.simpleCommand(commands.NewAddCatalog(network))
	factory.simpleCommand(commands.NewAddChildren(network))
	factory.simpleCommand(commands.NewApplication(network))
	//factory.simpleCommand(commands.NewApplications(network))
	factory.simpleCommand(commands.NewCatalog(network))
    factory.simpleCommand(commands.NewChildren(network))
    factory.simpleCommand(commands.NewConfig(network))
    factory.simpleCommand(commands.NewDeploy(network))
    factory.simpleCommand(commands.NewDelete(network))
    factory.simpleCommand(commands.NewDestroyPolicy(network))
    factory.simpleCommand(commands.NewEffector(network))
    factory.simpleCommand(commands.NewEntity(network))
    // NewList below is not used but we retain the code as an example of how to do a super command.
    //	factory.superCommand(commands.NewList(network))
	factory.simpleCommand(commands.NewLocations(network))
	factory.simpleCommand(commands.NewLogin(network, config))
	factory.simpleCommand(commands.NewPolicies(network))
	factory.simpleCommand(commands.NewPolicy(network))
	factory.simpleCommand(commands.NewRename(network))
	factory.simpleCommand(commands.NewSensor(network))
	factory.simpleCommand(commands.NewSetConfig(network))
	factory.simpleCommand(commands.NewSpec(network))
	factory.simpleCommand(commands.NewStartPolicy(network))
	factory.simpleCommand(commands.NewStopPolicy(network))
	factory.simpleCommand(commands.NewTree(network))
	factory.simpleCommand(commands.NewVersion(network))

	return factory
}


func (factory *concreteFactory) simpleCommand(cmd command.Command) {
	factory.cmdsByName[cmd.Metadata().Name] = cmd
}

func (factory *concreteFactory) superCommand(cmd command.SuperCommand)  {

    factory.simpleCommand(cmd)

	if nil == factory.subCommands[cmd.Metadata().Name] {
		factory.subCommands[cmd.Metadata().Name] = make(map[string]command.Command)
	}

	for _, sub := range cmd.SubCommandNames() {
		factory.subCommands[cmd.Metadata().Name][sub] = cmd.SubCommand(sub)
	}
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
