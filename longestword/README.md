## longestword

### Instructions

Write a function that takes a string and returns the longest word in the string.

- A word is a sequence of characters delimited by spaces.
- If there are multiple words with the same maximum length, return the first one.
- If the string is empty or contains only spaces, return an empty string.

### Expected function

```go
func LongestWord(s string) string {

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
	fmt.Println(LongestWord("hello world"))
	fmt.Println(LongestWord("the quick brown fox"))
	fmt.Println(LongestWord("a bb ccc dd eee"))
	fmt.Println(LongestWord(""))
	fmt.Println(LongestWord("   "))
	fmt.Println(LongestWord("single"))
}
```

And its output:

```console
$ go run .
hello
quick
ccc

single