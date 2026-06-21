# Prerequisites for 54-replaceall

## Before You Start

### 1. `strings.HasPrefix`

`strings.HasPrefix(s, prefix)` returns `true` if string `s` starts with `prefix`. You'll use this to check whether `old` appears at the current scan position.

```go
import "strings"

fmt.Println(strings.HasPrefix("banana", "ban"))  // true
fmt.Println(strings.HasPrefix("banana", "ana"))  // false
fmt.Println(strings.HasPrefix("nana", "na"))     // true
```

Combined: `strings.HasPrefix(text[i:], old)` checks if `old` starts at position `i`.

### 2. String Slicing from an Index: `s[i:]`

`s[i:]` gives the substring from index `i` to the end of the string.

```go
s := "Hello World"
fmt.Println(s[6:])  // "World"
fmt.Println(s[0:])  // "Hello World"
```

### 3. Manual Index Loop

For scan-and-build you control the index yourself so you can jump ahead by more than 1:

```go
i := 0
for i < len(s) {
    // process s[i]
    i++        // advance by 1 (normal case)
    // or: i += skip   // advance by skip (after a match)
}
```

This is different from `for...range` because you can jump ahead by an arbitrary amount.

### 4. Converting a Byte to a String

`s[i]` returns a `byte`, not a string. To concatenate it to a string result, wrap it:

```go
s := "Hello"
b := s[0]              // b is type byte (uint8), value 72
fmt.Println(string(b)) // "H"
```

### 5. String Concatenation with `+=`

```go
result := ""
result += "Hello"
result += " World"
fmt.Println(result)  // "Hello World"
```

Note: For large strings this is O(n²). You'll learn the efficient alternative in 65-stringbuilder.

## Review If Stuck

- [53-findsubstring](../53-findsubstring/skills.md) — covers locating a pattern at a given index within a string

## You're Ready When You Can...

- [ ] Use `strings.HasPrefix` to test if a string starts with a prefix
- [ ] Write a manual index loop `i := 0; for i < len(s) { ... }`
- [ ] Advance the index by a variable amount (`i += len(old)`)
- [ ] Convert a byte `s[i]` to string with `string(s[i])`

## Next Steps

- [55-split](../55-split/README.md)
