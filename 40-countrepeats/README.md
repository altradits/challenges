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
"hello"     → 1 repeat (the 'll' group)
"helloo"    → 2 repeats ('ll' and 'oo' groups)
"heelloo"   → 3 repeats ('ee', 'll' and 'oo' groups)
"aaa"       → 1 repeat (a repeats as a group)
```

### Groups vs. Individual Counts

| String | Groups of Repeats | Count |
|--------|-------------------|-------|
| `"hello"` | `ll` | 1 |
| `"helloo"` | `ll`, `oo` | 2 |
| `"heelloo"` | `ee`, `ll`, `oo` | 3 |
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

- Trace through `"heelloo"`: h(no prev) → e(no match) → e(match! count=1) → l(no match) → l(match! count=2) → o(no match) → o(match! count=3)
- The flag prevents counting `"aaa"` as 2 repeats

## Task

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
| `"hello"` | `1` | One group: `ll` |
| `"helloo"` | `2` | Two groups: `ll`, `oo` |
| `"heelloo"` | `3` | Three groups: `ee`, `ll`, `oo` |
| `"abc"` | `0` | No repeats |
| `""` | `0` | Empty string |
| `"aaa"` | `1` | One group: `aaa` |

### Usage Example

```go
package main

import "fmt"

func main() {
    fmt.Println(CountRepeats("hello"))    // 1
    fmt.Println(CountRepeats("helloo"))   // 2
    fmt.Println(CountRepeats("heelloo"))  // 3
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
- [41-retainfirsthalf](../41-retainfirsthalf/README.md) - Retainfirsthalf
- [42-wordcount](../42-wordcount/README.md) - Wordcount

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 40
```

or run the tests directly:

```bash
go test ./40-countrepeats/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 40` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
