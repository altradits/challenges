## grouping

### Instructions

Write a program that receives two strings and replicates the use of brackets in regular expressions. Brackets in regular expressions returns the words that contain the expression inside of it.

The program should handle the "`|`" operator, which searches for both strings on each side of the operator.

The output of the program should be, the results of the regular expression, numbered and displayed by the order of appearance in the string.

If the number of arguments is different from 2, if the regular expression is not valid, if the last argument is empty or if there are no matches, the program should print nothing.

### Usage

```console
$ go run . "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
1: heavy
2: steady
3: heavy
$ go run . "(e|n)" "I currently have 4 windows opened up… and I don't know why."
1: currently
2: currently
3: have
4: windows
5: opened
6: opened
7: and
8: don't
9: know
$ go run . "(hi)" "He swore he just saw his sushi move."
1: his
2: sushi
$ go run . "(s)" ""
$ go run . "i" "Something in the air"
$
```

---

## Time limit

**30 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 144
```

or run the tests directly:

```bash
go test ./144-grouping/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 144` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
