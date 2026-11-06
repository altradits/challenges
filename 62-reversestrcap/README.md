# 35. Reverse String Capitalization

## What You'll Learn

This exercise teaches you **selective case transformation**. By the end, you'll understand:
- How to transform only specific characters in a string
- Word boundary detection
- Combining multiple string operations
- Processing command-line arguments

## Theory: Selective Case Changes

### The Pattern

Transform each word by:
1. Converting all characters to lowercase
2. Uppercasing only the LAST character

```
"Hello" → "hellO"
"SMALL" → "smalL"
```

### Word Detection

Words are separated by spaces. Use `strings.Fields` or manual splitting.

## Step-by-Step Approach

1. **Split** input into words
2. **For each word**:
   - Convert to lowercase
   - Uppercase last character
3. **Join** words with spaces
4. **Return** result

## The Challenge

Write a function that uppercases the last letter of each word.

### Expected Function

```go
func ReverseStrCap(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `"First SMALL TesT"` | `"firsT smalL tesT"` |
| `"Hello"` | `"hellO"` |
| `""` | `""` |

## Next Steps

After completing this, you'll be ready for:
- [63-union](../63-union/README.md) - Union
- [64-inter](../64-inter/README.md) - Inter

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 62
```

or run the tests directly:

```bash
go test ./62-reversestrcap/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 62` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
