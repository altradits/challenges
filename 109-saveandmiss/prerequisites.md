# Prerequisites for 109-saveandmiss

## Before You Start

### 1. The Modulo Operator for Cycling

`i % n` gives a value that cycles from 0 to n-1, then repeats:

```go
for i := 0; i < 12; i++ {
    fmt.Printf("i=%d  i%%4=%d\n", i, i%4)
}
// i=0  i%4=0
// i=1  i%4=1
// i=2  i%4=2
// i=3  i%4=3
// i=4  i%4=0  ← resets
// ...
```

For save-and-miss with `n=3`: `i % 6` cycles 0,1,2,3,4,5,0,1,2,...  
Characters where `i%6 < 3` are in the save zone.

### 2. A Manual Cyclic Counter

An alternative to modulo: maintain a `pos` counter and reset it manually:

```go
pos := 0
for _, c := range s {
    // use pos
    pos++
    if pos == cycle {
        pos = 0  // reset at end of cycle
    }
}
```

This is equivalent to `i % cycle` but easier to read when the cycle length is complex.

### 3. Conditional Character Collection

Add a character to the result only when in the save zone:

```go
result := ""
for i, c := range s {
    if i%6 < 3 {
        result += string(c)
    }
    // else: skip (miss zone)
}
```

### 4. Edge Case: n <= 0

When `n` is 0 or negative, the cycle length would be 0 or negative, which causes division by zero or undefined behavior. Return the original string unchanged:

```go
if n <= 0 {
    return s
}
```

## Review If Stuck

- [107-thirdchar](../107-thirdchar/skills.md) — covers modulo-based periodic selection
- [108-zipstring](../108-zipstring/skills.md) — covers state tracking with a counter across iterations

## You're Ready When You Can...

- [ ] Use `i % (2*n) < n` to identify characters in the save zone
- [ ] Maintain a manual position counter and reset it when it reaches the cycle boundary
- [ ] Handle the edge case where `n <= 0`
- [ ] Trace through an example on paper: n=2, string="abcdef" → "ab ef"? No → "abef"

## Next Steps

- [110-reversestrcap](../110-reversestrcap/README.md)
