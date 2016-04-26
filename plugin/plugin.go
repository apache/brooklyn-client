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
	"os"
)

type Plugin interface {
	Run(connection Connection, args []string)
	GetMetadata() PluginMetadata
}

type Connection interface {
	CoreCommand(args ...string) []string
}

type PluginMetadata struct {
	Name     string
	Commands []Command
}

type Command struct {
	Name         string
	Alias        string
	HelpText     string
	UsageDetails Usage
}

type Usage struct {
	Usage   string
	Options map[string]string
}

func Start(cmd Plugin) {
	conn := NewConnection(os.Args[1])
	if len(os.Args) == 3 && os.Args[2] == "Install" {
		conn.installPlugin(cmd.GetMetadata())
	} else {
		cmd.Run(conn, os.Args[2:])
	}
}
