# 50. String Map

## Instructions

Write a program that applies a function to each character in a string.

### Function Signature

```go
func Map(s string, f func(rune) rune) string
```

### Examples

| Input | Output |
|-------|--------|
| `"abc", ToUpper` | `"ABC"` |
| `"Hello", ToLower` | `"hello"` |
| `"123", func(r rune) rune { return r + 1 }` | `"234"` |
| `""` | `""` |

### Constraints

- Apply the function to each rune
- Return the transformed string
- Handle empty strings

### Hints

- Use `strings.Builder` to build the result
- Iterate with `for...range` to handle runes correctly

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 77
```

or run the tests directly:

```bash
go test ./77-stringmap/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 77` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
