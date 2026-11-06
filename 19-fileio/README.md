# 19-fileio — File I/O

## Task

Implement in `package piscine`:

```go
// ReadLines returns all lines from a file (without newline characters).
func ReadLines(path string) ([]string, error)

// WriteLines writes lines to a file, one per line.
// Creates the file if it doesn't exist, overwrites if it does.
func WriteLines(path string, lines []string) error

// WordCount reads a file and returns a map of word → frequency.
func WordCount(path string) (map[string]int, error)
```

Read `skills.md` before you start.

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 19
```

or run the tests directly:

```bash
go test ./19-fileio/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 19` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
