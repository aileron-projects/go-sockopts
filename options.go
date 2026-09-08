package sockopts

import (
	"strconv"
)

// Option is the socket option.
type Option struct {
	Level int    // Level is the protocol level.
	Opt   int    // Opt is the option.
	Name  string // Name is the option name. Used for logging.
}

type (
	Bool    Option
	Int     Option
	Byte    Option
	String  Option
	Linger  Option
	Timeval Option
)

func levelString(level int) string {
	switch level {
	case SOL_SOCKET:
		return "SOL_SOCKET"
	case IPPROTO_IP:
		return "IPPROTO_IP"
	case IPPROTO_IPV6:
		return "IPPROTO_IPV6"
	case IPPROTO_TCP:
		return "IPPROTO_TCP"
	case IPPROTO_UDP:
		return "IPPROTO_UDP"
	default:
		return "level=" + "0x" + strconv.FormatInt(int64(level), 16)
	}
}
