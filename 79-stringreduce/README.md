# 52. String Reduce

## Instructions

Write a program that reduces a string to a single value by applying a function cumulatively.

### Function Signature

```go
func Reduce(s string, f func(rune, rune) rune) rune
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", Max` | `'c'` |
| `"hello", Min` | `'e'` |
| `"123", func(a, b rune) rune { return a + b }` | `'f'` (49+50+51) |
| `""` | `0` |

### Constraints

- Return 0 for empty string
- Apply function left-to-right
- Handle single character strings

### Hints

- Initialize with first rune, then apply to rest
- Use `for...range` to iterate runes

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 79
```

or run the tests directly:

```bash
go test ./79-stringreduce/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 79` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
