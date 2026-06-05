## Clean the list

### Instructions

Create a function `CleanList` that takes and returns a slice of strings and performs the following operations on each list item:

- Removes all spaces before and after (but not between words).
- Capitalizes the first letter (first letter only, other ones should be to lowercase).
- Adds its index number before a separator `x/ `.
- An empty slice as input should return an empty slice as output.
- And don't forget the milk! (add it at the end of the list if missing).

### Expected function

```go
func CleanList(lst []string) []string {

}
```

### Usage

Here is an example of how to use the `CleanList` function:

```go
package main

import (
	"fmt"
)

func main() {
	shoppingList := []string{"tomatoes", "pastas", "milk", "salt"}
	fmt.Println(CleanList(shoppingList))
}
```

This will output:

```console
$ go run .
[1/ Tomatoes 2/ Pastas 3/ Milk 4/ Salt]
$
```

### References

- [strings.TrimSpace](https://pkg.go.dev/strings#TrimSpace)
- [strings.Title](https://pkg.go.dev/strings#Title) (deprecated, use custom capitalization)
- [unicode.ToUpper](https://pkg.go.dev/unicode#ToUpper)
- [unicode.ToLower](https://pkg.go.dev/unicode#ToLower)