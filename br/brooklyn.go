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
	"fmt"
	"github.com/apache/brooklyn-client/app"
	"github.com/apache/brooklyn-client/command_factory"
	"github.com/apache/brooklyn-client/command_runner"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/io"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/plugin"
	"github.com/apache/brooklyn-client/scope"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	config := io.GetConfig()
	//target, username, password := "http://192.168.50.101:8081", "brooklyn", "Sns4Hh9j7l"
	target := config.Model.Target
	username := config.Model.Auth[target].Username
	password := config.Model.Auth[target].Password
	network := net.NewNetwork(target, username, password)
	cmdFactory := command_factory.NewFactory(network, config)

	args, scope := scope.ScopeArguments(os.Args)
	cmdRunner := command_runner.NewRunner(scope, cmdFactory)
	metadatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(args[0]), cmdRunner, metadatas...)

	// try looking for a plugin
	rpcServer, err := plugin.NewRpcServer(cmdFactory)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	if !runPluginIfExists(rpcServer, os.Args[1:], config.Model.Plugins) {
		theApp.Run(args)
	}
}

func runPluginIfExists(rpcService *plugin.RpcServer, args []string, pluginList map[string]*plugin.PluginMetadata) bool {
	for executableName, metadata := range pluginList {
		for _, command := range metadata.Commands {
			if command.Name == args[0] || command.Alias == args[0] {
				args[0] = command.Name

				rpcService.Start()
				defer rpcService.Stop()

				pluginArgs := append([]string{rpcService.Port()}, args...)

				cmd := exec.Command(executableName, pluginArgs...)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin

				defer stopPlugin(cmd)
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return true
			}
		}
	}
	return false
}

func stopPlugin(plugin *exec.Cmd) {
	plugin.Process.Kill()
	plugin.Wait()
}
