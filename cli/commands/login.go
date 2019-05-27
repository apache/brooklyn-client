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
	"encoding/base64"
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

const authorizationParam = "authorization"
const skipSslChecksParam = "skipSslChecks"

const BASIC_AUTH = "Basic"
const BEARER_AUTH = "Bearer"

type Login struct {
	network 		*net.Network
	config  		*io.Config
	brooklynUser	string
	brooklynPass	string
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
			cli.BoolFlag{Name: skipSslChecksParam, Usage: "Skip SSL Checks"},
			cli.StringFlag{Name: authorizationParam+", A", Usage: "Type of authentication header. Format: 'authorization=Basic'" +
				" or 'authorization=Bearer:xxxxx.yyyyyy.zzzzz'"},
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
	if !cmd.isBasicAuth() {
		return
	}
	if cmd.network.Credentials != ""{
		return
	}

	// Prompt for username if not supplied
	if cmd.brooklynUser == "" {
		cmd.brooklynUser = cmd.promptAndReadUsername()
	}

	// Prompt for password if not supplied (password is not echoed to screen
	if cmd.brooklynUser!= "" && cmd.brooklynPass == "" {
		cmd.brooklynPass = cmd.promptAndReadPassword()
	}

	cmd.network.Credentials = base64.StdEncoding.EncodeToString([]byte(cmd.brooklynUser + ":" + cmd.brooklynPass))
}

func (cmd *Login) getCredentialsFromAuthParamIfNeeded(authParam string) {
	if cmd.isBasicAuth() {
		return
	}
	if cmd.network.Credentials=="" {
		credentials := cmd.parseAuthorizationParam(authParam)
		if len(credentials) != 2 {
			error_handler.ErrorExit("Use authorization=type:value")
		} else {
			cmd.network.Credentials = credentials[1]
		}
	}
}

func (cmd *Login) getAuthorizationType(authorizationParamInput string) (authorizationType string)  {
	if authorizationParamInput !=""{
		if strings.EqualFold(BASIC_AUTH,authorizationParamInput){
			authorizationType = BASIC_AUTH
		}else {
			authorizationType = cmd.parseAuthorizationParam(authorizationParamInput)[0]
			if strings.EqualFold(authorizationType,BEARER_AUTH) {
				authorizationType=BEARER_AUTH
			}
		}
	}else {
		// default authentication
		authorizationType = BASIC_AUTH
	}
	return authorizationType
}

func (cmd *Login) Run(scope scope.Scope, c *cli.Context) {
	if !c.Args().Present() {
		error_handler.ErrorExit("A URL must be provided as the first argument", error_handler.CLIUsageErrorExitCode)
	}

	// If an argument was not supplied, it is set to empty string
	cmd.network.BrooklynUrl = c.Args().Get(0)
	cmd.brooklynUser = c.Args().Get(1)
	cmd.brooklynPass = c.Args().Get(2)
	cmd.network.SkipSslChecks = c.Bool("skipSslChecks")
	cmd.network.AuthorizationType = cmd.getAuthorizationType(c.String(authorizationParam))

	//clear Credentials credentials
	cmd.network.Credentials=""

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

	if cmd.network.BrooklynUrl != "" &&  cmd.brooklynUser == "" && c.String(authorizationParam)==""{
		// if only target supplied at command line see if it already exists in the config file
		if credentials, err := config.GetNetworkCredentialsForTarget(cmd.network.BrooklynUrl); err == nil {
			cmd.network.Credentials = credentials
		}
		if authorizationType, err := config.GetAuthType(cmd.network.BrooklynUrl); err == nil{
			// TODO remove
			//fmt.Printf("authorizationType from file %s\n",authorizationType)
			cmd.network.AuthorizationType = authorizationType
		}
	}

	cmd.getCredentialsFromCommandLineIfNeeded()
	cmd.getCredentialsFromAuthParamIfNeeded(c.String(authorizationParam))

	// now persist these credentials to the yaml file
	cmd.config.SetNetworkCredentials(cmd.network.BrooklynUrl, cmd.network.Credentials)
	cmd.config.SetAuthType(cmd.network.BrooklynUrl, cmd.network.AuthorizationType)
	cmd.config.SetSkipSslChecks(cmd.network.SkipSslChecks)
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

func (cmd *Login) isBasicAuth() bool {
	return strings.EqualFold(cmd.network.AuthorizationType,BASIC_AUTH)
}

func (cmd *Login) parseAuthorizationParam(authParam string) []string {
	return strings.SplitN(authParam, ":", 2)
}


