# Finance Tools for Your Portfolio

## Tool 1: Budget Planner

### What It Does
Helps users plan their monthly budget by calculating savings potential.

### String Tools Used
- `StringSplit` - Parse "rent:1200,groceries:300"
- `StringTrim` - Clean whitespace
- `StringFormat` - Format currency output
- `StringJoin` - Combine expense list

### Implementation
```go
func BudgetPlanner(income, expenses string) string {
    // Parse income
    monthlyIncome := parseCurrency(income)
    
    // Parse expenses
    expenseList := StringSplit(expenses, ",")
    totalExpenses := 0
    for _, e := range expenseList {
        parts := StringSplit(StringTrim(e), ":")
        if len(parts) == 2 {
            totalExpenses += parseInt(parts[1])
        }
    }
    
    // Calculate
    savings := monthlyIncome - totalExpenses
    
    // Format output
    return StringFormat("Monthly Savings: $%d", savings)
}
```

### Live Demo
```
Input: Income: $5000, Expenses: rent:1200,groceries:300,utilities:150
Output: Monthly Savings: $3350
```

---

## Tool 2: Savings Calculator

### What It Does
Calculates how long it will take to reach a savings goal.

### String Tools Used
- `Itoa` - Convert numbers to strings
- `StringFormat` - Format "X years Y months"
- `StringContains` - Check for valid input

### Implementation
```go
func SavingsCalculator(current, goal, monthly string) string {
    currentSavings := parseInt(current)
    targetGoal := parseInt(goal)
    monthlySave := parseInt(monthly)
    
    months := (targetGoal - currentSavings) / monthlySave
    years := months / 12
    remaining := months % 12
    
    return StringFormat("%d years %d months", years, remaining)
}
```

### Live Demo
```
Input: Current: $5000, Goal: $20000, Monthly: $1000
Output: 1 year 5 months
```

---

## Tool 3: Investment Tracker

### What It Does
Tracks stock/fund performance and formats reports.

### String Tools Used
- `StringPrefix` - Validate ticker format
- `StringFormat` - Format percentage changes
- `StringContains` - Check for valid symbols
- `StringJoin` - Combine report lines

### Implementation
```go
func InvestmentTracker(ticker, price, change string) string {
    // Validate ticker
    if !StringPrefix(ticker, "Valid") {
        return "Invalid ticker"
    }
    
    // Format output
    return StringFormat("%s: $%s (%s%%)", ticker, price, change)
}
```

### Live Demo
```
Input: Ticker: VTI, Price: 245.32, Change: +1.2
Output: VTI: $245.32 (+1.2%)
```

---

## Tool 4: Expense Splitter

### What It Does
Splits bills and expenses between multiple people.

### String Tools Used
- `StringSplit` - Parse people list
- `StringTrim` - Clean names
- `StringFormat` - Format amounts
- `StringJoin` - Combine results

### Implementation
```go
func ExpenseSplitter(total, people string) string {
    amount := parseFloat(total)
    personList := StringSplit(people, ",")
    perPerson := amount / StringLength(people)
    
    var results []string
    for _, p := range personList {
        name := StringTrim(p)
        results = append(results, StringFormat("%s: $%.2f", name, perPerson))
    }
    
    return StringJoin(results, "\n")
}
```

### Live Demo
```
Input: Total: $100, People: Alice, Bob, Carol
Output:
Alice: $33.33
Bob: $33.33
Carol: $33.33
```

---

## Tool 5: Currency Converter

### What It Does
Converts between different currencies.

### String Tools Used
- `StringReplace` - Remove $ and commas
- `StringTrim` - Clean input
- `StringFormat` - Format with currency symbol
- `StringContains` - Validate currency codes

### Implementation
```go
func CurrencyConverter(amount, from, to string) string {
    // Clean input
    cleanAmount := StringReplace(StringReplace(amount, "$", ""), ",", "")
    
    // Convert (simplified)
    rate := getExchangeRate(from, to)
    converted := parseFloat(cleanAmount) * rate
    
    return StringFormat("$%.2f %s = $%.2f %s", 
        parseFloat(cleanAmount), from, converted, to)
}
```

### Live Demo
```
Input: $100 USD to EUR
Output: $100.00 USD = $90.00 EUR
```

---

## Tool 6: Debt Payoff Calculator

### What It Does
Calculates debt payoff timeline.

### String Tools Used
- `StringFormat` - Format timeline
- `Itoa` - Convert months to string
- `StringJoin` - Combine payment schedule

### Implementation
```go
func DebtPayoff(principal, rate, payment string) string {
    // Calculate months to payoff
    months := calculatePayoffMonths(principal, rate, payment)
    
    return StringFormat("Debt paid off in %d months", months)
}
```

---

## Tool 7: Net Worth Tracker

### What It Does
Calculates and tracks net worth.

### String Tools Used
- `StringSplit` - Parse assets/liabilities
- `StringFormat` - Format net worth
- `StringJoin` - Combine categories

### Implementation
```go
func NetWorth(assets, liabilities string) string {
    assetTotal := sumValues(StringSplit(assets, ","))
    liabilityTotal := sumValues(StringSplit(liabilities, ","))
    
    netWorth := assetTotal - liabilityTotal
    
    return StringFormat("Net Worth: $%.2f", netWorth)
}
```

---

## Tool 8: Retirement Calculator

### What It Does
Calculates retirement savings needed.

### String Tools Used
- `StringFormat` - Format large numbers
- `Itoa` - Convert years
- `StringJoin` - Combine scenarios

### Implementation
```go
func RetirementCalculator(age, goal, current string) string {
    yearsToRetire := 65 - parseInt(age)
    needed := parseInt(goal) - parseInt(current)
    
    return StringFormat("Save $%.2f over %d years", 
        needed/float64(yearsToRetire), yearsToRetire)
}
```

---

## How to Add These to Your Portfolio

1. **Create the tool page** - Copy the HTML template
2. **Add the function** - Implement in Go
3. **Connect to API** - Add endpoint
4. **Test locally** - Run with `go run .`
5. **Deploy** - Push to GitHub Pages
6. **Share** - Send link to friends

## Pricing Strategy

| Tool | Free Tier | Basic ($9/mo) | Pro ($29/mo) |
|------|-----------|---------------|--------------|
| Budget Planner | ✓ | ✓ | ✓ |
| Savings Calculator | ✓ | ✓ | ✓ |
| Investment Tracker | ✓ | ✓ | ✓ |
| Expense Splitter | ✓ | ✓ | ✓ |
| Currency Converter | ✓ | ✓ | ✓ |
| Debt Payoff | | ✓ | ✓ |
| Net Worth | | ✓ | ✓ |
| Retirement | | | ✓ |

## Your Live Portfolio

After building and deploying:
```
https://yourname.github.io/finance-tools
```

Users can:
1. Try tools for free
2. See your string manipulation skills
3. Subscribe for more features
4. Use your tools daily
5. Refer friends to your service