# 32. Every Third Character

## What You'll Learn

This exercise teaches you **modular indexing and selective extraction**. By the end, you'll understand:
- Using modulo operator for periodic patterns
- Extracting characters at regular intervals
- Handling edge cases with short strings
- Building strings from selective characters

## Theory: Modulo for Patterns

### The Modulo Operator

`%` gives the remainder of division:

```
index % 3 == 0  // Every 3rd character (0, 3, 6, 9...)
index % 3 == 1  // 1st, 4th, 7th...
index % 3 == 2  // 2nd, 5th, 8th...
```

### Extracting Every Nth Character

```go
for i, c := range s {
    if i % n == n-1 {  // Every nth (1-indexed)
        result += string(c)
    }
}
```

## Step-by-Step Approach

1. **Iterate** with index using `for...range`
2. **Check** if `(index + 1) % 3 == 0` (every 3rd)
3. **Append** matching characters to result
4. **Return** result with newline

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Using `i % 3 == 0` | Gets 0, 3, 6... (1st, 4th, 7th) | Use `(i+1) % 3 == 0` for 3rd, 6th, 9th |
| Forgetting newline | Output missing `\n` | Add `\n` at end |
| Empty string handling | Should return just `\n` | Check length first |

## Task

Write a function that returns every 3rd character of a string.

### Expected Function

```go
func ThirdChar(s string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"123456789"` | `"369\n"` | 3rd, 6th, 9th |
| `""` | `"\n"` | Empty string |
| `"a b c d e f"` | `"b e\n"` | Every 3rd including spaces |
| `"12"` | `"\n"` | No 3rd character |

## Knowledge Check

Before coding, make sure you can answer:
1. What does `i % 3` give you?
2. How do you get the 3rd, 6th, 9th characters?
3. What should empty string return?

## Next Steps

After completing this, you'll be ready for:
- [60-zipstring](../60-zipstring/README.md) - Zipstring
- [61-saveandmiss](../61-saveandmiss/README.md) - Saveandmiss

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 59
```

or run the tests directly:

```bash
go test ./59-thirdchar/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 59` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
