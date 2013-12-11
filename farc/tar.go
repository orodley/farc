package farc

import (
	"archive/tar"
	"io"
	"os"
)

type TarArchive struct {
	tar.Reader
}

func newTarArchive(filename string) (Archive, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &TarArchive{*tar.NewReader(file)}, nil
}

// Methods satisfying Archive

func (tarArchive *TarArchive) NextFile() (io.Reader, FileInfo, error) {
	header, err := tarArchive.Next()
	if err != nil {
		return nil, nil, err
	}

	return tarArchive, makeAllFileInfo(header.FileInfo()), nil
}

func (tarArchive *TarArchive) NewFile(FileInfo) (io.Reader, error) {
	return nil, nil
}

func (tarArchive *TarArchive) Write(io.Writer) error {
	return nil
}
