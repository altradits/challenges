# 01. HTML Fundamentals

## What You'll Learn

This exercise teaches you **HTML basics for web development**. By the end, you'll understand:
- Semantic HTML elements
- Forms and input validation
- Accessibility best practices
- SEO-friendly structure

## Theory: HTML Fundamentals

### Semantic HTML

```html
<!-- Good: Semantic structure -->
<header>
  <nav>
    <ul>
      <li><a href="/budget">Budget Planner</a></li>
    </ul>
  </nav>
</header>

<main>
  <section>
    <h1>Budget Planner</h1>
    <form id="budget-form">
      <label for="category">Category</label>
      <input type="text" id="category" name="category" required>
      
      <label for="amount">Amount</label>
      <input type="number" id="amount" name="amount" required>
      
      <button type="submit">Add Budget</button>
    </form>
  </section>
</main>

<footer>
  <p>&copy; 2024 FinTech Portfolio</p>
</footer>
```

### Forms

```html
<form action="/api/budgets" method="POST">
  <fieldset>
    <legend>Budget Entry</legend>
    
    <label for="category">Category</label>
    <select id="category" name="category" required>
      <option value="food">Food</option>
      <option value="rent">Rent</option>
      <option value="utilities">Utilities</option>
    </select>
    
    <label for="amount">Amount ($)</label>
    <input type="number" id="amount" name="amount" min="0" step="0.01" required>
    
    <label for="date">Date</label>
    <input type="date" id="date" name="date" required>
  </fieldset>
  
  <button type="submit">Save Budget</button>
</form>
```

## The Challenge

Create the HTML structure for a FinTech tool.

### Expected Code

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Budget Planner - FinTech Portfolio</title>
</head>
<body>
  <header>
    <h1>FinTech Portfolio</h1>
    <nav>
      <a href="/">Home</a>
      <a href="/budget">Budget Planner</a>
      <a href="/savings">Savings Calculator</a>
    </nav>
  </header>
  
  <main>
    <section id="budget-section">
      <h2>Budget Planner</h2>
      <form id="budget-form">
        <label for="category">Category</label>
        <input type="text" id="category" name="category" required>
        
        <label for="amount">Amount</label>
        <input type="number" id="amount" name="amount" required>
        
        <button type="submit">Add Budget</button>
      </form>
      
      <table id="budget-table">
        <thead>
          <tr>
            <th>Category</th>
            <th>Amount</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody></tbody>
      </table>
    </section>
  </main>
</body>
</html>
```

## How This Helps Your FinTech Portfolio

This skill is used in:
- **Budget Planner** - Form for budget entry
- **Savings Calculator** - Input for savings goals
- **Investment Tracker** - Portfolio display
- **Net Worth Tracker** - Asset/liability forms

## Next Steps

After completing this, move to:
- [02-css](02-css/README.md) - Styling with CSS
- [03-js-basics](03-js-basics/README.md) - JavaScript interactivity