# Prerequisites for 36-cleanlist

## Before You Start

### 1. Slice operations — iterating and appending

```go
items := []string{"a", "b", "c"}
result := []string{}
for _, item := range items {
    result = append(result, item)
}
```

Review: [104-chunk](../104-chunk/skills.md)

### 2. Maps — key/value lookup

```go
m := make(map[string]bool)
m["hello"] = true
if m["hello"] {  // true — key exists
    fmt.Println("found")
}
if !m["world"] { // false — key not present, zero value
    fmt.Println("not found")
}
```

Review: [78-countrepeats](../78-countrepeats/skills.md) introduced maps for counting.

### 3. The map-as-set pattern

Using `map[T]bool` where you only care about membership (yes/no), not counts:

```go
seen := make(map[string]bool)
seen["go"] = true
// Later:
if !seen["python"] {  // not in set
    fmt.Println("new item")
}
```

## You're Ready When You Can...

- [ ] Iterate over a string slice with `for _, s := range slice`
- [ ] Create a map and check if a key exists
- [ ] Append to a slice only when a condition is met

## Next Steps

- [37-countwords](../37-countwords/README.md)
