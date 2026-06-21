## replaceall

### Instructions

Write a function that takes three strings as arguments: `text`, `old`, and `new`. The function should return a new string where all occurrences of `old` in `text` are replaced with `new`.

- If `old` is empty, return `text` unchanged.
- The replacement should be case-sensitive.
- Do not use the built-in `strings.ReplaceAll` function. Build the result manually.

### Expected function

```go
func ReplaceAll(text, old, new string) string {

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
	fmt.Println(ReplaceAll("Hello World", "l", "L"))
	fmt.Println(ReplaceAll("banana", "na", "NA"))
	fmt.Println(ReplaceAll("Hello World", "", "x"))
	fmt.Println(ReplaceAll("abcabc", "abc", "xyz"))
	fmt.Println(ReplaceAll("no match here", "z", "Z"))
}
```

And its output:

```console
$ go run .
HeLLo WorLd
baNANA
Hello World
xyzxyz
no match here
$
```
