# Prerequisites for RetainFirstHalf

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check string length**
   ```go
   if len(s) <= 1 {
       return s  // Too short to split
   }
   ```

2. **How to use integer division**
   ```go
   half := len(s) / 2  // Rounds down for odd lengths
   ```

3. **How to slice strings**
   ```go
   s[:half]  // First half of string
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Split strings in half**
2. **Handle odd-length strings**
3. **Use integer division**
4. **Build string partitioning tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Split long descriptions
- **Savings Calculator** - Truncate output
- **Investment Tracker** - Show partial tickers
- **Net Worth Tracker** - Display account previews

## Next Steps

After completing this exercise, move to:
- [15-wordcount](../15-wordcount/README.md) - Count words
- [16-findchar](../16-findchar/README.md) - Find character