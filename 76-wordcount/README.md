## wordcount

### Instructions

Write a function that takes a string as an argument and returns the number of words in the string.

- A word is defined as a sequence of non-space characters separated by one or more spaces.
- Leading and trailing spaces should be ignored.
- If the string is empty or contains only spaces, return 0.

### Expected function

```go
func WordCount(s string) int {

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
	fmt.Println(WordCount("Hello World"))
	fmt.Println(WordCount("Go is fun!"))
	fmt.Println(WordCount(""))
	fmt.Println(WordCount("   "))
	fmt.Println(WordCount("One  two   three"))
}
```

And its output:

```console
$ go run .
2
3
0
0
3
$
```
