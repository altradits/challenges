# 06. Redis

## What You'll Learn

This exercise teaches you **Redis for caching and sessions**. By the end, you'll understand:
- Key-value storage
- Caching patterns
- Session management
- Rate limiting

## Theory: Redis in Node.js

### Installation

```bash
npm install redis
```

### Basic Operations

```javascript
const redis = require('redis');
const client = redis.createClient();
await client.connect();

// Set and get
await client.set('user:1', JSON.stringify(user));
const data = await client.get('user:1');
const user = JSON.parse(data);

// TTL (expiration)
await client.setEx('session:abc', 3600, userId); // 1 hour

// Lists
await client.lPush('tasks', JSON.stringify(task));
const tasks = await client.lRange('tasks', 0, -1);

// Sets
await client.sAdd('user:1:categories', 'food', 'rent', 'utilities');
```

## The Challenge

Create a Redis cache for the FinTech portfolio.

### Expected Code

```javascript
// cache.js
class FinanceCache {
  constructor(redisClient) {
    this.client = redisClient;
  }

  async cacheUser(userId, userData) {
    await this.client.setEx(
      `user:${userId}`,
      3600,
      JSON.stringify(userData)
    );
  }

  async getCachedUser(userId) {
    const data = await this.client.get(`user:${userId}`);
    return data ? JSON.parse(data) : null;
  }

  async rateLimit(ip, limit = 100, window = 60) {
    const key = `rate:${ip}`;
    const current = await this.client.incr(key);
    if (current === 1) {
      await this.client.expire(key, window);
    }
    return current > limit;
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Cache user budgets for fast access
- **Savings Calculator** - Rate limit API calls
- **Investment Tracker** - Cache stock prices
- **Net Worth Tracker** - Session management

## Next Steps

After completing this, move to:
- [07-design](07-design/README.md) - Database design
- [08-transactions](08-transactions/README.md) - Transactions