## brainfuck

> Esoteric programming languages (esolang) are designed to test the boundaries of computer programming design, as proofs of concept, as software art, as a hacking interface or simply as a joke. One such esoteric language is `Brainfuck`. It was created by Urban Müller in 1993. It is a minimalist language consisting of 8 simple commands. It is Turing complete, but is not intended for practical use. It exists to amuse and challenge programmers.

### Instructions

Write a `Brainfuck` interpreter program.

The source code will be given as the first parameter, and will always be valid with fewer than 4096 operations.

Your `Brainfuck` interpreter will consist of an array of 2048 bytes, all initialized to 0, with a pointer to the first byte.

Every operator consists of a single character:

- `>`: increment the pointer.
- `<`: decrement the pointer.
- `+`: increment the pointed byte.
- `-`: decrement the pointed byte.
- `.`: print the pointed byte to standard output.
- `[`: If the byte at the current pointer is 0, skip forward to the command after the matching `]`.
- `]`: If the byte at the current pointer is not 0, jump back to the command after the matching `[`.
- Any other character is treated as a comment and ignored.

### Usage

```console
$ go run . "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>." | cat -e
Hello World!$
$ go run . "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>." | cat -e
Hi$
$ go run . "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++." | cat -e
abc$
$ go run .
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
go run ./cmd/drill check 143
```

or run the tests directly:

```bash
go test ./143-brainfuck/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 143` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
