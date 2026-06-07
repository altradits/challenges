# 06. Data Visualization

## What You'll Learn

This exercise teaches you **data visualization with Python**. By the end, you'll understand:
- Matplotlib basics
- Chart types
- Styling and customization
- Saving visualizations

## Theory: Visualization

### Installation

```bash
pip install matplotlib
```

### Basic Charts

```python
import matplotlib.pyplot as plt

# Line chart
plt.plot(months, savings)
plt.title('Savings Growth')
plt.xlabel('Month')
plt.ylabel('Amount ($)')
plt.show()

# Bar chart
categories = ['Food', 'Rent', 'Utilities']
amounts = [100, 500, 150]
plt.bar(categories, amounts)
plt.show()

# Pie chart
plt.pie(amounts, labels=categories, autopct='%1.1f%%')
plt.show()
```

### Multiple Charts

```python
fig, axes = plt.subplots(2, 2, figsize=(12, 8))

# Budget breakdown
axes[0, 0].pie(budget_amounts, labels=categories)
axes[0, 0].set_title('Budget Distribution')

# Monthly trend
axes[0, 1].plot(months, totals)
axes[0, 1].set_title('Monthly Spending')

# Category comparison
axes[1, 0].bar(categories, amounts)
axes[1, 0].set_title('Category Totals')

# Savings projection
axes[1, 1].plot(years, projected)
axes[1, 1].set_title('Savings Projection')

plt.tight_layout()
plt.savefig('fintech-dashboard.png')
```

## The Challenge

Create visualizations for the FinTech portfolio.

### Expected Code

```python
import matplotlib.pyplot as plt
import numpy as np

def create_budget_dashboard(budgets):
    """Create budget visualization dashboard."""
    categories = [b['category'] for b in budgets]
    amounts = [b['amount'] for b in budgets]
    
    fig, axes = plt.subplots(1, 2, figsize=(12, 5))
    
    # Pie chart
    axes[0].pie(amounts, labels=categories, autopct='%1.1f%%')
    axes[0].set_title('Budget Distribution')
    
    # Bar chart
    axes[1].bar(categories, amounts)
    axes[1].set_title('Spending by Category')
    axes[1].set_ylabel('Amount ($)')
    
    plt.tight_layout()
    plt.savefig('budget-dashboard.png')
    return fig

def create_savings_projection_chart(monthly, rate, years):
    """Create savings projection chart."""
    months = np.arange(1, years * 12 + 1)
    values = monthly * ((1 + rate) ** (months/12) - 1) / rate
    
    plt.figure(figsize=(10, 6))
    plt.plot(months, values)
    plt.title('Savings Growth Projection')
    plt.xlabel('Month')
    plt.ylabel('Total ($)')
    plt.grid(True)
    plt.savefig('savings-projection.png')
    return plt
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Spending charts
- **Savings Calculator** - Growth graphs
- **Investment Tracker** - Portfolio performance
- **Net Worth Tracker** - Asset trends

## Next Steps

After completing this, move to:
- [07-api](07-api/README.md) - API integration
- [08-analysis](08-analysis/README.md) - Data analysis