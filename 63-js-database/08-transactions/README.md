# 08. Transactions

## What You'll Learn

This exercise teaches you **database transactions**. By the end, you'll understand:
- ACID properties
- Transaction isolation
- Rollback handling
- Concurrent access

## Theory: Transactions

### ACID Properties

- **Atomicity** - All or nothing
- **Consistency** - Valid state before and after
- **Isolation** - Concurrent transactions don't interfere
- **Durability** - Committed data persists

### Transaction in PostgreSQL

```javascript
const { Pool } = require('pg');
const pool = new Pool();

// Transaction example
const transferMoney = async (fromId, toId, amount) => {
  const client = await pool.connect();
  try {
    await client.query('BEGIN');
    
    // Debit from account
    await client.query(
      'UPDATE accounts SET balance = balance - $1 WHERE id = $2',
      [amount, fromId]
    );
    
    // Credit to account
    await client.query(
      'UPDATE accounts SET balance = balance + $1 WHERE id = $2',
      [amount, toId]
    );
    
    await client.query('COMMIT');
  } catch (error) {
    await client.query('ROLLBACK');
    throw error;
  } finally {
    client.release();
  }
};
```

### Transaction in SQLite

```javascript
const db = new Database('app.db');

const transferMoney = (fromId, toId, amount) => {
  const transaction = db.transaction((transfer) => {
    transfer.prepare(
      'UPDATE accounts SET balance = balance - ? WHERE id = ?'
    ).run(amount, fromId);
    
    transfer.prepare(
      'UPDATE accounts SET balance = balance + ? WHERE id = ?'
    ).run(amount, toId);
  });
  
  return transaction();
};
```

## The Challenge

Implement transaction handling for the FinTech portfolio.

### Expected Code

```javascript
// transactions.js
class FinanceTransactions {
  async createBudgetWithTransaction(userId, category, amount) {
    const client = await pool.connect();
    try {
      await client.query('BEGIN');
      
      // Create budget
      const result = await client.query(
        'INSERT INTO budgets (user_id, category, amount) VALUES ($1, $2, $3) RETURNING *',
        [userId, category, amount]
      );
      
      // Log transaction
      await client.query(
        'INSERT INTO budget_logs (budget_id, action) VALUES ($1, $2)',
        [result.rows[0].id, 'created']
      );
      
      await client.query('COMMIT');
      return result.rows[0];
    } catch (error) {
      await client.query('ROLLBACK');
      throw error;
    } finally {
      client.release();
    }
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Atomic budget updates
- **Savings Calculator** - Consistent savings operations
- **Investment Tracker** - Portfolio transactions
- **Net Worth Tracker** - Asset transfers

## Next Steps

After completing this, move to:
- [09-performance](09-performance/README.md) - Query optimization
- [10-security](10-security/README.md) - Security best practices