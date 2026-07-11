# gorm/06-gorm-api — Full Gin + GORM REST API

## Challenge

Build a complete Blog API that combines Gin (from 01-gin) with GORM (from 04-gorm):

```
POST   /auth/register
POST   /auth/login

GET    /posts               → list posts (paginated)
POST   /posts               → create post (auth required)
GET    /posts/:id           → get post with author and comments
PUT    /posts/:id           → update post (auth: must be author)
DELETE /posts/:id           → delete post (auth: must be author)

POST   /posts/:id/comments  → add comment (auth required)
DELETE /posts/:id/comments/:commentId → delete comment (auth: must be comment author)

GET    /tags                → list all tags
GET    /tags/:name/posts    → posts with this tag
```

## Requirements

- Use the **Repository pattern**: separate DB logic from handlers
- Use **GORM transactions** for creating posts with tags
- Use **GORM preloads** for loading related data
- JWT authentication (from gin/05-auth-jwt)
- SQLite database (or PostgreSQL if you have it running)
- Proper error responses for 400, 401, 403, 404

This is the capstone challenge for the GORM + Gin track.

Read `skills.md` before you start.
