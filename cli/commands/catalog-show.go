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
	"github.com/urfave/cli"
	"github.com/apache/brooklyn-client/cli/models"
	"errors"
	"strings"
)

type CatalogShow struct {
	network *net.Network
}

func NewCatalogShow(network *net.Network) (cmd *CatalogShow) {
	cmd = new(CatalogShow)
	cmd.network = network
	return
}

const showCommandName = "show"

func (cmd *CatalogShow) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        showCommandName,
		Description: "Show a catalog item",
		Usage:       "BROOKLYN_NAME catalog " + showCommandName + " " + catalogItemTypesUsage + " ITEM",
	}
}

func (cmd *CatalogShow) Run(scope scope.Scope, c *cli.Context) {
	if err := net.VerifyLoginURL(cmd.network); err != nil {
		error_handler.ErrorExit(err)
	}
	err := cmd.show(c)
	if nil != err {
		error_handler.ErrorExit(err)
	}
}

func (cmd *CatalogShow) show(c *cli.Context) (error) {
	if len(c.Args()) != 2 {
		return errors.New(c.App.Name + " " + showCommandName + catalogItemTypesUsage + " ITEM[:VERSION]")
	}
	catalogType, err := GetCatalogType(c)
	if  err != nil {
		return err
	}
	item := c.Args().Get(1)
	var version string
	if strings.Contains(item, ":") {
		itemVersion := strings.Split(item, ":")
		item = itemVersion[0]
		version = itemVersion[1]
	}
	switch catalogType {
	case ApplicationsItemType:
		return cmd.showCatalogApplication(c, item, version)
	case EntitiesItemType:
		return cmd.showCatalogEntity(c, item, version)
	case LocationsItemType:
		return cmd.showCatalogLocation(c, item, version)
	case PoliciesItemType:
		return cmd.showPolicy(c, item, version)
	}
	return errors.New("Unrecognised argument")
}

func (cmd *CatalogShow) showPolicy(c *cli.Context, item string, version string) (error) {
	var summary models.CatalogItemSummary
	var err error
	if version == "" {
		summary, err = catalog.GetPolicy(cmd.network, item)
	} else {
		summary, err = catalog.GetPolicyWithVersion(cmd.network, item, version)
	}
	if err != nil {
		return err
	}

	return summary.Display(c)
}

func (cmd *CatalogShow) showCatalogLocation(c *cli.Context, item string, version string) (error) {
	var summary models.CatalogItemSummary
	var err error
	if version == "" {
		summary, err = catalog.GetLocation(cmd.network, item)
	} else {
		summary, err = catalog.GetLocationWithVersion(cmd.network, item, version)
	}
	if err != nil {
		return err
	}
	return summary.Display(c)
}

func (cmd *CatalogShow) showCatalogEntity(c *cli.Context, item string, version string) (error) {
	var summary models.CatalogEntitySummary
	var err error
	if version == "" {
		summary, err = catalog.GetEntity(cmd.network, item)
	} else {
		summary, err = catalog.GetEntityWithVersion(cmd.network, item, version)
	}
	if err != nil {
		return err
	}
	return summary.Display(c)
}

func (cmd *CatalogShow) showCatalogApplication(c *cli.Context, item string, version string) (error) {
	var summary models.CatalogEntitySummary
	var err error
	if version == "" {
		summary, err = catalog.GetApplication(cmd.network, item)
	} else {
		summary, err = catalog.GetApplicationWithVersion(cmd.network, item, version)
	}
	if err != nil {
		return err
	}

	return summary.Display(c)

}
