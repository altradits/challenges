package piscine

import (
	"bytes"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	cases := []struct {
		n    uint64
		want []byte
	}{
		{0, []byte{0x00}},
		{252, []byte{0xfc}},
		{253, []byte{0xfd, 0xfd, 0x00}},
		{0xffff, []byte{0xfd, 0xff, 0xff}},
		{0x10000, []byte{0xfe, 0x00, 0x00, 0x01, 0x00}},
		{0xffffffff, []byte{0xfe, 0xff, 0xff, 0xff, 0xff}},
		{0x100000000, []byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}},
		{0xffffffffffffffff, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	for _, tc := range cases {
		if got := EncodeVarint(tc.n); !bytes.Equal(got, tc.want) {
			t.Errorf("EncodeVarint(%d) = %x, want %x", tc.n, got, tc.want)
		}
	}
}

func TestDecodeVarint(t *testing.T) {
	cases := []struct {
		input []byte
		value uint64
		used  int
	}{
		{[]byte{0x00}, 0, 1},
		{[]byte{0xfc}, 252, 1},
		{[]byte{0xfd, 0xfd, 0x00}, 253, 3},
		{[]byte{0xfd, 0xff, 0xff}, 0xffff, 3},
		{[]byte{0xfe, 0x00, 0x00, 0x01, 0x00}, 0x10000, 5},
		{[]byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}, 0x100000000, 9},
	}
	for _, tc := range cases {
		v, n, err := DecodeVarint(tc.input)
		if err != nil {
			t.Errorf("DecodeVarint(%x) error: %v", tc.input, err)
			continue
		}
		if v != tc.value || n != tc.used {
			t.Errorf("DecodeVarint(%x) = (%d, %d), want (%d, %d)",
				tc.input, v, n, tc.value, tc.used)
		}
	}
}

func TestDecodeVarintRejects(t *testing.T) {
	bad := [][]byte{
		nil,
		{},
		{0xfd},
		{0xfd, 0x01},
		{0xfe, 0x01, 0x02},
		{0xff, 0x01, 0x02, 0x03},
		{0xfd, 0x00, 0x00},             // 0 must be 1 byte
		{0xfd, 0xfc, 0x00},             // 252 must be 1 byte
		{0xfe, 0xff, 0xff, 0x00, 0x00}, // 0xffff must be 0xfd form
		{0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0}, // 0xffffffff must be 0xfe form
	}
	for _, b := range bad {
		if v, n, err := DecodeVarint(b); err == nil {
			t.Errorf("DecodeVarint(%x) = (%d, %d), want error", b, v, n)
		}
	}
}

func TestVarintRoundTrip(t *testing.T) {
	for _, n := range []uint64{0, 1, 252, 253, 254, 0xffff, 0x10000,
		0xffffffff, 0x100000000, 1 << 63, 0xffffffffffffffff} {
		v, used, err := DecodeVarint(EncodeVarint(n))
		if err != nil {
			t.Errorf("round trip %d error: %v", n, err)
		} else if v != n || used != len(EncodeVarint(n)) {
			t.Errorf("round trip %d = (%d, %d)", n, v, used)
		}
	}
}
