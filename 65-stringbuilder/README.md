# 38. String Builder

## Instructions

Write a program that takes a string and returns a new string with all vowels removed.

Vowels are: a, e, i, o, u (both lowercase and uppercase).

### Function Signature

```go
func RemoveVowels(s string) string
```

### Examples

| Input | Output |
|-------|--------|
| `"Hello World"` | `"Hll Wrld"` |
| `"AEIOUaeiou"` | `""` |
| `"Go Programming"` | `"G Prgrmmng"` |
| `""` | `""` |

### Constraints

- Do not use the `strings.ReplaceAll` function
- Build the result using a `strings.Builder` for efficiency
- Preserve the original case of consonants
- Handle empty strings gracefully

### Hints

- Use `strings.Builder` to efficiently build the result string
- Iterate through each rune in the input string
- Check if each rune is a vowel before appending

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 65
```

or run the tests directly:

```bash
go test ./65-stringbuilder/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 65` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
