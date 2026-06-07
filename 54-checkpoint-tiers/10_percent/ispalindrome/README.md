## ispalindrome

### Instructions

Write a function that takes a string as an argument and returns `true` if the string is a palindrome, otherwise returns `false`.

- A palindrome reads the same forwards and backwards (ignoring case).
- Ignore spaces and punctuation. Only consider alphanumeric characters.
- If the string is empty, return `true` (an empty string is a palindrome).

### Expected function

```go
func IsPalindrome(s string) bool {

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
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("Hello"))
	fmt.Println(IsPalindrome("A man a plan a canal Panama"))
	fmt.Println(IsPalindrome(""))
}
```

And its output:

```console
$ go run .
true
false
true
true
$
```
