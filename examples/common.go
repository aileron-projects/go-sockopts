package examples

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"syscall"
	"time"

	"github.com/aileron-projects/go-sockopts"
)

func NewMsgHandler(msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	})
}

func RunServer(addr string, h http.Handler, control sockopts.ControlFunc) {
	lc := &net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return control(network, address, c)
		},
	}

	ln, err := lc.Listen(context.Background(), "tcp", addr)
	if err != nil {
		panic(err)
	}

	svr := &http.Server{
		Handler:     h,
		ReadTimeout: 30 * time.Second,
	}

	log.Println("starting server:", addr)
	err = svr.Serve(ln)
	log.Println("server closed: ", addr, ": ", err)
}
