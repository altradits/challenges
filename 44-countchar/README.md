# 17. Count Character

## What You'll Learn

This exercise teaches you **counting specific characters and accumulation patterns**. By the end, you'll understand:
- How to count occurrences of a specific value
- The difference between counting and detecting
- Accumulating results across iterations
- Handling edge cases in counting

## Theory: Counting Patterns

### Counting vs. Detecting

| Goal | Approach | Return |
|------|----------|--------|
| Detect if exists | Return on first match | `bool` |
| Count all matches | Count every match | `int` |

### The Accumulator Pattern

Counting requires an **accumulator** variable:

```go
count := 0           // Initialize
for _, c := range s {
    if c == target {
        count++      // Accumulate
    }
}
return count         // Return total
```

### Early Return vs. Full Scan

| Scenario | Approach |
|----------|----------|
| Just need to know IF it exists | Return on first match |
| Need to know HOW MANY | Scan entire string |

## Step-by-Step Approach

1. **Initialize** counter to 0
2. **Iterate** through each character
3. **Compare** with target character
4. **Increment** counter on match
5. **Return** final count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Returning on first match | Only counts 1 | Continue scanning |
| Forgetting to initialize | Counter starts at garbage | `count := 0` |
| Wrong comparison | `=` instead of `==` | Use `==` for comparison |

## Practice Tips

- Test with: `"Hello"`, count `'l'` → 2
- Test with: `"Hello"`, count `'z'` → 0
- Test with: `""`, count `'a'` → 0
- Remember: case-sensitive! `'A'` != `'a'`

## The Challenge

Write a function that counts occurrences of a specific character.

### Expected Function

```go
func CountChar(s string, c rune) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `("Hello World", 'l')` | `3` | Three 'l's |
| `("5  balloons", ' ')` | `3` | Three spaces |
| `("   ", ' ')` | `3` | All spaces |
| `("The 7 deadly sins", '7')` | `1` | One '7' |
| `("Hello", 'z')` | `0` | Not found |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CountChar("Hello World", 'l'))  // 3
    fmt.Println(CountChar("5  balloons", ' '))  // 3
    fmt.Println(CountChar("   ", ' '))           // 3
    fmt.Println(CountChar("The 7 deadly sins", '7')) // 1
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the difference between detecting and counting?
2. Why do we need an accumulator variable?
3. Should counting stop at the first match?

## Next Steps

After completing this, you'll be ready for:
- [45-findlastchar](../45-findlastchar/README.md) - Findlastchar
- [46-replacechar](../46-replacechar/README.md) - Replacechar