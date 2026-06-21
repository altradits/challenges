# Prerequisites for 51-stringsplit

## Before You Start

### 1. String Slices: `[]string`

A slice of strings is Go's standard way to hold a list of words or tokens:

```go
words := []string{"Hello", "World", "Go"}
fmt.Println(words[0])    // "Hello"
fmt.Println(len(words))  // 3
```

### 2. How Whitespace Works in Strings

Whitespace characters that separate words:
- `' '` — space (most common)
- `'\t'` — tab
- `'\n'` — newline
- `'\r'` — carriage return

```go
s := "  Go\tis\nfun  "
// Has leading spaces, a tab, and a newline inside
```

### 3. What `strings.Split` Does (and Why It's Not Enough Here)

`strings.Split("a  b", " ")` splits on the exact string `" "`. Two consecutive spaces create an empty element:

```go
parts := strings.Split("a  b", " ")
fmt.Println(parts)      // ["a" "" "b"]  ← empty element!
fmt.Println(len(parts)) // 3
```

This is NOT what you want for word splitting.

### 4. Appending to a Slice with `append`

Build a result slice dynamically:

```go
var result []string
result = append(result, "first")
result = append(result, "second")
fmt.Println(result)  // [first second]
```

### 5. `strings.Builder` from 50-stringbuilder

You may also need `strings.Builder` if you implement word splitting manually. The efficient O(n) approach to building up words character by character.

## Review If Stuck

- [50-stringbuilder](../50-stringbuilder/skills.md) — covers efficient string building
- [40-split](../40-split/skills.md) — covers the manual splitting algorithm with a start pointer

## You're Ready When You Can...

- [ ] Create and append to a `[]string` slice
- [ ] Explain why `strings.Split(s, " ")` fails on multiple spaces
- [ ] Identify whitespace characters: `' '`, `'\t'`, `'\n'`
- [ ] Describe the state-machine approach: `inWord` flag + `start` index

## Next Steps

- [52-stringjoin](../52-stringjoin/README.md)
