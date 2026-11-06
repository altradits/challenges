# 53. String Format

## Instructions

Write a program that formats a string using positional arguments, similar to `fmt.Sprintf`.

### Function Signature

```go
func Format(template string, args ...interface{}) string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello %s", "World"` | `"Hello World"` |
| `"%d + %d = %d", 1, 2, 3` | `"1 + 2 = 3"` |
| `"%s is %d years old", "Alice", 30` | `"Alice is 30 years old"` |
| `"No placeholders"` | `"No placeholders"` |

### Constraints

- Support `%s` (string), `%d` (integer), `%f` (float)
- Return original template if no placeholders
- Handle mismatched argument counts gracefully

### Hints

- Use `fmt.Sprintf` internally or implement manual replacement
- Count placeholders and match with args

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 80
```

or run the tests directly:

```bash
go test ./80-stringformat/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 80` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
