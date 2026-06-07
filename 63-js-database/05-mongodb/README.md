# 05. MongoDB

## What You'll Learn

This exercise teaches you **MongoDB with Mongoose** in Node.js. By the end, you'll understand:
- Document-based storage
- Mongoose schemas and models
- CRUD operations
- Aggregation pipelines

## Theory: MongoDB in Node.js

### Installation

```bash
npm install mongoose
```

### Schema Design

```javascript
const mongoose = require('mongoose');

const userSchema = new mongoose.Schema({
  email: { type: String, required: true, unique: true },
  name: String,
  budgets: [{
    category: String,
    amount: Number,
    date: { type: Date, default: Date.now }
  }],
  savings: [{
    goal: String,
    target: Number,
    current: Number
  }]
});

const User = mongoose.model('User', userSchema);
```

### Basic Operations

```javascript
// Connect
await mongoose.connect('mongodb://localhost:27017/finance');

// Create
const user = await User.create({ email: 'john@example.com', name: 'John' });

// Read
const users = await User.find({});
const user = await User.findById(1);

// Update
await User.findByIdAndUpdate(1, { name: 'Jane' });

// Delete
await User.findByIdAndDelete(1);
```

## The Challenge

Create a MongoDB schema for the FinTech portfolio.

### Expected Code

```javascript
// models.js
const budgetSchema = new mongoose.Schema({
  category: { type: String, required: true },
  amount: { type: Number, required: true },
  date: { type: Date, default: Date.now }
});

const investmentSchema = new mongoose.Schema({
  symbol: String,
  shares: Number,
  purchasePrice: Number,
  currentPrice: Number
});

const userSchema = new mongoose.Schema({
  email: { type: String, unique: true },
  name: String,
  budgets: [budgetSchema],
  investments: [investmentSchema]
});
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Flexible budget structure
- **Savings Calculator** - Dynamic savings goals
- **Investment Tracker** - Portfolio with varying assets
- **Net Worth Tracker** - Mixed asset types

## Next Steps

After completing this, move to:
- [06-redis](06-redis/README.md) - Caching and rate limiting
- [07-design](07-design/README.md) - Database design patterns