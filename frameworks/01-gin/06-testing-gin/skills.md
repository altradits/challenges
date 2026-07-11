# Skills for gin/06-testing-gin

## What You'll Learn

**Previous:** [gin/05-auth-jwt](../05-auth-jwt/skills.md) | **Next:** [gin/07-production-gin](../07-production-gin/skills.md)

**Challenge:** Use `httptest` to test Gin handlers without starting a real server.

## `net/http/httptest`

The `httptest` package lets you call HTTP handlers directly without a network:

```go
import "net/http/httptest"

// Create a fake request
req := httptest.NewRequest(http.MethodPost, "/tasks",
    strings.NewReader(`{"title":"Buy groceries"}`))
req.Header.Set("Content-Type", "application/json")

// Create a fake response recorder
w := httptest.NewRecorder()

// Call the handler (or router)
router.ServeHTTP(w, req)

// Inspect the result
w.Code            // HTTP status code (int)
w.Body.String()   // response body (string)
w.Header()        // response headers
```

## Test Structure

```go
func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)  // suppress debug output
    store := NewStore()
    return buildRouter(store)  // extract your router setup into a function
}

func TestCreateTask(t *testing.T) {
    r := setupRouter()

    tests := []struct {
        name     string
        body     string
        wantCode int
    }{
        {
            name:     "valid task",
            body:     `{"title":"Buy groceries"}`,
            wantCode: http.StatusCreated,
        },
        {
            name:     "missing title",
            body:     `{"description":"no title"}`,
            wantCode: http.StatusBadRequest,
        },
        {
            name:     "invalid json",
            body:     `{bad json`,
            wantCode: http.StatusBadRequest,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            req := httptest.NewRequest(http.MethodPost, "/tasks",
                strings.NewReader(tc.body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)
            if w.Code != tc.wantCode {
                t.Errorf("got %d, want %d\nbody: %s",
                    w.Code, tc.wantCode, w.Body.String())
            }
        })
    }
}
```

## Decode JSON Response

```go
var task Task
if err := json.NewDecoder(w.Body).Decode(&task); err != nil {
    t.Fatalf("failed to decode response: %v", err)
}
if task.Title != "Buy groceries" {
    t.Errorf("got title %q, want %q", task.Title, "Buy groceries")
}
```

## Testing Sequences (Create Then Get)

```go
func TestGetTask(t *testing.T) {
    r := setupRouter()

    // Step 1: create
    createReq := httptest.NewRequest(http.MethodPost, "/tasks",
        strings.NewReader(`{"title":"Test"}`))
    createReq.Header.Set("Content-Type", "application/json")
    createW := httptest.NewRecorder()
    r.ServeHTTP(createW, createReq)

    var created Task
    json.NewDecoder(createW.Body).Decode(&created)

    // Step 2: get
    getReq := httptest.NewRequest(http.MethodGet,
        fmt.Sprintf("/tasks/%d", created.ID), nil)
    getW := httptest.NewRecorder()
    r.ServeHTTP(getW, getReq)

    if getW.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", getW.Code)
    }
}
```

## Gin Test Mode

```go
gin.SetMode(gin.TestMode)
```

Suppresses the debug/route print output during tests.

## Documentation

- [net/http/httptest](https://pkg.go.dev/net/http/httptest)
- [Gin test examples](https://github.com/gin-gonic/gin/blob/master/gin_test.go)
- [169-testing-advanced](../../../169-testing-advanced/README.md) — table-driven tests

**Next:** [gin/07-production-gin](../07-production-gin/README.md)
