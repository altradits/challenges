package piscine

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	srv := NewServer("127.0.0.1:18080", h)
	if srv.Addr != "127.0.0.1:18080" {
		t.Errorf("Addr = %q", srv.Addr)
	}
	if srv.Handler == nil {
		t.Error("Handler is nil")
	}
}

func TestGracefulShutdownOnSignal(t *testing.T) {
	srv := NewServer("127.0.0.1:18081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	done := make(chan error, 1)
	go func() { done <- RunWithGracefulShutdown(srv) }()
	time.Sleep(300 * time.Millisecond) // let the server start
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	if err := proc.Signal(syscall.SIGTERM); err != nil {
		t.Fatal(err)
	}
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("RunWithGracefulShutdown: %v", err)
		}
	case <-time.After(12 * time.Second):
		t.Fatal("server did not shut down after SIGTERM")
	}
}
