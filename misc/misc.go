package misc

import (
	"fmt"
	"os"
)

// Filter filters an array of strings
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// RemoveEmpty is a predicate function for use in Filter to remove empty strings from array
func RemoveEmpty(s string) bool {
	if s == "" {
		return false
	}
	return true
}

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
