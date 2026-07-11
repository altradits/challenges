# fiber/04-fiber-api — Full Fiber REST API with Error Handling

## Challenge

Build an Inventory Management API with proper error handling:

```
POST   /api/items        → create item
GET    /api/items        → list items (pagination: ?page=1&limit=20)
GET    /api/items/:id    → get item
PUT    /api/items/:id    → full update
PATCH  /api/items/:id    → partial update
DELETE /api/items/:id    → delete item

GET    /api/items/low-stock   → items where quantity < 10
POST   /api/items/:id/restock → add quantity: { "quantity": 50 }
```

## Requirements

- Custom global error handler that returns consistent JSON errors
- Validation using `go-playground/validator`
- Pagination for list endpoint
- Thread-safe in-memory storage

## Item Model

```go
type Item struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"     validate:"required,min=1"`
    SKU      string  `json:"sku"      validate:"required"`
    Price    float64 `json:"price"    validate:"min=0"`
    Quantity int     `json:"quantity" validate:"min=0"`
}
```

Read `skills.md` before you start.
