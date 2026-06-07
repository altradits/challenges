# 06. State Management

## What You'll Learn

This exercise teaches you **state management in JavaScript**. By the end, you'll understand:
- LocalStorage for persistence
- Context API patterns
- State management libraries
- Global state patterns

## Theory: State Management

### LocalStorage

```javascript
// Save state
localStorage.setItem('budgets', JSON.stringify(budgets));

// Load state
const budgets = JSON.parse(localStorage.getItem('budgets') || '[]');

// Remove state
localStorage.removeItem('budgets');
```

### Context Pattern

```javascript
// context.js
class FinanceContext {
  constructor() {
    this.state = {
      user: null,
      budgets: [],
      savings: []
    };
    this.listeners = [];
  }
  
  subscribe(listener) {
    this.listeners.push(listener);
    return () => {
      this.listeners = this.listeners.filter(l => l !== listener);
    };
  }
  
  setState(newState) {
    this.state = { ...this.state, ...newState };
    this.listeners.forEach(listener => listener(this.state));
  }
  
  getState() {
    return this.state;
  }
}

export const context = new FinanceContext();
```

### Custom Hook Pattern

```javascript
// useBudgets.js
import { useState, useEffect } from 'react';

export function useBudgets() {
  const [budgets, setBudgets] = useState([]);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    fetch('/api/budgets')
      .then(res => res.json())
      .then(setBudgets)
      .finally(() => setLoading(false));
  }, []);
  
  const addBudget = async (budget) => {
    const response = await fetch('/api/budgets', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(budget)
    });
    const newBudget = await response.json();
    setBudgets([...budgets, newBudget]);
  };
  
  return { budgets, loading, addBudget };
}
```

## The Challenge

Implement state management for the FinTech portfolio.

### Expected Code

```javascript
// store.js
class FinanceStore {
  constructor() {
    this.state = {
      user: this.loadFromStorage('user'),
      budgets: this.loadFromStorage('budgets') || [],
      savings: this.loadFromStorage('savings') || []
    };
  }
  
  loadFromStorage(key) {
    try {
      return JSON.parse(localStorage.getItem(key));
    } catch {
      return null;
    }
  }
  
  saveToStorage(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
  }
  
  setUser(user) {
    this.state.user = user;
    this.saveToStorage('user', user);
    this.notify();
  }
  
  addBudget(budget) {
    this.state.budgets.push(budget);
    this.saveToStorage('budgets', this.state.budgets);
    this.notify();
  }
  
  notify() {
    window.dispatchEvent(new CustomEvent('stateChange', {
      detail: this.state
    }));
  }
}

export const store = new FinanceStore();
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Persist budgets across sessions
- **Savings Calculator** - Save goals and progress
- **Investment Tracker** - Cache portfolio data
- **Net Worth Tracker** - Track history

## Next Steps

After completing this, move to:
- [07-react](07-react/README.md) - React framework
- [08-typescript](08-typescript/README.md) - TypeScript