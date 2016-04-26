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
package plugin

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/apache/brooklyn-client/app"
	"github.com/apache/brooklyn-client/command"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/command_runner"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
	"net"
	"net/rpc"
	"path/filepath"
	"strconv"
	"strings"
)

type RpcServer struct {
	listener    net.Listener
	stopChannel chan struct{}
	RpcCommand  *RpcCommand
}

type Factory interface {
	GetByCmdName(cmdName string) (cmd command.Command, err error)
	GetBySubCmdName(cmdName string, subCmdName string) (cmd command.Command, err error)
	CommandMetadatas() []command_metadata.CommandMetadata
}

type Runner interface {
	RunCmdByName(cmdName string, c *cli.Context) (err error)
	RunSubCmdByName(cmdName string, subCommand string, c *cli.Context) (err error)
}

type RpcCommand struct {
	PluginMetadata *PluginMetadata
	commandFactory Factory
}

func NewRpcServer(commandFactory Factory) (*RpcServer, error) {
	rpcServer := &RpcServer{
		RpcCommand: &RpcCommand{
			PluginMetadata: &PluginMetadata{},
			commandFactory: commandFactory,
		},
	}
	return rpcServer, nil
}

func (rpcServer *RpcServer) Start() error {
	var err error

	rpcServer.stopChannel = make(chan struct{})
	rpc.Register(rpcServer.RpcCommand)
	rpcServer.listener, err = net.Listen("tcp", "127.0.0.1:4567")
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := rpcServer.listener.Accept()
			if err == nil {
				go rpc.ServeConn(conn)
			} else {
				select {
				case <-rpcServer.stopChannel:
					return
				default:
					fmt.Println(err)
				}
			}
		}
	}()

	return nil
}

func (cli *RpcServer) Stop() {
	close(cli.stopChannel)
	cli.listener.Close()
}

func (cli *RpcServer) Port() string {
	return strconv.Itoa(cli.listener.Addr().(*net.TCPAddr).Port)
}

func (cmd *RpcCommand) InstallPlugin(pluginMetadata PluginMetadata, retVal *bool) error {
	cmd.PluginMetadata = &pluginMetadata
	*retVal = true
	return nil
}

func (cmd *RpcCommand) ExecuteCoreCommand(args []string, retVal *[]string) error {
	args2, scope := scope.ScopeArguments(args)
	cmdRunner := command_runner.NewRunner(scope, cmd.commandFactory)
	metadatas := cmd.commandFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(args[0]), cmdRunner, metadatas...)
	var buf bytes.Buffer
	f := bufio.NewWriter(&buf)
	theApp.Writer = f
	theApp.Run(args2)
	f.Flush()
	*retVal = strings.Split(buf.String(), "\n")
	return nil
}
