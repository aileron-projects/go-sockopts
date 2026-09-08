package sockopts

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/aileron-projects/go-tester"
)

func TestError(t *testing.T) {
	t.Parallel()
	t.Run("unwrap", func(t *testing.T) {
		err := &Error{Inner: io.EOF}
		inner := err.Unwrap()
		tester.AssertEqualErr(t, io.EOF, inner)
	})
	t.Run("error message", func(t *testing.T) {
		err := &Error{Inner: io.EOF, Level: "level", Opt: "opt", Msg: "msg"}
		msg := err.Error()
		tester.AssertEqual(t, "go-sockopts/sockopts: level: opt: msg [EOF]", msg)
	})
	t.Run("empty message", func(t *testing.T) {
		err := &Error{Inner: io.EOF, Level: "level", Opt: "opt", Msg: ""}
		msg := err.Error()
		tester.AssertEqual(t, "go-sockopts/sockopts: level: opt: [EOF]", msg)
	})
	t.Run("nil error", func(t *testing.T) {
		var err *Error
		tester.AssertEqual(t, false, err.Is(nil))
	})
	t.Run("nil target", func(t *testing.T) {
		err := &Error{Inner: nil, Level: "level", Opt: "opt"}
		tester.AssertEqual(t, false, err.Is(nil))
	})
	t.Run("errors equal", func(t *testing.T) {
		target := &Error{Inner: nil, Level: "level", Opt: "opt"}
		err := &Error{Inner: io.EOF, Level: "level", Opt: "opt"}
		tester.AssertEqual(t, true, errors.Is(err, target))
	})
	t.Run("level not equal", func(t *testing.T) {
		target := &Error{Inner: nil, Level: "aaa", Opt: "opt"}
		err := &Error{Inner: nil, Level: "bbb", Opt: "opt"}
		tester.AssertEqual(t, false, errors.Is(err, target))
	})
	t.Run("opt not equal", func(t *testing.T) {
		target := &Error{Inner: nil, Level: "level", Opt: "aaa"}
		err := &Error{Inner: nil, Level: "level", Opt: "bbb"}
		tester.AssertEqual(t, false, errors.Is(err, target))
	})
	t.Run("wrapped error equal", func(t *testing.T) {
		target := &Error{Level: "level", Opt: "opt"}
		inner := &Error{Level: "level", Opt: "opt"}
		err := fmt.Errorf("outer error [%w]", inner)
		tester.AssertEqual(t, true, errors.Is(err, target))
	})
	t.Run("wrapped error not equal", func(t *testing.T) {
		target := &Error{Level: "level", Opt: "opt"}
		err := fmt.Errorf("outer error [%w]", io.EOF)
		tester.AssertEqual(t, false, errors.Is(err, target))
	})
	t.Run("wrapped errors equal", func(t *testing.T) {
		target := &Error{Level: "level", Opt: "opt"}
		inner := &Error{Level: "level", Opt: "opt"}
		err := fmt.Errorf("outer error [%w] [%w]", io.EOF, inner)
		tester.AssertEqual(t, true, errors.Is(err, target))
	})
	t.Run("wrapped error not equal", func(t *testing.T) {
		target := &Error{Level: "level", Opt: "opt"}
		err := fmt.Errorf("outer error [%w] [%w]", io.EOF, io.ErrUnexpectedEOF)
		tester.AssertEqual(t, false, errors.Is(err, target))
	})
}
