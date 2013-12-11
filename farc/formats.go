package farc

import (
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// NewArchive examines a path and returns a new Archive object for its format
func NewArchive(filename string) (Archive, error) {
	file, err := os.Open(filename)
	reader := io.Reader(file)
	if err != nil {
		return nil, err
	}

	// Walk the list of all of the extensions in the file. When we first find
	// a terminal one (e.g. zip, tar) we return the appropriate Archive.
	// If along the way we find a non-archive, compression format (e.g. gz,
	// bz2), we wrap the reader with the reader for that format and continue
	exts := Exts(filename)
	for _, ext := range exts {
		switch ext {
		case "gz":
			reader, err = gzip.NewReader(reader)
			if err != nil {
				return nil, err
			}
		case "bz2":
			reader = bzip2.NewReader(reader)
		case "tar":
			tarArchive, err := newTarArchive(reader)
			if err != nil {
				return nil, err
			} else {
				return tarArchive, nil
			}
		default:
			return nil, fmt.Errorf("Unsupported format `%s'", ext)
		}
	}

	// TODO: Handle non-archive compressed files here
	// We should create an archive type that handles this by returning the
	// reader when NextFile() is first called and then returning io.EOF
	return nil, errors.New("Non-archive compressed files are not supported yet")
}

// Exts is like path.Ext, but it detects if there are multiple extentions
// nested on top of eachother, and returns a slice of all of them.
// The slice is ordered from outer to inner, e.g.:
//  Exts("foo.tar.gz") => { "gz", "tar" }
func Exts(filename string) []string {
	base := path.Base(filename)
	exts := strings.Split(base, ".")[1:]
	for i := 0; i < len(exts)/2; i++ {
		otherIndex := len(exts) - i - 1

		tmp := exts[i]
		exts[i] = exts[otherIndex]
		exts[otherIndex] = tmp
	}

	return exts
}
