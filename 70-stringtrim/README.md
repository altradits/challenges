# 43. String Trim

## Instructions

Write a program that trims leading and trailing whitespace from a string.

### Function Signature

```go
func Trim(s string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"  hello  "` | `"hello"` |
| `"\t\nGo\n\t"` | `"Go"` |
| `"no spaces"` | `"no spaces"` |
| `"   "` | `""` |

### Constraints

- Remove spaces, tabs, and newlines
- Do not use `strings.TrimSpace`
- Return empty string if only whitespace

### Hints

- Use `unicode.IsSpace` to check for whitespace characters
- Find the first and last non-space indices

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 70
```

or run the tests directly:

```bash
go test ./70-stringtrim/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 70` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
