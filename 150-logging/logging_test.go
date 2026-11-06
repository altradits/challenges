package piscine

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "error")
	logger.Info("invisible")
	if buf.Len() != 0 {
		t.Errorf("info at error level should log nothing, got %q", buf.String())
	}
	logger.Error("visible")
	if !strings.Contains(buf.String(), "visible") {
		t.Errorf("error at error level should log, got %q", buf.String())
	}
	if !strings.Contains(buf.String(), `"level":"ERROR"`) {
		t.Errorf("logger should write JSON with level, got %q", buf.String())
	}
	var buf2 bytes.Buffer
	NewLogger(&buf2, "debug").Debug("d")
	if !strings.Contains(buf2.String(), "d") {
		t.Errorf("debug at debug level should log, got %q", buf2.String())
	}
	var buf3 bytes.Buffer
	NewLogger(&buf3, "bogus").Info("default-info")
	if !strings.Contains(buf3.String(), "default-info") {
		t.Errorf("unknown level should default to info, got %q", buf3.String())
	}
}

func TestRequestLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "info")
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	h := RequestLogger(logger, next)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/things?a=b", nil))
	if rec.Body.String() != "ok" {
		t.Errorf("wrapped handler body = %q, want ok", rec.Body.String())
	}
	logged := buf.String()
	for _, want := range []string{"GET", "/things", "duration"} {
		if !strings.Contains(logged, want) {
			t.Errorf("log line %q missing %q", logged, want)
		}
	}
}
