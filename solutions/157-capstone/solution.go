package piscine

import "time"

type Tool struct {
	Name        string
	Category    string
	UsageCount  int
	LastUsed    time.Time
	Description string
}

type Interaction struct {
	UserID    string
	ToolName  string
	Input     string
	Output    string
	Timestamp time.Time
	Duration  time.Duration
}

type Analytics struct {
	TotalUsers    int
	TotalUses     int
	PopularTools  map[string]int
	FeatureMatrix map[string][]string

	interactions []Interaction
}

// TrackToolUse records one tool invocation: appends the interaction,
// bumps TotalUses, and counts the tool in PopularTools.
func (a *Analytics) TrackToolUse(toolName, input, output string) {
	if a.PopularTools == nil {
		a.PopularTools = make(map[string]int)
	}
	a.interactions = append(a.interactions, Interaction{
		ToolName:  toolName,
		Input:     input,
		Output:    output,
		Timestamp: time.Now(),
	})
	a.TotalUses++
	a.PopularTools[toolName]++
}

// TopTool returns the most-used tool name, or "" when nothing is tracked.
func (a *Analytics) TopTool() string {
	best := ""
	for name, n := range a.PopularTools {
		if best == "" || n > a.PopularTools[best] {
			best = name
		}
	}
	return best
}
