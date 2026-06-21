## arrays

### Instructions

Write a function `SumArray` that takes a fixed-size array of 5 integers and returns the sum of all 5 values.

### Expected function

```go
func SumArray(arr [5]int) int {

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
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(piscine.SumArray(arr)) // 15

	arr2 := [5]int{0, 0, 0, 0, 0}
	fmt.Println(piscine.SumArray(arr2)) // 0

	arr3 := [5]int{-1, -2, 3, 4, 5}
	fmt.Println(piscine.SumArray(arr3)) // 9
}
```

### Test cases

| Input                        | Output |
|-----------------------------|--------|
| `[5]int{1, 2, 3, 4, 5}`     | `15`   |
| `[5]int{0, 0, 0, 0, 0}`     | `0`    |
| `[5]int{-1, -2, 3, 4, 5}`   | `9`    |
| `[5]int{10, 20, 30, 40, 50}`| `150`  |
| `[5]int{-5, -5, -5, -5, -5}`| `-25`  |
