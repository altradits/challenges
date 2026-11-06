# 147-http-basics — HTTP Server and Client

## Challenge

Implement in `package piscine`:

```go
// JSONHandler returns an http.HandlerFunc that responds with v as JSON.
func JSONHandler(v any) http.HandlerFunc

// Router builds an http.ServeMux with:
//   GET /health   → {"status":"ok"}
//   GET /echo     → echoes the "msg" query param as {"message":"..."}
//   POST /reverse → reads JSON body {"text":"..."}, responds {"result":"..."}
func Router() *http.ServeMux

// FetchJSON performs a GET request to url and decodes the JSON response into v.
func FetchJSON(url string, v any) error
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
go run ./cmd/drill check 147
```

or run the tests directly:

```bash
go test ./147-http-basics/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 147` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
