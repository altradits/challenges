## recursion

### Instructions

Write a function that takes a `uint` and returns its factorial using recursion.

The factorial of a non-negative integer `n` is defined as:

- `0! = 1` (base case)
- `1! = 1` (base case)
- `n! = n × (n-1)!` for `n > 1` (recursive case)

Your function **must** use recursion — that is, it must call itself.

### Expected function

```go
func Factorial(n uint) uint {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(10))
}
```

And its output:

```console
$ go run .
1
1
120
3628800
$
```

### Test Cases

| Input | Expected output |
|-------|----------------|
| `Factorial(0)` | `1` |
| `Factorial(1)` | `1` |
| `Factorial(5)` | `120` |
| `Factorial(10)` | `3628800` |
