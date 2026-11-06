# 36. Union of Characters

## What You'll Learn

This exercise teaches you **set union operations on strings**. By the end, you'll understand:
- How to compute the union of two strings
- Removing duplicate characters while preserving order
- Using maps for membership testing
- The difference between union and intersection

## Theory: Set Union

### What is Union?

Union combines elements from both sets, without duplicates:

```
"abc" ∪ "bcd" = "abcd"
"hello" ∪ "world" = "helowrd"
```

### Preserving Order

The union should preserve the order characters first appear:

```
"zpadinton" ∪ "paqefwtdjetyiytjneytjoeyjnejeyj"
= "zpadintoqefwjy"
```

### Using Maps for Lookup

```go
seen := make(map[rune]bool)
for _, c := s1 + s2 {
    if !seen[c] {
        seen[c] = true
        result += string(c)
    }
}
```

## Step-by-Step Approach

1. **Initialize** empty map and result string
2. **Iterate** through first string, add unique chars
3. **Iterate** through second string, add unique chars
4. **Return** result

## The Challenge

Write a function that returns the union of two strings.

### Expected Function

```go
func Union(s1, s2 string) string {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")` | `"zpadintoqefwjy"` |
| `("ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd")` | `"df6vewg4thras"` |

## Next Steps

After completing this, you'll be ready for:
- [64-inter](../64-inter/README.md) - Inter
- [65-stringbuilder](../65-stringbuilder/README.md) - Stringbuilder

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 63
```

or run the tests directly:

```bash
go test ./63-union/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 63` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
