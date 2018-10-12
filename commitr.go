package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"commitr/cli"
	"commitr/tasks"
)

var (
	commit, version string

	h, v bool
)

func init() {
	opts := []cli.Option{
		cli.Option{
			Name:  "help",
			Usage: "-h, --help show help",
			Dest:  &h,
		},
		cli.Option{
			Name:  "version",
			Usage: "-v, --version show version",
			Dest:  &v,
		},
	}

	cli.AddOptions(opts)
	flag.Usage = cli.PrintUsage

	flag.Parse()
}

func main() {
	version = "v.0.1.0"
	commit = "b7a5a6c"

	cli.CheckHelp(h)
	cli.CheckVersion(v, version, commit)

	reader := bufio.NewReader(os.Stdin)

	message := tasks.Ask("Enter commit message: ", reader)
	comment := tasks.Ask("Enter comment: ", reader)

	fmt.Printf("%v %v", message, comment)
}
