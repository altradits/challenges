# Skills for fiber/04-fiber-api

## What You'll Learn

**Previous:** [fiber/03-fiber-middleware](../03-fiber-middleware/skills.md) | **Next:** [fiber/05-fiber-realtime](../05-fiber-realtime/skills.md)

**Challenge:** Build a complete Fiber API with custom error handling, validation, and pagination.

## Custom Error Handler

```go
app := fiber.New(fiber.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        code := fiber.StatusInternalServerError
        message := "Internal Server Error"

        var e *fiber.Error
        if errors.As(err, &e) {
            code = e.Code
            message = e.Message
        }

        return c.Status(code).JSON(fiber.Map{
            "error":      message,
            "status":     code,
            "request_id": c.Locals("requestid"),
        })
    },
})
```

## Returning Errors from Handlers

```go
// Fiber built-in errors
return fiber.ErrBadRequest
return fiber.ErrNotFound
return fiber.NewError(fiber.StatusConflict, "SKU already exists")

// Custom message
return fiber.NewError(422, "validation failed: name is required")
```

## Pagination

```go
type PaginationParams struct {
    Page  int `query:"page"`
    Limit int `query:"limit"`
}

func listItems(c *fiber.Ctx) error {
    var p PaginationParams
    c.QueryParser(&p)
    if p.Page < 1 { p.Page = 1 }
    if p.Limit < 1 || p.Limit > 100 { p.Limit = 20 }

    offset := (p.Page - 1) * p.Limit
    // slice items[offset : offset+limit]

    return c.JSON(fiber.Map{
        "data":  items,
        "page":  p.Page,
        "limit": p.Limit,
        "total": total,
    })
}
```

## Validation in Fiber

Fiber doesn't have built-in validation (unlike Echo), but you can integrate `go-playground/validator`:

```go
validate := validator.New()

func (store *Store) Create(input CreateItemInput) (*Item, error) {
    if err := validate.Struct(input); err != nil {
        return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
    }
    // create...
}
```

## Documentation

- [Fiber error handling](https://docs.gofiber.io/guide/error-handling)
- [Fiber config](https://docs.gofiber.io/api/fiber#config)

**Next:** [fiber/05-fiber-realtime](../05-fiber-realtime/README.md)
