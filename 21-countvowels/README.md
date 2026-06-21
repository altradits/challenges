# 09. Count Vowels

### Instructions

Write a function that takes a string as an argument and returns the number of vowels (a, e, i, o, u) in the string.

- Count both lowercase and uppercase vowels.
- If the string is empty, return 0.

### Expected function

```go
func CountVowels(s string) int {

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
	fmt.Println(CountVowels("Hello World"))
	fmt.Println(CountVowels("Go is fun!"))
	fmt.Println(CountVowels(""))
	fmt.Println(CountVowels("AEIOUaeiou"))
}
```

And its output:

```console
$ go run .
3
3
0
10
$
```
