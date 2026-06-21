# 146-encoding-json — JSON Encoding and Decoding

## Challenge

Implement in `package piscine`:

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    Password  string    `json:"-"`           // never serialized
}

// MarshalUser serializes a User to a JSON byte slice.
func MarshalUser(u User) ([]byte, error)

// UnmarshalUser deserializes JSON bytes into a User.
func UnmarshalUser(data []byte) (User, error)

// PrettyPrint encodes v as indented JSON and writes it to w.
func PrettyPrint(w io.Writer, v any) error
```

Read `skills.md` before you start.
