# Prerequisites for 125-romannumbers

## Before You Start

### 1. Struct types — grouping related values

The Roman numeral table uses a struct to pair a value with its symbol:

```go
type romanPair struct {
    value  int
    symbol string
}
```

Review: [06-functions](../06-functions/skills.md) — function syntax; structs are the next natural step.

### 2. Slice of structs — ordered table

```go
table := []romanPair{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    // ...
    {1, "I"},
}
```

The table must be ordered largest to smallest for the greedy algorithm to work.

### 3. The greedy loop pattern

Keep subtracting the largest fitting value until the remainder is zero:

```go
for n > 0 {
    for _, pair := range table {
        for n >= pair.value {
            result += pair.symbol
            n -= pair.value
        }
    }
}
```

Review: [07-forloops](../07-forloops/skills.md) — nested for loops.

### 4. `os.Args` — reading the input number

```go
import (
    "fmt"
    "os"
    "strconv"
)

n, err := strconv.Atoi(os.Args[1])
if err != nil { return }
```

Review: [79-digitlen](../79-digitlen/skills.md) — strconv.Atoi usage.

### 5. `strings.Builder` — building the result string

```go
var b strings.Builder
b.WriteString("M")
b.WriteString("CM")
result := b.String()
```

Review: [08-stringspackage](../08-stringspackage/skills.md) — strings.Builder.

## You're Ready When You Can...

- [ ] Define a struct type with two fields
- [ ] Create a slice of structs as a lookup table, ordered largest to smallest
- [ ] Use a nested for loop to greedily consume values
- [ ] Convert a command-line argument string to int with `strconv.Atoi`
- [ ] Build a result string with `strings.Builder`

## Next Steps

- [126-brackets](../126-brackets/README.md)
