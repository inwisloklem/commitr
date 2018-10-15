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
	h, v            bool
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
	cli.CheckHelp(h)
	cli.CheckVersion(v, version, commit)

	r := bufio.NewReader(os.Stdin)
	m := tasks.Ask("Enter commit message: ", r)
	c := tasks.Ask("Enter comment: ", r)

	fmt.Printf("%v %v", m, c)
}
