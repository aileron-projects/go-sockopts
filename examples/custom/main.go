//go:build linux || darwin

package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/aileron-projects/go-sockopts"
	"golang.org/x/sys/unix"
)

// On linux, socket options can be checked with the command:
// 	strace -C -f -e trace=setsockopt go run ./

// Define custom options with appropriate type and values.
var SO_KEEPALIVE = sockopts.Bool{Level: unix.SOL_SOCKET, Opt: unix.SO_KEEPALIVE, Name: "SO_KEEPALIVE"}

func main() {
	lc := &net.ListenConfig{
		Control: SO_KEEPALIVE.SetFunc(true).Control,
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
