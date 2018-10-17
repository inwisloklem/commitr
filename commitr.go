package main

import (
	"bufio"
	"commitr/fs"
	"flag"
	"fmt"
	"os"

	"commitr/cli"
	"commitr/tasks"
)

const commitrList string = "./.commitr-list"

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

	reader := bufio.NewReader(os.Stdin)
	m := tasks.AskRequired("Enter commit message: ", reader)
	c := tasks.Ask("Enter comment (press <Enter> to skip): ", reader)

	cs := tasks.LoadCommands(fs.ReadFromFile(commitrList), m, c)
	tasks.ExecCommands(cs)

	fmt.Println("Success! All commands was executed as planned.")
}
