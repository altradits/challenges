# fiber/05-fiber-realtime — WebSocket Chat Server

## Challenge

Build a real-time chat server using Fiber's WebSocket support:

```
GET  /               → serve chat HTML page (inline HTML in Go string)
GET  /ws             → WebSocket upgrade — connects a client to the chat room
POST /broadcast      → send a message to all connected clients (HTTP endpoint for testing)
GET  /connections    → return number of active connections
```

## Message Format

```json
{ "type": "message", "user": "alice", "text": "hello everyone!" }
{ "type": "join",    "user": "alice" }
{ "type": "leave",   "user": "alice" }
```

## Install

```bash
go get github.com/gofiber/websocket/v2
```

## Requirements

- Hub pattern: a central goroutine manages all connections
- Broadcast to all connected clients when a message arrives
- Send `join` message when a client connects (ask for username via first message)
- Send `leave` message when a client disconnects
- Thread-safe client registry

Read `skills.md` before you start.
