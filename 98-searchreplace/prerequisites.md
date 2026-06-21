# Prerequisites for 98-searchreplace

## Before You Start

### 1. String slicing with `s[start:end]`

String slicing extracts a portion of a string by byte position.

```go
s := "Hello World"
s[:5]   // "Hello"  — first 5 bytes
s[6:]   // "World"  — from byte 6 to end
s[1:4]  // "ell"    — bytes 1,2,3
```

Review: [89-retainfirsthalf](../89-retainfirsthalf/skills.md)

### 2. Finding a substring's position

You need to find WHERE `old` starts inside `text`. Either manually with a loop or via `strings.Index`:

```go
import "strings"
idx := strings.Index("Hello World", "World")  // 6
idx  = strings.Index("Hello World", "xyz")    // -1 (not found)
```

Review: [101-findsubstring](../101-findsubstring/skills.md) for the manual approach.

### 3. String length with `len()`

You need `len(old)` to know how many bytes to skip after the match.

```go
len("World")  // 5
```

### 4. String concatenation with `+`

Joining three parts with `+`:

```go
before + replacement + after
```

## You're Ready When You Can...

- [ ] Slice a string to get the part before a given index
- [ ] Slice a string to get the part after a given range
- [ ] Check if `strings.Index` returned -1 (not found)

## Next Steps

- [99-cleanlist](../99-cleanlist/README.md)
