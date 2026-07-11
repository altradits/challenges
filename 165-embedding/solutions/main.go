package piscine

import "fmt"

type Logger struct {
	Prefix string
}

func (l Logger) Log(msg string) string {
	return fmt.Sprintf("[%s] %s", l.Prefix, msg)
}

type Server struct {
	Logger
	Host string
	Port int
}

func NewServer(host string, port int) Server {
	return Server{
		Logger: Logger{Prefix: "SERVER"},
		Host:   host,
		Port:   port,
	}
}

func (s Server) Address() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s Server) Start() string {
	return s.Log("Starting " + s.Address())
}

type Cache struct {
	store map[string]string
}

func NewCache() Cache {
	return Cache{store: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
	c.store[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.store[key]
	return v, ok
}

type CachedServer struct {
	Server
	Cache Cache
}

func NewCachedServer(host string, port int) CachedServer {
	return CachedServer{
		Server: NewServer(host, port),
		Cache:  NewCache(),
	}
}

func (cs *CachedServer) Store(key, value string) {
	cs.Cache.Set(key, value)
}

func (cs *CachedServer) Fetch(key string) (string, bool) {
	return cs.Cache.Get(key)
}
