<!-- markdownlint-disable MD033 MD041 -->

<div align="center">

[![Release](https://img.shields.io/github/v/release/aileron-projects/go-sockopts?sort=semver)](https://github.com/aileron-projects/go-sockopts/releases)
[![Reference](https://pkg.go.dev/badge/github.com/aileron-projects/go-sockopts.svg)](https://pkg.go.dev/github.com/aileron-projects/go-sockopts)
[![DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/aileron-projects/go-sockopts)
[![Test](https://github.com/aileron-projects/go-sockopts/actions/workflows/test.yaml/badge.svg)](https://github.com/aileron-projects/go-sockopts/actions/workflows/test.yaml)

[![Insights](https://badgen.net/badge/Insights/open%2Fsource%2Finsights/cyan)](https://deps.dev/go/github.com%2Faileron-projects%2Fgo-sockopts)
[![Insights](https://badgen.net/badge/Insights/OSS%2FInsight/orange)](https://ossinsight.io/analyze/aileron-projects/go-sockopts)

</div>

# go-sockopts

**Use network socket options easily on Go.**

## Features

- Simple
- Easy to use
- Apply multiple socket options
- Control function for net.Listener
- Custom option
- Multi platform support

## Usages

### Basic usage of single option

With this library, socket options can be easily set.

**With the library:**

```go
// Single socket option with this library.
lc := &net.ListenConfig{
    Control: sockopts.SO_REUSEPORT.SetFunc(true).Control,
}
```

**Without library:**

```go
// Single socket option without this library.
// Codes depends on platforms because the example uses `unix` package.
lc := &net.ListenConfig{
    Control: func(network, address string, c syscall.RawConn) error {
        var inner, outer error
        outer = c.Control(func(fd uintptr) {
            inner = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
        })
        return cmp.Or(inner, outer)
    },
}
```

### Basic usage of multiple options

**With the library:**

```go
// Multiple socket options with this library.
lc := &net.ListenConfig{
    Control: sockopts.Control(
        sockopts.SO_REUSEADDR.SetFunc(true),
        sockopts.SO_REUSEPORT.SetFunc(true),
    ),
}
```

**Without library:**

```go
// Multiple socket options without this library.
// Codes depends on platforms because the example uses `unix` package.
lc := &net.ListenConfig{
    Control: func(network, address string, c syscall.RawConn) error {
        var inner1, inner2, outer error
        outer = c.Control(func(fd uintptr) {
            inner1 = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
            inner2 = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
        })
        return cmp.Or(inner1, inner2, outer)
    },
}
```

### Run a server with SO_REUSEPORT

`SO_REUSEPORT` is one of the popular socket option to sahre the same port from multiple applications.

```go
lc := &net.ListenConfig{
    Control: sockopts.SO_REUSEPORT.SetFunc(true).Control,
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

if err := svr.Serve(ln); err != nil && err != http.ErrServerClosed {
    panic(err)
}
```

### Custom option

Custom options can be easily defined.
Choose appropriate data type and level/opt.

The Name field is used in error messages.

```go
var SO_KEEPALIVE = sockopts.Bool{
    Level: unix.SOL_SOCKET,
    Opt:   unix.SO_KEEPALIVE,
    Name:  "SO_KEEPALIVE",
}
```

## Docs & Examples

- GoDoc: <https://pkg.go.dev/github.com/aileron-projects/go-sockopts>
- Examples:
  - Use single option: [examples/single_opt/](./examples/single_opt/)
  - Use multiple options: [examples/multiple_opts/](./examples/multiple_opts/)
  - Custom option: [examples/custom/](./examples/custom/)
  - SO_REUSEADDR: [examples/reuseaddr/](./examples/reuseaddr/)
  - SO_REUSEPORT: [examples/reuseport/](./examples/reuseport/)

## References
