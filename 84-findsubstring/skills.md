# Skills for 84-findsubstring

## What You'll Learn

**Previous:** [83-missing-pieces](../83-missing-pieces/skills.md) | **Next:** [85-printif](../85-printif/skills.md)

**Challenge:** Write `FindSubstring(text, substring string) int` that returns the index of the first occurrence of `substring` in `text`, or `-1` if not found.

## Core Concept: Sliding Window Search Over a String

### What Is It?

Finding a substring means sliding a window of length `len(substring)` along `text`, and at each position checking whether the slice of `text` starting there matches `substring`. This is your first challenge that involves **string slicing** — extracting a portion of a string using `s[start:end]`.

### New Tool: String Slicing

In Go, `s[start:end]` returns the substring from index `start` up to (but not including) index `end`.

```go
s := "Hello World"
s[0:5]   // "Hello"
s[6:11]  // "World"
s[6:]    // "World"  (to the end)
s[:5]    // "Hello"  (from the start)
```

Indices are **byte** positions (not rune positions). For ASCII strings this is the same as character positions.

### New Tool: `len(s)` Returns the Byte Count

```go
len("Hello")  // 5
len("World")  // 5
len("")       // 0
```

### The Sliding Window Algorithm

```
text:      "banana"
substring: "ana"
len(sub):  3

position 0: "ban" == "ana"? No
position 1: "ana" == "ana"? YES -> return 1
```

In code:

```go
func FindSubstring(text, substring string) int {
    subLen := len(substring)

    // Edge case: empty substring is always found at position 0
    if subLen == 0 {
        return 0
    }

    textLen := len(text)
    for i := 0; i <= textLen-subLen; i++ {
        if text[i:i+subLen] == substring {
            return i
        }
    }
    return -1
}
```

Step by step:
1. Handle the empty substring edge case first
2. Loop `i` from `0` to `len(text) - len(substring)` (inclusive)
3. At each position, slice `text[i : i+subLen]` and compare to `substring`
4. If they match, return `i`
5. If the loop ends without a match, return `-1`

### Why `i <= textLen-subLen` (not `i < textLen`)?

If `text = "ab"` and `substring = "ab"`, then `textLen = 2` and `subLen = 2`. The only valid starting position is `0`. If you loop to `i < textLen` you would try `i=1` and slice `text[1:3]` — which goes out of bounds.

```
textLen - subLen = 2 - 2 = 0
So loop runs for i = 0 only. Correct.
```

### Visual: Sliding Window

```
text:      b  a  n  a  n  a
index:     0  1  2  3  4  5

window at 0: [b  a  n] vs "ana" -> no
window at 1: [a  n  a] vs "ana" -> no
window at 2: [n  a  n] vs "ana" -> no
window at 3: [a  n  a] vs "ana" -> no  (wait — "ana"[3:6] = "ana"?)

Actually for "banana":
i=0: text[0:3] = "ban" != "ana"
i=1: text[1:4] = "ana" == "ana" -> return 1
```

### Alternative: `strings.Index`

The standard library has this built in:

```go
import "strings"

func FindSubstring(text, substring string) int {
    return strings.Index(text, substring)
}
```

`strings.Index` returns `-1` when not found and `0` for empty substring — exactly matching the requirements. However, implementing it manually teaches you slicing and loop bounds.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `i < textLen` as loop bound | Index out of bounds when `i + subLen > textLen` | Use `i <= textLen - subLen` |
| Forgetting the empty substring case | Loop might return wrong result or index out of bounds | Check `if subLen == 0 { return 0 }` first |
| Not returning `-1` after the loop | Missing return — compile error | Add `return -1` after the loop |
| Case-sensitive mismatch | `"World" != "world"` — this is correct behavior | The search is intentionally case-sensitive |

## Solving This Challenge

### Algorithm

1. If `substring` is empty, return `0`
2. For each starting position `i` from `0` to `len(text)-len(substring)`:
   a. If `text[i:i+len(substring)] == substring`, return `i`
3. Return `-1`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [85-printif](../85-printif/README.md)
