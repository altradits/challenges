# 07. API Integration

## What You'll Learn

This exercise teaches you **API integration in Python**. By the end, you'll understand:
- HTTP requests with `requests`
- REST API consumption
- Authentication
- Error handling

## Theory: API Integration

### Installation

```bash
pip install requests
```

### HTTP Requests

```python
import requests

# GET request
response = requests.get('https://api.example.com/budgets')
budgets = response.json()

# POST request
response = requests.post(
    'https://api.example.com/budgets',
    json={'category': 'food', 'amount': 100}
)

# With headers
headers = {'Authorization': 'Bearer token'}
response = requests.get('https://api.example.com/profile', headers=headers)
```

### Error Handling

```python
try:
    response = requests.get('https://api.example.com/budgets')
    response.raise_for_status()
    budgets = response.json()
except requests.exceptions.HTTPError as e:
    print(f"HTTP error: {e}")
except requests.exceptions.RequestException as e:
    print(f"Request error: {e}")
```

## The Challenge

Create API integration for FinTech data.

### Expected Code

```python
import requests

class FinanceAPI:
    def __init__(self, base_url, token=None):
        self.base_url = base_url
        self.headers = {'Authorization': f'Bearer {token}'} if token else {}
    
    def get_budgets(self):
        response = requests.get(f'{self.base_url}/budgets', headers=self.headers)
        response.raise_for_status()
        return response.json()
    
    def create_budget(self, category, amount):
        response = requests.post(
            f'{self.base_url}/budgets',
            json={'category': category, 'amount': amount},
            headers=self.headers
        )
        response.raise_for_status()
        return response.json()
    
    def get_stock_price(self, symbol):
        response = requests.get(f'https://api.example.com/stocks/{symbol}')
        return response.json()
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Sync with backend
- **Savings Calculator** - Save to API
- **Investment Tracker** - Fetch stock data
- **Net Worth Tracker** - Update assets

## Next Steps

After completing this, move to:
- [08-analysis](08-analysis/README.md) - Data analysis
- [09-ml](09-ml/README.md) - Machine learning