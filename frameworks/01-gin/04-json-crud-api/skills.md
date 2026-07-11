# Skills for gin/04-json-crud-api

## What You'll Learn

**Previous:** [gin/03-middleware](../03-middleware/skills.md) | **Next:** [gin/05-auth-jwt](../05-auth-jwt/skills.md)

**Challenge:** Build a production-style CRUD API with validation, partial updates, filtering, and thread safety.

## Struct Tags for Validation

```go
type Task struct {
    Title  string `json:"title" binding:"required"`
    Status string `json:"status" binding:"omitempty,oneof=pending done"`
}
```

Gin uses the `validator` library under the hood. Common tags:

| Tag | Meaning |
|-----|---------|
| `binding:"required"` | Field must be present and non-zero |
| `binding:"omitempty"` | Skip validation if field is zero/empty |
| `binding:"min=1,max=100"` | String length or numeric range |
| `binding:"email"` | Must be a valid email format |
| `binding:"oneof=a b c"` | Must be one of the listed values |

## Partial Update (PATCH)

Use pointer fields — a pointer that is `nil` means the field was not provided:

```go
type UpdateTask struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Status      *string `json:"status"`
}

func patchHandler(c *gin.Context) {
    var update UpdateTask
    c.ShouldBindJSON(&update)

    if update.Title != nil {
        task.Title = *update.Title
    }
    if update.Status != nil {
        task.Status = *update.Status
    }
}
```

## Thread-Safe In-Memory Store

```go
type Store struct {
    mu     sync.RWMutex
    tasks  map[int]*Task
    nextID int
}

func (s *Store) Create(t *Task) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.nextID++
    t.ID = s.nextID
    s.tasks[t.ID] = t
}

func (s *Store) Get(id int) (*Task, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    t, ok := s.tasks[id]
    return t, ok
}
```

## Solving the Challenge

```go
package main

import (
    "net/http"
    "strconv"
    "sync"
    "time"
    "github.com/gin-gonic/gin"
)

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskInput struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
}

type UpdateTaskInput struct {
    Title       *string `json:"title"`
    Description *string `json:"description"`
    Status      *string `json:"status"`
}

type Store struct {
    mu     sync.RWMutex
    tasks  map[int]*Task
    nextID int
}

func NewStore() *Store {
    return &Store{tasks: make(map[int]*Task)}
}

func (s *Store) Create(input CreateTaskInput) *Task {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.nextID++
    t := &Task{
        ID:          s.nextID,
        Title:       input.Title,
        Description: input.Description,
        Status:      "pending",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    s.tasks[t.ID] = t
    return t
}

func (s *Store) Get(id int) (*Task, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    t, ok := s.tasks[id]
    return t, ok
}

func (s *Store) List(status string) []*Task {
    s.mu.RLock()
    defer s.mu.RUnlock()
    result := []*Task{}
    for _, t := range s.tasks {
        if status == "" || t.Status == status {
            result = append(result, t)
        }
    }
    return result
}

func (s *Store) Update(id int, input UpdateTaskInput) (*Task, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    t, ok := s.tasks[id]
    if !ok {
        return nil, false
    }
    if input.Title != nil { t.Title = *input.Title }
    if input.Description != nil { t.Description = *input.Description }
    if input.Status != nil { t.Status = *input.Status }
    t.UpdatedAt = time.Now()
    return t, true
}

func (s *Store) Delete(id int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, ok := s.tasks[id]; !ok {
        return false
    }
    delete(s.tasks, id)
    return true
}

func main() {
    store := NewStore()
    r := gin.Default()

    tasks := r.Group("/tasks")

    tasks.GET("/stats", func(c *gin.Context) {
        all := store.List("")
        pending, done := 0, 0
        for _, t := range all {
            if t.Status == "done" { done++ } else { pending++ }
        }
        c.JSON(http.StatusOK, gin.H{"total": len(all), "pending": pending, "done": done})
    })

    tasks.POST("", func(c *gin.Context) {
        var input CreateTaskInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, store.Create(input))
    })

    tasks.GET("", func(c *gin.Context) {
        c.JSON(http.StatusOK, store.List(c.Query("status")))
    })

    tasks.GET("/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        t, ok := store.Get(id)
        if !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusOK, t)
    })

    tasks.PATCH("/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var input UpdateTaskInput
        c.ShouldBindJSON(&input)
        t, ok := store.Update(id, input)
        if !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusOK, t)
    })

    tasks.DELETE("/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        if !store.Delete(id) {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.Status(http.StatusNoContent)
    })

    r.Run(":8080")
}
```

## Documentation

- [Gin model binding & validation](https://gin-gonic.com/docs/examples/binding-models/)
- [go-playground/validator](https://github.com/go-playground/validator)
- [REST API design best practices](https://restfulapi.net/)

**Next:** [gin/05-auth-jwt](../05-auth-jwt/README.md)
