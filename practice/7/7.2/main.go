package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	cw io.Writer
	C  int64
}

func (c *CountWriter) Write(p []byte) (int, error) {
	c.C = int64(len(p))
	n, err := c.cw.Write(p)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CountWriter{cw: w}
	return cw, &cw.C

}

func main() {
	cw, n := CountingWriter(os.Stdout)
	cw.Write([]byte("hello world\n"))
	fmt.Println(*n)
}
