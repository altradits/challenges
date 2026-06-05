## countrepeats

### Instructions

Write a function that takes a string and returns the number of consecutive repeated characters in the string.

- A "repeat" is when the same character appears consecutively two or more times.
- Count each group of consecutive repeats as one.
- If the string is empty or has no consecutive repeats, return 0.

### Expected function

```go
func CountRepeats(s string) int {

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
	fmt.Println(CountRepeats("hello"))
	fmt.Println(CountRepeats("helloo"))
	fmt.Println(CountRepeats("heelloo"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats(""))
	fmt.Println(CountRepeats("aaa"))
}
```

And its output:

```console
$ go run .
1
2
2
0
0
1