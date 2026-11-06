# 40. String Join

## Instructions

Write a program that joins a slice of strings with a separator.

### Function Signature

```go
func Join(words []string, sep string) string
```

### Examples

| Input | Output |
|-------|--------|
| `["a","b","c"], "-"` | `"a-b-c"` |
| `["hello","world"], " "` | `"hello world"` |
| `[], ","` | `""` |
| `["single"], "|"` | `"single"` |

### Constraints

- Return empty string for empty slice
- Do not use `strings.Join` - implement manually
- Handle single-element slices correctly

### Hints

- Use `strings.Builder` for efficient concatenation
- Add separator between elements, not after the last one

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 67
```

or run the tests directly:

```bash
go test ./67-stringjoin/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 67` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
