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
package command_runner

import (
	"github.com/apache/brooklyn-client/command"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Runner interface {
	RunCmdByName(cmdName string, c *cli.Context) (err error)
	RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) (err error)
}

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	GetBySubCmdName(cmdName string, subCmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type ConcreteRunner struct {
	cmdFactory Factory
	scope      scope.Scope
}

func NewRunner(scope scope.Scope, cmdFactory Factory) (runner Runner) {
	runner = &ConcreteRunner{
		cmdFactory: cmdFactory,
		scope:      scope,
	}
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
