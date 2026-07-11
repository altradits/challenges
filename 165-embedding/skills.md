# Skills for 165-embedding

## What You'll Learn

**Previous:** [164-constants-iota](../164-constants-iota/skills.md) | **Next:** [166-generics-intro](../166-generics-intro/skills.md)

**Challenge:** Use struct embedding to compose behaviour. Understand promotion, shadowing, and when to embed vs. contain.

## Struct Embedding

Go has no inheritance. Instead, you **embed** one struct inside another. The outer struct automatically gets all the exported fields and methods of the inner struct — this is called **promotion**.

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal        // embedded — no field name!
    Breed string
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Husky"}

// Promoted field and method — use as if they're on Dog directly
fmt.Println(d.Name)    // Rex     — Animal.Name promoted
fmt.Println(d.Speak()) // Rex makes a sound  — Animal.Speak promoted
fmt.Println(d.Breed)   // Husky
```

### Embedding vs Named Field

| Approach | Syntax | Access promoted fields/methods? |
|----------|--------|---------------------------------|
| Embedding | `type Dog struct { Animal }` | Yes — `d.Name`, `d.Speak()` |
| Named field | `type Dog struct { Pet Animal }` | No — must use `d.Pet.Name`, `d.Pet.Speak()` |

Use embedding when you want the inner type's interface to be part of the outer type's interface. Use a named field when you want a clear boundary.

### Method Shadowing

If the outer type defines the same method as an embedded type, the outer wins:

```go
func (d Dog) Speak() string {
    return d.Name + " barks"
}

d.Speak()          // "Rex barks"       — Dog's Speak shadows Animal's
d.Animal.Speak()   // "Rex makes a sound" — still accessible with full path
```

### Multiple Embeddings

You can embed multiple types:

```go
type Flyable struct{}
func (Flyable) Fly() string { return "flying" }

type Swimmable struct{}
func (Swimmable) Swim() string { return "swimming" }

type Duck struct {
    Animal
    Flyable
    Swimmable
}

d := Duck{Animal: Animal{Name: "Donald"}}
d.Fly()    // "flying"
d.Swim()   // "swimming"
d.Speak()  // "Donald makes a sound"
```

### Embedding Interfaces

You can embed an interface inside a struct. This makes the struct satisfy that interface, but you must populate the field yourself or the embedded methods will panic:

```go
type Handler struct {
    http.Handler  // embed the interface
}
```

This pattern appears in the standard library (e.g., `httptest.ResponseRecorder`).

## Pointer vs Value Embedding

```go
type Base struct{ ID int }
func (b *Base) SetID(id int) { b.ID = id }  // pointer receiver

type Extended struct {
    *Base   // pointer embedding
}

e := Extended{Base: &Base{}}
e.SetID(42)      // works — pointer methods are promoted
fmt.Println(e.ID) // 42
```

Embed `*T` when the embedded type has pointer-receiver methods you want to promote.

## Solving the Challenge

```go
package piscine

import "fmt"

type Logger struct {
    Prefix string
}

func (l Logger) Log(msg string) string {
    return fmt.Sprintf("[%s] %s", l.Prefix, msg)
}

type Server struct {
    Logger
    Host string
    Port int
}

func NewServer(host string, port int) Server {
    return Server{
        Logger: Logger{Prefix: "SERVER"},
        Host:   host,
        Port:   port,
    }
}

func (s Server) Address() string {
    return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s Server) Start() string {
    return s.Log("Starting " + s.Address())
}

type Cache struct {
    store map[string]string
}

func NewCache() Cache {
    return Cache{store: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
    c.store[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
    v, ok := c.store[key]
    return v, ok
}

type CachedServer struct {
    Server
    Cache Cache
}

func NewCachedServer(host string, port int) CachedServer {
    return CachedServer{
        Server: NewServer(host, port),
        Cache:  NewCache(),
    }
}

func (cs *CachedServer) Store(key, value string) {
    cs.Cache.Set(key, value)
}

func (cs *CachedServer) Fetch(key string) (string, bool) {
    return cs.Cache.Get(key)
}
```

## Documentation

- [Effective Go — Embedding](https://go.dev/doc/effective_go#embedding)
- [A Tour of Go — Type embedding](https://go.dev/tour/methods/10)
- [Go Spec — Struct types](https://go.dev/ref/spec#Struct_types)

**Next:** [166-generics-intro](../166-generics-intro/README.md)
