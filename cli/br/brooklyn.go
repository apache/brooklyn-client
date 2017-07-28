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
package main

import (
	"github.com/apache/brooklyn-client/cli/app"
	"github.com/apache/brooklyn-client/cli/command_factory"
	"github.com/apache/brooklyn-client/cli/command_runner"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/io"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"os"
	"path/filepath"
)

func main() {
	config := io.GetConfig()
	skipSslChecks := config.GetSkipSslChecks()
	target, username, password, err := config.GetNetworkCredentials()
	if err != nil && requiresLogin(os.Args) {
		error_handler.ErrorExit(err)
	}

	//target, username, password := "http://192.168.50.101:8081", "brooklyn", "Sns4Hh9j7l"
	network := net.NewNetwork(target, username, password, skipSslChecks, verbosity(os.Args))
	cmdFactory := command_factory.NewFactory(network, config)

	args, scope := scope.ScopeArguments(os.Args)
	cmdRunner := command_runner.NewRunner(scope, cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(args[0]), cmdRunner, metaDatas...)
	if err := theApp.Run(args); nil != err {
		error_handler.ErrorExit(err)
	}
}

func requiresLogin(args []string) bool {
	// global help or version commands
	if (contains(args, "-h") ||
		contains(args, "--help") ||
	    contains(args, "-v") ||
	    contains(args, "--version")) {
		return false
	}
	// br on its own
	if (len(args) == 1) {
		return false
	}
	//
	if (len(args) > 1) {
		if (args[1] == "login" || args[1] == "help") {
			return false
		}
	}
	return true
}

func verbosity(args []string) string {
	if contains(args, "--vverbose") {
		return "vverbose"
	} else if contains(args, "--verbose") {
		return "verbose"
	} else {
		return "normal"
	}
}

func contains(slice []string, val string) bool {
	for _, a := range slice {
		if a == val {
			return true
		}
	}
	return false
}