# 10. Security

## What You'll Learn

This exercise teaches you **database security best practices**. By the end, you'll understand:
- SQL injection prevention
- Data encryption
- Backup strategies
- Access control

## Theory: Database Security

### SQL Injection Prevention

```javascript
// Bad: Vulnerable to SQL injection
const query = `SELECT * FROM users WHERE email = '${email}'`;

// Good: Parameterized queries
const result = await pool.query(
  'SELECT * FROM users WHERE email = $1',
  [email]
);

// Also good: ORM (Prisma)
const user = await prisma.user.findUnique({
  where: { email }
});
```

### Data Encryption

```javascript
const crypto = require('crypto');

// Encrypt sensitive data
const encrypt = (text, key) => {
  const iv = crypto.randomBytes(16);
  const cipher = crypto.createCipheriv('aes-256-cbc', key, iv);
  let encrypted = cipher.update(text, 'utf8', 'hex');
  encrypted += cipher.final('hex');
  return { iv: iv.toString('hex'), encrypted };
};

// Decrypt
const decrypt = (data, key) => {
  const iv = Buffer.from(data.iv, 'hex');
  const decipher = crypto.createDecipheriv('aes-256-cbc', key, iv);
  let decrypted = decipher.update(data.encrypted, 'hex', 'utf8');
  decrypted += decipher.final('utf8');
  return decrypted;
};
```

### Backup Strategy

```bash
# PostgreSQL backup
pg_dump -U postgres finance > backup.sql

# Automated backup script
#!/bin/bash
DATE=$(date +%Y%m%d)
pg_dump -U postgres finance | gzip > backups/finance_$DATE.sql.gz
```

## The Challenge

Implement security measures for the FinTech database.

### Expected Code

```javascript
// security.js
class SecureFinanceDB {
  constructor(pool, encryptionKey) {
    this.pool = pool;
    this.key = encryptionKey;
  }

  // Secure user creation
  async createUser(email, name, ssn) {
    const encryptedSSN = encrypt(ssn, this.key);
    
    return await this.pool.query(
      'INSERT INTO users (email, name, ssn_encrypted) VALUES ($1, $2, $3) RETURNING id, email, name',
      [email, name, JSON.stringify(encryptedSSN)]
    );
  }

  // Input validation
  validateInput(data) {
    if (!data.email || !data.email.includes('@')) {
      throw new Error('Invalid email');
    }
    if (!data.amount || isNaN(data.amount)) {
      throw new Error('Invalid amount');
    }
    return true;
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Secure user data
- **Savings Calculator** - Protect financial data
- **Investment Tracker** - Secure portfolio info
- **Net Worth Tracker** - Protect asset data

## Next Steps

After completing this, you've mastered database operations!
Now build your FinTech portfolio with:
- [FinTech Database Capstone](capstone-fintech/README.md)
- [JavaScript Frontend](js-frontend/README.md)
- [Python for Data](python-data/README.md)