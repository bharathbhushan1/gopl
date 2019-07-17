package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// LimitReader reads bytes upto a limit.
type LimitReader struct {
	r     io.Reader
	limit int64
}

func (lr *LimitReader) Read(buf []byte) (int, error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(buf)) >= lr.limit {
		buf = buf[0:lr.limit]
	}
	n, err := lr.r.Read(buf)
	lr.limit -= int64(n)
	return n, err
}

func limitReader(r io.Reader, limit int64) io.Reader {
	return &LimitReader{r, limit}
}

func main() {
	r := strings.NewReader("hello, world!")
	lr := limitReader(r, 6)
	data, err := ioutil.ReadAll(lr)
	if err != nil && err != io.EOF {
		log.Fatalf("error reading: %s", err)
	}
	fmt.Printf("%q\n", data)
}
