## splitjoin

### Instructions

Write two functions:

1. `Split(s string, sep string) []string` - Splits a string by a separator and returns a slice of substrings.
2. `Join(arr []string, sep string) string` - Joins a slice of strings with a separator and returns a single string.

- If the separator is empty, each character becomes its own element in the slice.
- If the input string is empty, return a slice with one empty string element.
- Do not use the built-in `strings.Split` or `strings.Join` functions.

### Expected functions

```go
func Split(s string, sep string) []string {

}

func Join(arr []string, sep string) string {

}
```

### Usage

Here is a possible program to test your functions:

```go
package main

import (
	"fmt"
)

func main() {
	// Test Split
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("Hello World", " "))
	fmt.Printf("%q\n", Split("abc", ""))
	fmt.Printf("%q\n", Split("", ","))

	// Test Join
	fmt.Println(Join([]string{"a", "b", "c"}, ","))
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]string{"a", "b", "c"}, ""))
	fmt.Println(Join([]string{""}, ","))
}
```

And its output:

```console
$ go run .
["a" "b" "c"]
["Hello" "World"]
["a" "b" "c"]
[""]
a,b,c
Hello World
abc

$
```
