# Skills for slicesintro

## What You'll Learn

**Previous:** [../33-wordanatomy2/skills.md](../33-wordanatomy2/skills.md) | **Next:** [../35-clean-the-list/skills.md](../35-clean-the-list/skills.md)

**Challenge:** Return only the unique values from a slice of integers, preserving first-occurrence order.

## Core Concept: Slices — Go's Most-Used Data Structure

### What Is a Slice?

A slice is a dynamically-sized, ordered collection of elements of the same type. It is the most commonly used data structure in Go. Unlike arrays (which have a fixed size baked into the type), slices can grow and shrink at runtime.

```go
nums := []int{1, 2, 3}   // a slice of three ints
```

### Creating Slices

**Literal syntax** — write the values directly:
```go
nums := []int{1, 2, 3}
words := []string{"hello", "world"}
```

**`make()` — allocate with length and/or capacity:**
```go
nums := make([]int, 5)       // [0, 0, 0, 0, 0]  len=5, cap=5
nums := make([]int, 0, 10)   // []              len=0, cap=10
```
- The first number is the **length** (how many elements exist now).
- The second number is the **capacity** (how many can fit before reallocation).

**From an array:**
```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]  // [2, 3, 4]
```

### Appending to a Slice

Use `append` to add elements. Always assign the result back — `append` may return a new underlying array.

```go
s := []int{1, 2}
s = append(s, 3)           // [1, 2, 3]
s = append(s, 4, 5)        // [1, 2, 3, 4, 5]

other := []int{6, 7}
s = append(s, other...)    // [1, 2, 3, 4, 5, 6, 7]  (spread with ...)
```

### Slicing a Slice

You can take a sub-slice with `s[low:high]` — the result shares the same underlying array.

```go
s := []int{0, 1, 2, 3, 4}
s[1:3]  // [1, 2]    — indices 1 and 2
s[:2]   // [0, 1]    — from start up to (not including) index 2
s[3:]   // [3, 4]    — from index 3 to end
s[:]    // [0,1,2,3,4] — the whole slice
```

### `len()` and `cap()`

```go
s := make([]int, 3, 10)
len(s)  // 3  — number of elements currently in the slice
cap(s)  // 10 — total space allocated
```

### Iterating Over a Slice

```go
nums := []int{10, 20, 30}

// Index and value
for i, v := range nums {
    fmt.Println(i, v)
}

// Value only
for _, v := range nums {
    fmt.Println(v)
}

// Index only
for i := range nums {
    fmt.Println(i)
}
```

### Nil Slice vs Empty Slice

```go
var s []int      // nil slice   — len=0, s == nil is true
s := []int{}     // empty slice — len=0, s == nil is false
```

Both have length 0 and both work with `append` and `range`. The difference matters when checking `s == nil`.

### Useful Slice Tricks

**Delete element at index `i` (order not preserved — faster):**
```go
s[i] = s[len(s)-1]
s = s[:len(s)-1]
```

**Delete element at index `i` (preserve order):**
```go
s = append(s[:i], s[i+1:]...)
```

**Copy a slice** (independent copy — changes do not affect the original):
```go
src := []int{1, 2, 3}
dest := make([]int, len(src))
copy(dest, src)
```

### Why Slices Are Reference Types

A slice header contains a pointer to the underlying array. When you pass a slice to a function, the function receives a copy of the header — but both headers point to the **same** underlying array.

```go
func doubleFirst(s []int) {
    s[0] *= 2   // modifies the caller's underlying array
}

nums := []int{1, 2, 3}
doubleFirst(nums)
fmt.Println(nums)  // [2, 2, 3]
```

However, `append` may allocate a new array, so the caller's slice is unaffected if you only assign inside the function:

```go
func tryAppend(s []int) {
    s = append(s, 99)  // local reassignment — caller does NOT see the new element
}
```

If a function needs to return an extended slice, it should return the slice as a value.

## Solving This Challenge

### Algorithm

Use a map to track which numbers have already been seen:

```go
func Unique(nums []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, n := range nums {
        if !seen[n] {
            seen[n] = true
            result = append(result, n)
        }
    }
    return result
}
```

1. Start with an empty `result` slice and an empty `seen` map.
2. For each number in `nums`, check if it is already in `seen`.
3. If not, mark it as seen and append it to `result`.
4. Return `result`.

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| Not assigning `append` back | `append(result, n)` discards the new slice | Always write `result = append(result, n)` |
| Returning `nil` for empty input | Some callers distinguish `nil` from `[]int{}` | Return `[]int{}` explicitly, or just return `result` (it starts as `[]int{}`) |
| Using `len(nums)` without checking 0 | Fine here — `range` over an empty slice does nothing | No action needed, but good to be aware |

## The Challenge

See [README.md](README.md) for full description and test cases.

**Next:** [../35-clean-the-list/skills.md](../35-clean-the-list/skills.md)
