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

func getNetworkCredentialsFromConfig(yamlMap map[string]interface{}) (string, string, string, bool) {
	var target, username, password string
	var skipSslChecks bool
	target, found := yamlMap["target"].(string)
	if found {
		auth, found := yamlMap["auth"].(map[string]interface{})
		if found {
			creds := auth[target].(map[string]interface{})
			username, found = creds["username"].(string)
			if found {
				password, found = creds["password"].(string)
			}
		}
		skipSslChecks, _ = yamlMap["skipSslChecks"].(bool)
	}
	return target, username, password, skipSslChecks
}

func main() {
	config := io.GetConfig()
	target, username, password, skipSslChecks := getNetworkCredentialsFromConfig(config.Map)
	//target, username, password := "http://192.168.50.101:8081", "brooklyn", "Sns4Hh9j7l"
	network := net.NewNetwork(target, username, password, skipSslChecks)
	cmdFactory := command_factory.NewFactory(network, config)

	args, scope := scope.ScopeArguments(os.Args)
	cmdRunner := command_runner.NewRunner(scope, cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(args[0]), cmdRunner, metaDatas...)
	if err := theApp.Run(args); nil != err {
		error_handler.ErrorExit(err)
	}
}
