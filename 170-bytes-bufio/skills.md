# Skills for 170-bytes-bufio

## What You'll Learn

**Previous:** [169-testing-advanced](../169-testing-advanced/skills.md) | **Next:** [171-encoding-binary](../171-encoding-binary/skills.md)

**Challenge:** Process text with `bytes` and `bufio` instead of `strings`.

## Why bytes and bufio?

`strings` works on immutable strings; every concatenation allocates.
`bytes` works on `[]byte`, and `bufio` batches reads and writes so a
program that parses megabytes (logs, block data, test fixtures) stays fast.
Bitcoin Core-style code in Go leans on `bytes.Buffer` for message building.

## The bytes Package

`bytes` mirrors `strings` for byte slices:

```go
b := []byte("go  go test")
words := bytes.Fields(b)        // [][]byte{{'g','o'},{'g','o'},{'t',...}}
joined := bytes.Join(words, []byte(" ")) // []byte("go go test")
n := bytes.Count(b, []byte("go"))        // 2
upper := bytes.ToUpper([]byte("ab"))     // []byte("AB")
```

`bytes.Buffer` accumulates output without repeated allocation:

```go
var buf bytes.Buffer
buf.WriteString("a")
buf.WriteByte('\n')
buf.Write([]byte("b"))
out := buf.String() // "a\nb"
```

## bufio.Scanner — Line by Line

```go
sc := bufio.NewScanner(strings.NewReader("a\nb\n"))
for sc.Scan() {
    line := sc.Text() // "a", then "b"
}
if err := sc.Err(); err != nil {
    // handle read error
}
```

Notes: `Scan` returns false at EOF; the trailing newline is stripped; a
final line without `\n` is still returned. Default token buffer caps at
64KB per line — fine here.

## bufio.Writer — Batched Writes

```go
var buf bytes.Buffer
w := bufio.NewWriter(&buf)
w.WriteString("a\n")
w.WriteString("b")
w.Flush() // REQUIRED: pushes buffered bytes to buf
```

Forgetting `Flush` silently drops output — the classic bufio bug.

## Solving the Challenge

- `CountLines`: scan with `bufio.Scanner`, count `Scan()` successes.
- `WordFreq`: `bytes.Fields`, then count into a map. Return an empty
  (non-nil) map for empty input.
- `JoinLines`: `bytes.Join` with `[]byte("\n")`, or a `bytes.Buffer` loop.
- `ReverseWords`: `bytes.Fields` (collapses runs of whitespace), reverse
  the slice, `bytes.Join` with a single space.
