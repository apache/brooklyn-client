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
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	brIO "github.com/apache/brooklyn-client/io"
	"github.com/apache/brooklyn-client/plugin"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type InstallPlugin struct {
	config     *brIO.Config
	rpcService *plugin.RpcServer
}

func NewInstallPlugin(config *brIO.Config) (cmd *InstallPlugin) {
	cmd = new(InstallPlugin)
	cmd.config = config
	rpcService, err := plugin.NewRpcServer(nil)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	cmd.rpcService = rpcService
	return
}

func (cmd *InstallPlugin) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "install-plugin",
		Aliases:     []string{"install", "ip"},
		Description: "Install a plugin",
		Usage:       "BROOKLYN_NAME install-plugin [NAME] [LOCATION]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *InstallPlugin) Run(scope scope.Scope, c *cli.Context) {
	pluginExecutableName := c.Args()[0]
	pluginSrcFilepath := c.Args()[1]

	cmd.rpcService.Start()
	defer cmd.rpcService.Stop()
	_, err := exec.Command(pluginSrcFilepath, cmd.rpcService.Port(), "Install").Output()
	if err != nil {
		log.Fatal(err)
	}
	cmd.config.Model.Plugins = make(map[string]*plugin.PluginMetadata)
	cmd.config.Model.Plugins[pluginExecutableName] = cmd.rpcService.RpcCommand.PluginMetadata
	cmd.config.Write()
	pluginDestFilepath := filepath.Join(cmd.config.Model.PluginDir, pluginExecutableName)
	if err := copy(pluginSrcFilepath, pluginDestFilepath); err != nil {
		error_handler.ErrorExit(err)
	}
}

func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	closeErr := out.Close()
	if err != nil {
		return err
	}
	return closeErr
}
