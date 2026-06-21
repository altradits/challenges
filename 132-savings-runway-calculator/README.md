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
- [133-portfolio-website](../133-portfolio-website/README.md) - Portfolio Website
- [134-professional](../134-professional/README.md) - Professional