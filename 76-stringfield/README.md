# 49. String Field

## Instructions

Write a program that splits a string into fields using a delimiter.

### Function Signature

```go
func Split(s, sep string) []string
```

### Examples

| Input | Output |
|-------|--------|
| `"a,b,c", ","` | `["a", "b", "c"]` |
| `"a:b:c", ":"` | `["a", "b", "c"]` |
| `"abc", "x"` | `["abc"]` |
| `"", ","` | `[""]` |

### Constraints

- Return slice with one empty string if input is empty
- Do not use `strings.Split`
- Handle consecutive delimiters (empty fields)

### Hints

- Use `strings.Index` to find each delimiter
- Build result slice incrementally

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 76
```

or run the tests directly:

```bash
go test ./76-stringfield/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 76` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
