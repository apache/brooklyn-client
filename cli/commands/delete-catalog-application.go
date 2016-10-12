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
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/urfave/cli"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"strings"
	"github.com/apache/brooklyn-client/cli/api/catalog"
)

type DeleteCatalogApplication struct {
	network *net.Network
}

func NewDeleteCatalogApplication(network *net.Network) (cmd *DeleteCatalogApplication) {
	cmd = new(DeleteCatalogApplication)
	cmd.network = network
	return
}

func (cmd *DeleteCatalogApplication) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "delete",
		Description: "delete the given catalog application",
		Usage:       "BROOKLYN_NAME catalog delete [APPLICATION_ID:VERSION]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *DeleteCatalogApplication) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if len(c.Args()) != 1 {
		error_handler.ErrorExit("command requires single argument APPLICATION_ID:VERSION")
	}
	appVersion := strings.Split(c.Args().First(), ":")
	if len(appVersion) != 2 {
		error_handler.ErrorExit("command requires single argument APPLICATION_ID:VERSION")
	}
	appId := appVersion[0]
	version := appVersion[1]
	_, err := catalog.DeleteApplicationWithVersion(cmd.network, appId, version)
	if nil != err {
		error_handler.ErrorExit(err)
	}
}
