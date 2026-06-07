# 05. Build Tools

## What You'll Learn

This exercise teaches you **modern build tools**. By the end, you'll understand:
- Vite for fast development
- npm scripts
- Bundling and minification
- Development vs production

## Theory: Build Tools

### Vite Configuration

```javascript
// vite.config.js
import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  },
  build: {
    outDir: 'dist',
    minify: 'terser'
  }
});
```

### npm Scripts

```json
{
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview",
    "test": "jest",
    "lint": "eslint src --fix"
  }
}
```

### Project Structure

```
fintech-frontend/
├── index.html
├── vite.config.js
├── package.json
├── src/
│   ├── main.js
│   ├── api.js
│   ├── components/
│   │   ├── BudgetForm.js
│   │   └── BudgetList.js
│   └── styles/
│       └── main.css
└── dist/
```

## The Challenge

Set up a build system for the FinTech portfolio.

### Expected Code

```javascript
// vite.config.js
import { defineConfig } from 'vite';

export default defineConfig({
  root: 'src',
  publicDir: '../public',
  build: {
    outDir: '../dist',
    emptyOutDir: true,
    rollupOptions: {
      input: {
        main: 'src/index.html',
        budget: 'src/budget.html',
        savings: 'src/savings.html'
      }
    }
  },
  server: {
    port: 3000,
    open: true
  }
});
```

```json
// package.json
{
  "name": "fintech-portfolio",
  "version": "1.0.0",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview",
    "deploy": "npm run build && netlify deploy --dir=dist"
  },
  "dependencies": {
    "chart.js": "^4.0.0"
  },
  "devDependencies": {
    "vite": "^5.0.0"
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Fast development with hot reload
- **Savings Calculator** - Optimized production builds
- **Investment Tracker** - Multiple page setup
- **Net Worth Tracker** - Easy deployment

## Next Steps

After completing this, move to:
- [06-state](06-state/README.md) - State management
- [07-react](07-react/README.md) - React framework