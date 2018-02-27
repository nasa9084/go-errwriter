# go-errwriter
[![Build Status](https://travis-ci.org/nasa9084/go-errwriter.svg?branch=master)](https://travis-ci.org/nasa9084/go-errwriter)
-----

a simple wrapper for io.Writer to help simplify the error handling of Write()


## SYNOPSIS

``` go
package main

import (
    "bytes"

    "github.com/nasa9084/go-errwriter"
)

func main() {
    buf := bytes.Buffer{}
    ew := errwriter.New(&buf)

    ew.Write([]byte("something error will be occured"))
    ew.Write([]byte("this is not record"))
    ew.Write([]byte("this is not record"))

    if _, err := ew.Write(nil); err != nil {
        // error handling: error caused when "something error will be occured" is written
    }
}
```
