# Prerequisites for 55-split

## Before You Start

### 1. Slices of Strings: `[]string`

A slice is a dynamic list. You add to it with `append`:

```go
var result []string
result = append(result, "hello")
result = append(result, "world")
fmt.Println(result)        // [hello world]
fmt.Println(len(result))   // 2
```

### 2. String Slicing `s[a:b]`

Extract a substring from index `a` (inclusive) to `b` (exclusive):

```go
s := "a,b,c"
fmt.Println(s[0:1])  // "a"
fmt.Println(s[2:3])  // "b"
fmt.Println(s[4:])   // "c"  (from 4 to end)
fmt.Println(s[2:2])  // ""   (empty — start equals end)
```

### 3. Comparing Substrings

You can compare a slice of a string to another string with `==`:

```go
s := "hello,world"
fmt.Println(s[5:6] == ",")  // true
fmt.Println(s[0:5] == "hello")  // true
```

### 4. A `start` Pointer Pattern

Track where the current token started:

```go
start := 0
for i := 0; i < len(s); i++ {
    if s[i] == ',' {
        token := s[start:i]  // everything before the comma
        start = i + 1        // next token starts after the comma
        _ = token
    }
}
lastToken := s[start:]  // everything after the last comma
_ = lastToken
```

### 5. Multi-Character Separator Slicing

When the separator has length > 1, compare `s[i:i+len(sep)]` and skip ahead by `len(sep)`:

```go
sep := "::"
for i := 0; i <= len(s)-len(sep); i++ {
    if s[i:i+len(sep)] == sep {
        // found "::" at position i
    }
}
```

## Review If Stuck

- [54-replaceall](../54-replaceall/skills.md) — covers the scan-and-build pattern with a manual index

## You're Ready When You Can...

- [ ] Create a `[]string` slice and append to it
- [ ] Extract a substring with `s[a:b]`
- [ ] Compare a string slice like `s[i:i+2] == ",,"` 
- [ ] Use a `start` variable to track the beginning of the current token

## Next Steps

- [56-join](../56-join/README.md)
