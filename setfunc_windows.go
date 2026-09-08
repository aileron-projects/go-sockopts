//go:build windows

package sockopts

import (
	"golang.org/x/sys/windows"
)

const (
	SOL_SOCKET   = windows.SOL_SOCKET
	IPPROTO_IP   = windows.IPPROTO_IP
	IPPROTO_IPV6 = windows.IPPROTO_IPV6
	IPPROTO_TCP  = windows.IPPROTO_TCP
	IPPROTO_UDP  = windows.IPPROTO_UDP
)

var (
	SO_REUSEPORT = Bool{Level: windows.SOL_SOCKET, Opt: windows.SO_REUSEADDR, Name: "SO_REUSEPORT"}
)

func (opt Bool) SetFunc(value bool) SetFunc {
	v := 1
	if !value {
		v = 0
	}
	return func(fd uintptr) error {
		err := windows.SetsockoptInt(windows.Handle(fd), opt.Level, opt.Opt, v)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}

func (opt Int) SetFunc(value int) SetFunc {
	return func(fd uintptr) error {
		err := windows.SetsockoptInt(windows.Handle(fd), opt.Level, opt.Opt, value)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}

func (opt Linger) SetFunc(value int32) SetFunc {
	onoff := int32(1)
	if value < 0 {
		onoff = 0
	}
	l := &windows.Linger{Onoff: onoff, Linger: max(0, value)}
	return func(fd uintptr) error {
		err := windows.SetsockoptLinger(windows.Handle(fd), opt.Level, opt.Opt, l)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}
