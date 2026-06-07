# 08. TypeScript

## What You'll Learn

This exercise teaches you **TypeScript for type-safe frontend**. By the end, you'll understand:
- Type annotations
- Interfaces
- Generics
- Type inference

## Theory: TypeScript Fundamentals

### Type Annotations

```typescript
// Basic types
const name: string = 'John';
const age: number = 30;
const active: boolean = true;

// Arrays
const budgets: Budget[] = [];
const amounts: number[] = [100, 200, 300];

// Objects
interface User {
  id: number;
  email: string;
  name: string;
}

const user: User = {
  id: 1,
  email: 'john@example.com',
  name: 'John'
};
```

### Interfaces

```typescript
interface Budget {
  id: number;
  category: string;
  amount: number;
  date: Date;
}

interface BudgetFormProps {
  onAdd: (budget: Budget) => void;
}

function BudgetForm({ onAdd }: BudgetFormProps) {
  const [category, setCategory] = useState<string>('');
  const [amount, setAmount] = useState<number>(0);
  
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onAdd({ id: Date.now(), category, amount, date: new Date() });
  };
}
```

### Generics

```typescript
interface ApiResponse<T> {
  data: T;
  error?: string;
}

async function fetchBudgets(): Promise<ApiResponse<Budget[]>> {
  const res = await fetch('/api/budgets');
  return res.json();
}

// Custom hook with generics
function useApi<T>(url: string): {
  data: T | null;
  loading: boolean;
  error: string | null;
} {
  const [data, setData] = useState<T | null>(null);
  // ...
}
```

## The Challenge

Convert the FinTech app to TypeScript.

### Expected Code

```typescript
// types.ts
export interface Budget {
  id: number;
  category: string;
  amount: number;
  date: string;
}

export interface SavingsGoal {
  id: number;
  goal: string;
  target: number;
  current: number;
}

export interface Investment {
  id: number;
  symbol: string;
  shares: number;
  purchasePrice: number;
  currentPrice: number;
}

// api.ts
export async function fetchBudgets(): Promise<Budget[]> {
  const res = await fetch('/api/budgets');
  if (!res.ok) throw new Error('Failed to fetch');
  return res.json();
}

// BudgetForm.tsx
import { useState, FormEvent } from 'react';
import { Budget } from '../types';

interface Props {
  onAdd: (budget: Budget) => void;
}

export function BudgetForm({ onAdd }: Props) {
  const [category, setCategory] = useState<string>('');
  const [amount, setAmount] = useState<number>(0);
  
  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onAdd({
      id: Date.now(),
      category,
      amount,
      date: new Date().toISOString()
    });
  };
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Type-safe budget data
- **Savings Calculator** - Safe calculations
- **Investment Tracker** - Portfolio types
- **Net Worth Tracker** - Asset types

## Next Steps

After completing this, move to:
- [09-testing](09-testing/README.md) - Testing
- [10-deployment](10-deployment/README.md) - Deployment