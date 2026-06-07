# 10. Deployment

## What You'll Learn

This exercise teaches you **deploying frontend applications**. By the end, you'll understand:
- Netlify deployment
- Vercel deployment
- CI/CD with GitHub Actions
- Environment variables
- Custom domains

## Theory: Deployment

### Netlify

```bash
# Install Netlify CLI
npm install -g netlify-cli

# Deploy
netlify deploy --dir=dist
```

### Vercel

```bash
# Install Vercel CLI
npm install -g vercel

# Deploy
vercel --prod
```

### GitHub Actions

```yaml
# .github/workflows/deploy.yml
name: Deploy
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-node@v3
        with:
          node-version: '18'
      - run: npm ci
      - run: npm run build
      - uses: netlify/actions/cli@master
        with:
          args: deploy --dir=dist --prod
        env:
          NETLIFY_AUTH_TOKEN: ${{ secrets.NETLIFY_TOKEN }}
```

### Environment Variables

```bash
# .env
VITE_API_URL=https://api.yourdomain.com
VITE_GOOGLE_ANALYTICS=GA-XXXXX
```

```javascript
// Access in code
const apiUrl = import.meta.env.VITE_API_URL;
```

## The Challenge

Deploy the FinTech portfolio to production.

### Expected Code

```javascript
// vite.config.js
import { defineConfig } from 'vite';

export default defineConfig({
  define: {
    'process.env.API_URL': JSON.stringify(process.env.API_URL)
  },
  build: {
    outDir: 'dist',
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
          charts: ['chart.js']
        }
      }
    }
  }
});
```

```json
// netlify.toml
[build]
  publish = "dist"
  command = "npm run build"

[[redirects]]
  from = "/api/*"
  to = "https://api.yourdomain.com/api/:splat"
  status = 200
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Live budget tool
- **Savings Calculator** - Public calculator
- **Investment Tracker** - Portfolio dashboard
- **Net Worth Tracker** - Personal finance tool

## Next Steps

After completing this, you've mastered frontend development!
Build your complete FinTech portfolio:
- [FinTech Portfolio Capstone](capstone-fintech/README.md)
- [Python for Data](python-data/README.md)
- [Professional Presence](professional/README.md)