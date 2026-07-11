# Skills for 167-io-readers-writers

## What You'll Learn

**Previous:** [166-generics-intro](../166-generics-intro/skills.md) | **Next:** [168-sync-mutex](../168-sync-mutex/skills.md)

**Challenge:** Use Go's `io` package to read from and write to any data source uniformly.

## The `io.Reader` and `io.Writer` Interfaces

These two interfaces are the most widely used in the entire Go standard library. Everything that reads or writes bytes implements one or both.

```go
// From package io
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`Read` fills the buffer `p` and returns how many bytes were read. It returns `io.EOF` when there is no more data.

`Write` writes the contents of `p` and returns how many bytes were written.

### Everything Is a Reader or Writer

| Type | Implements |
|------|-----------|
| `os.File` | Reader, Writer |
| `bytes.Buffer` | Reader, Writer |
| `strings.NewReader(s)` | Reader |
| `net.Conn` | Reader, Writer |
| `http.Request.Body` | Reader |
| `http.ResponseWriter` | Writer |
| `bufio.Reader` | Reader |
| `gzip.Writer` | Writer |

Because all these types implement the same interface, you can write code that works with any of them.

## Key Functions

### `io.ReadAll`

Reads everything from a Reader into a byte slice:

```go
import "io"

data, err := io.ReadAll(r)
```

Equivalent to your `ReadAll` function.

### `io.WriteString`

Writes a string to a Writer:

```go
n, err := io.WriteString(w, "hello")
```

### `io.Copy`

Copies from a Reader to a Writer in chunks (efficient, handles large files):

```go
n, err := io.Copy(dst, src)
```

### `io.MultiWriter`

Writes to multiple writers simultaneously:

```go
w := io.MultiWriter(os.Stdout, logFile, &buffer)
io.WriteString(w, "logged everywhere at once")
```

### `io.TeeReader`

Reads from r while also writing everything read to w:

```go
tee := io.TeeReader(src, log)
data, _ := io.ReadAll(tee)  // src is read and log gets a copy
```

## Custom Readers and Writers

Any type that implements `Read(p []byte) (int, error)` is a Reader. This lets you wrap other sources:

```go
// An uppercase writer — wraps another writer
type upperWriter struct {
    w io.Writer
}

func (u upperWriter) Write(p []byte) (int, error) {
    upper := bytes.ToUpper(p)
    return u.w.Write(upper)
}

func UpperWriter(w io.Writer) io.Writer {
    return upperWriter{w: w}
}
```

## `bufio` — Buffered I/O

Raw `Read` and `Write` may be called many times for small chunks. `bufio` wraps any Reader or Writer with a buffer to reduce system calls:

```go
import "bufio"

// Buffered reader — read line by line
scanner := bufio.NewScanner(r)
for scanner.Scan() {
    line := scanner.Text()
}

// Buffered writer — flush when done
bw := bufio.NewWriter(w)
bw.WriteString("hello\n")
bw.Flush()  // must call Flush or data stays buffered
```

### Counting Lines

```go
func CountLines(r io.Reader) (int, error) {
    scanner := bufio.NewScanner(r)
    count := 0
    for scanner.Scan() {
        count++
    }
    return count, scanner.Err()
}
```

Wait — this counts lines (each call to `Scan()` gives one line). Number of newlines = number of lines when the file ends with a newline. See the solution below for an exact newline counter.

## Solving the Challenge

```go
package piscine

import (
    "bytes"
    "io"
    "bufio"
)

func ReadAll(r io.Reader) ([]byte, error) {
    return io.ReadAll(r)
}

func WriteString(w io.Writer, s string) (int, error) {
    return io.WriteString(w, s)
}

func Copy(dst io.Writer, src io.Reader) (int64, error) {
    return io.Copy(dst, src)
}

func CountLines(r io.Reader) (int, error) {
    data, err := io.ReadAll(r)
    if err != nil {
        return 0, err
    }
    return bytes.Count(data, []byte{'\n'}), nil
}

type upperWriter struct{ w io.Writer }

func (u upperWriter) Write(p []byte) (int, error) {
    return u.w.Write(bytes.ToUpper(p))
}

func UpperWriter(w io.Writer) io.Writer {
    return upperWriter{w: w}
}
```

## The Pipe Pattern

```go
// io.Pipe connects a writer to a reader
pr, pw := io.Pipe()

go func() {
    pw.Write([]byte("streamed data"))
    pw.Close()
}()

data, _ := io.ReadAll(pr)
```

## Documentation

- [io package](https://pkg.go.dev/io)
- [bufio package](https://pkg.go.dev/bufio)
- [bytes package](https://pkg.go.dev/bytes)
- [Go Blog — The Go image package (Reader pattern)](https://go.dev/blog/image-package)
- [Effective Go — Interfaces](https://go.dev/doc/effective_go#interfaces)

**Next:** [168-sync-mutex](../168-sync-mutex/README.md)
