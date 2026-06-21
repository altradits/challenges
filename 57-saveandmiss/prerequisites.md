# Prerequisites for saveandmiss

## Before You Start

To solve this challenge you need to understand:

### 1. `for range` Over a String
Iterate over every character as a rune.

```go
for _, c := range "hello" {
    fmt.Printf("%c\n", c)
}
```

### 2. Boolean Flag (Toggle)
Use a `bool` to track which mode you are in. Flip it with `!`.

```go
save := true
save = !save  // now false
save = !save  // now true again
```

### 3. Counter with Reset
Track your position within a group using an integer counter. Reset to 0 when the group is complete.

```go
count := 0
for _, c := range s {
    count++
    if count == num {
        count = 0  // start fresh for next group
        // flip mode
    }
}
```

### 4. Building a String with Concatenation
Append to a result string only when in "save" mode.

```go
result := ""
if save {
    result += string(c)
}
```

### 5. Early Return for Edge Cases
When `num <= 0`, the spec says return the original string.

```go
func SaveAndMiss(arg string, num int) string {
    if num <= 0 {
        return arg
    }
    // ... rest of logic
}
```

## Review If Stuck
- [56-reversestrcap](../56-reversestrcap/skills.md) — covers walking strings with a counter
- Prior `for range` challenges — covers iterating strings as runes

## You're Ready When You Can...
- [ ] Use `for _, c := range s` to iterate characters
- [ ] Toggle a `bool` variable with `!`
- [ ] Use a counter variable that resets to 0 at a boundary
- [ ] Conditionally append to a string only when a flag is true
- [ ] Return the input unchanged when a condition is met

## Next Steps
- [58-union](../58-union/README.md)
