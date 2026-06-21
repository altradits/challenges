# Skills for hiddenp

## What You'll Learn

**Previous:** [53-fprime](../53-fprime/skills.md) | **Next:** [55-inter](../55-inter/README.md)

**Challenge:** Check if every character of `s1` appears in `s2` in the same relative order (not necessarily consecutive). Print `1` if true, `0` if false.

## Core Concept: Subsequence Matching

### What Is It?

A string `s1` is a **subsequence** of `s2` if you can find all characters of `s1` in `s2`, in order, by only moving forward through `s2`. You do not need consecutive positions.

Examples:
- `"abc"` is a subsequence of `"2altrb53c.sse"` → find `a` (pos 4), then `b` (pos 5+), then `c` (pos 7+) — YES
- `"abc"` is NOT a subsequence of `"btarc"` → `a` at pos 1, `b` at pos 0 — `b` comes before `a`, fails

### How It Works

**Key idea:** Use one pointer into `s1` that only advances when you find a match in `s2`.

```go
func isHidden(s1, s2 string) bool {
    i := 0  // pointer into s1
    for _, c := range s2 {
        if i < len(s1) && c == rune(s1[i]) {
            i++  // matched the next character of s1
        }
    }
    return i == len(s1)  // did we match all of s1?
}
```

**Step by step for `"abc"` in `"2altrb53c.sse"`:**

| `s2` char | `s1[i]` | Match? | i |
|-----------|---------|--------|---|
| `2` | `a` | no | 0 |
| `a` | `a` | YES | 1 |
| `l` | `b` | no | 1 |
| `t` | `b` | no | 1 |
| `r` | `b` | no | 1 |
| `b` | `b` | YES | 2 |
| `5` | `c` | no | 2 |
| `3` | `c` | no | 2 |
| `c` | `c` | YES | 3 |
| ... | | | 3 |

Final: `i == len("abc") == 3` → return `true` → print `1`

**Main function:**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        return
    }
    s1 := os.Args[1]
    s2 := os.Args[2]

    i := 0
    for _, c := range s2 {
        if i < len(s1) && c == rune(s1[i]) {
            i++
        }
    }
    if i == len(s1) {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}
```

**Edge case — empty s1:** `i` starts at 0 and `len(s1) == 0`, so `i == len(s1)` is immediately true. Print `1`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Resetting `i` when you don't match | Treats every position independently, not in order | Only increment `i`, never reset it |
| Checking `s2[j]` as a byte against `s1[i]` as a byte | Works for ASCII but fails for multi-byte UTF-8 characters | Use `for _, c := range s2` and `rune(s1[i])` |
| Not handling `len(os.Args) != 3` | Panic or wrong output | Return immediately if arg count is wrong |
| Printing `0` for empty `s1` | Empty string is always hidden | Check `i == len(s1)` — which is true when `len(s1) == 0` |

## Solving This Challenge

### Algorithm
1. Require exactly 2 arguments; return otherwise.
2. Set `i = 0`.
3. For each character `c` in `s2`: if `i < len(s1)` and `c == rune(s1[i])`, increment `i`.
4. If `i == len(s1)`, print `1`; else print `0`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [55-inter](../55-inter/README.md)
