# 06-functions — Your First Function

## Task

Write a function called `Greet` in `package piscine` that takes a `name`
string and returns the greeting `"Hello, " + name + "!"`.

```go
func Greet(name string) string
```

## Input / Output

- Input: one string (`name`).
- Output: the greeting string `"Hello, <name>!"`.

## Examples

```go
piscine.Greet("World")  // "Hello, World!"
piscine.Greet("Go")     // "Hello, Go!"
piscine.Greet("Alice")  // "Hello, Alice!"
piscine.Greet("")       // "Hello, !"
```

## Constraints

- The solution file starts with `package piscine` (not `package main`).
- The function name is exported (`Greet`, capital G).
- Exact format: `Hello, `, then the name, then `!`. No trailing newline —
  the caller decides how to print.

## Time limit

**15 minutes** (round 1). Beat it on every repeat: aim for under 5 minutes by
round 4.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference
solution, turn off autocomplete if you can. Write every line from memory —
that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 06
```

or run the tests directly:

```bash
go test ./06-functions/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 06` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
