# 01. Database Basics

## What You'll Learn

This exercise teaches you **database fundamentals** in JavaScript. By the end, you'll understand:
- SQL vs NoSQL databases
- Tables, rows, and columns
- Basic SQL queries (SELECT, INSERT, UPDATE, DELETE)
- Database connections

## Theory: Database Fundamentals

### SQL vs NoSQL

| SQL | NoSQL |
|-----|-------|
| Structured data | Flexible schema |
| Tables with rows/columns | Documents, key-value, graphs |
| ACID transactions | Eventual consistency |
| PostgreSQL, MySQL, SQLite | MongoDB, Redis, DynamoDB |

### Basic SQL Queries

```sql
-- Create table
CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE
);

-- Insert data
INSERT INTO users (name, email) VALUES ('John', 'john@example.com');

-- Select data
SELECT * FROM users WHERE id = 1;

-- Update data
UPDATE users SET name = 'Jane' WHERE id = 1;

-- Delete data
DELETE FROM users WHERE id = 1;
```

## The Challenge

Create a simple database setup for a FinTech application.

### Expected Code

```javascript
// Create a database connection and table
const db = new Database('finance.db');

// Create users table
db.exec(`
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    name TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
  )
`);

// Create transactions table
db.exec(`
  CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    amount REAL,
    category TEXT,
    date DATE,
    FOREIGN KEY (user_id) REFERENCES users(id)
  )
`);
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Store user budgets and transactions
- **Savings Calculator** - Save calculation history
- **Investment Tracker** - Store portfolio data
- **Net Worth Tracker** - Track assets and liabilities

## Next Steps

After completing this, move to:
- [02-sqlite](02-sqlite/README.md) - SQLite with Node.js
- [03-postgres](03-postgres/README.md) - PostgreSQL integration