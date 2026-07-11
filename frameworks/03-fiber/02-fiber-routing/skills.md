# Skills for fiber/02-fiber-routing

## What You'll Learn

**Previous:** [fiber/01-hello-fiber](../01-hello-fiber/skills.md) | **Next:** [fiber/03-fiber-middleware](../03-fiber-middleware/skills.md)

**Challenge:** Use Fiber's route groups and body parser to build a product API.

## Route Groups

```go
api := app.Group("/api")
products := api.Group("/products")

products.Get("/", listProducts)
products.Post("/", createProduct)
products.Get("/:id", getProduct)
```

Note: Fiber uses `Get` (capitalised, not `GET`) — same style as Express.

## Body Parsing

```go
type CreateProductInput struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Category string  `json:"category"`
}

func createProduct(c *fiber.Ctx) error {
    var input CreateProductInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    // input is populated
}
```

## Query Params with Type Coercion

```go
minPrice := c.QueryFloat("min", 0)   // float64 with default 0
maxPrice := c.QueryFloat("max", 1e9)
q := c.Query("q")
category := c.Query("category")
```

## Params as Specific Types

```go
id, err := c.ParamsInt("id")
// or
id := c.ParamsInt("id", 0)  // with default
```

## Documentation

- [Fiber routing](https://docs.gofiber.io/guide/routing)
- [BodyParser](https://docs.gofiber.io/api/ctx#bodyparser)

**Next:** [fiber/03-fiber-middleware](../03-fiber-middleware/README.md)
