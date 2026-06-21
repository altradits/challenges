# Skills for 10-mapsintro

## What You'll Learn

**Previous:** [09-slicesintro](../09-slicesintro/skills.md) | **Next:** [11-arrays](../11-arrays/skills.md)

**Challenge:** Write `WordFrequency(s string) map[string]int` that counts how many times each word appears.

## Core Concept: Maps — Go's Built-in Key-Value Store

### What is a Map?

A map stores key-value pairs with O(1) average lookup time. Think of it as a dictionary, hash table, or associative array.

```go
// map[KeyType]ValueType
m := map[string]int{
    "go":  2,
    "is":  1,
}
fmt.Println(m["go"])  // 2
```

### Creating a Map

**Method 1: map literal** (when you know the initial values)
```go
capitals := map[string]string{
    "Nigeria": "Abuja",
    "Ghana":   "Accra",
}
```

**Method 2: `make`** (when you'll populate it dynamically)
```go
counts := make(map[string]int)  // empty map, ready to use
```

**Never use `var m map[string]int` for writing** — that creates a `nil` map and panics on write:
```go
var m map[string]int
m["key"] = 1  // PANIC: assignment to nil map
```

Always use `make` or a map literal.

### Reading and Writing

```go
m := make(map[string]int)

// Write
m["apple"] = 3
m["banana"] = 1

// Read (returns zero value if key missing)
count := m["apple"]   // 3
count = m["cherry"]   // 0 (key doesn't exist → zero value for int)
```

### The Zero Value Trick

Because a missing key returns 0, you can increment counts directly:

```go
m := make(map[string]int)
m["go"]++  // m["go"] was 0, now 1
m["go"]++  // now 2
```

This is the core pattern for counting anything.

### Checking If a Key Exists (comma-ok pattern)

```go
value, ok := m["apple"]
if ok {
    fmt.Println("found:", value)
} else {
    fmt.Println("not found")
}

// Short form: just check existence
_, exists := m["banana"]
```

The comma-ok pattern is how you distinguish "key missing" from "key exists with zero value."

### Deleting a Key

```go
delete(m, "apple")  // remove "apple" from the map
```

### Iterating Over a Map

```go
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}
```

**Note:** Map iteration order is RANDOM. Every run may print keys in a different order.

### Map as a Set

When you only need to track whether something exists (not a count), use `map[T]bool`:

```go
seen := make(map[string]bool)
words := []string{"go", "is", "go", "fast"}

unique := []string{}
for _, w := range words {
    if !seen[w] {
        seen[w] = true
        unique = append(unique, w)
    }
}
// unique = ["go", "is", "fast"]
```

### Solving the Challenge

```go
import "strings"

func WordFrequency(s string) map[string]int {
    freq := make(map[string]int)
    if s == "" {
        return freq
    }
    words := strings.Fields(s)   // split on whitespace
    for _, word := range words {
        freq[word]++              // zero-value trick: missing key → 0, then +1
    }
    return freq
}
```

### Common Mistakes

| Mistake | Fix |
|---------|-----|
| `var m map[string]int` then writing to it | Use `make(map[string]int)` |
| Assuming iteration order is stable | Maps are unordered — sort keys if you need sorted output |
| `m["key"] == nil` to check existence | Use `_, ok := m["key"]` — nil doesn't work for int/string values |

### Maps vs Slices

| | Map | Slice |
|---|-----|-------|
| Access by | Any comparable key | Integer index 0..n-1 |
| Order | None (random iteration) | Preserved |
| Lookup | O(1) average | O(n) for search |
| Use when | Key-based lookup | Sequential data |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [11-arrays](../11-arrays/README.md)
