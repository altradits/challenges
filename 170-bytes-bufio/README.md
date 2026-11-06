# 170-bytes-bufio — Bytes and Buffered I/O

## Task

In `package piscine`, write four text-processing helpers using only
`bytes` and `bufio` — no `strings` package, no manual loops over runes:

```go
func CountLines(s string) int
func WordFreq(s string) map[string]int
func JoinLines(lines []string) string
func ReverseWords(s string) string
```

Read `skills.md` before you start.

## Input / Output

- `CountLines` counts newline-separated lines (`"a\nb\n"` → 2).
- `WordFreq` maps each whitespace-separated word to its count.
- `JoinLines` joins lines with `"\n"` (no trailing newline).
- `ReverseWords` reverses word order, single spaces (`"a  b"` → `"b a"`).

## Examples

```
CountLines("a\nb\nc")          →  3
WordFreq("go go test")         →  map[go:2 test:1]
JoinLines([]string{"a","b"})   →  "a\nb"
ReverseWords("hello old world") →  "world old hello"
```

## Constraints

- Use `bytes.Fields`, `bytes.Join`/`bytes.Buffer`, `bufio.Scanner` or
  `bufio.Writer` where they fit; do not import `strings`.
- Empty input: `CountLines("")` → 0, `ReverseWords("")` → `""`,
  `JoinLines(nil)` → `""`, `WordFreq("")` → empty map.

## Time limit

**20 minutes** (round 1). Beat it on every repeat: each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

    go run ./cmd/drill check 170

or run the tests directly:

    go test ./170-bytes-bufio/

Green means done. Red means read the failure message, fix from memory, check again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, beat your previous time. If you cannot, re-read `skills.md` once,
then close it and go again.
