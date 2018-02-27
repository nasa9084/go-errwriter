package errwriter

import "io"

type errWriter struct {
	writer io.Writer
	err    error
}

// New creates a writer that writes until error is occured.
// Write writes given bytes if error has not occured before that time.
func New(w io.Writer) io.Writer {
	return &errWriter{writer: w}
}

func (ew *errWriter) Write(p []byte) (int, error) {
	if ew.err != nil {
		return -1, ew.err
	}
	n, err := ew.writer.Write(p)
	if err != nil {
		ew.err = err
	}
	return n, err
}
