# 08-stringspackage

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

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 8
```

or run the tests directly:

```bash
go test ./08-stringspackage/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 8` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
