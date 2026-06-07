# 08. Data Analysis

## What You'll Learn

This exercise teaches you **financial data analysis**. By the end, you'll understand:
- Statistical analysis
- Trend detection
- Correlation analysis
- Insights generation

## Theory: Data Analysis

### Statistical Functions

```python
import numpy as np
import pandas as pd

# Basic statistics
df = pd.DataFrame(budgets)
mean = df['amount'].mean()
median = df['amount'].median()
std = df['amount'].std()

# Percentiles
p25 = df['amount'].quantile(0.25)
p75 = df['amount'].quantile(0.75)
```

### Trend Analysis

```python
# Monthly trends
df['date'] = pd.to_datetime(df['date'])
df['month'] = df['date'].dt.to_period('M')
monthly = df.groupby('month')['amount'].sum()

# Growth rate
monthly.pct_change()

# Moving average
df['moving_avg'] = df['amount'].rolling(window=3).mean()
```

## The Challenge

Analyze FinTech data for insights.

### Expected Code

```python
import pandas as pd
import numpy as np

def analyze_spending_patterns(budgets):
    """Analyze spending patterns."""
    df = pd.DataFrame(budgets)
    df['date'] = pd.to_datetime(df['date'])
    
    return {
        'total_spent': df['amount'].sum(),
        'avg_monthly': df.groupby(df['date'].dt.to_period('M'))['amount'].sum().mean(),
        'top_category': df.groupby('category')['amount'].sum().idxmax(),
        'spending_trend': df.groupby(df['date'].dt.to_period('M'))['amount'].sum().pct_change().mean()
    }

def detect_anomalies(budgets, threshold=2):
    """Detect unusual spending."""
    df = pd.DataFrame(budgets)
    mean = df['amount'].mean()
    std = df['amount'].std()
    
    anomalies = df[
        (df['amount'] > mean + threshold * std) |
        (df['amount'] < mean - threshold * std)
    ]
    return anomalies.to_dict('records')
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Spending insights
- **Savings Calculator** - Progress tracking
- **Investment Tracker** - Performance analysis
- **Net Worth Tracker** - Growth trends

## Next Steps

After completing this, move to:
- [09-ml](09-ml/README.md) - Machine learning
- [10-deployment](10-deployment/README.md) - Deployment