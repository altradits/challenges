# Prerequisites for 48-union

## Before You Start

### 1. Maps in Go — `map[rune]bool`

A map associates keys with values. `map[rune]bool` is a set of runes — the value `true` means "present".

```go
seen := make(map[rune]bool)

seen['a'] = true  // add 'a'
seen['b'] = true  // add 'b'

fmt.Println(seen['a'])  // true  — present
fmt.Println(seen['c'])  // false — not present (zero value)
```

### 2. Checking Map Membership

```go
if seen['x'] {
    fmt.Println("x was seen")
} else {
    fmt.Println("x is new")
}
```

You do NOT need to check `ok` separately — accessing a missing key returns the zero value (`false` for bool), which is exactly what you want.

### 3. Adding to a Map While Iterating

```go
seen := make(map[rune]bool)
result := ""
for _, c := range "aabbcc" {
    if !seen[c] {
        seen[c] = true  // mark as seen
        result += string(c)  // add to output only once
    }
}
fmt.Println(result)  // "abc"
```

### 4. String Concatenation for Two Strings

You can concatenate two strings and iterate the combined result:

```go
s1 := "hello"
s2 := "world"
for _, c := range s1 + s2 {
    fmt.Printf("%c ", c)
}
// h e l l o w o r l d
```

### 5. What "Union" Means

Set union: all elements from both sets, no duplicates. If `s1 = "abc"` and `s2 = "bcd"`, the union is `"abcd"` — 'b' and 'c' appear in both but are listed once, in the order they first appear.

## Review If Stuck

- [37-countwords](../37-countwords/skills.md) — covers `for...range` iteration over strings
- [44-thirdchar](../44-thirdchar/skills.md) — covers conditional character accumulation

## You're Ready When You Can...

- [ ] Create a `map[rune]bool` with `make`
- [ ] Check if a rune is in the map with `if seen[c]`
- [ ] Add a rune to the map with `seen[c] = true`
- [ ] Iterate over the concatenation of two strings with `for _, c := range s1+s2`

## Next Steps

- [49-inter](../49-inter/README.md)
