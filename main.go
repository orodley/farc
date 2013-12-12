package main

import (
	"github.com/orodley/cli"
	"github.com/orodley/farc/farc"
	"os"
)

var commands = []cli.Command{
	{
		Name:      "extract",
		ShortName: "x",
		Usage:     "extract files from the archive",
		Action:    farc.ExtractArchive,
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
		Action:    farc.ListArchive,
	},
	{
		Name:      "add",
		ShortName: "a",
		Usage:     "add files to an existing archive",
	},
}

var flags = []cli.Flag{
	cli.BoolFlag{"verbose, v", "be more verbose"},
}

func main() {
	app := cli.NewApp()
	app.Name = "farc"
	app.Usage = "file archiver & compressor"
	app.Commands = commands
	app.Flags = flags
	app.Run(os.Args)
}
