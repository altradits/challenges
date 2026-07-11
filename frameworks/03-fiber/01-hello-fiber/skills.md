# Skills for fiber/01-hello-fiber

## What You'll Learn

**Previous:** [echo/05-echo-auth](../../02-echo/05-echo-auth/skills.md) | **Next:** [fiber/02-fiber-routing](../02-fiber-routing/skills.md)

**Challenge:** Build your first Fiber app. Notice the Express.js-inspired API.

## Fiber API: Express.js in Go

If you know Express.js, Fiber will feel immediately familiar:

```javascript
// Express.js
app.get('/', (req, res) => {
    res.json({ message: 'hello' });
});
```

```go
// Fiber
app.Get("/", func(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "hello"})
})
```

## Handler Pattern

```go
func(c *fiber.Ctx) error {
    // return a response — returning nil means "OK, already responded"
    return c.JSON(fiber.Map{"key": "value"})
}
```

`fiber.Map` is `map[string]interface{}` — like Gin's `gin.H`.

## Route Registration

```go
app.Get("/path", handler)
app.Post("/path", handler)
app.Put("/path", handler)
app.Delete("/path", handler)
app.Patch("/path", handler)
```

## Path Parameters

```go
app.Get("/hello/:name", func(c *fiber.Ctx) error {
    name := c.Params("name")
    return c.JSON(fiber.Map{"message": "Hello, " + name + "!"})
})
```

## Query Parameters

```go
app.Get("/search", func(c *fiber.Ctx) error {
    q := c.Query("q")
    page := c.QueryInt("page", 1)  // with default
    return c.JSON(fiber.Map{"q": q, "page": page})
})
```

## Sending Responses

```go
c.JSON(obj)                          // JSON response
c.SendString("plain text")           // text
c.Status(fiber.StatusCreated).JSON(obj) // with status code
c.SendStatus(fiber.StatusNoContent)  // status only
```

## Solving the Challenge

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello from Fiber!"})
    })

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "pong"})
    })

    app.Get("/hello/:name", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello, " + c.Params("name") + "!"})
    })

    app.Get("/echo", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"echo": c.Query("text")})
    })

    app.Listen(":8080")
}
```

## Fiber vs Gin vs Echo

| Feature | Fiber | Gin | Echo |
|---------|-------|-----|------|
| HTTP engine | fasthttp | net/http | net/http |
| stdlib compat | No | Yes | Yes |
| API style | Express.js | Custom | Custom |
| Zero alloc | Yes | Partial | Partial |

## Documentation

- [Fiber Quickstart](https://docs.gofiber.io/)
- [fiber.Ctx](https://pkg.go.dev/github.com/gofiber/fiber/v2#Ctx)
- [Fiber routing](https://docs.gofiber.io/guide/routing)

**Next:** [fiber/02-fiber-routing](../02-fiber-routing/README.md)
