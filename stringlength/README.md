## stringlength

### Instructions

Write a function that takes a string as an argument and returns its length.

- If the string is empty, return 0.
- Do not use the built-in `len()` function. Instead, iterate through the string and count each character manually.

### Expected function

```go
func StringLength(s string) int {

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
	fmt.Println(StringLength("Hello"))
	fmt.Println(StringLength(""))
	fmt.Println(StringLength("Go is fun"))
}
```

And its output:

```console
$ go run .
5
0
9
$
```
