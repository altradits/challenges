# Step-by-Step Application Building Guide

## How to Build Applications One Function at a Time

### The Pattern: Input → Process → Output

Every application follows this pattern. Let's build a simple budget planner:

## Step 1: Define the Input

```go
// budget/input.go
package budget

// BudgetInput represents user input
type BudgetInput struct {
    Income       string // "5000"
    Expenses     string // "rent:1200,groceries:300,utilities:150"
    SavingsGoal  string // "10000"
}

// ParseInput extracts values from string input
func ParseInput(s string) BudgetInput {
    // Use your string tools:
    // - Split to separate values
    // - Trim to clean whitespace
    // - StringFormat to validate
}
```

## Step 2: Process the Data

```go
// budget/calculator.go
package budget

// CalculateBudget processes the input
func CalculateBudget(input BudgetInput) BudgetResult {
    // Use your string tools:
    // - StringSplit to parse expenses
    // - StringMap to convert to numbers
    // - StringJoin to combine results
}

type BudgetResult struct {
    MonthlySavings int
    RunwayMonths   int
    Recommendations string
}
```

## Step 3: Format the Output

```go
// budget/output.go
package budget

// FormatResult creates user-friendly output
func FormatResult(result BudgetResult) string {
    // Use your string tools:
    // - StringFormat for currency
    // - StringJoin for multi-line output
    // - StringContains to add warnings
}
```

## Step 4: Create the API Endpoint

```go
// api/budget.go
func BudgetHandler(c *gin.Context) {
    input := c.PostForm("data")
    
    // Process
    parsed := budget.ParseInput(input)
    result := budget.CalculateBudget(parsed)
    output := budget.FormatResult(result)
    
    // Return
    c.JSON(200, gin.H{"result": output})
}
```

## Step 5: Add to Your Portfolio

```html
<!-- web/budget-planner.html -->
<div class="tool-container">
    <h2>Budget Planner</h2>
    <form id="budget-form">
        <input name="income" placeholder="Monthly Income">
        <input name="expenses" placeholder="Expenses (rent:1200,groceries:300)">
        <input name="goal" placeholder="Savings Goal">
        <button type="submit">Calculate</button>
    </form>
    <pre id="result"></pre>
</div>

<script>
document.getElementById('budget-form').onsubmit = async (e) => {
    e.preventDefault();
    const data = new FormData(e.target);
    const response = await fetch('/api/budget', {
        method: 'POST',
        body: data
    });
    const result = await response.json();
    document.getElementById('result').textContent = result.result;
};
</script>
```

## Building Your Finance Tools

### Tool 1: Budget Planner
**Functions needed:**
1. `ParseInput` - Split and clean input
2. `CalculateSavings` - Subtract expenses from income
3. `FormatOutput` - Create readable report

### Tool 2: Savings Calculator
**Functions needed:**
1. `ParseGoal` - Extract target amount
2. `CalculateTime` - Months to reach goal
3. `FormatTimeline` - "X years Y months"

### Tool 3: Investment Tracker
**Functions needed:**
1. `ParseTicker` - Extract stock symbol
2. `FormatChange` - "+1.2%" or "-0.5%"
3. `GenerateReport` - Multi-line performance

### Tool 4: Expense Splitter
**Functions needed:**
1. `ParsePeople` - Split names
2. `CalculateShares` - Divide amounts
3. `FormatBill` - "Alice: $25, Bob: $25"

### Tool 5: Currency Converter
**Functions needed:**
1. `ParseAmount` - Extract number
2. `ConvertCurrency` - Apply exchange rate
3. `FormatResult` - "$100 USD = €90 EUR"

## The Development Flow

```
1. Write the function signature
2. Add TODO comments
3. Implement with your string tools
4. Test with sample input
5. Add to portfolio website
6. Deploy and share
```

## Example: Building StringLength Tool

```go
// 1. Function signature
func StringLength(s string) int

// 2. TODO comments
// - Iterate through string
// - Count each character
// - Return count

// 3. Implementation
func StringLength(s string) int {
    count := 0
    for range s {
        count++
    }
    return count
}

// 4. Test
fmt.Println(StringLength("Hello")) // 5

// 5. Add to website
// Create web/stringlength.html with demo

// 6. Deploy
// git push to GitHub Pages
```

## Your Portfolio Website Structure

```
portfolio/
├── index.html          # Landing page
├── tools/
│   ├── stringlength.html
│   ├── toupper.html
│   └── ... (all 53 tools)
├── finance-tools/
│   ├── budget-planner.html
│   ├── savings-calculator.html
│   └── ... (5 finance tools)
├── pricing.html        # Subscription plans
├── dashboard.html        # User analytics
└── api/
    └── ... (Go backend)
```

## Monetization Integration

```html
<!-- Add to each tool page -->
<div class="upgrade-prompt" id="upgrade-prompt" style="display:none">
    <p>You've reached your limit. <a href="/pricing">Upgrade for more uses</a></p>
</div>
```

```go
// Check usage in backend
func CheckUsage(userID, toolName string) bool {
    // Query analytics database
    // Return true if under limit
}
```

## Next Steps

1. Start with one simple tool (StringLength)
2. Add it to your portfolio website
3. Deploy to GitHub Pages
4. Share the link
5. Add more tools one by one
6. Add payment integration
7. Launch your first product!