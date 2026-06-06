## isempty

### Instructions

Write a function that takes a string as an argument and returns `true` if the string is empty, otherwise returns `false`.

- An empty string has a length of 0.
- Do not use the built-in `len()` function. Instead, iterate through the string to check if it has any characters.

### Expected function

```go
func IsEmpty(s string) bool {

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
	fmt.Println(IsEmpty("Hello"))
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty(" "))
	fmt.Println(IsEmpty("G"))
}
```

And its output:

```console
$ go run .
false
true
false
false
$
```
