package sockopts

import (
	"cmp"
	"syscall"
)

// SetFunc is the function that set socket option to fd.
// SetFuncs provides Control method that can be used for [net.ListenConfig.Control].
type SetFunc func(fd uintptr) error

// ControlFunc controls connections.
// Socket options are configures through conn.
type ControlFunc func(network, address string, conn syscall.RawConn) error

// Control sets socket option to conn.
func (f SetFunc) Control(network, address string, conn syscall.RawConn) (err error) {
	conErr := conn.Control(func(fd uintptr) {
		if err = f(fd); err != nil {
			return // Fail fast.
		}
	})
	return cmp.Or(err, conErr)
}

// SetFuncs is the collection of [SetFunc].
// SetFuncs provides Control method that can be used for [net.ListenConfig.Control].
type SetFuncs []SetFunc

// Control sets socket options to conn.
func (fs SetFuncs) Control(network, address string, conn syscall.RawConn) (err error) {
	conErr := conn.Control(func(fd uintptr) {
		for _, f := range fs {
			if err = f(fd); err != nil {
				return // Fail fast.
			}
		}
	})
	return cmp.Or(err, conErr)
}

// Control returns a ControlFunc from given functions.
// Control is the alias for SetFuncs(fs).Control.
func Control(fs ...SetFunc) ControlFunc {
	return SetFuncs(fs).Control
}
