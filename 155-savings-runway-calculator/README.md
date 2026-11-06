# Savings Runway Calculator

## Objective

Build a service that calculates how long savings will last, applying string manipulation skills to process expense data and format results.

## Learning Curve

### Phase 1: Core Calculator
- Parse expense descriptions and amounts
- Calculate runway duration
- Format time periods

### Phase 2: String Processing
- Process budget category strings
- Validate date formats
- Generate milestone messages

### Phase 3: User Experience
- Create readable output
- Add progress tracking
- Build notification system

## Building Blocks

### 1. Expense Parser
```go
type Expense struct {
    Description string
    Amount      float64
    Category    string
    Date        string
}

func ParseExpense(s string) (Expense, error) {
    // Parse "Groceries: $150.50" format
}
```

### 2. Runway Calculator
```go
func CalculateRunway(savings, monthlyExpenses float64) (years, months int) {
    // Returns (2, 6) for 30 months
}

func FormatRunway(years, months int) string {
    // Returns "2 years 6 months"
}
```

### 3. Milestone Generator
```go
func GenerateMilestone(months int) string {
    // Returns "Milestone: 1 year of savings!"
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/calculate | Calculate savings runway |
| POST | /api/expenses | Add expense entry |
| GET | /api/milestones | Get upcoming milestones |
| GET | /api/report | Get formatted report |

## Technology Stack

- **Framework:** Echo or Fiber (Go web framework)
- **Database:** SQLite for local, PostgreSQL for production
- **Authentication:** JWT for user accounts
- **Notifications:** Email/SMS via Twilio
- **Deployment:** Docker + Fly.io

## How to Use Your String Tools

| Tool | Runway Calculator Application |
|------|-------------------------------|
| StringSplit | Parse "amount,category,description" |
| StringTrim | Clean expense input strings |
| StringContains | Validate category names |
| StringFormat | Format "X years Y months" |
| StringJoin | Combine report sections |
| StringPrefix | Check date format (YYYY-MM) |
| StringReplace | Clean currency symbols |
| StringMap | Normalize category names |

## Deployment Link

After deployment, your calculator will be available at:
```
https://savings-runway-calculator.yourname.com
```

## Next Steps

After completing this, you'll be ready for:
- [156-api-lesson](../156-api-lesson/README.md) - Go API Development
- [157-capstone](../157-capstone/README.md) - String Tools Analytics Dashboard

---

## Function contracts (what the tests pin)

Implement in `package piscine`:

- `ParseExpense(s string) (Expense, error)` — parse `"Groceries: $150.50"`;
  error when the colon or amount is missing or malformed.
- `CalculateRunway(savings, monthlyExpenses float64) (years, months int)` —
  `int(savings/monthlyExpenses)` split into years and months (`30` →
  `(2, 6)`); `(0, 0)` when expenses are not positive.
- `FormatRunway(years, months int) string` — `"2 years 6 months"`
  (singular "year"/"month" when the value is 1).
- `GenerateMilestone(months int) string` — `"Milestone: 6 months of
  savings!"` under a year, `"Milestone: 1 year of savings!"` /
  `"Milestone: 2 years of savings!"` after.

Those API, dashboard and deployment sections are stretch goals: build them
after the function core passes its tests.

## Time limit

**120 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 155
```

or run the tests directly:

```bash
go test ./155-savings-runway-calculator/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 155` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
