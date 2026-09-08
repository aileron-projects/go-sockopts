//go:build linux || darwin

package sockopts

import (
	"golang.org/x/sys/unix"
)

const (
	SOL_SOCKET   = unix.SOL_SOCKET
	IPPROTO_IP   = unix.IPPROTO_IP
	IPPROTO_IPV6 = unix.IPPROTO_IPV6
	IPPROTO_TCP  = unix.IPPROTO_TCP
	IPPROTO_UDP  = unix.IPPROTO_UDP
)

func (opt Bool) SetFunc(value bool) SetFunc {
	v := 1
	if !value {
		v = 0
	}
	return func(fd uintptr) error {
		err := unix.SetsockoptInt(int(fd), opt.Level, opt.Opt, v)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}

func (opt Int) SetFunc(value int) SetFunc {
	return func(fd uintptr) error {
		err := unix.SetsockoptInt(int(fd), opt.Level, opt.Opt, value)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}

func (opt Byte) SetFunc(value byte) SetFunc {
	return func(fd uintptr) error {
		err := unix.SetsockoptByte(int(fd), opt.Level, opt.Opt, value)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}

func (opt String) SetFunc(value string) SetFunc {
	return func(fd uintptr) error {
		err := unix.SetsockoptString(int(fd), opt.Level, opt.Opt, value)
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
	l := &unix.Linger{Onoff: onoff, Linger: max(0, value)}
	return func(fd uintptr) error {
		err := unix.SetsockoptLinger(int(fd), opt.Level, opt.Opt, l)
		if err != nil {
			return setError(err, levelString(opt.Level), opt.Name)
		}
		return nil
	}
}
