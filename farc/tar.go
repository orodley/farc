package farc

import (
	"archive/tar"
	"io"
)

type TarArchive struct {
	tar.Reader
	closer Closer
}

func newTarArchive(reader io.Reader, closer Closer) (Archive, error) {
	return &TarArchive{*tar.NewReader(reader), closer}, nil
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

func (tarArchive *TarArchive) Close() error {
	return tarArchive.closer.Close()
}
