package piscine

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func dsha(t *testing.T, s string) [32]byte {
	t.Helper()
	first := sha256.Sum256([]byte(s))
	return sha256.Sum256(first[:])
}

func mustHex32(t *testing.T, s string) [32]byte {
	t.Helper()
	raw, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("bad test vector: %v", err)
	}
	var out [32]byte
	copy(out[:], raw)
	return out
}

func TestMerkleRootEdges(t *testing.T) {
	if got := MerkleRoot(nil); got != [32]byte{} {
		t.Errorf("MerkleRoot(nil) = %x, want zero", got)
	}
	a := dsha(t, "tx0")
	if got := MerkleRoot([][32]byte{a}); got != a {
		t.Errorf("MerkleRoot([a]) = %x, want a itself", got)
	}
}

func TestMerkleRootPairs(t *testing.T) {
	a, b, c := dsha(t, "tx0"), dsha(t, "tx1"), dsha(t, "tx2")
	got := MerkleRoot([][32]byte{a, b})
	want := mustHex32(t, "f7ed2a56db7927143defd32e380fb083f74970e0d6f31c00b8aded86aefdf1d2")
	if got != want {
		t.Errorf("root2 = %x\nwant    %x", got, want)
	}
	got = MerkleRoot([][32]byte{a, b, c})
	want = mustHex32(t, "e723898188fe36f157a1a97e4a608823a4c64cc81ac658dba49c4b66eef59a1f")
	if got != want {
		t.Errorf("root3 = %x\nwant    %x", got, want)
	}
	// Order matters: swapping leaves changes the root.
	if swapped := MerkleRoot([][32]byte{b, a}); swapped == MerkleRoot([][32]byte{a, b}) {
		t.Error("swapped leaves gave the same root")
	}
}
