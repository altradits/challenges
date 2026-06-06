## reversestring

### Instructions

Write a function that takes a string as an argument and returns the string reversed.

- If the string is empty, return an empty string.
- Do not use any built-in reverse functions. Build the reversed string manually.

### Expected function

```go
func ReverseString(s string) string {

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
	fmt.Println(ReverseString("Hello"))
	fmt.Println(ReverseString("Go is fun!"))
	fmt.Println(ReverseString(""))
	fmt.Println(ReverseString("a"))
}
```

And its output:

```console
$ go run .
olleH
!nuf si oG

a
$
```
