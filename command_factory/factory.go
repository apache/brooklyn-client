package command_factory

import (
	"errors"
	"sort"
	"github.com/robertgmoss/brooklyn-cli/command"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/commands"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/io"
)

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type concreteFactory struct {
	cmdsByName map[string]command.Command
}

func NewFactory(network *net.Network, config *io.Config) (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.cmdsByName["login"] = commands.NewLogin(network, config)
	factory.cmdsByName["tree"] = commands.NewTree(network)
	factory.cmdsByName["entities"] = commands.NewEntities(network)
	factory.cmdsByName["entity-children"] = commands.NewChildren(network)
	factory.cmdsByName["add-children"] = commands.NewAddChildren(network)
	factory.cmdsByName["catalog"] = commands.NewCatalog(network)
	factory.cmdsByName["add-catalog"] = commands.NewAddCatalog(network)
	factory.cmdsByName["version"] = commands.NewVersion(network)
	factory.cmdsByName["create"] = commands.NewCreate(network)
	factory.cmdsByName["delete"] = commands.NewDelete(network)
	factory.cmdsByName["application"] = commands.NewApplication(network)
	factory.cmdsByName["applications"] = commands.NewApplications(network)
	factory.cmdsByName["sensors"] = commands.NewSensors(network)
	factory.cmdsByName["sensor"] = commands.NewSensor(network)
	factory.cmdsByName["effectors"] = commands.NewEffectors(network)
	factory.cmdsByName["policies"] = commands.NewPolicies(network)
	factory.cmdsByName["policy"] = commands.NewPolicy(network)
	factory.cmdsByName["start-policy"] = commands.NewStartPolicy(network)
	factory.cmdsByName["stop-policy"] = commands.NewStopPolicy(network)
	factory.cmdsByName["destroy-policy"] = commands.NewDestroyPolicy(network)
	factory.cmdsByName["config"] = commands.NewConfig(network)
	factory.cmdsByName["locations"] = commands.NewLocations(network)
	factory.cmdsByName["activity"] = commands.NewActivity(network)
	factory.cmdsByName["activity-children"] = commands.NewActivityChildren(network)
	factory.cmdsByName["activities"] = commands.NewActivities(network)
	factory.cmdsByName["spec"] = commands.NewSpec(network)
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
