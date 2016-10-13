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
	"github.com/apache/brooklyn-client/cli/api/catalog"
	"github.com/apache/brooklyn-client/cli/command_metadata"
	"github.com/apache/brooklyn-client/cli/error_handler"
	"github.com/apache/brooklyn-client/cli/net"
	"github.com/apache/brooklyn-client/cli/scope"
	"github.com/apache/brooklyn-client/cli/terminal"
	"github.com/urfave/cli"
	"github.com/apache/brooklyn-client/cli/models"
)

type CatalogList struct {
	network *net.Network
}

func NewCatalogList(network *net.Network) (cmd *CatalogList) {
	cmd = new(CatalogList)
	cmd.network = network
	return
}

func (cmd *CatalogList) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "list",
		Description: "* List the available catalog applications",
		Usage:       "BROOKLYN_NAME catalog list",
		Flags:       []cli.Flag{
			cli.BoolFlag{
				Name:  "applications, a",
				Usage: "list applications (default)",
			},
			cli.BoolFlag{
				Name:  "entities, e",
				Usage: "list entities",
			},
			cli.BoolFlag{
				Name:  "locations, l",
				Usage: "list locations",
			},
			cli.BoolFlag{
				Name:  "policies, p",
				Usage: "list policies",
			},
		},
	}
}

func (cmd *CatalogList) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	summary, err := cmd.list(c)
	if nil != err {
		error_handler.ErrorExit(err)
	}
	table := terminal.NewTable([]string{"Id", "Name", "Description"})
	for _, catalogItem := range summary {
		table.Add(catalogItem.Id, catalogItem.Name, catalogItem.Description)
	}
	table.Print()
}

func (cmd *CatalogList) list(c *cli.Context) ([]models.IdentityDetails, error) {
	if c.IsSet("entities") {
		list, err := cmd.listEntities(c)
		return list, err
	} else if c.IsSet("locations") {
		list, err := cmd.listLocations(c)
		return list, err
        } else if c.IsSet("policies") {
		list, err := cmd.listPolicies(c)
		return list, err
	}
	items, err := cmd.listCatalogItems(c)
	return items, err
}

func (cmd *CatalogList) listPolicies(c *cli.Context) ([]models.IdentityDetails, error) {
	policies, err := catalog.Policies(cmd.network)
	if err != nil {
		return nil, err
	}
	result := make([]models.IdentityDetails, len(policies))
	for i, policy := range policies {
		result[i] = policy.IdentityDetails
	}
	return result, nil
}

func (cmd *CatalogList) listLocations(c *cli.Context) ([]models.IdentityDetails, error) {
	locations, err := catalog.Locations(cmd.network)
	if err != nil {
		return nil, err
	}
	result := make([]models.IdentityDetails, len(locations))
	for i, location := range locations {
		result[i] = location.CatalogItemSummary.IdentityDetails
	}
	return result, nil
}

func (cmd *CatalogList) listEntities(c *cli.Context) ([]models.IdentityDetails, error) {
	entities, err := catalog.Entities(cmd.network)
	if err != nil {
		return nil, err
	}
	result := make([]models.IdentityDetails, len(entities))
	for i, ent := range entities {
		result[i] = ent.IdentityDetails
	}
	return result, nil
}

func (cmd *CatalogList) listCatalogItems(c *cli.Context) ([]models.IdentityDetails, error) {
	items, err := catalog.Catalog(cmd.network)
	if err != nil {
		return nil, err
	}
	result := make([]models.IdentityDetails, len(items))
	for i, item := range items {
		result[i] = item.IdentityDetails
	}
	return result, nil
}
