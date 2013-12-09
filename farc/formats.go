package farc

import "fmt"
import "path"

// NewArchive examines a path and returns a new Archive object for its format
func NewArchive(filename string) (Archive, error) {
	ext := path.Ext(filename)
	switch ext {
	case ".tar":
		tarArchive, err := newTarArchive(filename)
		if err != nil {
			return nil, err
		} else {
			return tarArchive, nil
		}
	default:
		return nil, fmt.Errorf("Unsupported format `%s'", ext)
	}
}
