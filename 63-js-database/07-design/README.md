# 07. Database Design

## What You'll Learn

This exercise teaches you **database design patterns**. By the end, you'll understand:
- Normalization
- Indexing
- Relationships
- Schema optimization

## Theory: Database Design

### Normalization

```sql
-- Good: Normalized (3NF)
users: id, email, name
budgets: id, user_id, category, amount, date
transactions: id, user_id, amount, type, date

-- Bad: Denormalized
users: id, email, name, budgets_json, transactions_json
```

### Indexing

```sql
-- Create index for fast lookups
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_transactions_user_date ON transactions(user_id, date);

-- Composite index
CREATE INDEX idx_budgets_user_category ON budgets(user_id, category);
```

### Relationships

```sql
-- One-to-Many
users (1) → budgets (many)

-- Many-to-Many
users ↔ categories (through user_categories table)
```

## The Challenge

Design a database schema for the FinTech portfolio.

### Expected Schema

```sql
-- Users table
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  name TEXT,
  password_hash TEXT,
  created_at TIMESTAMP
);

-- Budgets table
CREATE TABLE budgets (
  id INTEGER PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  category TEXT,
  amount REAL,
  month INTEGER,
  year INTEGER
);

-- Investments table
CREATE TABLE investments (
  id INTEGER PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  symbol TEXT,
  shares REAL,
  purchase_price REAL,
  purchase_date DATE
);

-- Indexes
CREATE INDEX idx_budgets_user ON budgets(user_id);
CREATE INDEX idx_investments_user ON investments(user_id);
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Efficient budget queries
- **Savings Calculator** - Fast savings goal lookups
- **Investment Tracker** - Portfolio performance queries
- **Net Worth Tracker** - Asset aggregation

## Next Steps

After completing this, move to:
- [08-transactions](08-transactions/README.md) - ACID transactions
- [09-performance](09-performance/README.md) - Query optimization