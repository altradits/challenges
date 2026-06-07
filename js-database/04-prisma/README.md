# 04. Prisma ORM

## What You'll Learn

This exercise teaches you **Prisma ORM** for type-safe database operations. By the end, you'll understand:
- Prisma schema design
- Migrations
- Type-safe queries
- Relations

## Theory: Prisma ORM

### Installation

```bash
npm install prisma
npx prisma init
```

### Schema Design

```prisma
// schema.prisma
model User {
  id        Int      @id @default(autoincrement())
  email     String   @unique
  name      String?
  tasks     Task[]
  createdAt DateTime @default(now())
}

model Task {
  id          Int     @id @default(autoincrement())
  title       String
  description String?
  completed   Boolean @default(false)
  user        User    @relation(fields: [userId], references: [id])
  userId      Int
}
```

### Basic Operations

```javascript
const { PrismaClient } = require('@prisma/client');
const prisma = new PrismaClient();

// Create
const user = await prisma.user.create({
  data: { email: 'john@example.com', name: 'John' }
});

// Read
const users = await prisma.user.findMany();
const user = await prisma.user.findUnique({ where: { id: 1 } });

// Update
const updated = await prisma.user.update({
  where: { id: 1 },
  data: { name: 'Jane' }
});

// Delete
await prisma.user.delete({ where: { id: 1 } });
```

## The Challenge

Create a Prisma schema for the FinTech portfolio.

### Expected Code

```prisma
// schema.prisma
model User {
  id        Int      @id @default(autoincrement())
  email     String   @unique
  name      String?
  budgets   Budget[]
  savings   Savings[]
  investments Investment[]
  netWorth  NetWorth[]
  createdAt DateTime @default(now())
}

model Budget {
  id       Int    @id @default(autoincrement())
  category String
  amount   Float
  user     User   @relation(fields: [userId], references: [id])
  userId   Int
}

model Savings {
  id       Int    @id @default(autoincrement())
  goal     String
  target   Float
  current  Float
  user     User   @relation(fields: [userId], references: [id])
  userId   Int
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Type-safe budget operations
- **Savings Calculator** - Safe savings goal tracking
- **Investment Tracker** - Portfolio data with relations
- **Net Worth Tracker** - Asset/liability relationships

## Next Steps

After completing this, move to:
- [05-mongodb](05-mongodb/README.md) - NoSQL with Mongoose
- [06-redis](06-redis/README.md) - Caching and sessions