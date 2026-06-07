# 05. NumPy

## What You'll Learn

This exercise teaches you **NumPy for numerical computing**. By the end, you'll understand:
- Arrays and operations
- Mathematical functions
- Statistical functions
- Financial calculations

## Theory: NumPy

### Installation

```bash
pip install numpy
```

### Basic Operations

```python
import numpy as np

# Create arrays
amounts = np.array([100, 200, 150, 300])
returns = np.array([0.05, -0.02, 0.08, 0.03])

# Mathematical operations
total = np.sum(amounts)
average = np.mean(amounts)
std_dev = np.std(amounts)

# Financial calculations
def compound_interest(principal, rate, time):
    return principal * np.power(1 + rate, time)

# Portfolio metrics
def portfolio_metrics(returns):
    return {
        'mean': np.mean(returns),
        'std': np.std(returns),
        'sharpe': np.mean(returns) / np.std(returns)
    }
```

## The Challenge

Create financial calculations with NumPy.

### Expected Code

```python
import numpy as np

def calculate_portfolio_returns(prices, weights):
    """Calculate portfolio returns."""
    returns = np.diff(prices) / prices[:-1]
    portfolio_return = np.sum(returns * weights)
    return portfolio_return

def calculate_volatility(returns):
    """Calculate volatility (standard deviation)."""
    return np.std(returns) * np.sqrt(252)  # Annualized

def calculate_sharpe_ratio(returns, risk_free_rate=0.02):
    """Calculate Sharpe ratio."""
    excess_returns = returns - risk_free_rate / 252
    return np.mean(excess_returns) / np.std(excess_returns) * np.sqrt(252)

def monte_carlo_simulation(initial, mean_return, volatility, years, simulations=1000):
    """Monte Carlo simulation for investment growth."""
    days = years * 252
    dt = 1 / 252
    
    # Random returns
    returns = np.random.normal(
        mean_return * dt, 
        volatility * np.sqrt(dt), 
        (simulations, days)
    )
    
    # Cumulative returns
    prices = initial * np.exp(np.cumsum(returns, axis=1))
    
    return prices
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Statistical analysis
- **Savings Calculator** - Future value calculations
- **Investment Tracker** - Risk metrics
- **Net Worth Tracker** - Growth projections

## Next Steps

After completing this, move to:
- [06-visualization](06-visualization/README.md) - Visualization
- [07-api](07-api/README.md) - API integration