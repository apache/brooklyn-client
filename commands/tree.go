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
	"github.com/apache/brooklyn-client/api/application"
	"github.com/apache/brooklyn-client/command_metadata"
	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
	"github.com/apache/brooklyn-client/scope"
	"github.com/codegangsta/cli"
)

type Tree struct {
	network *net.Network
	c       *cli.Context
}

func NewTree(network *net.Network) (cmd *Tree) {
	cmd = new(Tree)
	cmd.network = network
	return
}

func (cmd *Tree) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "tree",
		Description: "* Show the tree of all applications",
		Usage:       "BROOKLYN_NAME tree",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Tree) Run(scope scope.Scope, c *cli.Context) {
	cmd.c = c
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	trees, err := application.Tree(cmd.network)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	cmd.printTrees(trees, "")
}

func (cmd *Tree) printTrees(trees []models.Tree, indent string) {
	for i, app := range trees {
		cmd.printTree(app, indent, i == len(trees)-1)
	}
}

func (cmd *Tree) printTree(tree models.Tree, indent string, last bool) {
	fmt.Fprintln(cmd.c.App.Writer, indent+"|-", tree.Name)
	fmt.Fprintln(cmd.c.App.Writer, indent+"+-", tree.Type)

	if last {
		indent = indent + "  "
	} else {
		indent = indent + "| "
	}
	cmd.printTrees(tree.Children, indent)
}
