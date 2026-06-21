# Prerequisites for 96-printifnot

## Before You Start

### 1. `PrintIf` — you just completed it

This challenge is the exact inverse. If you completed [95-printif](../95-printif/README.md), you already know everything needed.

```go
// PrintIf returns "G" when len == 0 OR len >= 3
// PrintIfNot returns "G" when len < 3 (the inverse condition)
```

### 2. Inverting conditions

| Original | Inverse |
|----------|---------|
| `>= 3` | `< 3` |
| `== 0` | `!= 0` |
| `A or B` | `not A and not B` |

### 3. `len()` and `if/else`

```go
if len(str) < 3 {
    return "G\n"
}
return "Invalid Input\n"
```

## You're Ready When You Can...

- [ ] Flip the condition from PrintIf to its inverse
- [ ] Verify with: `""` → `"G\n"`, `"ab"` → `"G\n"`, `"abc"` → `"Invalid Input\n"`

## Next Steps

- [97-longestword](../97-longestword/README.md)
