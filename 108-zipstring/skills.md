# Skills for 108-zipstring

**Previous:** [107-thirdchar](../107-thirdchar/README.md) | **Next:** [109-saveandmiss](../109-saveandmiss/README.md)

**Challenge:** Implement run-length encoding: compress consecutive identical characters as count+character (e.g., `"YouuungFellllas"` → `"1Y1o3u1n1g1F1e4l1a1s"`).

## Core Concept: Run-Length Encoding with a "Previous Character" Tracker

### What Is It?

**Run-Length Encoding (RLE)** compresses consecutive identical characters by recording how many times each character repeats before changing. Instead of storing `"aaabbc"`, you store `"3a2b1c"`.

The technique: track the **current character** and a **count**. When the next character differs from the current one, output `count + currentChar` and start a new count of 1.

### The Algorithm

```go
func ZipString(s string) string {
    if len(s) == 0 {
        return ""
    }

    result := ""
    runes := []rune(s)
    current := runes[0]
    count := 1

    for i := 1; i < len(runes); i++ {
        if runes[i] == current {
            count++  // same character, extend the run
        } else {
            // different character: flush the current run
            result += fmt.Sprintf("%d%c", count, current)
            current = runes[i]
            count = 1  // start new run
        }
    }
    // flush the last run (never flushed inside the loop)
    result += fmt.Sprintf("%d%c", count, current)

    return result
}
```

### The Critical "Last Group" Problem

The loop only flushes a run when it finds a **different** character. The final run never triggers a different-character transition — so it would be silently lost without the post-loop flush.

```
"aabb":
i=1: 'a'=='a' → count=2
i=2: 'b'!='a' → flush "2a", current='b', count=1
i=3: 'b'=='b' → count=2
Loop ends → MUST flush "2b" here!
```

### Converting Count to String

`fmt.Sprintf("%d%c", count, current)` formats an integer and a rune into a string. Alternatively, using the `Itoa` function you wrote in challenge 106:

```go
result += Itoa(count) + string(current)
```

### Tracing Through `"Helloo"`

| i | runes[i] | == current? | count | current | result so far |
|---|---------|------------|-------|---------|---------------|
| start | — | — | 1 | 'H' | "" |
| 1 | 'e' | no ('H') | flush "1H", count=1 | 'e' | "1H" |
| 2 | 'l' | no ('e') | flush "1e", count=1 | 'l' | "1H1e" |
| 3 | 'l' | yes | 2 | 'l' | "1H1e" |
| 4 | 'o' | no ('l') | flush "2l", count=1 | 'o' | "1H1e2l" |
| 5 | 'o' | yes | 2 | 'o' | "1H1e2l" |
| end | — | — | flush "2o" | — | "1H1e2l2o" |

Result: `"1H1e2l2o"` — correct.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Forgetting the post-loop flush | Last group silently dropped | Always add `result += Sprintf(count, current)` after the loop |
| Not resetting `count = 1` on change | Count accumulates across runs | Set `count = 1` when starting a new run |
| Using `s[i]` instead of runes | Breaks on multi-byte Unicode characters | Convert to `[]rune` first |

## Algorithm

1. Handle empty string → return `""`
2. Convert `s` to `[]rune`, set `current = runes[0]`, `count = 1`
3. For `i` from 1 to len-1:
   - If `runes[i] == current`: `count++`
   - Else: append `count+current` to result, set `current = runes[i]`, `count = 1`
4. Append final `count+current` to result
5. Return result

## The Challenge

See [README.md](README.md).

**Next:** [109-saveandmiss](../109-saveandmiss/README.md)
