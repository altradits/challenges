# String Tools Analytics Dashboard

## Objective

Build a comprehensive analytics dashboard that tracks all the string manipulation tools you've created, monitors user interactions, and provides insights on which features are most popular.

This is your **final capstone project** that integrates all 53 string manipulation exercises into a unified platform.

---

## Project Structure

```
capstone/
├── main.go           # Entry point
├── analytics.go      # User tracking and metrics
├── tools/            # All your string tools
│   ├── stringlength/
│   ├── toupper/
│   └── ... (all 53 tools)
├── api/              # REST API endpoints
├── web/              # Web dashboard
└── data/             # User interaction data
```

---

## Learning Curve

### Phase 1: Data Collection
- Track which tools are being used
- Record user inputs and outputs
- Store interaction timestamps

### Phase 2: Analytics Engine
- Calculate usage statistics
- Identify popular features
- Generate usage reports

### Phase 3: Dashboard
- Visualize tool usage
- Show user engagement metrics
- Provide actionable insights

---

## Building Blocks

### 1. Tool Registry
```go
type Tool struct {
    Name        string
    Category    string
    UsageCount  int
    LastUsed    time.Time
    Description string
}
```

### 2. User Interaction Model
```go
type Interaction struct {
    UserID    string
    ToolName  string
    Input     string
    Output    string
    Timestamp time.Time
    Duration  time.Duration
}
```

### 3. Analytics Engine
```go
type Analytics struct {
    TotalUsers     int
    TotalUses      int
    PopularTools   map[string]int
    FeatureMatrix  map[string][]string
}
```

---

## Implementation Guide

### Technology Stack
- **Backend:** Go with Gin or Echo framework
- **Database:** SQLite (for simplicity) or PostgreSQL
- **Frontend:** HTML/CSS/JavaScript with Chart.js for visualization
- **Deployment:** Docker, deployed to cloud platform

### Step-by-Step Implementation

1. **Create Tool Registry**
   - Define all 53 tools in a registry
   - Assign categories (beginner, intermediate, advanced)
   - Track metadata for each tool

2. **Build Analytics API**
   - `POST /api/interact` - Record user interactions
   - `GET /api/stats` - Get usage statistics
   - `GET /api/tools/:name` - Get specific tool analytics

3. **Implement Web Dashboard**
   - Tool usage chart (bar graph)
   - User engagement timeline
   - Feature popularity heatmap

4. **Add User Tracking**
   - Generate unique user IDs
   - Track session duration
   - Record feature interactions

---

## Features to Track

| Tool | What to Track |
|------|---------------|
| StringLength | Input length, frequency |
| ToUpper/ToLower | Case conversion patterns |
| CountVowels | Vowel distribution in inputs |
| IsPalindrome | Palindrome detection rate |
| WordCount | Word count patterns |
| SearchReplace | Search patterns, replace frequency |

---

## Expected Output

```
Total Users: 150
Total Tool Uses: 1,247

Most Popular Tools:
1. StringLength - 342 uses
2. ToUpper - 215 uses
3. CountVowels - 187 uses

User Engagement:
- Active today: 23 users
- Active this week: 87 users
- Average session: 12 minutes
```

---

## How to Use Your Tools

Each tool from the 53-exercise course should be:
1. Registered in the tool registry
2. Wrapped with analytics tracking
3. Exposed via API endpoint
4. Displayed in the dashboard

Example integration:
```go
func (a *Analytics) TrackToolUse(toolName, input, output string) {
    a.interactions = append(a.interactions, Interaction{
        ToolName:  toolName,
        Input:     input,
        Output:    output,
        Timestamp: time.Now(),
    })
}
```

---

## Next Steps

After completing this, you'll be ready for:
- [158-bitcoin-oss](../158-bitcoin-oss/README.md) - Your first Bitcoin open-source PR
- [159-closures](../159-closures/README.md) - Closures and First-Class Functions

---

## Function contracts (what the tests pin)

Implement in `package piscine` (the analytics core; the dashboard, API and
deployment sections in this file are stretch goals after this passes):

- The `Tool`, `Interaction` and `Analytics` structs exactly as defined above.
- `func (a *Analytics) TrackToolUse(toolName, input, output string)` —
  appends an `Interaction` (with `Timestamp` set), increments `TotalUses`,
  and counts the tool in `PopularTools` (initializing the map when nil).
- `func (a *Analytics) TopTool() string` — the most-used tool name, or `""`
  when nothing is tracked.

## Time limit

**240 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 157
```

or run the tests directly:

```bash
go test ./157-capstone/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 157` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
