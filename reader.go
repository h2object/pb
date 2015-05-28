package pb

import (
	"io"
)

// It's proxy reader, implement io.Reader
type Reader struct {
	io.Reader
	bar *ProgressBar
}

func NewPbReader(rd io.Reader, pb *ProgressBar) *Reader {
	return &Reader{
		Reader: rd,
		bar: pb,
	}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.bar.Add(n)
	return
}
