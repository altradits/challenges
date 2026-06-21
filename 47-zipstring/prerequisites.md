# Prerequisites for zipstring

## Before You Start

To solve this challenge you need to understand:

### 1. Comparing Adjacent Characters
Access consecutive positions in a string by index to detect where runs change.
```go
s := "aabb"
for i := 1; i < len(s); i++ {
    if s[i] != s[i-1] {
        fmt.Println("change at", i)
    }
}
```

### 2. Tracking a Run Counter
Increment a counter while characters repeat, reset it when they change.
```go
count := 1
for i := 1; i < len(s); i++ {
    if s[i] == s[i-1] {
        count++
    } else {
        // emit count and s[i-1]
        count = 1
    }
}
// emit final run after loop
```

### 3. fmt.Sprintf With %d and %c
`%d` for integers, `%c` for a character (rune or byte). Use these together to format `"3u"`:
```go
result += fmt.Sprintf("%d%c", 3, 'u') // "3u"
result += fmt.Sprintf("%d%c", 1, 'Y') // "1Y"
```

### 4. The Final Flush Pattern
When iterating and comparing to the previous element, the last run is never "closed" inside the loop. Always emit it after the loop ends.
```go
// inside loop: emit when s[i] != s[i-1]
// after loop: emit the remaining run
result += fmt.Sprintf("%d%c", count, s[len(s)-1])
```

### 5. Empty String Guard
Check for empty string before accessing `s[len(s)-1]` to avoid a panic.
```go
if len(s) == 0 {
    return ""
}
```

## Review If Stuck

- [../21-countrepeats/skills.md](../21-countrepeats/skills.md) — covers the same "compare to previous, track state" pattern; this challenge adds counting each run's length
- [../41-itoa-35/skills.md](../41-itoa-35/skills.md) — covers converting integers to their character representation

## You're Ready When You Can...

- [ ] Compare `s[i]` to `s[i-1]` to detect run boundaries
- [ ] Maintain a run counter that increments on match and resets on change
- [ ] Use `fmt.Sprintf("%d%c", count, char)` to format each run
- [ ] Remember to flush the final run after the loop

## Next Steps

- [Next challenge](../48-addprimesum/README.md)
