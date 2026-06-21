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
