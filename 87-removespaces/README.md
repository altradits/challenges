## removespaces

### Instructions

Write a function that takes a string as an argument and returns a new string with all spaces removed.

- If the string is empty, return an empty string.
- Only remove space characters (' '), not other whitespace like tabs or newlines.

### Expected function

```go
func RemoveSpaces(s string) string {

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
	fmt.Println(RemoveSpaces("Hello World"))
	fmt.Println(RemoveSpaces("Go is fun!"))
	fmt.Println(RemoveSpaces(""))
	fmt.Println(RemoveSpaces("NoSpacesHere"))
}
```

And its output:

```console
$ go run .
HelloWorld
Goisfun!

NoSpacesHere
$
```
