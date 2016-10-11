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
	"github.com/apache/brooklyn-client/cli/api/version"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/io"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

type Login struct {
	network *net.Network
	config  *io.Config
}

func NewLogin(network *net.Network, config *io.Config) (cmd *Login) {
	cmd = new(Login)
	cmd.network = network
	cmd.config = config
	return
}

func (cmd *Login) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "login",
		Description: "Login to brooklyn",
		Usage:       "BROOKLYN_NAME login URL [USER [PASSWORD]]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Login) Run(scope scope.Scope, c *cli.Context) {
	if !c.Args().Present() {
		error_handler.ErrorExit("A URL must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}

	// If an argument was not supplied, it is set to empty string
	cmd.network.BrooklynUrl = c.Args().Get(0)
	cmd.network.BrooklynUser = c.Args().Get(1)
	cmd.network.BrooklynPass = c.Args().Get(2)
	cmd.network.SkipSslChecks = c.GlobalBool("skipSslChecks")

	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}

	// Strip off trailing '/' from URL if present.
	if cmd.network.BrooklynUrl[len(cmd.network.BrooklynUrl)-1] == '/' {
		if len(cmd.network.BrooklynUrl) == 1 {
			error_handler.ErrorExit("URL must not be a single \"/\" character", error_handler.CLIUsageErrorExitCode)
		}
		cmd.network.BrooklynUrl = cmd.network.BrooklynUrl[0 : len(cmd.network.BrooklynUrl)-1]
	}

	// Prompt for password if not supplied (password is not echoed to screen
	if cmd.network.BrooklynUser != "" && cmd.network.BrooklynPass == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			error_handler.ErrorExit(err)
		}
		fmt.Printf("\n")
		cmd.network.BrooklynPass = string(bytePassword)
	}

	if cmd.config.Map == nil {
		cmd.config.Map = make(map[string]interface{})
	}
	// now persist these credentials to the yaml file
	auth, ok := cmd.config.Map["auth"].(map[string]interface{})
	if !ok {
		auth = make(map[string]interface{})
		cmd.config.Map["auth"] = auth
	}

	auth[cmd.network.BrooklynUrl] = map[string]string{
		"username": cmd.network.BrooklynUser,
		"password": cmd.network.BrooklynPass,
	}

	cmd.config.Map["target"] = cmd.network.BrooklynUrl
	cmd.config.Map["skipSslChecks"] = cmd.network.SkipSslChecks
	cmd.config.Write()

	loginVersion, err := version.Version(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	fmt.Printf("Connected to Brooklyn version %s at %s\n", loginVersion.Version, cmd.network.BrooklynUrl)
}
