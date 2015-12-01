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

	factory.cmdsByName["access"] = commands.NewAccess(network)
	//factory.cmdsByName["activities"] = commands.NewActivities(network)
	factory.cmdsByName["activity"] = commands.NewActivity(network)
	factory.cmdsByName["activity-children"] = commands.NewActivityChildren(network)
	factory.cmdsByName["activity-stream"] = commands.NewActivityStream(network)
	factory.cmdsByName["add-catalog"] = commands.NewAddCatalog(network)
	factory.cmdsByName["add-children"] = commands.NewAddChildren(network)
	factory.cmdsByName["application"] = commands.NewApplication(network)
	factory.cmdsByName["catalog"] = commands.NewCatalog(network)
	factory.cmdsByName["config"] = commands.NewConfig(network)
	factory.cmdsByName["create"] = commands.NewCreate(network)
	factory.cmdsByName["delete"] = commands.NewDelete(network)
	factory.cmdsByName["destroy-policy"] = commands.NewDestroyPolicy(network)
	factory.cmdsByName["effectors"] = commands.NewEffectors(network)
	factory.cmdsByName["entities"] = commands.NewEntities(network)
	factory.cmdsByName["entity-children"] = commands.NewChildren(network)
	factory.cmdsByName["list"] = commands.NewList(network)
    factory.subCommand("list", "application", commands.NewListApplication(network))
	factory.cmdsByName["locations"] = commands.NewLocations(network)
	factory.cmdsByName["login"] = commands.NewLogin(network, config)
	factory.cmdsByName["policies"] = commands.NewPolicies(network)
	factory.cmdsByName["policy"] = commands.NewPolicy(network)
	factory.cmdsByName["rename-entity"] = commands.NewRename(network)
	factory.cmdsByName["sensor"] = commands.NewSensor(network)
	factory.cmdsByName["sensors"] = commands.NewSensors(network)
	factory.cmdsByName["set-config"] = commands.NewSetConfig(network)
	factory.cmdsByName["spec"] = commands.NewSpec(network)
	factory.cmdsByName["start-policy"] = commands.NewStartPolicy(network)
	factory.cmdsByName["stop-policy"] = commands.NewStopPolicy(network)
	factory.cmdsByName["tree"] = commands.NewTree(network)
	factory.cmdsByName["version"] = commands.NewVersion(network)

	return factory
}

func (factory concreteFactory) subCommand(commandName string, subCommandName string, subCommand command.Command)  {
	if nil == factory.subCommands[commandName] {
		factory.subCommands[commandName] = make(map[string]command.Command)
	}
	factory.subCommands[commandName][subCommandName] = subCommand
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
