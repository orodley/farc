package farc

import (
	"fmt"
	"github.com/orodley/cli"
	"io"
	"os"
	"path/filepath"
)

// ExtractArchive implements the "extract"/"x" command
func ExtractArchive(c *cli.Context) {
	if len(c.Args()) != 1 {
		cli.ShowCommandHelp(c, "extract")
		return
	}

	archive, err := NewArchive(c.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	verbose := c.GlobalBool("verbose")

	if verbose {
		fmt.Println("Extracting", c.Args()[0])
	}

	for {
		reader, fi, err := archive.NextFile()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		if verbose {
			fmt.Println(" ", fi.Name())
		}

		if IsDir(fi) {
			// It's possible we already created it in order to have somewhere
			// to put child files
			if ex, err := exists(fi.Name()); ex {
				os.Chmod(fi.Name(), fi.Mode())
			} else if err != nil {
				fmt.Println(err)
				return
			} else {
				os.MkdirAll(fi.Name(), fi.Mode())
			}
		} else {
			flags := os.O_WRONLY | os.O_CREATE | os.O_EXCL
			dir := filepath.Dir(fi.Name())
			ex, err := exists(dir)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Make sure all its parent directories exist
			if !ex {
				os.MkdirAll(dir, 0770)
			}

			file, err := os.OpenFile(fi.Name(), flags, fi.Mode())
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			_, err = io.Copy(file, reader)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
