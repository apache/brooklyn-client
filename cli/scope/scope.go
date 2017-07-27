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
	"strings"
)

type Scope struct {
	Application string
	Entity      string
	Effector    string
	Config      string
	Activity    string
}

func (scope Scope) String() string {
	return strings.Join([]string{
		"{Application: ", scope.Application,
		", Entity: ", scope.Entity,
		", Effector: ", scope.Effector,
		", Config: ", scope.Config,
		", Activity: ", scope.Activity,
		"}",
	}, "")
}

func application(scope *Scope, id string) {
	scope.Application = id
}

func entity(scope *Scope, id string) {
	scope.Entity = id
}

func effector(scope *Scope, id string) {
	scope.Effector = id
}

func config(scope *Scope, id string) {
	scope.Config = id
}

func activity(scope *Scope, id string) {
	scope.Activity = id
}

var scopeSpecifier = map[string]func(scope *Scope, id string){
	"application": application,
	"app":         application,
	"a":           application,
	"entity":      entity,
	"ent":         entity,
	"e":           entity,
	"effector":    effector,
	"eff":         effector,
	"f":           effector,
	"config":      config,
	"conf":        config,
	"con":         config,
	"c":           config,
	"activity":    activity,
	"act":         activity,
	"v":           activity,
}

// Scopes the arguments.
// Assumes the arguments are a copy of the program args, including the first member that defines the program name.
// Removes the scope arguments from the array and applies them to a scope object.
// Returns the remaining arguments with the program name restored to first argument.
// For example with input
//      br application 1 entity 2 doSomething
// the function will return ([]string{"br", "doSomething"}, Scope{Application:1, Entity:2})
func ScopeArguments(args []string) ([]string, Scope) {
	scope := Scope{}
	if len(args) < 2 {
		return args, scope
	}
	args = defineScope(args, &scope)
	return args, scope
}

func defineScope(args []string, scope *Scope) []string {
	// args is: name [global flags] [scope and id pair] command [flags] [arguments]
	// Skip name and global flags
	idx := 1
	for isFlag(args[idx]) {
		idx += 1
	}
	newArgs := args[:idx]
	args = args[idx:]

	for len(args) > 2 && !isFlag(args[1]) {
		if scopeFn, nameOfAScope := scopeSpecifier[args[0]]; nameOfAScope {
			scopeFn(scope, args[1])
			args = args[2:]
		} else {
			break
		}
	}
	setDefaultEntityIfRequired(scope)
	return append(newArgs, args...)
}

// true if the first character of arg is -
func isFlag(arg string) bool {
	return arg[0] == '-'
}

func setDefaultEntityIfRequired(scope *Scope) {
	if "" == scope.Entity {
		scope.Entity = scope.Application
	}
}
