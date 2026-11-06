# 47. String Prefix

## Instructions

Write a program that checks if a string starts with a given prefix.

### Function Signature

```go
func HasPrefix(s, prefix string) bool
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "hello"` | `true` |
| `"hello world", "world"` | `false` |
| `"", ""` | `true` |
| `"abc", "abcd"` | `false` |

### Constraints

- Return true if prefix is empty
- Return false if s is shorter than prefix
- Do not use `strings.HasPrefix`

### Hints

- Compare the first len(prefix) characters of s

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 74
```

or run the tests directly:

```bash
go test ./74-stringprefix/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 74` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
