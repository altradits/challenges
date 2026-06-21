# Portfolio Website - Finance Tools

## Objective

Build a portfolio website that showcases and sells your string-based finance tools.

## Learning Curve

### Phase 1: Basic Website
- Create static HTML pages
- Add CSS styling
- Deploy to GitHub Pages

### Phase 2: Tool Integration
- Embed your 53 string tools
- Create interactive demos
- Add user input forms

### Phase 3: E-commerce Features
- Add payment integration (Stripe)
- Create user accounts
- Track tool usage

## Building Blocks

### 1. Tool Showcase Page
```html
<!-- Each tool gets its own page -->
<div class="tool-card">
    <h3>StringLength</h3>
    <p>Count characters in any string</p>
    <form onsubmit="return runTool('StringLength')">
        <input type="text" id="input" placeholder="Enter text...">
        <button type="submit">Run</button>
    </form>
    <div id="output"></div>
</div>
```

### 2. Pricing Page
```html
<div class="pricing-card">
    <h3>Basic Plan</h3>
    <p>$9/month</p>
    <ul>
        <li>10 string tools</li>
        <li>100 API calls/day</li>
        <li>Email support</li>
    </ul>
    <button onclick="checkout('basic')">Subscribe</button>
</div>
```

### 3. User Dashboard
```html
<div class="dashboard">
    <h2>Your Tools</h2>
    <div id="tool-list">
        <!-- Dynamically populated based on subscription -->
    </div>
    <div id="usage-stats">
        <!-- Usage statistics from analytics -->
    </div>
</div>
```

## Step-by-Step Application Building

### Step 1: Create Tool Pages
For each of your 53 tools, create:
1. A dedicated HTML page
2. JavaScript to run the tool
3. Example inputs/outputs
4. Copy-to-clipboard button

### Step 2: Add Navigation
```html
<nav>
    <a href="/">Home</a>
    <a href="/tools">Tools</a>
    <a href="/pricing">Pricing</a>
    <a href="/dashboard">Dashboard</a>
</nav>
```

### Step 3: Integrate with Analytics
```javascript
// Track tool usage on your website
function trackUsage(toolName, input, output) {
    fetch('/api/interact', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({tool: toolName, input, output})
    });
}
```

## Finance Tools for Your Portfolio

### 1. Budget Planner
- Uses: `WordCount`, `Split`, `Join`, `StringFormat`
- Features: Monthly budget tracking, category splitting

### 2. Savings Calculator
- Uses: `Itoa`, `StringFormat`, `StringReplace`
- Features: Calculate savings goals, format currency

### 3. Investment Tracker
- Uses: `StringContains`, `StringPrefix`, `StringMap`
- Features: Track stock tickers, format performance

### 4. Expense Splitter
- Uses: `StringSplit`, `StringJoin`, `StringTrim`
- Features: Split bills, format amounts

### 5. Currency Converter
- Uses: `StringReplace`, `StringTrim`, `StringFormat`
- Features: Convert currencies, format results

## Technology Stack

| Component | Technology | Why |
|-----------|------------|-----|
| Frontend | HTML/CSS/JavaScript | No build step, easy to deploy |
| Backend | Go (Gin) | Your existing skills |
| Database | SQLite | Simple, file-based |
| Payments | Stripe | Easy integration |
| Hosting | GitHub Pages + Railway | Free tier available |
| Analytics | Your capstone analytics | Track user behavior |

## Deployment

### GitHub Pages (Frontend)
```bash
# Push to gh-pages branch
git checkout -b gh-pages
git push origin gh-pages
```

### Railway (Backend)
```bash
railway init
railway up
```

## Your Live Portfolio

After deployment:
```
https://yourname.github.io/string-tools-portfolio
https://string-tools-api.railway.app
```

## Monetization Strategies

1. **Subscription Tiers**
   - Free: 5 tools, 100 uses/day
   - Basic: 25 tools, 1000 uses/day ($9/month)
   - Pro: All 53 tools, unlimited ($29/month)

2. **One-time Purchase**
   - Buy individual tools
   - Lifetime access

3. **API Access**
   - Pay-per-use API
   - Bulk discounts

## Next Steps

After completing this, you'll be ready for:
- [134-professional](../134-professional/README.md) - Professional
- [135-python-data](../135-python-data/README.md) - Python Data