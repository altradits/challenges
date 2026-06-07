# 04. ES6+ Features

## What You'll Learn

This exercise teaches you **modern JavaScript features**. By the end, you'll understand:
- Modules and imports/exports
- Arrow functions
- Destructuring
- Async/await
- Classes

## Theory: ES6+ Features

### Modules

```javascript
// math.js
export const add = (a, b) => a + b;
export const multiply = (a, b) => a * b;

// app.js
import { add, multiply } from './math.js';
```

### Arrow Functions

```javascript
// Traditional
const add = function(a, b) { return a + b; };

// Arrow
const add = (a, b) => a + b;

// With object
const createBudget = (category, amount) => ({
  category,
  amount,
  date: new Date()
});
```

### Destructuring

```javascript
// Object destructuring
const { category, amount } = budget;

// Array destructuring
const [first, second] = budgets;

// Function parameters
function processBudget({ category, amount }) {
  console.log(category, amount);
}
```

### Async/Await

```javascript
// Promise chain
fetch('/api/budgets')
  .then(res => res.json())
  .then(data => console.log(data));

// Async/await
const loadBudgets = async () => {
  const response = await fetch('/api/budgets');
  const budgets = await response.json();
  return budgets;
};
```

### Classes

```javascript
class BudgetManager {
  constructor(apiUrl) {
    this.apiUrl = apiUrl;
    this.budgets = [];
  }
  
  async load() {
    const response = await fetch(`${this.apiUrl}/budgets`);
    this.budgets = await response.json();
  }
  
  async add(category, amount) {
    const budget = { category, amount, date: new Date() };
    const response = await fetch(`${this.apiUrl}/budgets`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(budget)
    });
    this.budgets.push(await response.json());
  }
}
```

## The Challenge

Refactor the FinTech app with modern JavaScript.

### Expected Code

```javascript
// api.js
export const api = {
  get: async (endpoint) => {
    const res = await fetch(endpoint);
    return res.json();
  },
  
  post: async (endpoint, data) => {
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    });
    return res.json();
  }
};

// budget-manager.js
export class BudgetManager {
  constructor(api) {
    this.api = api;
    this.budgets = [];
  }
  
  async loadAll() {
    this.budgets = await this.api.get('/api/budgets');
    return this.budgets;
  }
  
  async create({ category, amount }) {
    const budget = await this.api.post('/api/budgets', {
      category,
      amount,
      date: new Date().toISOString()
    });
    this.budgets.push(budget);
    return budget;
  }
  
  getTotal() {
    return this.budgets.reduce((sum, b) => sum + b.amount, 0);
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Modular budget management
- **Savings Calculator** - Clean calculation logic
- **Investment Tracker** - Portfolio classes
- **Net Worth Tracker** - Asset management

## Next Steps

After completing this, move to:
- [05-buildtools](05-buildtools/README.md) - Build tools
- [06-state](06-state/README.md) - State management