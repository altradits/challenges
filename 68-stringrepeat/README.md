# 41. String Repeat

## Instructions

Write a program that repeats a string n times.

### Function Signature

```go
func Repeat(s string, n int) string
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", 3` | `"abcabcabc"` |
| `"x", 5` | `"xxxxx"` |
| `"go", 0` | `""` |
| `"", 10` | `""` |

### Constraints

- Return empty string if n <= 0
- Do not use `strings.Repeat`
- Use `strings.Builder` for efficiency

### Hints

- Loop n times and append the string each iteration
- Handle negative n values

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 68
```

or run the tests directly:

```bash
go test ./68-stringrepeat/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 68` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
