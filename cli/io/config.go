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

	"bytes"
	"encoding/base64"
	"errors"
	"github.com/apache/brooklyn-client/cli/error_handler"
)

// Deprecated: support old style of .brooklyn_cli format for version <= 0.11.0
const authKey = "auth"

const credentialsKey = "credentials"
const usernameKey = "username"
const passwordKey = "password"
const targetKey = "target"
const skipSslChecksKey = "skipSslChecks"
const authTypeKey = "authType"

type Config struct {
	FilePath string
	Map      map[string]interface{}
}

func GetConfig() (config *Config) {
	// check to see if $BRCLI_HOME/.brooklyn_cli or $HOME/.brooklyn_cli exists
	// Then parse it to get user credentials
	config = new(Config)
	if os.Getenv("BRCLI_HOME") != "" {
		config.FilePath = filepath.Join(os.Getenv("BRCLI_HOME"), ".brooklyn_cli")
	} else {
		config.FilePath = filepath.Join(os.Getenv("HOME"), ".brooklyn_cli")
	}
	if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
		config.Map = make(map[string]interface{})
		config.Write()
	}
	config.read()
	return
}


func (config *Config) Delete() (err error) {
	if _, err := os.Stat(config.FilePath); err == nil {
		err = os.Remove(config.FilePath)
	}
	return err
}

func (config *Config) Write() {
	// Create file as read/write by user (but does not change perms of existing file)
	fileToWrite, err := os.OpenFile(config.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer fileToWrite.Close()
	if err != nil {
		error_handler.ErrorExit(err)
	}

	marshalledMap, err := json.Marshal(config.Map)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	var formatted bytes.Buffer
	if err = json.Indent(&formatted, marshalledMap, "", "    "); err != nil {
		error_handler.ErrorExit(err)
	}
	fileToWrite.Write(formatted.Bytes())
}

func (config *Config) read() {
	fileToRead, err := os.Open(config.FilePath)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	defer fileToRead.Close()

	dec := json.NewDecoder(fileToRead)
	dec.Decode(&config.Map)
}


// getCredentials reads credentials from .brooklyn_cli data formatted for versions > 0.11.0
// Note that the password is base64 encoded to avoid json formatting problems
//{
//    "credentials": {
//        "http://geoffs-macbook-pro.local:8081": "Z2VvZmY6cGFzc3dvcmQ=",
//        "http://localhost:8081": "Z2VvZmY6cGFzc3dvcmQ="
//    },
//    "skipSslChecks": false,
//    "target": "http://geoffs-macbook-pro.local:8081"
//}

func (config *Config) getCredentials(target string) (credentials string, err error) {
	credentialsMap, found := config.Map[credentialsKey].(map[string]interface{})
	if !found {
		err = errors.New("No credentials found in configuration")
		return
	}

	credentials, found = credentialsMap[target].(string)
	if !found {
		err = errors.New("No credentials found in configuration for " + target)
		return
	}

	return credentials, err
}

// Deprecated:
// getCredentialsOldStyle provides backward support for .brooklyn_cli format for version <= 0.11.0:
// {
//  "auth": {
//    "http://geoffs-macbook-pro.local:8081": {
//      "password": "password",
//      "username": "geoff"
//    }
//  },
//  "skipSslChecks": false,
//  "target": "http://geoffs-macbook-pro.local:8081"
//}
func (config *Config) getCredentialsOldStyle(target string) (username string, password string, err error) {
	auth, found := config.Map[authKey].(map[string]interface{})
	if !found {
		err = errors.New("No credentials for " + target)
		return
	}

	creds, found := auth[target].(map[string]interface{})
	if !found {
		err = errors.New("No credentials found for " + target)
		return
	}

	if username, found = creds[usernameKey].(string); !found {
		err = errors.New("No credentials for " + username)
		return
	}

	if password, found = creds[passwordKey].(string); !found {
		err = errors.New("No credentials for " + username)
		return
	}

	return username, password, err
}

func (config *Config) initialize() {
	if config.Map == nil {
		config.Map = make(map[string]interface{})
	}
	if _, found := config.Map[credentialsKey]; !found {
		config.Map[credentialsKey] = make(map[string]interface{})
	}
	if _, found := config.Map[authTypeKey]; !found {
		config.Map[authTypeKey] = make(map[string]interface{})
	}
}

func (config *Config) setCredential(target string, credentials string) {
	credentialsMap := config.Map[credentialsKey].(map[string]interface{})
	credentialsMap[target] = credentials
}

func (config *Config) SetNetworkCredentials(target string, credentials string) {
	config.initialize()
	config.adaptLegacyCredentialFormat()
	config.setCredential(target, credentials)
	config.Map[targetKey] = target

	// Overwrite old style format from version <= 0.11.0
	delete(config.Map, authKey)
}

func (config *Config) adaptLegacyCredentialFormat() {
	auth, found := config.Map[authKey].(map[string]interface{})
	if !found {
		return
	}
	for url, credMap := range auth {
		creds := credMap.(map[string]interface{})
		var username, password string
		username, found := creds[usernameKey].(string)
		if found {
			password, found = creds[passwordKey].(string)
		}
		if found {
			credentials := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
			config.setCredential(url, credentials)
		}
	}
}

func (config *Config) GetNetworkCredentialsForTarget(target string) (credentials string, err error) {
	if credentials, err = config.getCredentials(target); err != nil {
		var username, password string
		username, password, err = config.getCredentialsOldStyle(target)
		if err == nil{
			credentials = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		}
	}
	return
}

func (config *Config) GetNetworkCredentials() (target string, credentials string, err error) {
	target, found := config.Map[targetKey].(string)
	if found {
		credentials, err = config.GetNetworkCredentialsForTarget(target)
	} else {
		err = errors.New("Not logged in")
	}
	return target, credentials, err
}

func (config *Config) GetSkipSslChecks() bool {
	if config.Map == nil {
		config.Map = make(map[string]interface{})
	}
	skipSslChecks, _ := config.Map[skipSslChecksKey].(bool)
	return skipSslChecks
}

func (config *Config) SetSkipSslChecks(skipSslChecks bool) {
	config.Map[skipSslChecksKey] = skipSslChecks
}

func (config *Config) GetAuthType(target string) (authType string, err error){
	authTypeMap, found := config.Map[authTypeKey].(map[string]interface{})
	if found{
		authType, found = authTypeMap[target].(string)
	}
	// default behaviour
	if! found{
		authType = "Basic"
	}
	return authType, err
}

func (config *Config) SetAuthType(target string, authType string) {
	authTypeMap := config.Map[authTypeKey].(map[string]interface{})
	authTypeMap[target] = authType
}