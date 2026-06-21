# 04. Pandas

## What You'll Learn

This exercise teaches you **Pandas for data analysis**. By the end, you'll understand:
- DataFrames and Series
- Filtering and grouping
- Aggregation
- Time series

## Theory: Pandas

### Installation

```bash
pip install pandas
```

### Basic Operations

```python
import pandas as pd

# Create DataFrame
df = pd.DataFrame([
    {'category': 'food', 'amount': 100, 'date': '2024-01-01'},
    {'category': 'rent', 'amount': 500, 'date': '2024-01-01'}
])

# Read CSV
df = pd.read_csv('budgets.csv')

# Filter
food = df[df['category'] == 'food']

# Group and aggregate
totals = df.groupby('category')['amount'].sum()

# Add column
df['year'] = pd.to_datetime(df['date']).dt.year

# Sort
df = df.sort_values('amount', ascending=False)
```

## The Challenge

Analyze FinTech data with Pandas.

### Expected Code

```python
import pandas as pd

def analyze_budgets(budgets):
    """Analyze budget data."""
    df = pd.DataFrame(budgets)
    
    # Monthly totals
    df['date'] = pd.to_datetime(df['date'])
    monthly = df.groupby(df['date'].dt.to_period('M'))['amount'].sum()
    
    # Category totals
    by_category = df.groupby('category')['amount'].agg(['sum', 'mean', 'count'])
    
    # Spending trends
    df['month'] = df['date'].dt.month
    monthly_avg = df.groupby('month')['amount'].mean()
    
    return {
        'monthly_totals': monthly.to_dict(),
        'by_category': by_category.to_dict(),
        'monthly_avg': monthly_avg.to_dict()
    }

def calculate_savings_projection(monthly_savings, rate, years):
    """Project savings growth."""
    months = years * 12
    months_list = range(1, months + 1)
    values = [monthly_savings * ((1 + rate) ** (m/12) - 1) / rate for m in months_list]
    
    return pd.DataFrame({
        'month': months_list,
        'total': values
    })
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Spending analysis
- **Savings Calculator** - Growth projections
- **Investment Tracker** - Portfolio analysis
- **Net Worth Tracker** - Trend visualization

## Next Steps

After completing this, move to:
- [05-numpy](05-numpy/README.md) - NumPy
- [06-visualization](06-visualization/README.md) - Visualization