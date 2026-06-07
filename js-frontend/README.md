# JavaScript Frontend Mastery Course

## What You'll Learn

This course teaches you **frontend development in JavaScript** from beginner to mastery, focusing on building the FinTech portfolio website.

## Course Structure

### Beginner Level (01-03) - HTML, CSS, JavaScript Basics
| # | Topic | Key Concepts |
|---|-------|--------------|
| 01 | [HTML Fundamentals](01-html/README.md) | Semantic HTML, forms, accessibility |
| 02 | [CSS Styling](02-css/README.md) | Flexbox, Grid, responsive design |
| 03 | [JavaScript Basics](03-js-basics/README.md) | DOM, events, fetch API |

### Intermediate Level (04-06) - Modern Frontend
| # | Topic | Key Concepts |
|---|-------|--------------|
| 04 | [ES6+ Features](04-es6/README.md) | Modules, async/await, destructuring |
| 05 | [Build Tools](05-buildtools/README.md) | Vite, Webpack, npm scripts |
| 06 | [State Management](06-state/README.md) | LocalStorage, Context API |

### Advanced Level (07-10) - Frameworks & Production
| # | Topic | Key Concepts |
|---|-------|--------------|
| 07 | [React](07-react/README.md) | Components, hooks, routing |
| 08 | [TypeScript](08-typescript/README.md) | Types, interfaces, generics |
| 09 | [Testing](09-testing/README.md) | Jest, React Testing Library |
| 10 | [Deployment](10-deployment/README.md) | Netlify, Vercel, CI/CD |

## How This Helps Your FinTech Portfolio

### Finance Tools UI Integration
- **Budget Planner** - Interactive budget forms
- **Savings Calculator** - Real-time calculations
- **Investment Tracker** - Portfolio visualization
- **Net Worth Tracker** - Charts and graphs

## Key JavaScript Frontend Concepts

### DOM Manipulation
```javascript
// Select elements
const form = document.querySelector('#budget-form');
const results = document.querySelector('#results');

// Add event listener
form.addEventListener('submit', async (e) => {
  e.preventDefault();
  const data = new FormData(form);
  const response = await fetch('/api/budgets', {
    method: 'POST',
    body: JSON.stringify(Object.fromEntries(data))
  });
});
```

### React Component
```jsx
function BudgetForm({ onSubmit }) {
  const [category, setCategory] = useState('');
  const [amount, setAmount] = useState(0);

  return (
    <form onSubmit={onSubmit}>
      <input 
        value={category} 
        onChange={(e) => setCategory(e.target.value)} 
      />
      <input 
        type="number" 
        value={amount} 
        onChange={(e) => setAmount(e.target.value)} 
      />
    </form>
  );
}
```

## Progress Tracking

```
Beginner:     HTML → CSS → JavaScript
              ↓
Intermediate: ES6+ → Build Tools → State
              ↓
Advanced:     React → TypeScript → Testing → Deployment
              ↓
Capstone:     FinTech Portfolio Website
```

## Next Steps

After completing this course, you'll be able to:
- Build responsive, accessible web interfaces
- Create interactive financial dashboards
- Deploy production-ready frontend applications
- Integrate with your Go and JavaScript APIs