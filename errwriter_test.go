package errwriter_test

import (
	"bytes"
	"fmt"
	"testing"

	errwriter "github.com/nasa9084/go-errwriter"
)

func TestErrWriter(t *testing.T) {
	testErrWriterSuccess(t)
}

func testErrWriterSuccess(t *testing.T) {
	buf := &bytes.Buffer{}
	ew := errwriter.New(buf)
	ew.Write([]byte("foo"))
	ew.Write([]byte("bar"))
	if _, err := ew.Write(nil); err != nil {
		t.Errorf("error should not be occured, but: %s\n", err)
		return
	}
	if s := buf.String(); s != "foobar" {
		t.Errorf("written string does not match: %s != foobar\n", s)
		return
	}
}

type ErrWriter struct {
	n int
}

func (e *ErrWriter) Write([]byte) (int, error) {
	e.n++
	return -1, fmt.Errorf("something error %d", e.n)
}

func testErrWriterError(t *testing.T) {
	ew := errwriter.New(&ErrWriter{})
	ew.Write([]byte("foo"))
	ew.Write([]byte("bar"))
	_, err := ew.Write(nil)
	if err == nil {
		t.Error("error should be occured, but nil")
		return
	}
	if err.Error() != "something error 1" {
		t.Errorf("error number does not match: %s != something error 1\n", err.Error())
		return
	}
}
