# Prerequisites for 52-stringjoin

## Before You Start

### 1. `strings.Builder` (from 50-stringbuilder)

The key tool for this challenge. Writes to an internal buffer, returns the result with `.String()`:

```go
var b strings.Builder
b.WriteString("hello")
b.WriteString(" ")
b.WriteString("world")
fmt.Println(b.String())  // "hello world"
```

### 2. Iterating a Slice by Index

```go
words := []string{"a", "b", "c"}
for i := 0; i < len(words); i++ {
    fmt.Println(words[i])
}
// a
// b
// c
```

### 3. The "Start at Index 1" Pattern

To put the separator between elements (not before the first or after the last):

```go
var b strings.Builder
b.WriteString(words[0])          // first: no separator
for i := 1; i < len(words); i++ {
    b.WriteString(sep)           // separator BEFORE each subsequent element
    b.WriteString(words[i])
}
```

### 4. Empty Slice Guard

Accessing `words[0]` on an empty slice panics. Always check first:

```go
if len(words) == 0 {
    return ""
}
```

### 5. `b.Grow(n)` — Optional Pre-Allocation

If you know the approximate output size, you can hint to Builder to pre-allocate:

```go
b.Grow(100)  // allocate space for at least 100 bytes upfront
```

This reduces buffer doublings and improves performance further. It's optional.

## Review If Stuck

- [41-join](../41-join/skills.md) — covers the exact same algorithm using `+=` (this challenge upgrades it with Builder)
- [50-stringbuilder](../50-stringbuilder/skills.md) — covers why Builder is O(n) vs `+=` O(n²)

## You're Ready When You Can...

- [ ] Declare `var b strings.Builder` and use `b.WriteString()`
- [ ] Get the final string with `b.String()`
- [ ] Write the separator-between pattern (start at index 1)
- [ ] Guard against empty slice with `len(words) == 0`

## Next Steps

- [53-stringrepeat](../53-stringrepeat/README.md)
