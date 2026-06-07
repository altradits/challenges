# 01. Python Basics

## What You'll Learn

This exercise teaches you **Python fundamentals**. By the end, you'll understand:
- Variables and data types
- Functions
- Control flow (if, for, while)
- Error handling

## Theory: Python Fundamentals

### Variables and Types

```python
# Numbers
amount = 100.50
count = 5

# Strings
name = "John Doe"
email = 'john@example.com'

# Booleans
is_active = True
is_verified = False

# None (null equivalent)
user = None
```

### Functions

```python
def calculate_savings(monthly, rate, months):
    """Calculate future savings with compound interest."""
    return monthly * ((1 + rate) ** months - 1) / rate

def get_budget_total(budgets):
    """Sum all budget amounts."""
    total = 0
    for budget in budgets:
        total += budget['amount']
    return total
```

### Control Flow

```python
# If/else
if amount > 1000:
    print("High budget")
elif amount > 500:
    print("Medium budget")
else:
    print("Low budget")

# For loop
for budget in budgets:
    print(f"{budget['category']}: ${budget['amount']}")

# While loop
while True:
    user_input = input("Enter amount (or 'q' to quit): ")
    if user_input == 'q':
        break
```

## The Challenge

Create basic Python functions for FinTech calculations.

### Expected Code

```python
def calculate_compound_interest(principal, rate, time, compounds_per_year=12):
    """Calculate compound interest."""
    return principal * (1 + rate / compounds_per_year) ** (compounds_per_year * time)

def calculate_monthly_payment(principal, rate, years):
    """Calculate monthly loan payment."""
    monthly_rate = rate / 12
    months = years * 12
    return principal * monthly_rate * (1 + monthly_rate) ** months / ((1 + monthly_rate) ** months - 1)

def categorize_amount(amount):
    """Categorize spending amount."""
    if amount < 100:
        return "low"
    elif amount < 500:
        return "medium"
    else:
        return "high"

# Test
if __name__ == "__main__":
    print(calculate_compound_interest(1000, 0.05, 10))
    print(calculate_monthly_payment(10000, 0.04, 5))
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Categorize spending
- **Savings Calculator** - Compound interest
- **Investment Tracker** - Return calculations
- **Net Worth Tracker** - Growth projections

## Next Steps

After completing this, move to:
- [02-structures](02-structures/README.md) - Data structures
- [03-fileio](03-fileio/README.md) - File I/O