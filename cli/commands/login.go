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
	"syscall"
	"bufio"
	"os"
	"strings"
	"net/http"
	"errors"

	"github.com/apache/brooklyn-client/cli/api/version"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/io"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
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
		Flags:       []cli.Flag{
			cli.BoolFlag{Name: "skipSslChecks", Usage: "Skip SSL Checks"},
			cli.BoolFlag{Name: "noCredentials", Usage: "No user/password needed"},
			cli.StringSliceFlag{Name: "header, H", Usage: "Optional headers"},
		},
	}
}


func (cmd *Login) promptAndReadUsername() (username string) {
	for username == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Username: ")
		user, err := reader.ReadString('\n')
		if err != nil {
			error_handler.ErrorExit(err)
		}
		username = strings.TrimSpace(user)
	}
	return username
}

func (cmd *Login) promptAndReadPassword() (password string) {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		error_handler.ErrorExit(err)
	}
	fmt.Printf("\n")
	return string(bytePassword)
}

func (cmd *Login) getCredentialsFromCommandLineIfNeeded() {
	if !cmd.network.CredentialsRequired {
		return
	}
	// Prompt for username if not supplied
	if cmd.network.BrooklynUser == "" {
		cmd.network.BrooklynUser = cmd.promptAndReadUsername()
	}

	// Prompt for password if not supplied (password is not echoed to screen
	if cmd.network.BrooklynUser != "" && cmd.network.BrooklynPass == "" {
		cmd.network.BrooklynPass = cmd.promptAndReadPassword()
	}
}

func parseHeaders(parsedHeaders []string) (headerMap map[string]interface{})  {
	if len(parsedHeaders)>0{
		headerMap = make(map[string]interface{})
		for _, header:=range parsedHeaders  {
			if strings.Contains(header,"="){
				keyValue := strings.SplitN(header,"=",2)
				headerMap[keyValue[0]]=keyValue[1]
			}else{
				headerMap[header]=""
			}
		}
	}
	return
}

func (cmd *Login) Run(scope scope.Scope, c *cli.Context) {
	if !c.Args().Present() {
		error_handler.ErrorExit("A URL must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}

	// If an argument was not supplied, it is set to empty string
	cmd.network.BrooklynUrl = c.Args().Get(0)
	cmd.network.BrooklynUser = c.Args().Get(1)
	cmd.network.BrooklynPass = c.Args().Get(2)
	cmd.network.SkipSslChecks = c.Bool("skipSslChecks")
	cmd.network.CredentialsRequired = !c.Bool("noCredentials")
	cmd.network.UserHeaders = parseHeaders(c.StringSlice("header"))

	config := io.GetConfig()

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

	if cmd.network.BrooklynUrl != "" &&  cmd.network.BrooklynUser == "" && cmd.network.CredentialsRequired {
		// if only target supplied at command line see if it already exists in the config file
		if username, password, err := config.GetNetworkCredentialsForTarget(cmd.network.BrooklynUrl); err == nil {
			cmd.network.BrooklynUser = username
			cmd.network.BrooklynPass = password
		}
	}
	cmd.getCredentialsFromCommandLineIfNeeded()

	// now persist these credentials to the yaml file
	cmd.config.SetNetworkCredentials(cmd.network.BrooklynUrl, cmd.network.BrooklynUser, cmd.network.BrooklynPass)
	cmd.config.SetSkipSslChecks(cmd.network.SkipSslChecks)
	cmd.config.SetCredentialsRequired(cmd.network.CredentialsRequired)
	cmd.config.SetUserHeaders(cmd.network.UserHeaders)
	cmd.config.Write()

	loginVersion, code, err := version.Version(cmd.network)
	if nil != err {
		if code == http.StatusUnauthorized {
			err = errors.New("Unauthorized")
		}
		error_handler.ErrorExit(err)
	}
	fmt.Printf("Connected to Brooklyn version %s at %s\n", loginVersion.Version, cmd.network.BrooklynUrl)
}
