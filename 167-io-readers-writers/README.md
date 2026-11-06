# 167-io-readers-writers — The io Package: Readers and Writers

## Task

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

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 167
```

or run the tests directly:

```bash
go test ./167-io-readers-writers/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 167` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
