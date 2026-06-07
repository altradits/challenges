# JavaScript Database Mastery Course

## What You'll Learn

This course teaches you **database operations in JavaScript** from beginner to mastery, focusing on building the backend for your FinTech portfolio website.

## Course Structure

### Beginner Level (01-03) - Database Fundamentals
| # | Topic | Key Concepts |
|---|-------|--------------|
| 01 | [Database Basics](01-basics/README.md) | SQL vs NoSQL, tables, queries, connections |
| 02 | [SQLite with Node.js](02-sqlite/README.md) | `better-sqlite3`, queries, prepared statements |
| 03 | [PostgreSQL](03-postgres/README.md) | `pg` library, connection pooling, migrations |

### Intermediate Level (04-06) - ORM and Advanced Patterns
| # | Topic | Key Concepts |
|---|-------|--------------|
| 04 | [Prisma ORM](04-prisma/README.md) | Schema design, migrations, type safety |
| 05 | [MongoDB](05-mongodb/README.md) | Documents, collections, Mongoose ODM |
| 06 | [Redis](06-redis/README.md) | Caching, sessions, rate limiting |

### Advanced Level (07-10) - Production & Optimization
| # | Topic | Key Concepts |
|---|-------|--------------|
| 07 | [Database Design](07-design/README.md) | Normalization, indexing, relationships |
| 08 | [Transactions](08-transactions/README.md) | ACID, isolation levels, rollbacks |
| 09 | [Performance](09-performance/README.md) | Query optimization, connection pooling |
| 10 | [Security](10-security/README.md) | SQL injection, encryption, backups |

## How This Helps Your FinTech Portfolio

### Finance Tools Database Integration
- **Budget Planner** - Store user budgets, categories, transactions
- **Savings Calculator** - Save calculation history, goals
- **Investment Tracker** - Store portfolio data, stock prices
- **Net Worth Tracker** - Track assets, liabilities, history

## Key JavaScript Database Concepts

### Connecting to SQLite
```javascript
const Database = require('better-sqlite3');
const db = new Database('app.db');
```

### Connecting to PostgreSQL
```javascript
const { Pool } = require('pg');
const pool = new Pool({
  host: 'localhost',
  database: 'finance',
  user: 'user',
  password: 'password'
});
```

### Using Prisma
```javascript
// schema.prisma
model User {
  id    Int    @id @default(autoincrement())
  email String @unique
  name  String?
  tasks Task[]
}
```

## Progress Tracking

```
Beginner:     SQL Basics → SQLite → PostgreSQL
              ↓
Intermediate: ORM → MongoDB → Redis
              ↓
Advanced:     Design → Transactions → Performance → Security
              ↓
Capstone:     FinTech Database Integration
```

## Next Steps

After completing this, you'll be ready for:
- [138-js-frontend](../138-js-frontend/README.md) - Js Frontend