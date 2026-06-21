# Skills for 97-longestword

## What You'll Learn

**Previous:** [96-printifnot](../96-printifnot/skills.md) | **Next:** [98-searchreplace](../98-searchreplace/skills.md)

**Challenge:** Return the longest word in a string.

## Core Concept: Tracking a Running Maximum

### What Is It?

Finding the "maximum" element while iterating means keeping a variable that holds the best answer seen so far and updating it whenever you find something better.

### How It Works

```go
func LongestWord(s string) string {
    words := strings.Fields(s)  // split on any whitespace
    longest := ""
    for _, word := range words {
        if len(word) > len(longest) {
            longest = word
        }
    }
    return longest
}
```

Walk-through with `"Go is awesome"`:
- Start: `longest = ""`
- `"Go"` (2) > `""` (0) → `longest = "Go"`
- `"is"` (2) > `"Go"` (2) → false, keep `"Go"`
- `"awesome"` (7) > `"Go"` (2) → `longest = "awesome"`
- Return `"awesome"`

### Why Initialize to `""`?

Initializing to empty string means any real word is longer and wins on the first comparison. This cleanly handles the empty-input case too — if there are no words, you return `""`.

### Ties

When two words have equal length, the **first** one wins because `>` (strictly greater) is used. Change to `>=` if you want the last one to win.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Comparing strings with `>` | Lexicographic, not length | Use `len(word) > len(longest)` |
| Initializing `longest` to `words[0]` | Panics on empty input | Initialize to `""` |
| Using byte indexing instead of Fields | Misses multi-byte chars | Use `strings.Fields` |

## Algorithm

1. Split `s` into words with `strings.Fields(s)`
2. Initialize `longest = ""`
3. For each word: if `len(word) > len(longest)`, update `longest`
4. Return `longest`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [98-searchreplace](../98-searchreplace/skills.md)
