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
