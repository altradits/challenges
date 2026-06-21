# 07. React

## What You'll Learn

This exercise teaches you **React for frontend development**. By the end, you'll understand:
- Components and props
- State and lifecycle
- Hooks (useState, useEffect, useContext)
- React Router

## Theory: React Fundamentals

### Components

```jsx
// BudgetForm.jsx
import { useState } from 'react';

export function BudgetForm({ onAdd }) {
  const [category, setCategory] = useState('');
  const [amount, setAmount] = useState(0);
  
  const handleSubmit = (e) => {
    e.preventDefault();
    onAdd({ category, amount });
    setCategory('');
    setAmount(0);
  };
  
  return (
    <form onSubmit={handleSubmit}>
      <input
        value={category}
        onChange={(e) => setCategory(e.target.value)}
        placeholder="Category"
      />
      <input
        type="number"
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        placeholder="Amount"
      />
      <button type="submit">Add Budget</button>
    </form>
  );
}
```

### Hooks

```jsx
// useApi.js
import { useState, useEffect } from 'react';

export function useApi(url) {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    fetch(url)
      .then(res => res.json())
      .then(setData)
      .catch(setError)
      .finally(() => setLoading(false));
  }, [url]);
  
  return { data, loading, error };
}
```

### React Router

```jsx
// App.jsx
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { BudgetPlanner } from './pages/BudgetPlanner';
import { SavingsCalculator } from './pages/SavingsCalculator';

function App() {
  return (
    <BrowserRouter>
      <nav>
        <Link to="/budget">Budget</Link>
        <Link to="/savings">Savings</Link>
      </nav>
      
      <Routes>
        <Route path="/budget" element={<BudgetPlanner />} />
        <Route path="/savings" element={<SavingsCalculator />} />
      </Routes>
    </BrowserRouter>
  );
}
```

## The Challenge

Build a React app for the FinTech portfolio.

### Expected Code

```jsx
// App.jsx
import { useState, useEffect } from 'react';
import { BudgetList } from './components/BudgetList';
import { BudgetForm } from './components/BudgetForm';

export default function App() {
  const [budgets, setBudgets] = useState([]);
  
  useEffect(() => {
    fetch('/api/budgets')
      .then(res => res.json())
      .then(setBudgets);
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
  
  return (
    <div className="container">
      <h1>FinTech Portfolio</h1>
      <BudgetForm onAdd={addBudget} />
      <BudgetList budgets={budgets} />
    </div>
  );
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Interactive budget forms
- **Savings Calculator** - Real-time calculations
- **Investment Tracker** - Portfolio dashboard
- **Net Worth Tracker** - Asset visualization

## Next Steps

After completing this, move to:
- [08-typescript](08-typescript/README.md) - TypeScript
- [09-testing](09-testing/README.md) - Testing