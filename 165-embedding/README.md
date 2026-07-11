# 165-embedding — Struct Embedding and Composition

## Challenge

Implement in `package piscine`:

```go
// Base logging capability
type Logger struct {
    Prefix string
}

func (l Logger) Log(msg string) string   // returns "[Prefix] msg"

// Server embeds Logger — gets Log for free
type Server struct {
    Logger
    Host string
    Port int
}

func NewServer(host string, port int) Server  // uses "SERVER" as prefix
func (s Server) Address() string              // returns "host:port"
func (s Server) Start() string                // returns "[SERVER] Starting host:port"

// Cache with a named field (not embedded)
type Cache struct {
    store map[string]string
}

func NewCache() Cache
func (c *Cache) Set(key, value string)
func (c *Cache) Get(key string) (string, bool)

// CachedServer embeds both Server and has a Cache
type CachedServer struct {
    Server
    Cache Cache
}

func NewCachedServer(host string, port int) CachedServer
func (cs *CachedServer) Store(key, value string)   // wraps Cache.Set
func (cs *CachedServer) Fetch(key string) (string, bool)  // wraps Cache.Get
```

## Examples

```
l := Logger{Prefix: "APP"}
l.Log("hello")  →  "[APP] hello"

s := NewServer("localhost", 8080)
s.Address()    →  "localhost:8080"
s.Log("boot")  →  "[SERVER] boot"   // promoted from Logger
s.Start()      →  "[SERVER] Starting localhost:8080"

cs := NewCachedServer("0.0.0.0", 9000)
cs.Store("name", "Alice")
cs.Fetch("name")  →  "Alice", true
cs.Fetch("age")   →  "", false
cs.Address()      →  "0.0.0.0:9000"  // promoted from Server
```

Read `skills.md` before you start.
