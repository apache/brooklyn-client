/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package command_factory

import (
	"errors"
	"github.com/apache/brooklyn-client/cli/command"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/commands"
	"github.com/apache/brooklyn-client/cli/io"
	"github.com/apache/brooklyn-client/cli/net"
	"sort"
	"strings"
)

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	GetBySubCmdName(cmdName string, subCmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type concreteFactory struct {
	cmdsByName  map[string]command.Command
	subCommands map[string]map[string]command.Command
}

func NewFactory(network *net.Network, config *io.Config) (factory concreteFactory) {
	factory.cmdsByName = make(map[string]command.Command)
	factory.subCommands = make(map[string]map[string]command.Command)

	factory.simpleCommand(commands.NewAccess(network))
	//factory.command(commands.NewActivities(network))
	factory.simpleCommand(commands.NewActivity(network))
	factory.simpleCommand(commands.NewActivityStreamEnv(network))
	factory.simpleCommand(commands.NewActivityStreamStderr(network))
	factory.simpleCommand(commands.NewActivityStreamStdin(network))
	factory.simpleCommand(commands.NewActivityStreamStdout(network))
	factory.simpleCommand(commands.NewAddCatalog(network))
	factory.simpleCommand(commands.NewAddChildren(network))
	factory.simpleCommand(commands.NewApplication(network))
	//factory.simpleCommand(commands.NewApplications(network))
	factory.superCommand(commands.NewCatalog(network))
	factory.simpleCommand(commands.NewConfig(network))
	factory.simpleCommand(commands.NewDeploy(network))
	factory.simpleCommand(commands.NewDelete(network))
	factory.simpleCommand(commands.NewDestroyPolicy(network))
	factory.simpleCommand(commands.NewEffector(network))
	factory.simpleCommand(commands.NewEntity(network))
	factory.simpleCommand(commands.NewInvoke(network))
	factory.simpleCommand(commands.NewInvokeRestart(network))
	factory.simpleCommand(commands.NewInvokeStart(network))
	factory.simpleCommand(commands.NewInvokeStop(network))
	// NewList below is not used but we retain the code as an example of how to do a super command.
	//	factory.superCommand(commands.NewList(network))
	factory.simpleCommand(commands.NewLocations(network))
	factory.simpleCommand(commands.NewLogin(network, config))
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

func (factory *concreteFactory) superCommand(cmd command.SuperCommand) {

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
