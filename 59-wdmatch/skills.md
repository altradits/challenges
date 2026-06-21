# Skills for wdmatch

## What You'll Learn

**Previous:** [58-union](../58-union/skills.md) | **Next:** [60-fifthandskip](../60-fifthandskip/README.md)

**Challenge:** Check if you can spell `word` using characters from `s` in the same order they appear in `s`. If yes, print `word`; otherwise print nothing.

## Core Concept: Subsequence Matching (Order-Respecting Search)

### What Is It?

This is the same subsequence matching pattern as `hiddenp` (54), but applied in a different context — here you print the word itself (not 1/0) and nothing on failure.

`word` is constructable from `s` if every character of `word` can be found in `s` in order, advancing through `s` only forward.

### How It Works

The algorithm is identical to hiddenp's single-pointer approach:

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
    word := os.Args[1]
    s    := os.Args[2]

    i := 0  // pointer into word
    for _, c := range s {
        if i < len(word) && c == rune(word[i]) {
            i++
        }
    }
    if i == len(word) {
        fmt.Println(word)
    }
    // print nothing on failure
}
```

**Trace `"faya"` in `"fgvvfdxcacpolhyghbreda"`:**

| s char | word[i] | Match? | i |
|--------|---------|--------|---|
| f | f | YES | 1 |
| g | a | no | 1 |
| v | a | no | 1 |
| v | a | no | 1 |
| f | a | no | 1 |
| d | a | no | 1 |
| x | a | no | 1 |
| c | a | no | 1 |
| a | a | YES | 2 |
| c | y | no | 2 |
| p | y | no | 2 |
| o | y | no | 2 |
| l | y | no | 2 |
| h | y | no | 2 |
| y | y | YES | 3 |
| g | a | no | 3 |
| h | a | no | 3 |
| b | a | no | 3 |
| r | a | no | 3 |
| e | a | no | 3 |
| d | a | no | 3 |
| a | a | YES | 4 |

Final: `i == len("faya") == 4` → print `faya`

**Difference from hiddenp:**

| Feature | hiddenp (54) | wdmatch (59) |
|---------|-------------|-------------|
| s1 is in s2? | Print `1` or `0` | Print `word` or nothing |
| Arg order | `s1` then `s2` | `word` then `s` |
| Empty s1 | Print `1` | Print empty string |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Printing the empty string for empty `word` | `fmt.Println("")` prints a newline | `if len(word) == 0` you can still print it — or check if the challenge intends silence |
| Printing `word` even when only partially matched | Off-by-one: `i < len(word)` | Check `i == len(word)` (exact match), not `i > 0` |

## Solving This Challenge

### Algorithm
1. Require exactly 2 arguments; return otherwise.
2. Set `i = 0`.
3. For each character `c` in `s`: if `i < len(word)` and `c == rune(word[i])`, increment `i`.
4. If `i == len(word)`, print `word` followed by newline.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [60-fifthandskip](../60-fifthandskip/README.md)
