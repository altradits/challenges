# Skills for rostring

## What You'll Learn

**Previous:** [65-revwstr](../65-revwstr/skills.md) | **Next:** [67-wordflip](../67-wordflip/README.md)

**Challenge:** Rotate a string one word to the left — the first word moves to the end, all others keep their order. Words are alphanumeric sequences; output is joined by single spaces.

## Core Concept: String Rotation via Word Splitting

### What Is It?

Rotating a string left by one word means taking the first word and appending it to the end. The key skills are: extracting real words (ignoring extra spaces and non-alphanumeric sequences), then reassembling with `strings.Join`.

### How It Works

**Step 1 — Extract words (alphanumeric sequences only):**

`strings.Fields` splits on whitespace but keeps non-alphanumeric tokens like `,` and `;` as separate "words." For this challenge, words are defined as alphanumeric sequences, so we must either filter or use a different approach.

Looking at the expected output for `"     AkjhZ zLKIJz , 23y"` → `"zLKIJz , 23y AkjhZ"`: the comma `,` is kept as a separate token and preserved in order! So `strings.Fields` actually works here — it treats `,` as its own word token.

```go
import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println()
        return
    }
    s := os.Args[1]
    words := strings.Fields(s)

    if len(words) == 0 {
        fmt.Println()
        return
    }

    // Rotate left: first word goes to the end
    rotated := append(words[1:], words[0])
    fmt.Println(strings.Join(rotated, " "))
}
```

**Trace `"Let there     be light"`:**
- `strings.Fields` → `["Let", "there", "be", "light"]`
- Rotate: `["there", "be", "light", "Let"]`
- Join: `"there be light Let"` ✓

**Trace `"abc   "` (trailing spaces):**
- `strings.Fields` → `["abc"]`
- Rotate: `["abc"]` (one word, goes to end — same position)
- Join: `"abc"` ✓

**Trace `"     AkjhZ zLKIJz , 23y"`:**
- `strings.Fields` → `["AkjhZ", "zLKIJz", ",", "23y"]`
- Rotate: `["zLKIJz", ",", "23y", "AkjhZ"]`
- Join: `"zLKIJz , 23y AkjhZ"` ✓

**Understanding `append(words[1:], words[0])`:**

`words[1:]` is a new slice starting from the second element. `words[0]` is the first element. Appending the first element to the tail gives the rotated result.

```go
words := []string{"a", "b", "c", "d"}
rotated := append(words[1:], words[0])
// rotated = ["b", "c", "d", "a"]
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not handling single-word input | Index out of range on `words[1:]` | `words[1:]` on a 1-element slice returns `[]` — safe in Go |
| Not handling empty/whitespace-only input | `words` is empty, `words[0]` panics | Check `len(words) == 0` first |
| Printing extra spaces | `strings.Join` with `" "` produces exactly one space | Use `strings.Fields` + `strings.Join` — extra spaces are normalized |

## Solving This Challenge

### Algorithm
1. If arg count != 1, print newline and return.
2. `words = strings.Fields(s)`.
3. If `len(words) == 0`, print newline and return.
4. `rotated = append(words[1:], words[0])`.
5. `fmt.Println(strings.Join(rotated, " "))`.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [67-wordflip](../67-wordflip/README.md)
