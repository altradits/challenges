# Skills for 62-reversestrcap

**Previous:** [61-saveandmiss](../61-saveandmiss/skills.md) | **Next:** [63-union](../63-union/skills.md)

**Challenge:** For each word in the input string, lowercase all letters then uppercase the last letter. Return the transformed words joined by spaces.

## Core Concept: Word-Level Transformation with Rune Slice Manipulation

### What Is It?

This challenge combines two skills: **splitting into words** and **transforming each word**. The per-word transformation — lowercase everything, then uppercase the last character — requires treating each word as a `[]rune` so you can modify the last element by index.

### Step 1: Split into Words

`strings.Fields(s)` splits on any whitespace and handles multiple consecutive spaces:

```go
words := strings.Fields("First SMALL TesT")
// ["First", "SMALL", "TesT"]
```

### Step 2: Transform Each Word

For each word:
1. Convert to `[]rune` (to enable index-based modification)
2. Lowercase all characters with `unicode.ToLower`
3. Uppercase the last character with `unicode.ToUpper`

```go
func transformWord(word string) string {
    if len(word) == 0 {
        return ""
    }
    runes := []rune(word)

    // Step 1: lowercase everything
    for i, c := range runes {
        runes[i] = unicode.ToLower(c)
    }

    // Step 2: uppercase the last character
    last := len(runes) - 1
    runes[last] = unicode.ToUpper(runes[last])

    return string(runes)
}
```

### Step 3: Rejoin with Spaces

After transforming all words, join them back with a single space:

```go
func ReverseStrCap(s string) string {
    words := strings.Fields(s)
    for i, word := range words {
        words[i] = transformWord(word)
    }
    return strings.Join(words, " ")
}
```

### Why `[]rune` Instead of Indexing the String?

Strings in Go are immutable. You cannot do `s[0] = 'x'`. You must convert to a mutable type first. `[]rune` lets you assign to individual positions:

```go
runes := []rune("Hello")
runes[0] = unicode.ToLower(runes[0])  // works!
// string[0] = ... would be a compile error
```

### Tracing Through `"First SMALL TesT"`

After `strings.Fields`: `["First", "SMALL", "TesT"]`

| Word | After lowercase | After uppercase last | Result |
|------|----------------|---------------------|--------|
| "First" | "first" | "firsT" | "firsT" |
| "SMALL" | "small" | "smalL" | "smalL" |
| "TesT"  | "test"  | "tesT"  | "tesT" |

Joined: `"firsT smalL tesT"` — correct.

### Edge Cases

- Empty string: `strings.Fields("")` returns `[]` (empty slice). `strings.Join([], " ")` returns `""`. Handled automatically.
- Single-character word: lowercased then uppercased → uppercase. `"a"` → `"A"`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Trying to modify `s[i]` directly | Compile error — strings are immutable | Convert to `[]rune` first |
| Using `strings.Split(s, " ")` | Fails with multiple spaces | Use `strings.Fields` instead |
| Uppercasing before lowercasing | First char gets lowercased then uppercased redundantly | Lowercase all first, then uppercase only last |

## Algorithm

1. If `s == ""`, return `""`
2. Split `s` into words with `strings.Fields`
3. For each word:
   a. Convert to `[]rune`
   b. Lowercase all runes
   c. Uppercase last rune
   d. Convert back to string
4. Join transformed words with `" "`
5. Return result

## The Challenge

See [README.md](README.md).

**Next:** [63-union](../63-union/README.md)
