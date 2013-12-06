package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "farc"
	app.Usage = "file archiver & compressor"
	app.Action = func(c *cli.Context) {
		println("foo")
	}

	app.Run(os.Args)
}
