package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	terminalSizeCh = make(chan os.Signal, 1)
)

func init() {
	signal.Notify(terminalSizeCh, syscall.SIGWINCH)
}

func client(conn net.Conn) error {
	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(os.Stdin, conn) }()
	terminalSizeCh <- syscall.SIGWINCH // Initial resize.
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
