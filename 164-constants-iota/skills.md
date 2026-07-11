# Skills for 164-constants-iota

## What You'll Learn

**Previous:** [163-type-conversions](../163-type-conversions/skills.md) | **Next:** [165-embedding](../165-embedding/skills.md)

**Challenge:** Define typed constants with `iota` for enumerations and bit flags.

## Constants — `const`

A constant is a value fixed at **compile time**. The compiler rejects any attempt to assign to it:

```go
const Pi      = 3.14159
const AppName = "GoSchool"
const MaxRetries = 3

Pi = 4  // compile error: cannot assign to Pi
```

### Typed vs Untyped Constants

```go
// Untyped — adapts to context
const Hundred = 100
var f float64 = Hundred   // works; 100 is used as float64
var n int     = Hundred   // works; 100 is used as int

// Typed — fixed type
const TypedHundred int = 100
var f float64 = TypedHundred  // compile error: cannot use int as float64
```

Untyped constants are preferred unless you need to prevent accidental cross-type use.

## `iota` — Auto-Incrementing Constants

`iota` starts at `0` inside a `const` block and increments by `1` for each constant in the block:

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
    D         // 3
)
```

`iota` resets to 0 at the start of each new `const` block.

### Typed Enumerations with `iota`

```go
type Direction int

const (
    North Direction = iota  // 0
    East                    // 1
    South                   // 2
    West                    // 3
)
```

Now `North`, `East`, `South`, `West` are all of type `Direction`, not plain `int`. This prevents mixing them with unrelated integers.

### Bit Flags with `iota`

Combine `iota` with bit shifting to create flag constants:

```go
type Permission uint

const (
    Read    Permission = 1 << iota  // 1  (binary: 001)
    Write                           // 2  (binary: 010)
    Execute                         // 4  (binary: 100)
)
```

`1 << iota` shifts 1 left by `iota` positions, giving unique powers of two.

Combine flags with `|` (bitwise OR):

```go
rwx := Read | Write | Execute  // 7 (binary: 111)
rw  := Read | Write            // 3 (binary: 011)
```

Check a flag with `&` (bitwise AND):

```go
func HasPermission(p, flag Permission) bool {
    return p&flag != 0
}

HasPermission(rw, Read)     // true
HasPermission(rw, Execute)  // false
```

### Skipping Values with Blank Identifier

```go
const (
    _   = iota  // skip 0 (use blank identifier)
    KB  = 1 << (10 * iota)  // 1024
    MB                       // 1048576
    GB                       // 1073741824
)
```

### The `String()` Method — `fmt.Stringer` Interface

If your type implements `String() string`, `fmt.Println` and `fmt.Sprintf("%v", ...)` will use it automatically:

```go
func (d Direction) String() string {
    switch d {
    case North:
        return "North"
    case East:
        return "East"
    case South:
        return "South"
    case West:
        return "West"
    default:
        return "Unknown"
    }
}

fmt.Println(North)  // "North" instead of "0"
```

This is the `fmt.Stringer` interface:

```go
type Stringer interface {
    String() string
}
```

Any type that has `String() string` satisfies it and gets pretty-printing for free.

## Solving the Challenge

```go
package piscine

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func TurnRight(d Direction) Direction {
    return (d + 1) % 4
}

func TurnLeft(d Direction) Direction {
    return (d + 3) % 4  // +3 mod 4 == -1 mod 4
}

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

type Permission uint

const (
    Read    Permission = 1 << iota
    Write
    Execute
)

func HasPermission(p, flag Permission) bool {
    return p&flag != 0
}
```

The `[...]string{...}[d]` trick creates a compile-time array and indexes into it — clean and fast.

## Documentation

- [A Tour of Go — Constants](https://go.dev/tour/basics/15)
- [A Tour of Go — Numeric Constants](https://go.dev/tour/basics/16)
- [Go Blog — Constants](https://go.dev/blog/constants)
- [fmt.Stringer](https://pkg.go.dev/fmt#Stringer)

**Next:** [165-embedding](../165-embedding/README.md)
