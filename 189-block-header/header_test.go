package piscine

import (
	"bytes"
	"encoding/hex"
	"testing"
)

const genesisHex = "01000000" +
	"0000000000000000000000000000000000000000000000000000000000000000" +
	"3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a" +
	"29ab5f49" + "ffff001d" + "1dac2b7c"

func genesisBytes(t *testing.T) []byte {
	t.Helper()
	raw, err := hex.DecodeString(genesisHex)
	if err != nil || len(raw) != 80 {
		t.Fatalf("bad genesis vector: %v", err)
	}
	return raw
}

func TestParseGenesis(t *testing.T) {
	h, err := ParseHeader(genesisBytes(t))
	if err != nil {
		t.Fatalf("ParseHeader error: %v", err)
	}
	if h.Version != 1 {
		t.Errorf("Version = %d, want 1", h.Version)
	}
	if h.PrevBlock != [32]byte{} {
		t.Errorf("PrevBlock = %x, want zero", h.PrevBlock)
	}
	if got := hex.EncodeToString(h.MerkleRoot[:]); got != "3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a" {
		t.Errorf("MerkleRoot = %s", got)
	}
	if h.Timestamp != 1231006505 {
		t.Errorf("Timestamp = %d, want 1231006505", h.Timestamp)
	}
	if h.Bits != 0x1d00ffff {
		t.Errorf("Bits = %#x, want 0x1d00ffff", h.Bits)
	}
	if h.Nonce != 2083236893 {
		t.Errorf("Nonce = %d, want 2083236893", h.Nonce)
	}
}

func TestHeaderRoundTrip(t *testing.T) {
	raw := genesisBytes(t)
	h, err := ParseHeader(raw)
	if err != nil {
		t.Fatalf("ParseHeader error: %v", err)
	}
	if got := SerializeHeader(h); !bytes.Equal(got, raw) {
		t.Errorf("round trip mismatch:\n%x\n%x", got, raw)
	}
}

func TestHeaderHash(t *testing.T) {
	h, err := ParseHeader(genesisBytes(t))
	if err != nil {
		t.Fatalf("ParseHeader error: %v", err)
	}
	got := HeaderHash(h)
	want := "6fe28c0ab6f1b372c1a6a246ae63f74f931e8365e15a089c68d6190000000000"
	if hex.EncodeToString(got[:]) != want {
		t.Errorf("HeaderHash = %x\nwant     %s", got, want)
	}
}

func TestParseHeaderRejects(t *testing.T) {
	for _, n := range []int{0, 79, 81, 160} {
		if _, err := ParseHeader(make([]byte, n)); err == nil {
			t.Errorf("ParseHeader(%d bytes) should error", n)
		}
	}
}
