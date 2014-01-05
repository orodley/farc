package farc

import (
	"io"
	"os"
	"time"
)

const (
	defaultPerm = 0770 // Permissions to use when none specified
)

// Archive provides a uniform format for manipulating archives
type Archive interface {
	// NextFile advances to the next file, returning a reader that
	// can be used to read data from it, and an os.FileInfo object
	// containing metadata about it.
	// Once the last file in the archive is reached, error should be io.EOF
	NextFile() (io.Reader, FileInfo, error)

	// NewFile adds a new file to the archive with metadata as given in
	// a FileInfo object, returning a writer that can be used to write
	// data to it
	NewFile(FileInfo) (io.Reader, error)

	// Write writes the archive out to a Writer
	Write(io.Writer) error
}

// FileInfo is the os.FileInfo - Sys()
// For archive types that don't contain any of these pieces of information,
// the method on their implementation should return a sensible default
type FileInfo interface {
	Name() string
	Size() int64
	Mode() os.FileMode
	ModTime() time.Time
}

func IsDir(f FileInfo) bool {
	return f.Mode().IsDir()
}

// AllFileInfo is a simple struct that contains all the values from FileInfo,
// along with methods that just return them. Useful if an archive type
// contains all the required data in a header; you can just use this struct
type AllFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// makeAllFileInfo takes an os.FileInfo and converts it into a farc.FileInfo.
// All that needs to be done is copy over all methods but Sys() into the
// appropriate fields
func makeAllFileInfo(ofi os.FileInfo) *AllFileInfo {
	return &AllFileInfo{
		name:    ofi.Name(),
		size:    ofi.Size(),
		mode:    ofi.Mode(),
		modTime: ofi.ModTime(),
	}
}

func (a *AllFileInfo) Name() string {
	return a.name
}

func (a *AllFileInfo) Size() int64 {
	return a.size
}

func (a *AllFileInfo) Mode() os.FileMode {
	return a.mode
}

func (a *AllFileInfo) ModTime() time.Time {
	return a.modTime
}
