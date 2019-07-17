package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp6", "localhost:8000")
	if err != nil {
		log.Fatalf("listen failed: %s", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "accept failed: %s", err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t1. ", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t2. ", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t3. ", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 2*time.Second)
	}
	c.Close()
}
