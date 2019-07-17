package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// ByteCounter counts the number of bytes written to it
type ByteCounter int

func (c *ByteCounter) Write(buf []byte) (int, error) {
	*c += ByteCounter(len(buf))
	return len(buf), nil
}

// WordCounter counts the number of words written to it
type WordCounter int

func (c *WordCounter) Write(buf []byte) (int, error) {
	reader := bytes.NewReader(buf)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return count, fmt.Errorf("error counting words %s", err)
	}
	*c += WordCounter(count)
	return len(buf), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	fmt.Fprintf(&c, "%s\n", "world!")
	fmt.Println(c)

	fmt.Println("------------------------")
	var wc WordCounter
	fmt.Println(wc)
	fmt.Fprintf(&wc, "%s %s", "hello", "world")
	fmt.Println(wc)
	fmt.Fprintf(&wc, "%s", "hello")
	fmt.Fprintf(&wc, "%s %s\n%s %s", "a", "b", "c", "d")
	fmt.Println(wc)
	fmt.Println("------------------------")

	cw, countPtr := createCountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "hello")
	fmt.Println(*countPtr)
}

type bytecountingWriter struct {
	writer  io.Writer
	written int64
}

func (c *bytecountingWriter) Write(buf []byte) (int, error) {
	len, err := c.writer.Write(buf)
	c.written += int64(len)
	return len, err
}

func createCountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &bytecountingWriter{w, 0}
	return cw, &cw.written
}
