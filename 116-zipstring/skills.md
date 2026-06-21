# Skills for zipstring

## What You'll Learn

**Previous:** [115-weareunique](../115-weareunique/skills.md) | **Next:** [117-addprimesum](../117-addprimesum/skills.md)

**Challenge:** Replace each character with its consecutive-run count followed by the character itself (run-length encoding).

## Core Concept: Run-Length Encoding

### What Is Run-Length Encoding?
Count consecutive runs of the same character and represent each run as `count + character`. For example: `"YouuungFellllas"` → `"1Y1o3u1n1g1F1e4l1a1s"` (the `uuu` run becomes `3u`, the `llll` run becomes `4l`).

### How It Works
Use the same "compare to previous" pattern from countrepeats, but track the current run length:

```go
func ZipString(s string) string {
    if len(s) == 0 {
        return ""
    }
    result := ""
    count := 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            count++
        } else {
            // flush the previous run
            result += fmt.Sprintf("%d%c", count, s[i-1])
            count = 1
        }
    }
    // flush the last run
    result += fmt.Sprintf("%d%c", count, s[len(s)-1])
    return result
}
```

### Step-by-Step for `"YouuungFellllas"`
- Y: count=1
- o≠Y: flush → "1Y", count=1
- u≠o: flush → "1Y1o", count=1
- u=u: count=2
- u=u: count=3
- n≠u: flush → "1Y1o3u", count=1
- g≠n: flush → "1Y1o3u1n", count=1
- ... and so on
- Final flush: all remaining chars appended
- Result: `"1Y1o3u1n1g1F1e4l1a1s"`

### The Flush Pattern
You process a run when the *next* character is different (or you've reached the end). This means you always flush *one run behind* where you are:
- When `s[i] != s[i-1]`: emit the count and character for position `i-1`
- After the loop: emit the final run for the last character

### fmt.Sprintf With %d%c
`%d` formats an integer; `%c` formats a rune as its character:
```go
fmt.Sprintf("%d%c", 3, 'u') // "3u"
fmt.Sprintf("%d%c", 1, 'Y') // "1Y"
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting the final flush after the loop | Last run is missing from output | Add `result += fmt.Sprintf(...)` after the loop |
| Flushing `s[i]` instead of `s[i-1]` | Emits the wrong character | When `s[i] != s[i-1]`, flush `s[i-1]` (the character that just ended) |
| Not handling empty string | Panic on `s[len(s)-1]` | Guard with `if len(s) == 0 { return "" }` |

## Solving This Challenge

### Algorithm
1. Return `""` if string is empty.
2. Initialize `count = 1`.
3. Loop from i=1: if `s[i] == s[i-1]`, increment count; else flush (`count + s[i-1]`) and reset count=1.
4. After loop, flush the final run.
5. Return result.

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [117-addprimesum](../117-addprimesum/README.md)
