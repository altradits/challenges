# Skills for concatslice

## What You'll Learn

**Previous:** [120-concatalternate](../120-concatalternate/skills.md) | **Next:** [122-fprime](../122-fprime/skills.md)

**Challenge:** Concatenate two int slices into one, preserving order from slice1 followed by slice2.

## Core Concept: Concatenating Slices with `append`

### What Is It?

Concatenating slices means joining them end-to-end so all elements of the first slice appear before all elements of the second. In Go there are two idiomatic ways to do this: using `append` with the spread operator `...`, or manually looping.

### How It Works

**Method 1 — The idiomatic Go way (spread operator):**

The `...` operator expands a slice into individual arguments. When passed to `append`, it appends every element of `slice2` onto `slice1`.

```go
func ConcatSlice(slice1, slice2 []int) []int {
    return append(slice1, slice2...)
}
```

This works even when either slice is empty:
```go
append([]int{1, 2, 3}, []int{}...)     // [1 2 3]
append([]int{}, []int{4, 5, 6}...)     // [4 5 6]
```

**Method 2 — Manual loop (builds understanding):**

```go
func ConcatSlice(slice1, slice2 []int) []int {
    result := []int{}
    for _, v := range slice1 {
        result = append(result, v)
    }
    for _, v := range slice2 {
        result = append(result, v)
    }
    return result
}
```

**Why the spread operator matters:**

The `...` unpacks a slice for use wherever variadic arguments are expected. This is different from passing the slice as a whole:

```go
a := []int{1, 2}
b := []int{3, 4}

append(a, b)    // COMPILE ERROR: cannot use b (type []int) as type int
append(a, b...) // CORRECT: [1 2 3 4]
```

### Common Mistakes

| Mistake | Problem | Fix |
|---------|---------|-----|
| `append(slice1, slice2)` | Compile error — `append` expects `int`, not `[]int` | Use `append(slice1, slice2...)` |
| Modifying `slice1` in place | The original slice may be changed | Create a new slice: `result := make([]int, len(slice1)); copy(result, slice1)` then append |
| Returning `nil` for empty inputs | `nil` and `[]int{}` print differently | `append(nil, []int{}...)` returns `[]int{}` which is safe |

## Solving This Challenge

### Algorithm
1. Use `append(slice1, slice2...)` to join both slices.
2. Return the result.
3. Handle edge cases: both empty, one empty — `append` handles these naturally.

## The Challenge
See [README.md](README.md) for full description.

**Next:** [122-fprime](../122-fprime/README.md)
