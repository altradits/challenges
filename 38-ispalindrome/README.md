# 11. Is Palindrome

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

---

## Time limit

**10 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 38
```

or run the tests directly:

```bash
go test ./38-ispalindrome/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 38` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
