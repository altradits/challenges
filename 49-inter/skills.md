# Skills for 49-inter

**Previous:** [48-union](../48-union/skills.md) | **Next:** [50-stringbuilder](../50-stringbuilder/skills.md)

**Challenge:** Return a string containing only the characters that appear in BOTH `s1` and `s2`, preserving the order they appear in `s1`, with no duplicates.

## Core Concept: Set Intersection Using Two Maps

### What Is It?

**Intersection** is the opposite of union: instead of combining, you filter. Keep only the characters from s1 that also exist in s2. You need two maps:

1. A **membership set** built from s2: "which characters does s2 contain?"
2. A **seen set**: "have I already added this character to the result?"

### The Algorithm

```go
func Intersection(s1, s2 string) string {
    // Step 1: build a set from s2
    inS2 := make(map[rune]bool)
    for _, c := range s2 {
        inS2[c] = true
    }

    // Step 2: iterate s1, keep only chars that are in s2 and not yet added
    seen := make(map[rune]bool)
    result := ""
    for _, c := range s1 {
        if inS2[c] && !seen[c] {
            seen[c] = true
            result += string(c)
        }
    }
    return result
}
```

### Why Two Maps?

- `inS2` answers: "does this character appear anywhere in s2?"
- `seen` answers: "have I already added this character to the result?"

Without `seen`, a character that appears multiple times in s1 would be added to the result multiple times (as long as it's in s2).

### Contrast with Union

| | Union (111) | Intersection (112) |
|---|---|---|
| Input | s1="abc", s2="bcd" | s1="abc", s2="bcd" |
| Keep | chars in s1 OR s2 | chars in s1 AND s2 |
| Result | "abcd" | "bc" |
| Order | s1 first, then new from s2 | s1 order |

### Tracing Through `Intersection("padinton", "paqefwtd...")`

s2 = `"paqefwtdjetyiytjneytjoeyjnejeyj"` — build `inS2` from this.

Then iterate s1 = `"padinton"`:

| c | inS2[c]? | seen[c]? | Action | result |
|---|---------|---------|--------|--------|
| 'p' | yes | no | add | "p" |
| 'a' | yes | no | add | "pa" |
| 'd' | yes | no | add | "pad" |
| 'i' | yes | no | add | "padi" |
| 'n' | yes | no | add | "padin" |
| 't' | yes | no | add | "padint" |
| 'o' | yes | no | add | "padinto" |
| 'n' | yes | yes | skip (already added) | "padinto" |

Result: `"padinto"` — correct.

### Alternative: Simpler with One Map

If s2 has no duplicates, you only need one map. But using two is safer and works in all cases.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Iterating s2 instead of s1 | Wrong order (s2's order, not s1's) | Always iterate s1; use s2 only to build the lookup set |
| No `seen` map | Duplicate characters in result | Track which characters have already been added |
| Building set from s1 | Reversed logic | Set should be from s2; iteration should be over s1 |

## Algorithm

1. Build `inS2 = map[rune]bool` by iterating s2
2. Create `seen = map[rune]bool{}`, `result = ""`
3. For each rune `c` in s1:
   - If `inS2[c]` and `!seen[c]`: set `seen[c] = true`, append `string(c)` to result
4. Return `result`

## The Challenge

See [README.md](README.md).

**Next:** [50-stringbuilder](../50-stringbuilder/README.md)
