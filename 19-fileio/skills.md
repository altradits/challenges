# Skills for 34-fileio

## What You'll Learn

**Previous:** [18-typeassertions](../18-typeassertions/skills.md) | **Next:** [20-regexp](../20-regexp/skills.md)

**Challenge:** Read and write files using the `os` and `bufio` packages.

## Reading a Whole File

```go
import "os"

data, err := os.ReadFile("hello.txt")
if err != nil {
    return err
}
fmt.Println(string(data))
```

`os.ReadFile` reads the entire file into a `[]byte`. Convert to string to work with it as text.

## Writing a Whole File

```go
content := []byte("Hello, file!\n")
err := os.WriteFile("output.txt", content, 0644)
```

`0644` is the file permission: owner can read+write, group and others can read.

## Opening a File for Streaming

For large files, read line-by-line instead of loading everything at once:

```go
import (
    "bufio"
    "os"
)

f, err := os.Open("data.txt")   // opens for reading only
if err != nil {
    return err
}
defer f.Close()   // ALWAYS close the file when done

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()   // line without the trailing newline
    fmt.Println(line)
}
if err := scanner.Err(); err != nil {
    return err
}
```

**`defer f.Close()`** ensures the file is closed even if the function returns early (error, panic, etc.).

## Creating / Writing a File

```go
f, err := os.Create("output.txt")  // creates or truncates
if err != nil {
    return err
}
defer f.Close()

writer := bufio.NewWriter(f)
fmt.Fprintln(writer, "line 1")
fmt.Fprintln(writer, "line 2")
writer.Flush()  // must flush buffered writer before closing
```

## `io.Reader` and `io.Writer`

These are the two most important interfaces in Go's standard library:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`*os.File`, `*bytes.Buffer`, `*strings.Reader`, `*bufio.Reader` — all implement `io.Reader`.
`*os.File`, `*bytes.Buffer`, `*bufio.Writer` — all implement `io.Writer`.

Functions that accept `io.Reader` or `io.Writer` work with any of these transparently.

### `io.Copy`

```go
import "io"

// Copy from src to dst
n, err := io.Copy(dst, src)
```

Example: copy a file:
```go
src, _ := os.Open("source.txt")
dst, _ := os.Create("dest.txt")
defer src.Close()
defer dst.Close()
io.Copy(dst, src)
```

## Appending to a File

Use `os.OpenFile` when you need more control than `os.Create` or `os.Open` give you:

```go
f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    return err
}
defer f.Close()
fmt.Fprintln(f, "new log entry")
```

| Flag | Meaning |
|------|---------|
| `os.O_RDONLY` | read only |
| `os.O_WRONLY` | write only |
| `os.O_RDWR` | read + write |
| `os.O_CREATE` | create if not exists |
| `os.O_APPEND` | append to end of file |
| `os.O_TRUNC` | truncate on open |

## `bufio.Scanner` — Controlling the Buffer

By default, `bufio.Scanner` splits on newlines and has a 64KB line limit. For larger lines:

```go
scanner := bufio.NewScanner(f)
scanner.Buffer(make([]byte, 1024*1024), 1024*1024)  // 1MB buffer

// Split by words instead of lines:
scanner.Split(bufio.ScanWords)

// Split by bytes:
scanner.Split(bufio.ScanBytes)
```

Always check `scanner.Err()` after the loop — it distinguishes EOF (normal) from real errors.

## Working with `strings.Reader` and `bytes.Buffer`

Any `io.Reader` works wherever a file does — useful for tests:

```go
import (
    "strings"
    "bytes"
)

// Read from an in-memory string
r := strings.NewReader("line1\nline2\n")
scanner := bufio.NewScanner(r)

// Write to an in-memory buffer
var buf bytes.Buffer
fmt.Fprintln(&buf, "hello")
fmt.Fprintln(&buf, "world")
fmt.Println(buf.String())  // hello\nworld\n
```

This is the pattern used in tests: pass a `bytes.Buffer` instead of a real file.

## Solving the Challenge

```go
package piscine

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func ReadLines(path string) ([]string, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var lines []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func WriteLines(path string, lines []string) error {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()

    w := bufio.NewWriter(f)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

func WordCount(path string) (map[string]int, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    freq := make(map[string]int)
    for _, word := range strings.Fields(string(data)) {
        freq[word]++
    }
    return freq, nil
}
```

**Next:** [20-regexp](../20-regexp/README.md)
