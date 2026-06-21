# Prerequisites for 121-stringcount

## Before You Start

### 1. Finding a substring position with `strings.Index`

```go
import "strings"
idx := strings.Index("hello world", "world")  // 6
idx  = strings.Index("hello world", "xyz")    // -1
```

Review: [120-stringindex](../120-stringindex/skills.md)

### 2. Incrementing a counter in a loop

```go
count := 0
for _, c := range s {
    if c == target {
        count++
    }
}
```

Review: [92-countchar](../92-countchar/skills.md)

### 3. Slicing a string to advance past a match

```go
s := "hello hello"
idx := strings.Index(s, "hello")  // 0
s = s[idx+len("hello"):]          // " hello" — rest after match
```

### 4. The infinite loop pattern with `break`

```go
for {
    // do something
    if donecondition {
        break
    }
}
```

## You're Ready When You Can...

- [ ] Use `strings.Index` to find a substring
- [ ] Use slicing to advance past a found position
- [ ] Write a loop that counts matches until none are left

## Next Steps

- [122-stringprefix](../122-stringprefix/README.md)
