# 164-constants-iota — Constants and Iota

## Challenge

Implement in `package piscine`:

```go
// Define a Direction type and its four cardinal values using iota.
// North=0, East=1, South=2, West=3
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

// TurnRight returns the direction 90° clockwise from d.
func TurnRight(d Direction) Direction

// TurnLeft returns the direction 90° counter-clockwise from d.
func TurnLeft(d Direction) Direction

// String returns the name of the direction ("North", "East", "South", "West").
func (d Direction) String() string

// Define file permission bits using iota with bit shifting.
type Permission uint

const (
    Read    Permission = 1 << iota  // 1
    Write                           // 2
    Execute                         // 4
)

// HasPermission returns true if p includes the given permission flag.
func HasPermission(p, flag Permission) bool
```

## Examples

```
TurnRight(North)  →  East
TurnRight(West)   →  North
TurnLeft(North)   →  West
North.String()    →  "North"

HasPermission(Read|Write, Read)     →  true
HasPermission(Read|Write, Execute)  →  false
```

Read `skills.md` before you start.

---

## Time limit

**20 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 164
```

or run the tests directly:

```bash
go test ./164-constants-iota/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 164` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
