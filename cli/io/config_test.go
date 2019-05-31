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
	"encoding/base64"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {

	testFileFormat(t, "testconfig.json", true)
	testFileFormat(t, "legacyConfig.json", false)

	testAuthType(t, "testconfig.json", true)
	testAuthType(t, "legacyConfig.json", false)

}

func getConfigFromFile(t *testing.T, testFile string)(config *Config) {
	config = new(Config)
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
	return
}

func testFileFormat(t *testing.T, testFile string, testBearer bool) {
	config := getConfigFromFile(t, testFile)

	_, credentials, err := config.GetNetworkCredentials()
	assertCredentials(err, t, credentials, base64.StdEncoding.EncodeToString([]byte("user1:password1")))

	if testBearer{
		credentials, err = config.GetNetworkCredentialsForTarget("http://another.one:8081")
		assertCredentials(err, t, credentials, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	}

}

func testAuthType(t *testing.T, testFile string, checkDefaultBehaviour bool) {
	config := getConfigFromFile(t, testFile)

	authType, err := config.GetAuthType("http://some.site:8081")
	assertAuthType(err, t, authType, "Basic")

	if checkDefaultBehaviour {
		authType, err = config.GetAuthType("http://another.one:8081")
		assertAuthType(err, t, authType, "Bearer")
	}
}

func assertCredentials(err error, t *testing.T, credentials string, expectedCredentials string) {
	if err != nil {
		t.Error(err)
	}
	if credentials != expectedCredentials {
		t.Errorf("credentials != %s: %s", credentials, expectedCredentials)
	}
}

func assertAuthType(err error, t *testing.T, authType string, expectedAuthType string) {
	if err != nil {
		t.Error(err)
	}
	if authType != expectedAuthType {
		t.Errorf("authType != %s: %s", authType, expectedAuthType)
	}
}