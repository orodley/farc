package farc

import (
	"io"
	"os"
)

// Archive provides a uniform format for manipulating archives
type Archive interface {
	// NextFile advances to the next file, returning a reader that
	// can be used to read data from it, and an os.FileInfo object
	// containing metadata about it
	NextFile() (io.Reader, os.FileInfo, error)

	// NewFile adds a new file to the archive with metadata as given in
	// an os.FileInfo object, returning a writer that can be used to write
	// data to it
	NewFile(os.FileInfo) (io.Reader, error)

	// Write writes the archive out to a Writer
	Write(io.Writer) error
}
