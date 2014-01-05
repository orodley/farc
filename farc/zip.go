package farc

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

type ZipArchive struct {
	zip.Reader
	currFile int
}

func newZipArchive(reader io.Reader, size int64) (Archive, error) {
	zipReader, err := zip.NewReader(makeReaderAt(reader), size)
	if err != nil {
		return nil, err
	}

	return &ZipArchive{*zipReader, 0}, nil
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
	mode := os.FileMode(0770) // default value - zip doesn't save permissions
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
