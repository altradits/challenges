## functions

### Instructions

Write a function called `Greet` that takes a `name` string as an argument and returns a greeting string in the format `"Hello, " + name + "!"`.

### Expected function

```go
func Greet(name string) string {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Greet("World"))
	fmt.Println(piscine.Greet("Go"))
	fmt.Println(piscine.Greet("Alice"))
}
```

And its output:

```console
$ go run .
Hello, World!
Hello, Go!
Hello, Alice!
$
```
