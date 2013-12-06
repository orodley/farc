package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var commands = []cli.Command{
	{
		Name:      "extract",
		ShortName: "x",
		Usage:     "extract files from the archive",
	},
	{
		Name:      "make",
		ShortName: "m",
		Usage:     "create a new archive",
	},
	{
		Name:      "list",
		ShortName: "ls",
		Usage:     "list the contents of an archive",
	},
	{
		Name:      "add",
		ShortName: "a",
		Usage:     "add files to an existing archive",
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "farc"
	app.Usage = "file archiver & compressor"
	app.Commands = commands
	app.Action = func(c *cli.Context) {
		println("foo")
	}

	app.Run(os.Args)
}
