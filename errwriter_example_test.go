package errwriter_test

import (
	"bytes"

	errwriter "github.com/nasa9084/go-errwriter"
)

func ExampleErrWriter() {
	buf := bytes.Buffer{}
	ew := errwriter.New(&buf)

	ew.Write([]byte("something error will be occured"))
	ew.Write([]byte("this is not record"))
	ew.Write([]byte("this is not record"))

	if _, err := ew.Write(nil); err != nil {
		// error handling: error caused when "something error will be occured" is written
	}
}
