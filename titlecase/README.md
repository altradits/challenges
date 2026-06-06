## Title Case

### Instructions

Write a function that takes a string as an argument and returns the string converted to title case.

- Title case means the first letter of each word is uppercase, and all other letters are lowercase.
- A word is defined as a sequence of characters separated by spaces.
- If the string is empty, return an empty string.

### Expected function

```go
func TitleCase(s string) string {

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
	fmt.Println(TitleCase("hello world"))
	fmt.Println(TitleCase("Go is fun!"))
	fmt.Println(TitleCase(""))
	fmt.Println(TitleCase("multiple   spaces   here"))
	fmt.Println(TitleCase("ALL CAPS"))
}
```

And its output:

```console
$ go run .
Hello World
Go Is Fun!

Multiple   Spaces   Here
All Caps
$
```
