# Prerequisites for clean-the-list

## Before You Start

To solve this challenge you need to understand:

### 1. Iterating Over a Slice with Index
Use `for i, item := range slice` to get both the position and value.
```go
lst := []string{"apples", "milk", "bread"}
for i, item := range lst {
    fmt.Printf("%d: %s\n", i, item)
}
// 0: apples
// 1: milk
// 2: bread
```

### 2. strings.TrimSpace
Removes all leading and trailing whitespace (spaces, tabs, newlines).
```go
import "strings"

fmt.Println(strings.TrimSpace("  hello  ")) // "hello"
fmt.Println(strings.TrimSpace("  "))        // ""
```

### 3. Capitalizing Only the First Character
Convert the string to a `[]rune`, uppercase `runes[0]`, then convert back.
```go
import "unicode"

s := "hello world"
runes := []rune(s)
runes[0] = unicode.ToUpper(runes[0])
result := string(runes) // "Hello world"
```

### 4. fmt.Sprintf for Building Formatted Strings
Build a string with a format verb without printing it immediately.
```go
import "fmt"

item := fmt.Sprintf("%d/ %s", 3, "Bread")
// item == "3/ Bread"
```

### 5. append to Build the Result Slice
Collect processed items into a new slice using `append`.
```go
result := []string{}
result = append(result, "1/ Tomatoes")
result = append(result, "2/ Milk")
```

### 6. Checking for a Specific Item
Use a boolean flag to track whether a required item has been seen.
```go
hasMilk := false
for _, item := range lst {
    if strings.ToLower(strings.TrimSpace(item)) == "milk" {
        hasMilk = true
    }
}
if !hasMilk {
    // add milk
}
```

## Review If Stuck

- [09-slicesintro](../09-slicesintro/skills.md) — comprehensive introduction to Go slices: make, append, range, and slice tricks
- [../102-splitjoin/skills.md](../102-splitjoin/skills.md) — covers building a `[]string` slice with `append`
- [../97-hashcode/skills.md](../97-hashcode/skills.md) — covers character-level operations using runes

## You're Ready When You Can...

- [ ] Iterate over a slice with both index and value using `for i, item := range`
- [ ] Trim spaces with `strings.TrimSpace`
- [ ] Capitalize only the first letter using `[]rune` and `unicode.ToUpper`
- [ ] Format strings with `fmt.Sprintf` using `%d` and `%s`
- [ ] Append items to a `[]string` result slice

## Next Steps

- [106-cleanstr](../106-cleanstr/README.md)
