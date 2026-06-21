# Prerequisites for weareunique

## Before You Start

To solve this challenge you need to understand:

### 1. Maps in Go
A map stores key-value pairs. `map[rune]bool` maps characters to boolean values — ideal for a "set" that tracks which characters have been seen.
```go
seen := map[rune]bool{}  // create an empty map
seen['a'] = true          // add 'a' to the set
seen['b'] = true          // add 'b' to the set
fmt.Println(seen['a'])    // true — 'a' is in the set
fmt.Println(seen['c'])    // false — 'c' is not in the set (zero value)
```

### 2. Map as a Set for Deduplication
Setting a key to `true` multiple times has no effect — the map keeps only one entry per key.
```go
seen := map[rune]bool{}
for _, c := range "foo" {
    seen[c] = true
}
// seen == {'f': true, 'o': true} — only 2 keys, not 3
fmt.Println(len(seen)) // 2
```

### 3. Checking Map Membership
`seen[key]` returns `false` if the key is not present (Go returns the zero value for missing keys).
```go
seen := map[rune]bool{'a': true}
if seen['a'] {
    fmt.Println("a is in the set")
}
if !seen['b'] {
    fmt.Println("b is NOT in the set")
}
```

### 4. Iterating Over a Map
Use `for key := range m` to visit every key in a map.
```go
set := map[rune]bool{'a': true, 'b': true}
for c := range set {
    fmt.Println(string(c))
}
```

### 5. Counting Elements Not in Another Map
```go
count := 0
for c := range set1 {
    if !set2[c] {
        count++
    }
}
```

## Review If Stuck

- [../93-countrepeats/skills.md](../93-countrepeats/skills.md) — covers tracking state while iterating over a string

## You're Ready When You Can...

- [ ] Create a `map[rune]bool{}` and add entries to it
- [ ] Use a map as a set to deduplicate characters from a string
- [ ] Check whether a key is absent from a map with `!m[key]`
- [ ] Iterate over all keys in a map with `for k := range m`
- [ ] Count characters in one set that are absent from another

## Next Steps

- [116-zipstring](../116-zipstring/README.md)
