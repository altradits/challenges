# 51. String Filter

## Instructions

Write a program that filters characters from a string based on a predicate function.

### Function Signature

```go
func Filter(s string, f func(rune) bool) string
```

### Examples

| Input | Output |
|-------|--------|
| `"hello", IsVowel` | `"hll"` |
| `"abc123", IsDigit` | `"123"` |
| `"Go Lang", IsUpper` | `"GL"` |
| `""` | `""` |

### Constraints

- Keep characters where f(r) returns true
- Return empty string if no characters match
- Handle empty input

### Hints

- Use `strings.Builder` to build result
- Apply predicate to each rune

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 78
```

or run the tests directly:

```bash
go test ./78-stringfilter/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 78` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
