package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().String())
		if err != nil {
			return
		}
		_, err = io.WriteString(c, "\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
