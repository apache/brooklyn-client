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
	"testing"
	"path/filepath"
)

func TestConfig(t *testing.T) {

	testFileFormat(t, "testConfig.json")
	testFileFormat(t, "legacyConfig.json")
}

func testFileFormat(t *testing.T, testFile string) {

	config := new(Config)
	expectedTarget := "http://some.site:8081"

	path, err := filepath.Abs(testFile)
	if err != nil {
		t.Error(err)
	}
	config.FilePath = path
	config.read()
	if config.Map["target"] != expectedTarget {
		t.Errorf("target != %s: %s", expectedTarget, config.Map["target"])
	}
	_, username, password, err := config.GetNetworkCredentials()
	assertUserPassword(err, t, username, "user1", password, "password1")

	username, password, err = config.GetNetworkCredentialsForTarget("http://another.one:8081")
	assertUserPassword(err, t, username, "user2", password, "password2")
}

func assertUserPassword(err error, t *testing.T, username string, expectedUser string, password string, expectedPassword string) {
	if err != nil {
		t.Error(err)
	}
	if username != expectedUser {
		t.Errorf("username != %s: %s", expectedUser, username)
	}
	if password != expectedPassword {
		t.Errorf("password != %s: %s", expectedPassword, username)
	}
}