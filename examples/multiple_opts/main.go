//go:build linux || darwin || windows

package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/aileron-projects/go-sockopts"
)

// On linux, socket options can be checked with the command:
// 	strace -C -f -e trace=setsockopt go run ./

func main() {
	lc := &net.ListenConfig{
		Control: sockopts.Control(
			sockopts.SO_REUSEADDR.SetFunc(true),
			sockopts.SO_REUSEPORT.SetFunc(true),
		),
		// Control: func(network, address string, c syscall.RawConn) error {
		// 	var inner1, inner2, outer error
		// 	outer = c.Control(func(fd uintptr) {
		// 		inner1 = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		// 		inner2 = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		// 	})
		// 	return cmp.Or(inner1, inner2, outer)
		// },
	}

	ln, err := lc.Listen(context.Background(), "tcp", ":8080")
	if err != nil {
		panic(err)
	}

	svr := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Gopher !!"))
		}),
	}

	log.Println("starting server")
	if err := svr.Serve(ln); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	log.Println("server closed")
}
