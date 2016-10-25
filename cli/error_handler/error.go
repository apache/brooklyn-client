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
package error_handler

import (
	"fmt"
	"os"
	"github.com/apache/brooklyn-client/cli/net"
)

const CLIUsageErrorExitCode int = 1
const CliGenericErrorExitCode int = 2
const CLITrapErrorCode int = 3

func ErrorExit(errorvalue interface{}, errorcode ...int) {
	switch errorvalue.(type) {
	case net.HttpError:
		httpError := errorvalue.(net.HttpError)
		fmt.Fprintln(os.Stderr, httpError.Body)
	case error:
		fmt.Fprintln(os.Stderr, errorvalue)
	case string:
		fmt.Fprintln(os.Stderr, errorvalue)
	case nil:
		fmt.Fprintln(os.Stderr, "No error message provided")
	default:
		fmt.Fprintln(os.Stderr, "Unknown Error Type: ", errorvalue)
	}
	if len(errorcode) > 0 {
		os.Exit(errorcode[0])
	} else {
		os.Exit(CliGenericErrorExitCode)
	}
}
