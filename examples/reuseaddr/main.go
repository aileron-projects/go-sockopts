//go:build linux || windows || darwin

package main

import (
	"log"

	"github.com/aileron-projects/go-sockopts"
	"github.com/aileron-projects/go-sockopts/examples"
)

// On linux, socket options can be checked with the command:
// 	strace -C -f -e trace=setsockopt go run ./

func main() {
	addr := ":8080"
	control := sockopts.SO_REUSEADDR.SetFunc(true).Control
	handler1 := examples.NewMsgHandler("Hello, from alice !!")
	handler2 := examples.NewMsgHandler("Hello, from bob !!")

	closed := make(chan struct{}, 1)
	go func() {
		examples.RunServer(addr, handler1, control)
		closed <- struct{}{}
	}()
	go func() {
		examples.RunServer(addr, handler2, control)
		closed <- struct{}{}
	}()

	log.Panicln(<-closed)
	log.Panicln(<-closed)
	log.Println("example exit")
}
