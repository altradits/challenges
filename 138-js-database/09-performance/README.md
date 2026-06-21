# 09. Performance

## What You'll Learn

This exercise teaches you **database performance optimization**. By the end, you'll understand:
- Query optimization
- Connection pooling
- Caching strategies
- Index optimization

## Theory: Performance Optimization

### Connection Pooling

```javascript
const { Pool } = require('pg');

const pool = new Pool({
  host: 'localhost',
  database: 'finance',
  user: 'postgres',
  password: 'password',
  // Performance settings
  max: 20,           // Max connections
  idleTimeoutMillis: 30000,
  connectionTimeoutMillis: 2000,
});
```

### Query Optimization

```javascript
// Bad: N+1 queries
const users = await db.query('SELECT * FROM users');
for (const user of users) {
  user.budgets = await db.query(
    'SELECT * FROM budgets WHERE user_id = ?', [user.id]
  );
}

// Good: Single query with JOIN
const usersWithBudgets = await db.query(`
  SELECT u.*, b.* 
  FROM users u 
  LEFT JOIN budgets b ON u.id = b.user_id
`);
```

### Caching Strategy

```javascript
const redis = require('redis');
const client = redis.createClient();

// Cache-aside pattern
async function getCachedUser(id) {
  const cached = await client.get(`user:${id}`);
  if (cached) return JSON.parse(cached);
  
  const user = await db.query('SELECT * FROM users WHERE id = ?', [id]);
  await client.setEx(`user:${id}`, 3600, JSON.stringify(user));
  return user;
}
```

## The Challenge

Optimize database queries for the FinTech portfolio.

### Expected Code

```javascript
// performance.js
class OptimizedFinanceDB {
  constructor(pool, cache) {
    this.pool = pool;
    this.cache = cache;
  }

  // Paginated query
  async getBudgets(userId, page = 1, limit = 20) {
    const offset = (page - 1) * limit;
    const result = await this.pool.query(
      'SELECT * FROM budgets WHERE user_id = $1 ORDER BY date DESC LIMIT $2 OFFSET $3',
      [userId, limit, offset]
    );
    return result.rows;
  }

  // Batch insert
  async createMultipleBudgets(budgets) {
    const query = `
      INSERT INTO budgets (user_id, category, amount, date) 
      VALUES ${budgets.map((_, i) => `($${i*4+1}, $${i*4+2}, $${i*4+3}, $${i*4+4})`).join(',')}
    `;
    const values = budgets.flatMap(b => [b.userId, b.category, b.amount, b.date]);
    return await this.pool.query(query, values);
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Fast budget loading
- **Savings Calculator** - Efficient calculations
- **Investment Tracker** - Quick portfolio queries
- **Net Worth Tracker** - Fast asset aggregation

## Next Steps

After completing this, move to:
- [10-security](10-security/README.md) - Security best practices
- [capstone-fintech](capstone-fintech/README.md) - FinTech database integration