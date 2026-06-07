# 02. Data Structures

## What You'll Learn

This exercise teaches you **Python data structures**. By the end, you'll understand:
- Lists and list operations
- Dictionaries and key-value access
- Sets for unique values
- Tuples for immutable data

## Theory: Python Data Structures

### Lists

```python
# Create list
budgets = [
    {"category": "food", "amount": 100},
    {"category": "rent", "amount": 500}
]

# Add item
budgets.append({"category": "utilities", "amount": 150})

# Access item
first = budgets[0]
last = budgets[-1]

# List comprehension
total = sum(b['amount'] for b in budgets)
categories = [b['category'] for b in budgets]
```

### Dictionaries

```python
# Create dictionary
user = {
    "id": 1,
    "name": "John",
    "email": "john@example.com",
    "budgets": []
}

# Access values
name = user['name']
amount = user.get('amount', 0)  # Default value

# Add/update
user['budgets'].append({"category": "food", "amount": 100})

# Iterate
for key, value in user.items():
    print(f"{key}: {value}")
```

### Sets

```python
# Unique categories
categories = {"food", "rent", "utilities"}
categories.add("food")  # Won't duplicate

# Set operations
all_categories = {"food", "rent", "utilities", "entertainment"}
used = {"food", "rent"}
unused = all_categories - used
```

## The Challenge

Create data structures for FinTech portfolio.

### Expected Code

```python
# Create user portfolio structure
def create_user_portfolio():
    return {
        "user": {
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com"
        },
        "budgets": [],
        "savings": [],
        "investments": [],
        "net_worth": 0
    }

# Add budget
def add_budget(portfolio, category, amount):
    portfolio["budgets"].append({
        "id": len(portfolio["budgets"]) + 1,
        "category": category,
        "amount": amount,
        "date": "2024-01-01"
    })
    return portfolio

# Calculate totals
def get_budget_totals(budgets):
    totals = {}
    for budget in budgets:
        category = budget["category"]
        totals[category] = totals.get(category, 0) + budget["amount"]
    return totals

# Get unique categories
def get_categories(budgets):
    return {b["category"] for b in budgets}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Store and organize budgets
- **Savings Calculator** - Track savings goals
- **Investment Tracker** - Portfolio data
- **Net Worth Tracker** - Asset categories

## Next Steps

After completing this, move to:
- [03-fileio](03-fileio/README.md) - File I/O
- [04-pandas](04-pandas/README.md) - Pandas