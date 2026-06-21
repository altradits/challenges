## forloops

### Instructions

Write a function called `CountDown` that takes an integer `n` as an argument and prints all integers from `n` down to `0`, each on its own line.

### Expected function

```go
func CountDown(n int) {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	piscine.CountDown(5)
	fmt.Println("---")
	piscine.CountDown(0)
}
```

And its output:

```console
$ go run .
5
4
3
2
1
0
---
0
$
```
