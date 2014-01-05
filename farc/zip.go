package farc

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

type ZipArchive struct {
	zip.Reader
	closer   Closer
	currFile int
}

func newZipArchive(reader io.Reader, closer Closer, size int64) (Archive, error) {
	zipReader, err := zip.NewReader(makeReaderAt(reader), size)
	if err != nil {
		return nil, err
	}

	return &ZipArchive{*zipReader, closer, 0}, nil
}

// Methods satisfying Archive

func (zipArchive *ZipArchive) NextFile() (io.Reader, FileInfo, error) {
	if zipArchive.currFile >= len(zipArchive.File) {
		return nil, nil, io.EOF
	}

	file := zipArchive.File[zipArchive.currFile]
	zipArchive.currFile++

	reader, err := file.Open()
	if err != nil {
		return nil, nil, err
	}

	modTime := convertDosTime(file.ModifiedDate, file.ModifiedTime)
	mode := os.FileMode(defaultPerm)
	// This seems to be the only way to tell if a file is a directory in zips
	if strings.HasSuffix(file.Name, "/") {
		mode |= os.ModeDir
	}
	afi := &AllFileInfo{
		name:    file.Name,
		size:    int64(file.UncompressedSize64),
		mode:    mode,
		modTime: modTime,
	}

	return reader, afi, nil
}

func (zipArchive *ZipArchive) NewFile(FileInfo) (io.Reader, error) {
	return nil, nil
}

func (zipArchive *ZipArchive) Write(io.Writer) error {
	return nil
}

func (zipArchive *ZipArchive) Close() error {
	return zipArchive.closer.Close()
}
