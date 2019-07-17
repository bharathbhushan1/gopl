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
	conn, err := net.Dial("tcp6", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("DONE")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("STDIN read. SLEEPING FOR 5 seconds")
	time.Sleep(5 * time.Second)
	tcpConn := conn.(*net.TCPConn)
	fmt.Println("CLOSING TCP CONN")
	tcpConn.CloseRead()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
