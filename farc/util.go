package farc

import (
	"io"
	"time"
)

const (
	initialSize  = 1024 // Start at 1KiB
	growthFactor = 2    // Multiply the current capacity by this when growing
)

type readerAt struct {
	io.Reader
	buf []byte
}

func (r *readerAt) ReadAt(p []byte, off int64) (int, error) {
	end := off + int64(len(p))
	if end <= int64(len(r.buf)) {
		// We've already read it and it's sitting in the buffer
		return copy(p, r.buf[off:]), nil
	} else {
		oldLen := len(r.buf)
		// We need to read some more data
		if end > int64(cap(r.buf)) {
			newBuf := make([]byte, end, end*growthFactor)
			copy(newBuf, r.buf)
			r.buf = newBuf
		} else {
			r.buf = r.buf[:end]
		}

		r.Read(r.buf[oldLen:])
		return copy(p, r.buf[off:]), nil
	}
}

// makeReaderAt returns a ReaderAt wrapping a Reader
func makeReaderAt(r io.Reader) io.ReaderAt {
	return &readerAt{
		r,
		make([]byte, 0, initialSize),
	}
}

// convertDosTime takes two uint16s containing MS-DOS format time and date
// values and converts them to a time.Time
func convertDosTime(dosDate, dosTime uint16) time.Time {
	year := int(dosDate >> 9)
	month := time.Month((dosDate >> 5) & 0xF)
	day := int(dosDate & 0x1F)
	hour := int((dosTime >> 11) & 0x1F)
	minute := int((dosTime >> 5) & 0x3F)
	second := int(dosTime & 0x1F)

	return time.Date(year, month, day, hour, minute, second, 0, time.Local)
}
