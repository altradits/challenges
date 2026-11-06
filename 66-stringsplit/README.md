# 39. String Split

## Instructions

Write a program that splits a string into words and returns them as a slice.

### Function Signature

```go
func SplitWords(s string) []string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello World"` | `["Hello", "World"]` |
| `"  Go   is  fun  "` | `["Go", "is", "fun"]` |
| `""` | `[]` |
| `"single"` | `["single"]` |

### Constraints

- Handle multiple consecutive spaces correctly
- Trim leading and trailing spaces
- Return an empty slice for empty input
- Do not use `strings.Split` - implement the logic manually

### Hints

- Use `strings.Fields` or implement manual splitting
- Consider edge cases with whitespace

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 66
```

or run the tests directly:

```bash
go test ./66-stringsplit/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 66` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
