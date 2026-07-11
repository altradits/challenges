# Skills for fiber/05-fiber-realtime

## What You'll Learn

**Previous:** [fiber/04-fiber-api](../04-fiber-api/skills.md) | **Next:** [frameworks/04-gorm](../../04-gorm/README.md)

**Challenge:** Build a WebSocket chat server with a hub that broadcasts to all connected clients.

## WebSocket with Fiber

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/websocket/v2"
)

app.Get("/ws", websocket.New(func(c *websocket.Conn) {
    // c is a *websocket.Conn
    for {
        mt, msg, err := c.ReadMessage()
        if err != nil {
            break  // client disconnected
        }
        c.WriteMessage(mt, msg)  // echo back
    }
}))
```

## The Hub Pattern

A hub is a long-running goroutine that owns the client registry and broadcasts:

```go
type Hub struct {
    clients    map[*websocket.Conn]bool
    broadcast  chan []byte
    register   chan *websocket.Conn
    unregister chan *websocket.Conn
    mu         sync.RWMutex
}

func (h *Hub) Run() {
    for {
        select {
        case conn := <-h.register:
            h.mu.Lock()
            h.clients[conn] = true
            h.mu.Unlock()
        case conn := <-h.unregister:
            h.mu.Lock()
            delete(h.clients, conn)
            conn.Close()
            h.mu.Unlock()
        case msg := <-h.broadcast:
            h.mu.RLock()
            for conn := range h.clients {
                conn.WriteMessage(websocket.TextMessage, msg)
            }
            h.mu.RUnlock()
        }
    }
}
```

## WebSocket Message Types

```go
websocket.TextMessage    // 1 — text (JSON usually goes here)
websocket.BinaryMessage  // 2 — binary data
websocket.CloseMessage   // 8 — connection close
websocket.PingMessage    // 9 — keep-alive ping
websocket.PongMessage    // 10 — keep-alive pong
```

## Reading and Writing JSON

```go
type Message struct {
    Type string `json:"type"`
    User string `json:"user"`
    Text string `json:"text"`
}

// Read
var msg Message
if err := c.ReadJSON(&msg); err != nil {
    break
}

// Write
c.WriteJSON(Message{Type: "message", User: "alice", Text: "hello"})
```

## Client-Side (JavaScript)

```javascript
const ws = new WebSocket("ws://localhost:8080/ws");
ws.onmessage = (e) => console.log(JSON.parse(e.data));
ws.send(JSON.stringify({ type: "message", user: "alice", text: "hi!" }));
```

## Documentation

- [gofiber/websocket](https://github.com/gofiber/websocket)
- [WebSocket protocol (RFC 6455)](https://tools.ietf.org/html/rfc6455)
- [Go nhooyr.io/websocket — alternative](https://github.com/nhooyr/websocket)

**Next:** [frameworks/04-gorm](../../04-gorm/README.md)
