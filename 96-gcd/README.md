# 96-gcd — Greatest Common Divisor

## Task

Write a function called `Gcd` in `package piscine` that takes two `uint`
values and returns their greatest common divisor: the largest positive
integer that divides both. (See the [definition](https://en.wikipedia.org/wiki/Greatest_common_divisor).)

```go
func Gcd(a, b uint) uint
```

## Input / Output

- Input: two `uint` values, expected to be strictly positive.
- Output: their GCD as a `uint`; `0` if either input is `0`.

## Examples

```go
piscine.Gcd(42, 10)  // 2
piscine.Gcd(42, 12)  // 6
piscine.Gcd(14, 77)  // 7
piscine.Gcd(17, 3)   // 1
piscine.Gcd(0, 5)    // 0
piscine.Gcd(6, 6)    // 6
```

## Constraints

- Either input `0` → return `0` (do not divide by zero, do not loop forever).
- Equal inputs → return that input.
- Coprime inputs → return `1`.
- Must terminate fast for large values (a modulo-based Euclidean loop, not
  repeated subtraction all the way down).

## Time limit

**20 minutes** (round 1). Beat it on every repeat: aim for under 8 minutes by
round 4.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference
solution, turn off autocomplete if you can. Write every line from memory —
that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 96
```

or run the tests directly:

```bash
go test ./96-gcd/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 96` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
