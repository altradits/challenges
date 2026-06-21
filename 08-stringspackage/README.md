## stringspackage

### Instructions

Write a function that takes a string `s` and returns three values:

- `words` — the number of words in the string (words are separated by whitespace)
- `chars` — the number of non-space characters in the string
- `hasHello` — `true` if the string contains the word `"hello"` (case-insensitive)

### Expected function

```go
func SentenceStats(s string) (words int, chars int, hasHello bool) {

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
	w, c, h := SentenceStats("Hello world")
	fmt.Println(w, c, h)

	w, c, h = SentenceStats("say HELLO to everyone")
	fmt.Println(w, c, h)

	w, c, h = SentenceStats("no greeting here")
	fmt.Println(w, c, h)

	w, c, h = SentenceStats("")
	fmt.Println(w, c, h)
}
```

And its output:

```console
$ go run .
2 10 true
4 18 true
3 14 false
0 0 false
$
```

### Test Cases

| Input | `words` | `chars` | `hasHello` |
|-------|---------|---------|------------|
| `"Hello world"` | 2 | 10 | `true` |
| `"say HELLO to everyone"` | 4 | 18 | `true` |
| `"no greeting here"` | 3 | 14 | `false` |
| `""` | 0 | 0 | `false` |
