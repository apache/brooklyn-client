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
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {

	testFileFormat(t, "testconfig.json")
	testFileFormat(t, "legacyConfig.json")
	testCredentialsRequired(t, "testConfigHeaders.json")
	testHeaders(t, "testConfigHeaders.json")
}

func testCredentialsRequired(t *testing.T,testFile string) {
	config :=getConfigFromFile(t,testFile)
	isCredentialsRequired := config.GetCredentialsRequired()
	assertBool(nil,t,"isCredentialsRequired",isCredentialsRequired,false)
}

func testHeaders(t *testing.T, testFile string) {
	config :=getConfigFromFile(t,testFile)
	userHeaders := config.GetUserHeaders()
	expectedHeaders:=make(map[string]interface{})
	expectedHeaders["Header1"]="Header one"
	expectedHeaders["Header2"]="Header 2"
	expectedHeaders["header3"]=""

	assertHeaders(nil, t, userHeaders, expectedHeaders)

}

func assertHeaders(err error, t *testing.T, actualHeaders map[string]interface{}, expectedHeaders map[string]interface{}) {
	var expectedLen=len(expectedHeaders)

	if len(actualHeaders) != expectedLen{
		t.Errorf("Headers len != %d: %d ",expectedLen, len(actualHeaders))
	}else{
		i := 0
		keysExpected := make([]string, len(expectedHeaders))
		for k:= range expectedHeaders {
			keysExpected[i] = k
			i++
			if val, found := actualHeaders[k]; found{
				if expectedHeaders[k] != val{
					t.Errorf("Header value for %s != %s:%s ",k,expectedHeaders[k],val)
				}
			}else{
				t.Errorf("Header not found %s ",k)
			}
		}
	}
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

func testFileFormat(t *testing.T, testFile string) {
	config :=getConfigFromFile(t,testFile)

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

func assertBool(err error, t *testing.T, paramDescription string, expectedValue bool, actualValue bool) {
	if err != nil {
		t.Error(err)
	}
	if actualValue != expectedValue {
		t.Errorf("%s != %t: %t",paramDescription, expectedValue, actualValue )
	}
}