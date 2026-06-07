# 03. JavaScript Basics

## What You'll Learn

This exercise teaches you **JavaScript fundamentals for the web**. By the end, you'll understand:
- DOM manipulation
- Event handling
- Fetch API for HTTP requests
- Async/await patterns

## Theory: JavaScript Fundamentals

### DOM Manipulation

```javascript
// Select elements
const form = document.querySelector('#budget-form');
const table = document.querySelector('#budget-table tbody');
const input = document.getElementById('amount');

// Create elements
const row = document.createElement('tr');
const cell = document.createElement('td');
cell.textContent = 'Food';
row.appendChild(cell);

// Add to DOM
table.appendChild(row);
```

### Event Handling

```javascript
// Form submission
form.addEventListener('submit', async (e) => {
  e.preventDefault();
  
  const formData = new FormData(form);
  const data = Object.fromEntries(formData);
  
  // Send to API
  const response = await fetch('/api/budgets', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  
  const result = await response.json();
  console.log(result);
});
```

### Fetch API

```javascript
// GET request
const budgets = await fetch('/api/budgets')
  .then(res => res.json());

// POST request
const createBudget = async (budget) => {
  const response = await fetch('/api/budgets', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(budget)
  });
  return response.json();
};

// Error handling
try {
  const data = await fetch('/api/budgets');
  if (!data.ok) throw new Error('Failed to fetch');
} catch (error) {
  console.error(error);
}
```

## The Challenge

Create interactive JavaScript for the FinTech portfolio.

### Expected Code

```javascript
// app.js
document.addEventListener('DOMContentLoaded', () => {
  const form = document.querySelector('#budget-form');
  const table = document.querySelector('#budget-table tbody');
  
  // Load existing budgets
  loadBudgets();
  
  // Handle form submission
  form.addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const category = form.category.value;
    const amount = parseFloat(form.amount.value);
    
    if (!category || !amount) {
      alert('Please fill all fields');
      return;
    }
    
    // Create budget via API
    const response = await fetch('/api/budgets', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ category, amount })
    });
    
    if (response.ok) {
      form.reset();
      loadBudgets();
    }
  });
});

async function loadBudgets() {
  const response = await fetch('/api/budgets');
  const budgets = await response.json();
  
  const table = document.querySelector('#budget-table tbody');
  table.innerHTML = '';
  
  budgets.forEach(budget => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${budget.category}</td>
      <td>$${budget.amount.toFixed(2)}</td>
      <td>${new Date(budget.date).toLocaleDateString()}</td>
    `;
    table.appendChild(row);
  });
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Add/remove budgets dynamically
- **Savings Calculator** - Real-time calculations
- **Investment Tracker** - Update portfolio values
- **Net Worth Tracker** - Calculate totals

## Next Steps

After completing this, move to:
- [04-es6](04-es6/README.md) - Modern JavaScript features
- [05-buildtools](05-buildtools/README.md) - Build tools