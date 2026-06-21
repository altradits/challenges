# Prerequisites for wordflip

## Before You Start

To solve this challenge you need to understand:

### 1. `strings.TrimSpace` — Remove Leading and Trailing Whitespace
Returns a copy of the string with all leading and trailing whitespace removed.

```go
import "strings"

strings.TrimSpace("  hello  ")  // "hello"
strings.TrimSpace("   ")        // "" (empty — all spaces trimmed)
strings.TrimSpace("")           // ""
```

### 2. `strings.Fields` — Split on Any Whitespace
Handles multiple consecutive spaces as a single separator. Returns an empty slice for all-whitespace input.

```go
strings.Fields(" hello  all  of  you! ")
// ["hello", "all", "of", "you!"]

strings.Fields("   ")
// [] (empty slice)
```

### 3. In-Place Slice Reversal (Two-Pointer)
Swap elements from both ends moving inward.

```go
words := []string{"a", "b", "c", "d"}
for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
    words[i], words[j] = words[j], words[i]
}
// words = ["d", "c", "b", "a"]
```

### 4. `strings.Join` — Reassemble with Single Spaces
Joins a word slice with exactly one space between words.

```go
strings.Join([]string{"you!", "of", "all", "hello"}, " ")
// "you! of all hello"
```

### 5. Special Return Value for Invalid Input
Return a specific string (not just an empty string) when the input is empty or all spaces.

```go
if strings.TrimSpace(str) == "" {
    return "Invalid Output"
}
```

## Review If Stuck
- [65-revwstr](../65-revwstr/skills.md) — teaches `strings.Fields`, `strings.Join`, and two-pointer word reversal
- [66-rostring](../66-rostring/skills.md) — reinforces word splitting and joining patterns

## You're Ready When You Can...
- [ ] Use `strings.TrimSpace` to detect all-whitespace strings
- [ ] Use `strings.Fields` to split ignoring multiple/leading/trailing spaces
- [ ] Reverse a `[]string` in-place with a two-pointer swap
- [ ] Use `strings.Join(words, " ")` to reassemble
- [ ] Return a special string (`"Invalid Output"`) when input is blank

## Next Steps
- [68-itoabase](../68-itoabase/README.md)
