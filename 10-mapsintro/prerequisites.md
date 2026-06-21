# Prerequisites for 10-mapsintro

## Before You Start

### 1. Slices and `make` — from 09-slicesintro

You already know slices use `make` to create them. Maps use the same function:

```go
s := make([]int, 5)          // slice
m := make(map[string]int)    // map — same pattern
```

Review: [09-slicesintro](../09-slicesintro/skills.md)

### 2. `for...range` over a slice — from 07-forloops

You'll iterate over words with range:

```go
for _, word := range words {
    freq[word]++
}
```

Review: [07-forloops](../07-forloops/skills.md)

### 3. `strings.Fields` splits on whitespace — from 08-stringspackage

```go
words := strings.Fields("go is go")
// ["go", "is", "go"]
```

Review: [08-stringspackage](../08-stringspackage/skills.md)

### 4. Zero value of `int` is `0`

When you access a map key that doesn't exist, you get the zero value for the value type:

```go
m := make(map[string]int)
fmt.Println(m["missing"])  // 0 (not an error)
m["missing"]++             // now m["missing"] == 1
```

This is why you can use `m[key]++` to count — you don't need to check if the key exists first.

## You're Ready When You Can...

- [ ] Create a map with `make(map[string]int)`
- [ ] Set a value: `m["key"] = value`
- [ ] Increment a count: `m["word"]++`
- [ ] Iterate over the map with `for k, v := range m`
- [ ] Use the comma-ok pattern: `v, ok := m["key"]`

## Next Steps

- [11-arrays](../11-arrays/README.md)
