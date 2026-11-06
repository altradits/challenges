package piscine

import (
	"encoding/hex"
	"testing"
)

const (
	helloInternal = "9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50"
	helloDisplay  = "503d8319a48348cdc610a582f7bf754b5833df65038606eb48510790dfc99595"
)

func mustDecode(t *testing.T, s string) [32]byte {
	t.Helper()
	raw, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("bad test vector: %v", err)
	}
	var out [32]byte
	copy(out[:], raw)
	return out
}

func TestDoubleSHA(t *testing.T) {
	if got := DoubleSHA([]byte("hello")); got != mustDecode(t, helloInternal) {
		t.Errorf("DoubleSHA(hello) = %x, want %s", got, helloInternal)
	}
}

func TestTxidDisplay(t *testing.T) {
	h := mustDecode(t, helloInternal)
	if got := TxidDisplay(h); got != helloDisplay {
		t.Errorf("TxidDisplay =\n%s\nwant\n%s", got, helloDisplay)
	}
	// Display must not mutate the input array.
	if h != mustDecode(t, helloInternal) {
		t.Error("TxidDisplay modified its argument")
	}
}

func TestTxidParse(t *testing.T) {
	h, err := TxidParse(helloDisplay)
	if err != nil {
		t.Fatalf("TxidParse error: %v", err)
	}
	if h != mustDecode(t, helloInternal) {
		t.Errorf("TxidParse = %x, want %s", h, helloInternal)
	}
	for _, bad := range []string{"", "zz", helloDisplay[:63], helloDisplay + "00", "zz" + helloDisplay[2:]} {
		if _, err := TxidParse(bad); err == nil {
			t.Errorf("TxidParse(%q) should error", bad)
		}
	}
}

func TestTxidRoundTrip(t *testing.T) {
	var h [32]byte
	for i := range h {
		h[i] = byte(i * 7)
	}
	back, err := TxidParse(TxidDisplay(h))
	if err != nil {
		t.Fatalf("round trip error: %v", err)
	}
	if back != h {
		t.Errorf("round trip = %x, want %x", back, h)
	}
}
