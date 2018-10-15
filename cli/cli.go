package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Option struct for command line options
type Option struct {
	Name  string
	Usage string
	Dest  *bool
}

// AddOptions adds options for use via flag package
func AddOptions(opts []Option) {
	for _, v := range opts {
		flag.BoolVar(v.Dest, v.Name, false, v.Usage)
		flag.BoolVar(v.Dest, string(v.Name[0]), false, "")
	}
}

// CheckHelp shows usage info if help option provided
func CheckHelp(f bool) {
	if f {
		flag.Usage()
		os.Exit(0)
	}
}

// CheckVersion shows version and last commit if version option provided
func CheckVersion(f bool, version, commit string) {
	if f {
		fmt.Printf("%v (last commit %v)\n", version, commit)
		os.Exit(0)
	}
}

// PrintUsage prints usage info based on provided options
func PrintUsage() {
	fmt.Println("usage: commitr [-h | --help] [-v | --version]")

	var max int
	flag.VisitAll(func(f *flag.Flag) {
		if f.Usage != "" {
			split := strings.Split(f.Usage, " ")
			length := len(strings.Join([]string{split[0], split[1]}, " "))

			if length > max {
				max = length
			}
		}
	})

	flag.VisitAll(func(f *flag.Flag) {
		if f.Usage != "" {
			split := strings.Split(f.Usage, " ")
			str := strings.Join([]string{split[0], split[1]}, " ")

			for len(str) < max {
				str += " "
			}

			fmt.Printf("%v %v\n", str, strings.Join(split[2:], " "))
		}
	})
}
