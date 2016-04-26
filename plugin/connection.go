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
	"net/rpc"
)

type connection struct {
	port string
}

func NewConnection(port string) *connection {
	return &connection{
		port: port,
	}
}

func (c *connection) installPlugin(metadata PluginMetadata) bool {
	client, err := rpc.Dial("tcp", "127.0.0.1:"+c.port)
	if err != nil {
		return false
	}
	defer client.Close()
	var reply int
	client.Call("RpcCommand.InstallPlugin", metadata, &reply)
	return reply == 0
}

func (c *connection) CoreCommand(args ...string) []string {
	client, err := rpc.Dial("tcp", "127.0.0.1:"+c.port)
	if err != nil {
		return []string{}
	}
	defer client.Close()
	var cmdOutput []string
	client.Call("RpcCommand.ExecuteCoreCommand", args, &cmdOutput)
	return cmdOutput
}
