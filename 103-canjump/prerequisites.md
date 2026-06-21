# Prerequisites for canjump

## Before You Start

To solve this challenge you need to understand:

### 1. Slice Indexing With []uint
A slice of unsigned integers is accessed like any other slice. Use `len(s)` for bounds checking.
```go
nums := []uint{2, 3, 1, 1, 4}
fmt.Println(nums[0])      // 2
fmt.Println(len(nums))    // 5
fmt.Println(nums[len(nums)-1]) // 4 — last element
```

### 2. Type Conversion Between uint and int
When doing arithmetic with `uint` values and `int` position variables, convert explicitly.
```go
var jump uint = nums[0]  // type is uint
pos := 0                  // type is int
pos += int(jump)          // convert uint to int for the addition
```

### 3. Infinite Loop With for {}
A loop with no condition runs forever. Use `return` inside to exit.
```go
pos := 0
for {
    // do work
    if someCondition {
        return true
    }
    if anotherCondition {
        return false
    }
}
```

### 4. Bounds Checking
Before using `pos` as an index, ensure it is within `[0, len(nums)-1]`.
```go
if pos >= len(nums) {
    return false  // jumped past the end
}
if pos == len(nums)-1 {
    return true   // reached the last index
}
```

### 5. Detecting a Stuck State
If `nums[pos] == 0` and `pos` is not the last index, you cannot advance.
```go
if nums[pos] == 0 {
    return false
}
```

## Review If Stuck

- [../93-findprevprime/skills.md](../93-findprevprime/skills.md) — covers loops that terminate based on computed conditions
- [../90-clean-the-list/skills.md](../90-clean-the-list/skills.md) — covers working with slice values

## You're Ready When You Can...

- [ ] Access elements of a `[]uint` slice by index
- [ ] Convert `uint` to `int` for arithmetic with a position counter
- [ ] Write a `for {}` loop that uses `return` to exit on success or failure
- [ ] Check bounds correctly: `pos == len-1` for success, `pos >= len` for out-of-bounds

## Next Steps

- [104-chunk](../104-chunk/README.md)
