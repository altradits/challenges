## tolower

### Instructions

Write a function that takes a string as an argument and returns a new string with all uppercase letters converted to lowercase.

- Non-alphabetic characters (numbers, spaces, punctuation) should remain unchanged.
- If the string is empty, return an empty string.

### Expected function

```go
func ToLower(s string) string {

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
	fmt.Println(ToLower("Hello"))
	fmt.Println(ToLower("Go is fun!"))
	fmt.Println(ToLower("123ABC"))
	fmt.Println(ToLower(""))
}
```

And its output:

```console
$ go run .
hello
go is fun!
123abc

$
```
