# 03. PostgreSQL

## What You'll Learn

This exercise teaches you **PostgreSQL integration in Node.js**. By the end, you'll understand:
- Using `pg` library for PostgreSQL
- Connection pooling
- Async/await patterns
- Transactions

## Theory: PostgreSQL in Node.js

### Installation

```bash
npm install pg
```

### Basic Operations

```javascript
const { Pool } = require('pg');

const pool = new Pool({
  host: 'localhost',
  database: 'finance',
  user: 'postgres',
  password: 'password',
  port: 5432
});

// Query with async/await
const getUsers = async () => {
  const result = await pool.query('SELECT * FROM users');
  return result.rows;
};

// Insert with parameters
const createUser = async (email, name) => {
  const result = await pool.query(
    'INSERT INTO users (email, name) VALUES ($1, $2) RETURNING *',
    [email, name]
  );
  return result.rows[0];
};
```

## The Challenge

Create a PostgreSQL database module for production use.

### Expected Code

```javascript
// pg-database.js
class PostgresFinanceDB {
  constructor(config) {
    this.pool = new Pool(config);
  }

  async query(text, params) {
    const result = await this.pool.query(text, params);
    return result.rows;
  }

  async createUser(email, name) {
    return await this.query(
      'INSERT INTO users (email, name) VALUES ($1, $2) RETURNING *',
      [email, name]
    );
  }

  async getUser(id) {
    const users = await this.query('SELECT * FROM users WHERE id = $1', [id]);
    return users[0];
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Production database for user data
- **Savings Calculator** - Scalable storage for calculations
- **Investment Tracker** - Handle large portfolio datasets
- **Net Worth Tracker** - Concurrent user access

## Next Steps

After completing this, move to:
- [04-prisma](04-prisma/README.md) - ORM patterns
- [05-mongodb](05-mongodb/README.md) - NoSQL databases