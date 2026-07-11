package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// TODO: register your routes here
	// GET /         → { "message": "Hello, World!" }
	// GET /ping     → { "message": "pong" }
	// GET /hello/:name → { "message": "Hello, <name>!" }

	r.Run(":8080")
}
