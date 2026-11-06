# 42. String Replace

## Instructions

Write a program that replaces all occurrences of a substring with another substring.

### Function Signature

```go
func Replace(s, old, new string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"hello world", "l", "L"` | `"heLLo worLd"` |
| `"foo bar foo", "foo", "baz"` | `"baz bar baz"` |
| `"abc", "x", "y"` | `"abc"` |
| `"", "a", "b"` | `""` |

### Constraints

- Replace all occurrences, not just the first
- Handle empty strings gracefully
- Do not use `strings.ReplaceAll`

### Hints

- Use `strings.Index` to find each occurrence
- Build the result incrementally

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 69
```

or run the tests directly:

```bash
go test ./69-stringreplace/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 69` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
