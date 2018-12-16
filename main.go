package main

import (
	"io"
	"log"
	"net"
	"os"
)

func client(conn net.Conn) error {
	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(os.Stdin, conn) }()
	_, _ = io.Copy(conn, os.Stdout)

	return nil
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	if err := client(conn); err != nil {
		log.Fatal(err)
	}
}
