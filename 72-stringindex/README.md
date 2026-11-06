# 45. String Index

## Instructions

Write a program that finds the index of the first occurrence of a substring.

### Function Signature

```go
func Index(s, substr string) int
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "world"` | `6` |
| `"hello world", "xyz"` | `-1` |
| `"abcabc", "b"` | `1` |
| `"", "a"` | `-1` |

### Constraints

- Return -1 if substring not found
- Return 0 if substring is empty
- Do not use `strings.Index`

### Hints

- Iterate through the string checking for matches
- Use rune-based iteration for Unicode safety

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 72
```

or run the tests directly:

```bash
go test ./72-stringindex/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 72` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
