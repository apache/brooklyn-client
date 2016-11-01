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
	"errors"
)

type CatalogList struct {
	network *net.Network
}

func NewCatalogList(network *net.Network) (cmd *CatalogList) {
	cmd = new(CatalogList)
	cmd.network = network
	return
}

const listCommandName = "list"

func (cmd *CatalogList) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        listCommandName,
		Description: "* List the available catalog applications",
		Usage:       "BROOKLYN_NAME catalog " + listCommandName + " " + catalogItemTypesUsage + " (may be abbreviated)",
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
	if len(c.Args()) != 1 {
		return nil, errors.New(c.App.Name + " " + listCommandName + catalogItemTypesUsage + " (may be abbreviated)")
	}
	catalogType, err := GetCatalogType(c)
	if  err != nil {
		return nil, err
	}
	switch catalogType {
	case ApplicationsItemType:
		items, err := cmd.listCatalogApplications(c)
		return items, err
	case EntitiesItemType:
		items, err := cmd.listEntities(c)
		return items, err
	case LocationsItemType:
		items, err := cmd.listLocations(c)
		return items, err
	case PoliciesItemType:
		items, err := cmd.listPolicies(c)
		return items, err
	}
	return nil, errors.New("Unrecognised argument")
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
		result[i] = location.IdentityDetails
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

func (cmd *CatalogList) listCatalogApplications(c *cli.Context) ([]models.IdentityDetails, error) {
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
