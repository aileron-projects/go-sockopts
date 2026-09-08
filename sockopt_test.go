package sockopts

import (
	"io"
	"syscall"
	"testing"

	"github.com/aileron-projects/go-tester"
)

type testRawConn struct {
	syscall.RawConn
	fd  uintptr
	err error
}

func (c *testRawConn) Control(f func(fd uintptr)) error {
	f(c.fd)
	return c.err
}

func TestSetFunc(t *testing.T) {
	t.Parallel()
	t.Run("no error", func(t *testing.T) {
		count := 0
		f := func(fd uintptr) error { count++; return nil }
		err := SetFunc(f).Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 1, count)
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("error", func(t *testing.T) {
		count := 0
		f := func(fd uintptr) error { count++; return io.EOF }
		err := SetFunc(f).Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 1, count)
		tester.AssertEqualErr(t, io.EOF, err)
	})
}

func TestSetFuncs(t *testing.T) {
	t.Parallel()
	t.Run("nil", func(t *testing.T) {
		fs := SetFuncs(nil)
		err := fs.Control("network", "address", &testRawConn{})
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("1 func", func(t *testing.T) {
		count := 0
		fs := SetFuncs([]SetFunc{
			func(fd uintptr) error { count++; return nil },
			func(fd uintptr) error { count++; return nil },
		})
		err := fs.Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 2, count)
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("2 funcs", func(t *testing.T) {
		count := 0
		fs := SetFuncs([]SetFunc{
			func(fd uintptr) error { count++; return nil },
			func(fd uintptr) error { count++; return nil },
		})
		err := fs.Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 2, count)
		tester.AssertEqualErr(t, nil, err)
	})
	t.Run("error at first", func(t *testing.T) {
		count := 0
		fs := SetFuncs([]SetFunc{
			func(fd uintptr) error { count++; return io.EOF },
			func(fd uintptr) error { count++; return nil },
		})
		err := fs.Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 1, count)
		tester.AssertEqualErr(t, io.EOF, err)
	})
	t.Run("error at last", func(t *testing.T) {
		count := 0
		fs := SetFuncs([]SetFunc{
			func(fd uintptr) error { count++; return nil },
			func(fd uintptr) error { count++; return io.EOF },
		})
		err := fs.Control("network", "address", &testRawConn{})
		tester.AssertEqual(t, 2, count)
		tester.AssertEqualErr(t, io.EOF, err)
	})
}
