package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "far"
	app.Usage = "file archiver"
	app.Action = func(c *cli.Context) {
		println("foo")
	}

	app.Run(os.Args)
}
