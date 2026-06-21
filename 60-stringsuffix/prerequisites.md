# Prerequisites for 60-stringsuffix

## Before You Start

### 1. `strings.HasPrefix` and string slicing from the front

```go
s[:n]  // first n bytes
```

Review: [59-stringprefix](../59-stringprefix/skills.md) — HasSuffix is the mirror image.

### 2. Slicing from the end: `s[len(s)-n:]`

```go
s := "hello"
s[len(s)-2:]  // "lo" — last 2 bytes
s[len(s)-5:]  // "hello" — whole string
```

Review: [15-lastchar](../15-lastchar/skills.md) — `s[len(s)-1]` is the same idea, one byte from the end.

### 3. Length comparison before slicing

You must ensure `len(s) >= len(suffix)` before taking `s[len(s)-len(suffix):]`, otherwise you get a negative index which panics.

```go
if len(s) < len(suffix) {
    return false
}
```

## You're Ready When You Can...

- [ ] Slice the last N characters from a string
- [ ] Guard against negative indices with a length check
- [ ] Compare a string slice to another string

## Next Steps

- [61-stringfield](../61-stringfield/README.md)
