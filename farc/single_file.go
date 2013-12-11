package farc

import (
	"io"
)

// SingleFileArchive is used for treating a single compressed file as an
// archive, for uniform handling of compressed files and archives
// It wraps an io.Reader and FileInfo which it returns on the first call to
// NextFile(). Any subsequent calls return io.EOF
type SingleFileArchive struct {
	reader   io.Reader
	fileInfo FileInfo
	done     bool
}

func newSFA(reader io.Reader, fileInfo FileInfo) *SingleFileArchive {
	return &SingleFileArchive{
		reader:   reader,
		fileInfo: fileInfo,
		done:     false,
	}
}

// Methods satisfying Archive

func (sfa *SingleFileArchive) NextFile() (io.Reader, FileInfo, error) {
	if sfa.done {
		return nil, nil, io.EOF
	}

	sfa.done = true
	return sfa.reader, sfa.fileInfo, nil
}

func (sfa *SingleFileArchive) NewFile(FileInfo) (io.Reader, error) {
	return nil, nil
}

func (sfa *SingleFileArchive) Write(writer io.Writer) error {
	_, err := io.Copy(writer, sfa.reader)
	return err
}
