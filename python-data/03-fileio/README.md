# 03. File I/O

## What You'll Learn

This exercise teaches you **file operations in Python**. By the end, you'll understand:
- Reading and writing files
- CSV file handling
- JSON serialization
- Error handling

## Theory: File I/O

### Reading/Writing Files

```python
# Write to file
with open('budgets.txt', 'w') as f:
    f.write('Food: 100\n')
    f.write('Rent: 500\n')

# Read from file
with open('budgets.txt', 'r') as f:
    content = f.read()
    print(content)
```

### CSV Files

```python
import csv

# Read CSV
with open('budgets.csv', 'r') as f:
    reader = csv.DictReader(f)
    budgets = [row for row in reader]

# Write CSV
with open('budgets.csv', 'w', newline='') as f:
    writer = csv.DictWriter(f, fieldnames=['category', 'amount', 'date'])
    writer.writeheader()
    writer.writerow({'category': 'food', 'amount': 100, 'date': '2024-01-01'})
```

### JSON

```python
import json

# Read JSON
with open('portfolio.json', 'r') as f:
    portfolio = json.load(f)

# Write JSON
with open('portfolio.json', 'w') as f:
    json.dump(portfolio, f, indent=2)
```

## The Challenge

Create file operations for FinTech data.

### Expected Code

```python
import json
import csv
from datetime import datetime

def save_budgets(budgets, filename='budgets.json'):
    """Save budgets to JSON file."""
    with open(filename, 'w') as f:
        json.dump(budgets, f, indent=2)

def load_budgets(filename='budgets.json'):
    """Load budgets from JSON file."""
    try:
        with open(filename, 'r') as f:
            return json.load(f)
    except FileNotFoundError:
        return []

def export_budgets_csv(budgets, filename='budgets.csv'):
    """Export budgets to CSV."""
    with open(filename, 'w', newline='') as f:
        writer = csv.DictWriter(f, fieldnames=['category', 'amount', 'date'])
        writer.writeheader()
        for budget in budgets:
            writer.writerow(budget)

def import_budgets_csv(filename='budgets.csv'):
    """Import budgets from CSV."""
    budgets = []
    with open(filename, 'r') as f:
        reader = csv.DictReader(f)
        for row in reader:
            budgets.append({
                'category': row['category'],
                'amount': float(row['amount']),
                'date': row['date']
            })
    return budgets
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Export/import budgets
- **Savings Calculator** - Save calculation history
- **Investment Tracker** - Import portfolio data
- **Net Worth Tracker** - Backup financial data

## Next Steps

After completing this, move to:
- [04-pandas](04-pandas/README.md) - Pandas
- [05-numpy](05-numpy/README.md) - NumPy