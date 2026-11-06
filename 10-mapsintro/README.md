# 10. Maps Introduction

## What You'll Learn

This lesson teaches Go's built-in key-value data structure: the **map**.

## Task

Write `WordFrequency(s string) map[string]int` that returns a map showing how many times each word appears in the string.

### Expected Function

```go
func WordFrequency(s string) map[string]int {
    // Your code here
}
```

### Test Cases

| Input | Expected Output |
|-------|-----------------|
| `"go is go"` | `{"go": 2, "is": 1}` |
| `"hello"` | `{"hello": 1}` |
| `""` | `{}` (empty map) |
| `"a b a c a"` | `{"a": 3, "b": 1, "c": 1}` |

### Usage Example

```go
package main

import "fmt"

func main() {
    freq := WordFrequency("go is go")
    fmt.Println(freq["go"])  // 2
    fmt.Println(freq["is"])  // 1
}
```

## Next Steps

After completing this, you'll be ready for:
- [11-arrays](../11-arrays/README.md) - Fixed-size arrays

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 10
```

or run the tests directly:

```bash
go test ./10-mapsintro/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 10` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
