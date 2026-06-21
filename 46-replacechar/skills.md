# Skills for 46-replacechar

## What You'll Learn

**Previous:** [45-findlastchar](../45-findlastchar/skills.md) | **Next:** [47-printif](../47-printif/skills.md)

**Challenge:** Write a function `ReplaceChar(s string, old, new rune) string` that replaces every occurrence of character `old` with character `new`.

## Core Concept: Conditional String Building — Transform or Pass Through

### What Is Character Replacement?

Replacement combines the filter pattern (from [39-removespaces](../39-removespaces/skills.md)) with the transform pattern (from [32-toupper](../32-toupper/skills.md)). Instead of removing a character or changing its case, you substitute it with a different specific character.

### The Implementation

```go
func ReplaceChar(s string, old, new rune) string {
    result := ""
    for _, c := range s {
        if c == old {
            result += string(new)   // substitute
        } else {
            result += string(c)     // keep original
        }
    }
    return result
}
```

Every character is appended — either the replacement or the original. Nothing is dropped.

### Three-Parameter Functions

This challenge uses three parameters — the first multi-rune-parameter function in the series:

```go
func ReplaceChar(s string, old, new rune) string
```

When two parameters share the same type, Go lets you write them together: `old, new rune` instead of `old rune, new rune`. Both mean the same thing.

Called with:

```go
ReplaceChar("Hello", 'l', 'L')   // "HeLLo"
```

### `new` Is a Reserved Word in Some Contexts

`new` is a built-in function in Go (it allocates memory). Using it as a parameter name shadows the built-in. This is legal and is done in the README's expected function signature, so follow the README. If you prefer to avoid the ambiguity, you can name it `replacement` locally.

### Comparing with Earlier Challenges

| Challenge | When match? | When no match? | Result |
|-----------|------------|----------------|--------|
| 39-removespaces | Drop the character | Keep | Filtered string |
| 32-toupper | Change the character | Keep | Transformed string |
| 46-replacechar | Substitute with `new` | Keep `old` value | Substituted string |

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `result += string(old)` in the else | Appends `old`, not the current char | Use `string(c)` in the else |
| Only replacing the first occurrence | Early return or wrong logic | Loop over ALL characters; no return inside |
| Dropping non-matching characters | Like RemoveSpaces behaviour | Use `else { result += string(c) }` |

## Solving This Challenge

### Algorithm

1. `result := ""`
2. For each rune `c` in `s`:
   - If `c == old`: append `string(new)`
   - Else: append `string(c)`
3. Return `result`

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [47-printif](../47-printif/README.md)
