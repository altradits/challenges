# echo/02-echo-routing — Routing and Groups

## Challenge

Build a Blog REST API with Echo:

```
POST   /api/v1/posts          → create post { "title", "body", "author_id" }
GET    /api/v1/posts          → list posts
GET    /api/v1/posts/:id      → get post
PUT    /api/v1/posts/:id      → update post
DELETE /api/v1/posts/:id      → delete post

POST   /api/v1/posts/:id/comments → add comment { "text", "author" }
GET    /api/v1/posts/:id/comments → list comments for post
```

Use sub-routing to nest comments under posts.

Read `skills.md` before you start.
