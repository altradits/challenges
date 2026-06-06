## toupper

### Instructions

Write a function that takes a string as an argument and returns a new string with all lowercase letters converted to uppercase.

- Non-alphabetic characters (numbers, spaces, punctuation) should remain unchanged.
- If the string is empty, return an empty string.

### Expected function

```go
func ToUpper(s string) string {

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
	fmt.Println(ToUpper("Hello"))
	fmt.Println(ToUpper("Go is fun!"))
	fmt.Println(ToUpper("123abc"))
	fmt.Println(ToUpper(""))
}
```

And its output:

```console
$ go run .
HELLO
GO IS FUN!
123ABC

$
```
