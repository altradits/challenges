# 09. Machine Learning

## What You'll Learn

This exercise teaches you **machine learning for finance**. By the end, you'll understand:
- Scikit-learn basics
- Predictive modeling
- Time series forecasting
- Risk assessment

## Theory: Machine Learning

### Installation

```bash
pip install scikit-learn
```

### Basic Model

```python
from sklearn.linear_model import LinearRegression
import numpy as np

# Simple prediction
X = np.array([[1], [2], [3], [4]])  # Months
y = np.array([100, 150, 200, 250])  # Savings

model = LinearRegression()
model.fit(X, y)

# Predict next month
next_month = model.predict([[5]])
```

### Time Series

```python
from sklearn.ensemble import RandomForestRegressor

# Stock price prediction
model = RandomForestRegressor()
model.fit(historical_prices, future_prices)

prediction = model.predict([[current_price]])
```

## The Challenge

Create ML models for FinTech.

### Expected Code

```python
from sklearn.linear_model import LinearRegression
import numpy as np

def predict_savings_growth(historical_data):
    """Predict future savings."""
    months = np.array([[i] for i in range(len(historical_data))])
    amounts = np.array(historical_data)
    
    model = LinearRegression()
    model.fit(months, amounts)
    
    # Predict next 12 months
    future_months = np.array([[i] for i in range(len(historical_data), len(historical_data) + 12)])
    predictions = model.predict(future_months)
    
    return predictions.tolist()

def assess_investment_risk(returns):
    """Assess investment risk level."""
    from sklearn.cluster import KMeans
    
    returns_array = np.array(returns).reshape(-1, 1)
    kmeans = KMeans(n_clusters=3)
    kmeans.fit(returns_array)
    
    return {
        'risk_level': 'low' if np.mean(returns) > 0 else 'high',
        'volatility': np.std(returns)
    }
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Spending predictions
- **Savings Calculator** - Future value forecasts
- **Investment Tracker** - Risk assessment
- **Net Worth Tracker** - Growth predictions

## Next Steps

After completing this, move to:
- [10-deployment](10-deployment/README.md) - Deployment
- [capstone-fintech](capstone-fintech/README.md) - FinTech data services