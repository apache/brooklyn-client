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
package scope

import (
	"testing"
)

type testCase struct{
	Args []string
	ExpectedArgs []string
	ExpectedScope Scope
}

func TestScope(t *testing.T) {
	testCases := []testCase {
		{
			Args: []string{"br", "application"},
			ExpectedArgs: []string{"br", "application"},
			ExpectedScope: Scope{},
		}, {
			Args: []string{"br", "application", "appid", "entity"},
			ExpectedArgs: []string{"br", "entity"},
			ExpectedScope: Scope{Application: "appid", Entity: "appid"},
		}, {
			Args: []string{"br", "application", "appid", "entity", "entity"},
			ExpectedArgs: []string{"br", "entity", "entity"},
			// ScopeArguments sets Entity incorrectly when the command is entity.
			ExpectedScope: Scope{Application: "appid", Entity: "appid"},
		}, {
			Args: []string{"br", "application", "appid", "entity", "entity", "sensor"},
			ExpectedArgs: []string{"br", "sensor"},
			ExpectedScope: Scope{Application: "appid", Entity: "entity"},
		}, {
			Args: []string{"br", "--verbose", "application", "appid", "entity", "entityId", "sensor"},
			ExpectedArgs: []string{"br", "--verbose", "sensor"},
			ExpectedScope: Scope{Application: "appid", Entity: "entityId"},
		}, {
			Args: []string{"br", "--verbose", "application", "appid", "entity", "entityId", "sensor", "http.port"},
			ExpectedArgs: []string{"br", "--verbose", "sensor", "http.port"},
			ExpectedScope: Scope{Application: "appid", Entity: "entityId"},
		}, {
			Args: []string{"br", "--verbose", "a", "appid", "e", "entityId", "v", "--children", "activityId"},
			ExpectedArgs: []string{"br", "--verbose", "v", "--children", "activityId"},
			ExpectedScope: Scope{Application: "appid", Entity: "entityId"},
		},{
			Args: []string{"br", "-v"},
			ExpectedArgs: []string{"br", "-v"},
			ExpectedScope: Scope{},
		},
	}
	for _, elem := range testCases {
		argsOut, scope := ScopeArguments(elem.Args)
		assertArgs(t, argsOut, elem.ExpectedArgs)
		assertScope(t, scope, elem.ExpectedScope)
	}
}

func assertArgs(t *testing.T, actual []string, expected []string) {
	if len(actual) != len(expected) {
		t.Errorf("%q != %q", actual, expected)
		t.FailNow()
	}
	for idx, act := range actual {
		exp := expected[idx]
		if act != exp {
			t.Errorf("mismatch at index %d: %q != %q", idx, actual, expected)
		}
	}
}

func assertScope(t *testing.T, actual Scope, expected Scope) {
	if actual.Application != expected.Application ||
		actual.Activity != expected.Activity ||
		actual.Config != expected.Config ||
		actual.Effector != expected.Effector ||
		actual.Entity != expected.Entity {
		t.Errorf("%v != %v", actual, expected)
	}
}
