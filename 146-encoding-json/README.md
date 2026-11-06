# 146-encoding-json — JSON Encoding and Decoding

## Task

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

---

## Time limit

**25 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 146
```

or run the tests directly:

```bash
go test ./146-encoding-json/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 146` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
