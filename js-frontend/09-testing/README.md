# 09. Testing

## What You'll Learn

This exercise teaches you **testing frontend applications**. By the end, you'll understand:
- Jest for unit testing
- React Testing Library
- Mocking API calls
- Test coverage

## Theory: Testing in JavaScript

### Jest Basics

```javascript
// math.test.js
import { add, multiply } from './math';

describe('Math functions', () => {
  test('add two numbers', () => {
    expect(add(2, 3)).toBe(5);
  });
  
  test('multiply two numbers', () => {
    expect(multiply(2, 3)).toBe(6);
  });
});
```

### React Testing Library

```jsx
// BudgetForm.test.jsx
import { render, screen, fireEvent } from '@testing-library/react';
import { BudgetForm } from './BudgetForm';

test('submits budget form', () => {
  const mockOnAdd = jest.fn();
  render(<BudgetForm onAdd={mockOnAdd} />);
  
  fireEvent.change(screen.getByPlaceholderText('Category'), {
    target: { value: 'Food' }
  });
  
  fireEvent.click(screen.getByText('Add Budget'));
  
  expect(mockOnAdd).toHaveBeenCalledWith({
    category: 'Food',
    amount: expect.any(Number)
  });
});
```

### Mocking API Calls

```javascript
// api.test.js
import { fetchBudgets } from './api';

global.fetch = jest.fn();

test('fetches budgets from API', async () => {
  const mockBudgets = [
    { id: 1, category: 'Food', amount: 100 }
  ];
  
  fetch.mockResolvedValue({
    ok: true,
    json: () => Promise.resolve(mockBudgets)
  });
  
  const budgets = await fetchBudgets();
  expect(budgets).toEqual(mockBudgets);
});
```

## The Challenge

Write tests for the FinTech frontend.

### Expected Code

```javascript
// budget-utils.test.js
import { calculateTotal, filterByCategory } from './budget-utils';

describe('Budget utilities', () => {
  const budgets = [
    { id: 1, category: 'food', amount: 100 },
    { id: 2, category: 'rent', amount: 500 },
    { id: 3, category: 'food', amount: 50 }
  ];
  
  test('calculate total', () => {
    expect(calculateTotal(budgets)).toBe(650);
  });
  
  test('filter by category', () => {
    const food = filterByCategory(budgets, 'food');
    expect(food).toHaveLength(2);
    expect(food[0].category).toBe('food');
  });
});
```

```jsx
// BudgetList.test.jsx
import { render, screen } from '@testing-library/react';
import { BudgetList } from './BudgetList';

test('renders budget list', () => {
  const budgets = [
    { id: 1, category: 'Food', amount: 100 },
    { id: 2, category: 'Rent', amount: 500 }
  ];
  
  render(<BudgetList budgets={budgets} />);
  
  expect(screen.getByText('Food')).toBeInTheDocument();
  expect(screen.getByText('Rent')).toBeInTheDocument();
});
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Test budget calculations
- **Savings Calculator** - Test savings formulas
- **Investment Tracker** - Test portfolio logic
- **Net Worth Tracker** - Test asset calculations

## Next Steps

After completing this, move to:
- [10-deployment](10-deployment/README.md) - Deployment
- [capstone-fintech](capstone-fintech/README.md) - FinTech portfolio