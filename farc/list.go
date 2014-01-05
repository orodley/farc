package farc

import (
	"fmt"
	"github.com/orodley/cli"
	"io"
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
	defer archive.Close()

	for {
		_, fi, err := archive.NextFile()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(fi.Name())
	}
}
