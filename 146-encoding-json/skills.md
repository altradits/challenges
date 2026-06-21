# Skills for 146-encoding-json

## What You'll Learn

**Previous:** [145-time](../145-time/skills.md) | **Next:** [147-http-basics](../147-http-basics/skills.md)

**Challenge:** Marshal Go structs to JSON and unmarshal JSON into Go structs.

## The `encoding/json` Package

```go
import "encoding/json"
```

### `json.Marshal` — Go → JSON

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

u := User{Name: "Alice", Email: "alice@example.com", Age: 30}
data, err := json.Marshal(u)
// data = {"name":"Alice","email":"alice@example.com","age":30}
```

### `json.Unmarshal` — JSON → Go

```go
data := []byte(`{"name":"Alice","email":"alice@example.com","age":30}`)
var u User
err := json.Unmarshal(data, &u)
// u.Name = "Alice"
```

### Struct Tags

Tags control how fields are serialized:

```go
type Product struct {
    ID          int     `json:"id"`              // rename field
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description,omitempty"` // omit if zero value
    internal    string  `json:"-"`               // never include
    UpdatedAt   *time.Time `json:"updated_at"`   // pointer: omit if nil
}
```

`omitempty` omits: `""`, `0`, `false`, `nil`, empty slice/map.

### Handling Unknown Fields: `map[string]any`

```go
var raw map[string]any
err := json.Unmarshal(data, &raw)
name := raw["name"].(string)
```

Use when the JSON structure is unknown at compile time.

### Streaming: `json.Encoder` / `json.Decoder`

For HTTP bodies, files, or any `io.Reader`/`io.Writer`:

```go
// Encode (write to writer)
enc := json.NewEncoder(w)        // w is io.Writer (e.g. http.ResponseWriter)
enc.SetIndent("", "  ")          // pretty print
err := enc.Encode(user)          // writes JSON + newline

// Decode (read from reader)
dec := json.NewDecoder(r.Body)   // r.Body is io.Reader
var user User
err := dec.Decode(&user)
```

Prefer Encoder/Decoder over Marshal/Unmarshal for HTTP handlers.

### Pretty Printing

```go
data, err := json.MarshalIndent(v, "", "  ")
fmt.Println(string(data))
```

### Custom Marshaling

Implement `MarshalJSON() ([]byte, error)` and `UnmarshalJSON([]byte) error`:

```go
type Date struct{ time.Time }

func (d Date) MarshalJSON() ([]byte, error) {
    return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(b []byte) error {
    var s string
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    t, err := time.Parse("2006-01-02", s)
    d.Time = t
    return err
}
```

### Solving the Challenge

```go
package piscine

import (
    "encoding/json"
    "io"
    "time"
)

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    Password  string    `json:"-"`
}

func MarshalUser(u User) ([]byte, error) {
    return json.Marshal(u)
}

func UnmarshalUser(data []byte) (User, error) {
    var u User
    err := json.Unmarshal(data, &u)
    return u, err
}

func PrettyPrint(w io.Writer, v any) error {
    enc := json.NewEncoder(w)
    enc.SetIndent("", "  ")
    return enc.Encode(v)
}
```

**Next:** [147-http-basics](../147-http-basics/README.md)
