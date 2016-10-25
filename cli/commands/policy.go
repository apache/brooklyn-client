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
	"github.com/apache/brooklyn-client/cli/api/entity_policies"
	"github.com/apache/brooklyn-client/cli/api/entity_policy_config"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/models"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"sort"
)

type Policy struct {
	network *net.Network
}

type policyConfigList []models.PolicyConfigList

// Len is the number of elements in the collection.
func (configs policyConfigList) Len() int {
	return len(configs)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (configs policyConfigList) Less(i, j int) bool {
	return configs[i].Name < configs[j].Name
}

// Swap swaps the elements with indexes i and j.
func (configs policyConfigList) Swap(i, j int) {
	temp := configs[i]
	configs[i] = configs[j]
	configs[j] = temp
}

func NewPolicy(network *net.Network) (cmd *Policy) {
	cmd = new(Policy)
	cmd.network = network
	return
}

func (cmd *Policy) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "policy",
		Aliases:     []string{"policies", "pol", "pols"},
		Description: "Show the policies for an application or entity",
		Usage:       "BROOKLYN_NAME SCOPE policy [NAME]",
		Flags:       []cli.Flag{},
	}
}

func (cmd *Policy) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	if c.Args().Present() {
		cmd.show(scope.Application, scope.Entity, c.Args().First())
	} else {
		cmd.list(scope.Application, scope.Entity)
	}
}

func (cmd *Policy) show(application, entity, policy string) {
	configs, err := entity_policy_config.GetAllConfigValues(cmd.network, application, entity, policy)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Name", "Value", "Description"})
	var theConfigs policyConfigList = configs
	sort.Sort(theConfigs)

	for _, config := range theConfigs {
		value, err := entity_policy_config.GetConfigValue(cmd.network, application, entity, policy, config.Name)
		if nil != err {
			error_handler.ErrorExit(err)
		}
		table.Add(config.Name, value, config.Description)
	}
	table.Print()
}

func (cmd *Policy) list(application, entity string) {
	policies, err := entity_policies.PolicyList(cmd.network, application, entity)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "State"})
	for _, policy := range policies {
		table.Add(policy.Id, policy.Name, string(policy.State))
	}
	table.Print()
}
