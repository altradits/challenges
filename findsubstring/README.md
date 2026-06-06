## findsubstring

### Instructions

Write a function that takes two strings as arguments: a `text` and a `substring`. The function should return the index of the first occurrence of `substring` in `text`.

- If the `substring` is not found, return -1.
- If the `substring` is empty, return 0.
- The search should be case-sensitive.

### Expected function

```go
func FindSubstring(text, substring string) int {

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
	fmt.Println(FindSubstring("Hello World", "World"))
	fmt.Println(FindSubstring("Hello World", "hello"))
	fmt.Println(FindSubstring("Hello World", ""))
	fmt.Println(FindSubstring("Hello World", "xyz"))
	fmt.Println(FindSubstring("banana", "ana"))
}
```

And its output:

```console
$ go run .
6
-1
0
-1
1
$
```
