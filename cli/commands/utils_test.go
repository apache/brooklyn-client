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
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestDivertStdoutToString(t *testing.T) {
	is := is.New(t)

	result, _ := divertStdoutToString(func() error {
		fmt.Println("Hello Brooklyn")
		return nil
	})
	if result != "Hello Brooklyn" {
		t.Fatalf("Result was %s\n", result)
	}
	is.Equal("Hello Brooklyn", result)
}
