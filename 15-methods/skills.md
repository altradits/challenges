# Skills for 30-methods

## What You'll Learn

**Previous:** [14-pointers](../14-pointers/skills.md) | **Next:** [16-interfaces](../16-interfaces/skills.md)

**Challenge:** Use pointer receivers to mutate struct state; value receivers to read it.

## Value Receivers vs Pointer Receivers

A method's receiver can be either a value copy or a pointer.

### Value Receiver — reads only

```go
func (c Counter) Value() int {
    return c.count   // c is a copy — mutations here don't affect the original
}
```

### Pointer Receiver — can mutate

```go
func (c *Counter) Increment() {
    c.count++   // c is the real Counter — this change persists
}
```

### Why the Distinction Matters

```go
type Counter struct{ count int }

c := Counter{}
c.Increment()  // Go automatically takes &c when calling a pointer receiver method
               // equivalent to (&c).Increment()
fmt.Println(c.Value())  // 1
```

Go will **automatically** convert `c.Increment()` to `(&c).Increment()` when `c` is addressable (a variable, not a map value or return value).

### The Rule: Be Consistent

If *any* method on a type uses a pointer receiver, make *all* methods use pointer receivers. This ensures the method set is consistent and works with interfaces.

```go
// Good — all pointer receivers
func (c *Counter) Increment() { c.count++ }
func (c *Counter) Add(n int)   { c.count += n }
func (c *Counter) Reset()      { c.count = 0 }
func (c *Counter) Value() int  { return c.count }

// Acceptable — mixing, but then Value() is available on both T and *T
func (c Counter) Value() int  { return c.count }
```

### When to Use Each

| Pointer receiver `*T` | Value receiver `T` |
|---|---|
| Method modifies the struct | Method only reads |
| Struct is large (avoid copy) | Struct is tiny (1-2 fields) |
| Must match other pointer receivers for interface | Only reads, no interface requirement |

### Methods on Non-Struct Types

You can add methods to any named type in the same package:

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}

temp := Celsius(100)
fmt.Println(temp.ToFahrenheit())  // 212
```

### Solving the Challenge

```go
package piscine

type Counter struct{ count int }

func (c *Counter) Increment()   { c.count++ }
func (c *Counter) Add(n int)    { c.count += n }
func (c *Counter) Reset()       { c.count = 0 }
func (c Counter) Value() int    { return c.count }
```

**Next:** [16-interfaces](../16-interfaces/README.md)
