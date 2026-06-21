## slicesintro

### Instructions

Write a function `Unique` that takes a slice of integers and returns a new slice containing only the unique values, in the order they first appeared.

### Expected function

```go
func Unique(nums []int) []int {

}
```

### Usage

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Unique([]int{1, 2, 2, 3, 1, 4})) // [1 2 3 4]
	fmt.Println(piscine.Unique([]int{}))                  // []
	fmt.Println(piscine.Unique([]int{5}))                 // [5]
}
```

### Test cases

| Input                      | Output        |
|---------------------------|---------------|
| `[]int{1, 2, 2, 3, 1, 4}` | `[]int{1, 2, 3, 4}` |
| `[]int{}`                  | `[]int{}`     |
| `[]int{5}`                 | `[]int{5}`    |
| `[]int{7, 7, 7}`           | `[]int{7}`    |
| `[]int{3, 1, 2, 1, 3}`    | `[]int{3, 1, 2}` |
