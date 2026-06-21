# 13. Count Repeats

## What You'll Learn

This exercise teaches you **pattern detection and consecutive character analysis**. By the end, you'll understand:
- How to detect consecutive repeated characters
- Tracking state across iterations
- Counting groups vs. counting individual characters
- The difference between "repeat" and "occurrence"

## Theory: Consecutive Patterns

### What is a "Repeat"?

A repeat is when the **same character appears consecutively**:

```
"hello"     → no repeats (all different consecutively)
"helloo"    → 1 repeat (the 'o' repeats)
"heelloo"   → 2 repeats (e repeats, o repeats)
"aaa"       → 1 repeat (a repeats as a group)
```

### Groups vs. Individual Counts

| String | Groups of Repeats | Count |
|--------|-------------------|-------|
| `"hello"` | None | 0 |
| `"helloo"` | `oo` | 1 |
| `"heelloo"` | `ee`, `oo` | 2 |
| `"aaa"` | `aaa` (one group) | 1 |

**Important:** We count GROUPS of consecutive repeats, not individual repeated characters!

### State Tracking

To detect repeats, we need to **remember the previous character**:

```go
prevChar := '\0'  // Null character (nothing seen yet)
repeatCount := 0
inRepeat := false

for _, c := range s {
    if c == prevChar {
        if !inRepeat {
            repeatCount++  // First time we see this repeat
            inRepeat = true
        }
    } else {
        inRepeat = false  // Different character, reset
    }
    prevChar = c
}
```

## Step-by-Step Approach

1. **Initialize**: previous char, repeat count, in-repeat flag
2. **Iterate** through each character
3. **Compare** current with previous
4. **If same AND not already in repeat**: increment count, set flag
5. **If different**: reset flag
6. **Update** previous character
7. **Return** count

## Common Pitfalls

| Mistake | Why It's Wrong | Correct Approach |
|---------|---------------|------------------|
| Counting each repeated char | `"aaa"` would give 2, not 1 | Count groups only |
| Forgetting the flag | `"aaa"` would count 2 repeats | Track if already in a repeat |
| Not handling first char | No previous to compare | Initialize with sentinel value |

## Practice Tips

- Trace through `"heelloo"`: h(no prev) → e(no match) → e(match! count=1) → l(no match) → l(match! count=2) → o(no match) → o(match! already in repeat, skip)
- The flag prevents counting `"aaa"` as 2 repeats

## The Challenge

Write a function that counts groups of consecutive repeated characters.

### Expected Function

```go
func CountRepeats(s string) int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output | Why |
|-------|-----------------|-----|
| `"hello"` | `0` | No consecutive repeats |
| `"helloo"` | `1` | One group: `oo` |
| `"heelloo"` | `2` | Two groups: `ee`, `oo` |
| `"abc"` | `0` | No repeats |
| `""` | `0` | Empty string |
| `"aaa"` | `1` | One group: `aaa` |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CountRepeats("hello"))    // 0
    fmt.Println(CountRepeats("helloo"))   // 1
    fmt.Println(CountRepeats("heelloo"))  // 2
    fmt.Println(CountRepeats("abc"))      // 0
    fmt.Println(CountRepeats(""))         // 0
    fmt.Println(CountRepeats("aaa"))      // 1
}
```

## Knowledge Check

Before coding, make sure you can answer:
1. What's the difference between counting repeats and counting repeated characters?
2. Why do we need a "in repeat" flag?
3. What is a sentinel value and why use it?

## Next Steps

After completing this, you'll be ready for:
- [26-retainfirsthalf](../26-retainfirsthalf/README.md) - Retainfirsthalf
- [27-wordcount](../27-wordcount/README.md) - Wordcount