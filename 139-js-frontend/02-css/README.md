# 02. CSS Styling

## What You'll Learn

This exercise teaches you **CSS for modern web design**. By the end, you'll understand:
- Flexbox and Grid layouts
- Responsive design
- CSS variables
- Modern styling techniques

## Theory: CSS Fundamentals

### Flexbox

```css
.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
}

.card {
  flex: 1;
  min-width: 250px;
}
```

### Grid

```css
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}

.budget-card {
  grid-column: span 1;
}
```

### Responsive Design

```css
/* Mobile first */
.container {
  padding: 1rem;
}

/* Tablet */
@media (min-width: 768px) {
  .container {
    max-width: 750px;
    margin: 0 auto;
  }
}

/* Desktop */
@media (min-width: 1024px) {
  .container {
    max-width: 1200px;
  }
}
```

### CSS Variables

```css
:root {
  --primary: #2563eb;
  --secondary: #64748b;
  --success: #22c55e;
  --danger: #ef4444;
  --font-main: 'Inter', sans-serif;
}

.button {
  background: var(--primary);
  color: white;
  font-family: var(--font-main);
}
```

## The Challenge

Style the FinTech portfolio with modern CSS.

### Expected Code

```css
/* styles.css */
:root {
  --primary: #2563eb;
  --secondary: #64748b;
  --success: #22c55e;
  --danger: #ef4444;
  --background: #f8fafc;
  --text: #1e293b;
  --border: #e2e8f0;
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
  background: var(--background);
  color: var(--text);
  line-height: 1.6;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
  border-bottom: 1px solid var(--border);
}

.nav a {
  margin-left: 1rem;
  text-decoration: none;
  color: var(--primary);
}

.budget-form {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
  margin: 2rem 0;
}

.input-group {
  display: flex;
  flex-direction: column;
}

.button {
  background: var(--primary);
  color: white;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.375rem;
  cursor: pointer;
}

.button:hover {
  opacity: 0.9;
}

.budget-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
}

.budget-table th,
.budget-table td {
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid var(--border);
}

.budget-table th {
  background: var(--background);
  font-weight: 600;
}

@media (max-width: 768px) {
  .container {
    padding: 1rem;
  }
  
  .header {
    flex-direction: column;
    gap: 1rem;
  }
}
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Clean, responsive forms
- **Savings Calculator** - Visual savings progress
- **Investment Tracker** - Portfolio cards
- **Net Worth Tracker** - Dashboard layout

## Next Steps

After completing this, move to:
- [03-js-basics](03-js-basics/README.md) - JavaScript interactivity
- [04-es6](04-es6/README.md) - Modern JavaScript