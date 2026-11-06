# 44. String Contains

## Instructions

Write a program that checks if a string contains a substring.

### Function Signature

```go
func Contains(s, substr string) bool
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "world"` | `true` |
| `"hello world", "xyz"` | `false` |
| `"", "a"` | `false` |
| `"abc", ""` | `true` |

### Constraints

- Return true if substr is empty
- Return false if s is empty and substr is not
- Do not use `strings.Contains`

### Hints

- Use `strings.Index` to check for substring presence

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 71
```

or run the tests directly:

```bash
go test ./71-stringcontains/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 71` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
