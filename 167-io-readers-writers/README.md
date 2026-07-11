# 167-io-readers-writers — The io Package: Readers and Writers

## Challenge

Implement in `package piscine`:

```go
// ReadAll reads all bytes from r and returns them.
func ReadAll(r io.Reader) ([]byte, error)

// WriteString writes s to w and returns bytes written and any error.
func WriteString(w io.Writer, s string) (int, error)

// Copy copies all bytes from src to dst, returns bytes copied.
func Copy(dst io.Writer, src io.Reader) (int64, error)

// CountLines counts the number of newline characters in r.
func CountLines(r io.Reader) (int, error)

// UpperWriter returns a Writer that uppercases all bytes before writing to w.
func UpperWriter(w io.Writer) io.Writer
```

## Examples

```go
// ReadAll from strings.NewReader
data, _ := ReadAll(strings.NewReader("hello world"))
// data == []byte("hello world")

// WriteString to bytes.Buffer
var buf bytes.Buffer
n, _ := WriteString(&buf, "hello")
// n == 5, buf.String() == "hello"

// CountLines
n, _ := CountLines(strings.NewReader("a\nb\nc"))
// n == 2  (two newlines)

// UpperWriter
var buf bytes.Buffer
w := UpperWriter(&buf)
w.Write([]byte("hello"))
// buf.String() == "HELLO"
```

Read `skills.md` before you start.
