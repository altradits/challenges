# Index Fund Tracker

## Objective

Create a backend service for tracking index fund performance with string manipulation for parsing and formatting.

## Learning Curve

### Phase 1: Data Parsing
- Parse fund ticker symbols
- Extract performance data from strings
- Validate fund data formats

### Phase 2: Comparison Engine
- Compare fund performance
- Generate comparison strings
- Format percentage changes

### Phase 3: Reporting
- Create performance reports
- Generate time-series data
- Build user dashboards

## Building Blocks

### 1. Ticker Parser
```go
type Fund struct {
    Ticker     string
    Name       string
    Price      float64
    Change     float64
    ChangePct  float64
}

func ParseTicker(s string) (Fund, error) {
    // Extract ticker from "VTI: $245.32 (+1.2%)"
}
```

### 2. Performance Formatter
```go
func FormatChange(change float64) string {
    // Returns "+1.2%" or "-0.5%"
}

func FormatComparison(f1, f2 Fund) string {
    // Returns "VTI outperforms VOO by 2.3%"
}
```

### 3. Report Generator
```go
func GeneratePerformanceReport(funds []Fund) string {
    // Returns formatted multi-line report
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/funds | List all tracked funds |
| GET | /api/funds/:ticker | Get specific fund |
| POST | /api/compare | Compare two funds |
| GET | /api/report | Get performance report |

## Technology Stack

- **Framework:** Gin (Go web framework)
- **Data Source:** Alpha Vantage or Yahoo Finance API
- **Database:** PostgreSQL with timescaledb extension
- **Frontend:** React or Vue.js dashboard
- **Deployment:** Docker + AWS/Azure

## How to Use Your String Tools

| Tool | Fund Tracking Application |
|------|--------------------------|
| StringPrefix | Validate ticker format (3-5 letters) |
| StringContains | Search funds by name |
| StringReplace | Clean API response data |
| StringSplit | Parse CSV fund lists |
| StringJoin | Combine report sections |
| StringFormat | Format percentage changes |
| StringTrim | Clean whitespace from data |
| StringMap | Transform ticker to uppercase |

## Deployment Link

After deployment, your tracker will be available at:
```
https://index-fund-tracker.yourname.com
```

## Next Steps

1. Integrate with financial data API
2. Implement ticker parsing
3. Add performance comparison logic
4. Create web dashboard
5. Deploy and share with investors