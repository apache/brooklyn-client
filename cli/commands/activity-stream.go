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
package commands

import (
	"fmt"
	"github.com/apache/brooklyn-client/cli/api/activities"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
)

type ActivityStreamEnv struct {
	network *net.Network
}

type ActivityStreamStderr struct {
	network *net.Network
}

type ActivityStreamStdin struct {
	network *net.Network
}

type ActivityStreamStdout struct {
	network *net.Network
}

func NewActivityStreamEnv(network *net.Network) (cmd *ActivityStreamEnv) {
	cmd = new(ActivityStreamEnv)
	cmd.network = network
	return
}

func NewActivityStreamStderr(network *net.Network) (cmd *ActivityStreamStderr) {
	cmd = new(ActivityStreamStderr)
	cmd.network = network
	return
}

func NewActivityStreamStdin(network *net.Network) (cmd *ActivityStreamStdin) {
	cmd = new(ActivityStreamStdin)
	cmd.network = network
	return
}

func NewActivityStreamStdout(network *net.Network) (cmd *ActivityStreamStdout) {
	cmd = new(ActivityStreamStdout)
	cmd.network = network
	return
}

func (cmd *ActivityStreamEnv) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "env",
		Description: "Show the ENV stream for a given activity",
		Usage:       "BROOKLYN_NAME ACTIVITY-SCOPE env",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStderr) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stderr",
		Description: "Show the STDERR stream for a given activity",
		Usage:       "BROOKLYN_NAME ACTIVITY-SCOPE stderr",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStdin) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stdin",
		Description: "Show the STDIN stream for a given activity",
		Usage:       "BROOKLYN_NAME ACTIVITY-SCOPE ] stdin",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamStdout) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "stdout",
		Description: "Show the STDOUT stream for a given activity",
		Usage:       "BROOKLYN_NAME ACTIVITY-SCOPE stdout",
		Flags:       []cli.Flag{},
	}
}

func (cmd *ActivityStreamEnv) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "env")
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStderr) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stderr")
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStdin) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stdin")
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(activityStream)
}

func (cmd *ActivityStreamStdout) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	activityStream, err := activities.ActivityStream(cmd.network, scope.Activity, "stdout")
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Println(activityStream)
}
