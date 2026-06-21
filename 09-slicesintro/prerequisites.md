# Prerequisites for slicesintro

## Before You Start

To solve this challenge you need to understand:

### 1. For Loops
Use `for _, v := range slice` to iterate over every element in a slice.
```go
nums := []int{1, 2, 3}
for _, n := range nums {
    fmt.Println(n)
}
// 1
// 2
// 3
```

See [../07-forloops/skills.md](../07-forloops/skills.md) for a full introduction.

### 2. Maps (for an efficient solution)
A map lets you check in O(1) time whether a value has already been seen.
```go
seen := make(map[int]bool)
seen[42] = true

if seen[42] {
    fmt.Println("already seen")
}
```

See [../93-countrepeats/skills.md](../93-countrepeats/skills.md) for a full introduction to maps.

### 3. Defining Functions
The challenge asks you to write a named function that accepts a `[]int` parameter and returns a `[]int`.
```go
func Unique(nums []int) []int {
    // your code here
    return result
}
```

See [../06-functions/skills.md](../06-functions/skills.md) for function syntax.

## You're Ready When You Can...

- [ ] Declare a slice literal: `nums := []int{1, 2, 3}`
- [ ] Iterate over a slice with `for _, v := range nums`
- [ ] Append to a slice: `result = append(result, v)`
- [ ] Create a map and use it to track seen values
- [ ] Write a function that takes a `[]int` parameter and returns a `[]int`

## Next Steps

- [10-mapsintro](../10-mapsintro/README.md)
