package piscine

import "testing"

func TestDirectionValues(t *testing.T) {
	if North != 0 || East != 1 || South != 2 || West != 3 {
		t.Errorf("directions = %d %d %d %d, want 0 1 2 3", North, East, South, West)
	}
}

func TestTurn(t *testing.T) {
	if got := TurnRight(North); got != East {
		t.Errorf("TurnRight(North) = %v", got)
	}
	if got := TurnRight(West); got != North {
		t.Errorf("TurnRight(West) = %v, want North (wrap)", got)
	}
	if got := TurnLeft(North); got != West {
		t.Errorf("TurnLeft(North) = %v, want West (wrap)", got)
	}
	if got := TurnLeft(East); got != North {
		t.Errorf("TurnLeft(East) = %v", got)
	}
}

func TestDirectionString(t *testing.T) {
	cases := map[Direction]string{North: "North", East: "East", South: "South", West: "West"}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("(%d).String() = %q, want %q", int(d), got, want)
		}
	}
}

func TestPermissions(t *testing.T) {
	if Read != 1 || Write != 2 || Execute != 4 {
		t.Errorf("perms = %d %d %d, want 1 2 4", Read, Write, Execute)
	}
	p := Read | Write
	if !HasPermission(p, Read) || !HasPermission(p, Write) {
		t.Error("should have Read and Write")
	}
	if HasPermission(p, Execute) {
		t.Error("should not have Execute")
	}
}
