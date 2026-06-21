# Prerequisites for fifthandskip

## Before You Start

To solve this challenge you need to understand:

### 1. Skipping Characters with `continue`
Use `continue` in a `for` loop to skip the rest of the loop body for the current iteration.

```go
for _, c := range str {
    if c == ' ' {
        continue  // skip spaces
    }
    // only non-space characters reach here
}
```

### 2. The Modulo Operator `%` for Cyclical Position
`count % 6` tells you where you are within a repeating pattern of 6 positions (0, 1, 2, 3, 4, 5, 0, 1, ...).

```go
count := 0
for count < 12 {
    fmt.Println(count % 6)  // 0, 1, 2, 3, 4, 5, 0, 1, 2, 3, 4, 5
    count++
}
```

### 3. Conditional String Building
Append to the result based on conditions. Add a separator before new groups.

```go
if pos == 0 && count > 0 {
    result += " "  // space between groups
}
result += string(c)
```

### 4. Edge Case Handling — Return Early
Check for empty strings and strings shorter than 5 non-space characters before processing.

```go
if str == "" {
    return "\n"
}
if nonSpaceCount < 5 {
    return "Invalid Input\n"
}
```

### 5. Counting Characters Matching a Condition
Count non-space characters separately from the main loop.

```go
nonSpace := 0
for _, c := range str {
    if c != ' ' {
        nonSpace++
    }
}
```

## Review If Stuck
- [111-saveandmiss](../111-saveandmiss/skills.md) — covers periodic selection with a counter
- [113-wdmatch](../113-wdmatch/skills.md) — covers walking strings while tracking a separate counter

## You're Ready When You Can...
- [ ] Use `continue` to skip characters in a `for range` loop
- [ ] Use `count % 6` to determine position in a repeating 6-item cycle
- [ ] Add a separator only between groups (not before the first)
- [ ] Count non-space characters before the main loop
- [ ] Return special strings for empty input and too-short input

## Next Steps
- [115-notdecimal](../115-notdecimal/README.md)
