package sockopts

// Error is the socket option error.
type Error struct {
	Inner error  // Inner is the inner error.
	Level string // Level is the socket level. e.g. "SOL_SOCKET"
	Opt   string // Opt is the option name. e.g. "SO_BINDTOIFINDEX"
	Msg   string // Msg is the error message.
}

func (e *Error) Unwrap() error {
	return e.Inner
}

func (e *Error) Error() string {
	msg := "go-sockopts/sockopts: " + e.Level + ": " + e.Opt + ":"
	if e.Msg != "" {
		msg += " " + e.Msg
	}
	if e.Inner != nil {
		msg += " [" + e.Inner.Error() + "]"
	}
	return msg
}

func (e *Error) Is(target error) bool {
	ee, ok := target.(*Error)
	if ok {
		return e.Level == ee.Level && e.Opt == ee.Opt
	}
	return false
}

func setError(inner error, level, opt string) *Error {
	return &Error{
		Inner: inner,
		Level: level,
		Opt:   opt,
		Msg:   "failed to set socket option",
	}
}
