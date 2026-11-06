## revwstr

### Instructions

Write a program that takes a `string` as a parameter, and prints its words in reverse, followed by a newline.

- A word is a sequence of **alphanumerical** characters.

- If the number of arguments is different from 1, the program will display nothing.

- In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the `string` and that words will always be separated by exactly one space).

### Usage

```console
$ go run . "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
$ go run . "abcdefghijklm"
abcdefghijklm
$ go run . "he stared at the mountain"
mountain the at stared he
$ go run . "" | cat-e
$
$
```

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 134
```

or run the tests directly:

```bash
go test ./134-revwstr/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 134` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
