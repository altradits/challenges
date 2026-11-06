package piscine

import "testing"

func TestTrackToolUse(t *testing.T) {
	var a Analytics
	a.TrackToolUse("StringLength", "hi", "2")
	a.TrackToolUse("StringLength", "hey", "3")
	a.TrackToolUse("ToUpper", "a", "A")
	if a.TotalUses != 3 {
		t.Errorf("TotalUses = %d, want 3", a.TotalUses)
	}
	if a.PopularTools["StringLength"] != 2 || a.PopularTools["ToUpper"] != 1 {
		t.Errorf("PopularTools = %v", a.PopularTools)
	}
	if len(a.interactions) != 3 {
		t.Errorf("interactions = %d, want 3", len(a.interactions))
	}
	if a.interactions[0].ToolName != "StringLength" || a.interactions[0].Input != "hi" {
		t.Errorf("interaction[0] = %+v", a.interactions[0])
	}
	if got := a.TopTool(); got != "StringLength" {
		t.Errorf("TopTool = %q, want StringLength", got)
	}
	var fresh Analytics
	if got := fresh.TopTool(); got != "" {
		t.Errorf("TopTool empty = %q, want empty", got)
	}
}
