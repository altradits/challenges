# fiber/02-fiber-routing — Fiber Routing and Body Parsing

## Challenge

Build a Product API with Fiber:

```
POST   /api/products          → create product
GET    /api/products          → list all (with optional ?category=electronics)
GET    /api/products/:id      → get one product
PUT    /api/products/:id      → update product
DELETE /api/products/:id      → delete product

GET    /api/products/search?q=laptop&min=100&max=2000 → filtered search
```

## Product Model

```go
type Product struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Category string  `json:"category"`
    InStock  bool    `json:"in_stock"`
}
```

Read `skills.md` before you start.
