# Skills for 36-sort-and-math

## What You'll Learn

**Previous:** [20-regexp](../20-regexp/skills.md) | **Next:** [22-modules-and-tooling](../22-modules-and-tooling/skills.md)

**Challenge:** Sort slices and use mathematical functions.

## The `sort` Package

```go
import "sort"
```

### Sorting Built-in Types

```go
ints    := []int{5, 2, 8, 1}
floats  := []float64{3.14, 1.41, 2.72}
strings := []string{"banana", "apple", "cherry"}

sort.Ints(ints)        // [1 2 5 8]
sort.Float64s(floats)  // [1.41 2.72 3.14]
sort.Strings(strings)  // [apple banana cherry]
```

Sorts in-place. The original slice is modified.

### Custom Sort: `sort.Slice`

```go
type Person struct{ Name string; Age int }

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Carol", 35},
}

// Sort by age ascending
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age
})
// [{Bob 25} {Alice 30} {Carol 35}]

// Sort by name descending
sort.Slice(people, func(i, j int) bool {
    return people[i].Name > people[j].Name
})
```

### Checking if a Slice is Sorted

```go
sort.IntsAreSorted([]int{1, 2, 3})   // true
sort.IntsAreSorted([]int{3, 1, 2})   // false
```

### Binary Search: `sort.Search`

```go
nums := []int{1, 3, 5, 7, 9}
target := 5
i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
// i = 2 (index of 5)
```

`sort.Search` requires the slice to be sorted.

## The `math` Package

```go
import "math"
```

### Common Functions

```go
math.Sqrt(16)         // 4 — square root
math.Pow(2, 10)       // 1024 — 2^10
math.Abs(-5.5)        // 5.5 — absolute value
math.Min(3.0, 5.0)    // 3 — minimum
math.Max(3.0, 5.0)    // 5 — maximum
math.Floor(3.7)       // 3 — round down
math.Ceil(3.2)        // 4 — round up
math.Round(3.5)       // 4 — round to nearest
math.Log(math.E)      // 1 — natural logarithm
math.Log2(8)          // 3 — log base 2
math.Log10(100)       // 2 — log base 10
math.Pi               // 3.141592... — constant
math.E                // 2.718281... — Euler's number
math.MaxInt64         // max int64 value
math.Inf(1)           // +Infinity
math.IsNaN(x)         // check for not-a-number
```

### Integer Math

For integers, Go has no `math` functions — use operators or `sort`:

```go
// No math.Min for int — write your own or use sort
func minInt(a, b int) int {
    if a < b { return a }
    return b
}
```

Go 1.21+ has `min()` and `max()` as built-in functions.

### Solving the Challenge

```go
package piscine

import (
    "math"
    "sort"
)

func MedianSorted(nums []float64) float64 {
    if len(nums) == 0 {
        return 0
    }
    sort.Float64s(nums)
    n := len(nums)
    if n%2 == 0 {
        return (nums[n/2-1] + nums[n/2]) / 2
    }
    return nums[n/2]
}

func TopN(nums []int, n int) []int {
    cp := make([]int, len(nums))
    copy(cp, nums)
    sort.Sort(sort.Reverse(sort.IntSlice(cp)))
    if n > len(cp) {
        n = len(cp)
    }
    return cp[:n]
}

func Hypotenuse(a, b float64) float64 {
    return math.Sqrt(a*a + b*b)
}

func RoundToNearest(f float64) float64 {
    return math.Round(f*2) / 2
}
```

**Next:** [22-modules-and-tooling](../22-modules-and-tooling/README.md)
