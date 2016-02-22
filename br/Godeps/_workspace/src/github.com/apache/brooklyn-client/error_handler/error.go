package error_handler

import (
	"fmt"
	"os"
)

const CLIUsageErrorExitCode int = 1
const CliGenericErrorExitCode int = 2
const CLITrapErrorCode int = 3

func ErrorExit(errorvalue interface{}, errorcode ...int) {
	switch errorvalue.(type) {
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
