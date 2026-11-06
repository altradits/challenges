package piscine

import (
	"reflect"
	"testing"
)

func TestConversions(t *testing.T) {
	if got := IntToFloat(5); got != 5.0 {
		t.Errorf("IntToFloat(5) = %v", got)
	}
	if got := FloatToInt(3.99); got != 3 {
		t.Errorf("FloatToInt(3.99) = %d, want 3 (truncate)", got)
	}
	if got := FloatToInt(-3.99); got != -3 {
		t.Errorf("FloatToInt(-3.99) = %d, want -3 (toward zero)", got)
	}
	if got := RuneToString('é'); got != "é" {
		t.Errorf("RuneToString = %q", got)
	}
	if got := StringToRunes("héllo"); !reflect.DeepEqual(got, []rune("héllo")) {
		t.Errorf("StringToRunes = %q", got)
	}
	if got := BytesToString([]byte{'a', 'b'}); got != "ab" {
		t.Errorf("BytesToString = %q", got)
	}
	if got := CountRunes("héllo"); got != 5 {
		t.Errorf("CountRunes = %d, want 5 (runes, not bytes)", got)
	}
	if got := CountRunes(""); got != 0 {
		t.Errorf("CountRunes empty = %d", got)
	}
}
