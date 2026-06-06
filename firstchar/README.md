## firstchar

### Instructions

Write a function that takes a string as an argument and returns the first character of the string.

- If the string is empty, return an empty string.
- If the string has only one character, return that character.

### Expected function

```go
func FirstChar(s string) string {

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
	fmt.Println(FirstChar("Hello"))
	fmt.Println(FirstChar(""))
	fmt.Println(FirstChar("G"))
	fmt.Println(FirstChar("Go is fun"))
}
```

And its output:

```console
$ go run .
H

G
G
$
```
