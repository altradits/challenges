# 10. Deployment

## What You'll Learn

This exercise teaches you **deploying Python applications**. By the end, you'll understand:
- Flask for web APIs
- FastAPI for modern APIs
- Docker containerization
- Environment configuration

## Theory: Deployment

### Flask

```python
from flask import Flask, jsonify, request

app = Flask(__name__)

@app.route('/api/budgets', methods=['GET'])
def get_budgets():
    return jsonify(budgets)

@app.route('/api/budgets', methods=['POST'])
def create_budget():
    budget = request.json
    budgets.append(budget)
    return jsonify(budget), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
```

### FastAPI

```python
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class Budget(BaseModel):
    category: str
    amount: float

@app.get('/api/budgets')
async def get_budgets():
    return budgets

@app.post('/api/budgets')
async def create_budget(budget: Budget):
    budgets.append(budget.dict())
    return budget
```

### Docker

```dockerfile
FROM python:3.11-slim

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
EXPOSE 5000

CMD ['python', 'app.py']
```

## The Challenge

Deploy a Python API for FinTech.

### Expected Code

```python
# main.py
from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn

app = FastAPI(title="FinTech Data API")

class Budget(BaseModel):
    category: str
    amount: float
    date: str

budgets = []

@app.get('/budgets')
async def get_budgets():
    return budgets

@app.post('/budgets')
async def create_budget(budget: Budget):
    budgets.append(budget.dict())
    return budget

if __name__ == '__main__':
    uvicorn.run(app, host='0.0.0.0', port=5000)
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Data API
- **Savings Calculator** - Calculation service
- **Investment Tracker** - Stock data API
- **Net Worth Tracker** - Analytics API

## Next Steps

After completing this, you've mastered Python for data!
Build your complete FinTech portfolio:
- [FinTech Data Services Capstone](capstone-fintech/README.md)
- [Professional Presence](professional/README.md)