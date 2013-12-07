package farc

import (
	"fmt"
	"github.com/codegangsta/cli"
)

// ListArchive implements the "list"/"ls" command
func ListArchive(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "list")
		return
	}

	archive, err := NewArchive(c.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(archive)
}
