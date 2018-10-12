package misc

import (
	"fmt"
	"os"
)

// HandleError prints error message
func HandleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}

// HandleFatalError prints error message and exits
func HandleFatalError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
