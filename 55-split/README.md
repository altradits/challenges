# 28. Split

## What You'll Learn

This exercise teaches you **tokenizing strings and slice manipulation**. By the end, you'll understand:
- How to split a string into parts
- Working with slices of strings
- Handling edge cases in splitting
- The difference between split and manual parsing

## Theory: String Splitting

### What is Splitting?

Splitting breaks a string into a slice of substrings:

```
"a,b,c" split by "," → ["a", "b", "c"]
"Hello World" split by " " → ["Hello", "World"]
```

### The Split Algorithm

```go
func split(s, sep string) []string {
    var result []string
    start := 0
    for i := 0; i <= len(s); i++ {
        if i == len(s) || s[i:i+len(sep)] == sep {
            result = append(result, s[start:i])
            start = i + len(sep)
            i += len(sep) - 1  // Skip past separator
        }
    }
    return result
}
```

## Step-by-Step Approach

1. **Initialize** empty result slice and start index
2. **Loop** through the string
3. **When separator found**: extract substring, add to result
4. **Update** start index past separator
5. **Return** result

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Off-by-one in slicing | Missing or extra chars | Careful with indices |
| Forgetting last segment | Trailing text missed | Handle end of string |
| Empty separator | Infinite loop | Check for empty sep |

## Task

Write a function that splits a string by a separator.

### Expected Function

```go
func Split(s, sep string) []string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("a,b,c", ",")` | `["a", "b", "c"]` | Basic split |
| `("Hello World", " ")` | `["Hello", "World"]` | Space separator |
| `("abc", "")` | `["a", "b", "c"]` | Empty sep = split by char |
| `("", ",")` | `[""]` | Empty string |

## Knowledge Check

Before coding, make sure you can answer:
1. How do you extract a substring using slicing?
2. What happens when the separator is at the start or end?
3. How do you handle an empty separator?

## Next Steps

After completing this, you'll be ready for:
- [56-join](../56-join/README.md) - Join
- [57-cameltosnakecase](../57-cameltosnakecase/README.md) - Cameltosnakecase

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 55
```

or run the tests directly:

```bash
go test ./55-split/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 55` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
