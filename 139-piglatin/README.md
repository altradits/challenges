## piglatin

### Instructions

Write a **program** that transforms a string passed as argument in its `Pig Latin` version.

The rules used by Pig Latin are as follows:

- If a word begins with a vowel, just add "ay" to the end.
- If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
- Only the latin vowels will be considered as vowel(aeiou).

If the word has no vowels, the program should print "No vowels".

If the number of arguments is different from one, the program prints nothing.

### Usage

```console
$ go run .
$ go run . pig | cat -e
igpay$
$ go run . Is | cat -e
Isay$
$ go run . crunch | cat -e
unchcray$
$ go run . crnch | cat -e
No vowels$
$ go run . something else | cat -e
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
go run ./cmd/drill check 139
```

or run the tests directly:

```bash
go test ./139-piglatin/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 139` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
