package farc

import (
	"archive/tar"
	"io"
	"os"
)

type TarArchive struct {
	reader tar.Reader
}

func newTarArchive(filename string) (Archive, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &TarArchive{*tar.NewReader(file)}, nil
}

// Methods satisfying Archive

func (tarArchive *TarArchive) NextFile() (io.Reader, os.FileInfo, error) {
	return nil, nil, nil
}

func (tarArchive *TarArchive) NewFile(os.FileInfo) (io.Reader, error) {
	return nil, nil
}

func (tarArchive *TarArchive) Write(io.Writer) error {
	return nil
}
