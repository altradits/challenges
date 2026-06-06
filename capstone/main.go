package main

import (
	"fmt"
	"time"
)

// Tool represents a string manipulation tool
type Tool struct {
	Name        string
	Category    string
	UsageCount  int
	LastUsed    time.Time
	Description string
}

// Interaction represents a user interaction with a tool
type Interaction struct {
	UserID    string
	ToolName  string
	Input     string
	Output    string
	Timestamp time.Time
	Duration  time.Duration
}

// Analytics tracks all user interactions and tool usage
type Analytics struct {
	TotalUsers    int
	TotalUses     int
	PopularTools  map[string]int
	Interactions  []Interaction
	ToolRegistry  map[string]Tool
}

// NewAnalytics creates a new analytics instance
func NewAnalytics() *Analytics {
	return &Analytics{
		PopularTools: make(map[string]int),
		ToolRegistry: make(map[string]Tool),
	}
}

// RegisterTool adds a tool to the registry
func (a *Analytics) RegisterTool(name, category, description string) {
	a.ToolRegistry[name] = Tool{
		Name:        name,
		Category:    category,
		Description: description,
	}
}

// TrackInteraction records a user interaction
func (a *Analytics) TrackInteraction(userID, toolName, input, output string, duration time.Duration) {
	a.Interactions = append(a.Interactions, Interaction{
		UserID:    userID,
		ToolName:  toolName,
		Input:     input,
		Output:    output,
		Timestamp: time.Now(),
		Duration:  duration,
	})
	a.PopularTools[toolName]++
	a.TotalUses++
}

// GetStats returns usage statistics
func (a *Analytics) GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	stats["total_users"] = a.TotalUsers
	stats["total_uses"] = a.TotalUses
	stats["popular_tools"] = a.PopularTools
	return stats
}

func main() {
	// Initialize analytics
	analytics := NewAnalytics()

	// Register all 53 string tools
	tools := []struct {
		name string
		cat  string
		desc string
	}{
		{"StringLength", "Beginner", "Count characters in a string"},
		{"FirstChar", "Beginner", "Get first character"},
		{"LastChar", "Beginner", "Get last character"},
		{"IsEmpty", "Beginner", "Check if string is empty"},
		{"ToUpper", "Beginner", "Convert to uppercase"},
		{"ToLower", "Beginner", "Convert to lowercase"},
		{"CountAlpha", "Beginner", "Count alphabetic characters"},
		{"CheckNumber", "Beginner", "Check for digits"},
		{"CountVowels", "Beginner-Intermediate", "Count vowels in string"},
		{"ReverseString", "Beginner-Intermediate", "Reverse a string"},
		{"IsPalindrome", "Beginner-Intermediate", "Check palindrome"},
		{"RemoveSpaces", "Beginner-Intermediate", "Remove space characters"},
		{"CountRepeats", "Beginner-Intermediate", "Count consecutive repeats"},
		{"RetainFirstHalf", "Beginner-Intermediate", "Keep first half of string"},
		{"WordCount", "Intermediate", "Count words in string"},
		{"FindChar", "Intermediate", "Find character index"},
		{"CountChar", "Intermediate", "Count character occurrences"},
		{"FindLastChar", "Intermediate", "Find last character index"},
		{"ReplaceChar", "Intermediate", "Replace characters"},
		{"PrintIf", "Intermediate", "Conditional printing"},
		{"PrintIfNot", "Intermediate", "Inverse conditional printing"},
		{"LongestWord", "Intermediate", "Find longest word"},
		{"SearchReplace", "Advanced", "Search and replace substring"},
		{"CleanList", "Advanced", "Clean string list"},
		{"CountWordsAdvanced", "Advanced", "Advanced word counting"},
		{"FindSubstring", "Advanced", "Find substring index"},
		{"ReplaceAll", "Advanced", "Replace all occurrences"},
		{"Split", "Advanced", "Split string by separator"},
		{"Join", "Advanced", "Join strings with separator"},
		{"CamelToSnakeCase", "Advanced", "Convert naming convention"},
		{"Itoa", "Expert", "Integer to string"},
		{"ThirdChar", "Expert", "Extract every 3rd character"},
		{"ZipString", "Expert", "Run-length encoding"},
		{"SaveAndMiss", "Expert", "Periodic save/skip"},
		{"ReverseStrCap", "Expert", "Reverse capitalization"},
		{"Union", "Expert", "String union operation"},
		{"Intersection", "Expert", "String intersection"},
		{"StringBuilder", "Utilities", "Efficient string building"},
		{"StringSplit", "Utilities", "Split words"},
		{"StringJoin", "Utilities", "Join strings"},
		{"StringRepeat", "Utilities", "Repeat string"},
		{"StringReplace", "Utilities", "Replace substring"},
		{"StringTrim", "Utilities", "Trim whitespace"},
		{"StringContains", "Utilities", "Check substring"},
		{"StringIndex", "Utilities", "Find substring index"},
		{"StringCount", "Utilities", "Count substring"},
		{"StringPrefix", "Utilities", "Check prefix"},
		{"StringSuffix", "Utilities", "Check suffix"},
		{"StringField", "Utilities", "Split by delimiter"},
		{"StringMap", "Utilities", "Map function over string"},
		{"StringFilter", "Utilities", "Filter string characters"},
		{"StringReduce", "Utilities", "Reduce string to value"},
		{"StringFormat", "Utilities", "Format with placeholders"},
	}

	for _, t := range tools {
		analytics.RegisterTool(t.name, t.cat, t.desc)
	}

	// Display initial stats
	fmt.Println("String Tools Analytics Dashboard")
	fmt.Println("================================")
	fmt.Println()
	fmt.Printf("Total Tools Registered: %d\n", len(analytics.ToolRegistry))
	fmt.Printf("Categories: Beginner, Beginner-Intermediate, Intermediate, Advanced, Expert, Utilities\n")
	fmt.Println()
	fmt.Println("Ready to track user interactions!")
