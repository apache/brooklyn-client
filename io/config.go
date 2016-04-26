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
package io

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/apache/brooklyn-client/error_handler"
	"github.com/apache/brooklyn-client/plugin"
)

type Config struct {
	FilePath string
	Model    ConfigModel
}

type ConfigModel struct {
	Auth      map[string]Credentials
	PluginDir string
	Plugins   map[string]*plugin.PluginMetadata
	Target    string
}

type Url struct {
	Target string
}

type Credentials struct {
	Username string
	Password string
}

func GetConfig() (config *Config) {
	// check to see if $BRCLI_HOME/.brooklyn_cli or $HOME/.brooklyn_cli exists
	// Then parse it to get user credentials
	config = new(Config)
	var pluginDir string
	if os.Getenv("BRCLI_HOME") != "" {
		pluginDir = filepath.Join(os.Getenv("BRCLI_HOME"), "plugins")
		config.FilePath = filepath.Join(os.Getenv("BRCLI_HOME"), ".brooklyn_cli")
	} else {
		pluginDir = os.Getenv("HOME")
		config.FilePath = filepath.Join(os.Getenv("HOME"), ".brooklyn_cli")
	}
	if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
		config.Model = ConfigModel{}
		config.Write()
	}

	config.Read()

	if config.Model.PluginDir == "" {
		config.Model.PluginDir = pluginDir
		config.Write()
	}
	return
}

func (config *Config) Write() {
	// Create file as read/write by user (but does not change perms of existing file)
	fileToWrite, err := os.OpenFile(config.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	defer fileToWrite.Close()

	enc := json.NewEncoder(fileToWrite)
	enc.Encode(config.Model)
}

func (config *Config) Read() {
	fileToRead, err := os.Open(config.FilePath)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	defer fileToRead.Close()

	dec := json.NewDecoder(fileToRead)
	dec.Decode(&config.Model)
}
