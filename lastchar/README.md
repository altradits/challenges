## lastchar

### Instructions

Write a function that takes a string as an argument and returns the last character of the string.

- If the string is empty, return an empty string.
- If the string has only one character, return that character.

### Expected function

```go
func LastChar(s string) string {

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
	fmt.Println(LastChar("Hello"))
	fmt.Println(LastChar(""))
	fmt.Println(LastChar("G"))
	fmt.Println(LastChar("Go is fun"))
}
```

And its output:

```console
$ go run .
o

G
n
$
```
