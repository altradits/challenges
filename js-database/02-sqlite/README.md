# 02. SQLite with Node.js

## What You'll Learn

This exercise teaches you **SQLite integration in Node.js**. By the end, you'll understand:
- Using `better-sqlite3` for database operations
- Prepared statements for security
- Query execution and result handling
- Database migrations

## Theory: SQLite in Node.js

### Installation

```bash
npm install better-sqlite3
```

### Basic Operations

```javascript
const Database = require('better-sqlite3');
const db = new Database('app.db');

// Execute a query
const stmt = db.prepare('SELECT * FROM users WHERE id = ?');
const user = stmt.get(1);

// Insert data
const insert = db.prepare('INSERT INTO users (name, email) VALUES (?, ?)');
const info = insert.run('John', 'john@example.com');

// Update data
const update = db.prepare('UPDATE users SET name = ? WHERE id = ?');
update.run('Jane', 1);

// Delete data
const del = db.prepare('DELETE FROM users WHERE id = ?');
del.run(1);
```

## The Challenge

Create a database module for the FinTech portfolio.

### Expected Code

```javascript
// database.js
class FinanceDB {
  constructor(dbPath) {
    this.db = new Database(dbPath);
    this.init();
  }

  init() {
    this.db.exec(`
      CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        name TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
      )
    `);
  }

  createUser(email, name) {
    const stmt = this.db.prepare(
      'INSERT INTO users (email, name) VALUES (?, ?)'
    );
    return stmt.run(email, name);
  }

  getUser(id) {
    const stmt = this.db.prepare('SELECT * FROM users WHERE id = ?');
    return stmt.get(id);
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Store and retrieve budget data
- **Savings Calculator** - Save user calculations
- **Investment Tracker** - Store portfolio history
- **Net Worth Tracker** - Track net worth over time

## Next Steps

After completing this, move to:
- [03-postgres](03-postgres/README.md) - PostgreSQL integration
- [04-prisma](04-prisma/README.md) - ORM patterns