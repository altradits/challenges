# Skills for fiber/03-fiber-middleware

## What You'll Learn

**Previous:** [fiber/02-fiber-routing](../02-fiber-routing/skills.md) | **Next:** [fiber/04-fiber-api](../04-fiber-api/skills.md)

**Challenge:** Stack Fiber built-in middleware and write a custom one.

## Fiber Built-In Middleware

All middleware is in `github.com/gofiber/fiber/v2/middleware/*`:

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/gofiber/fiber/v2/middleware/requestid"
    "github.com/gofiber/fiber/v2/middleware/compress"
)

app := fiber.New()
app.Use(logger.New())
app.Use(recover.New())
app.Use(cors.New())
app.Use(requestid.New())
app.Use(compress.New())
```

## Custom Middleware in Fiber

```go
app.Use(func(c *fiber.Ctx) error {
    c.Set("X-API-Version", "v1")
    return c.Next()  // call the next middleware/handler
})
```

Or as a named function:

```go
func APIVersion(version string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        c.Set("X-API-Version", version)
        return c.Next()
    }
}

app.Use(APIVersion("v1"))
```

## Middleware Order

Fiber processes middleware in registration order. Order matters:

```go
app.Use(recover.New())   // must be first to catch panics from later middleware
app.Use(logger.New())    // logs after recover is in place
app.Use(cors.New())
// ... routes
```

## Accessing Request ID in Handlers

```go
func myHandler(c *fiber.Ctx) error {
    id := c.Locals("requestid").(string)
    return c.JSON(fiber.Map{"request_id": id})
}
```

`c.Locals` is like a request-scoped context.

## Documentation

- [Fiber middleware](https://docs.gofiber.io/category/-middleware)
- [fiber/middleware/logger](https://docs.gofiber.io/api/middleware/logger)
- [fiber/middleware/cors](https://docs.gofiber.io/api/middleware/cors)

**Next:** [fiber/04-fiber-api](../04-fiber-api/README.md)
