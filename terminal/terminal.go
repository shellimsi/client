package terminal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kr/pty"
)

var (
	terminalSizeCh = make(chan os.Signal, 1)
)

func init() {
	signal.Notify(terminalSizeCh, syscall.SIGWINCH)
}

type ResizeHandler interface {
	Handle(size *pty.Winsize) error
}

// func handleSize(conn net.Conn) error {

// 	for range terminalSizeCh {
// 		if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
// 			log.Printf("error resizing pty: %s", err)
// 		}
// 	}

// }
