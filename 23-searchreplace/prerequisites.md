# Prerequisites for SearchReplace

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to find substring index**
   ```go
   for i := 0; i <= len(text)-len(old); i++ {
       if text[i:i+len(old)] == old {
           // Found at position i
       }
   }
   ```

2. **How to combine string parts**
   ```go
   result := text[:index] + new + text[index+len(old):]
   ```

3. **How to handle not found case**
   ```go
   // If not found, return original string
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Find and replace substrings**
2. **Build string combinations**
3. **Handle search edge cases**
4. **Create text editing tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Replace category names
- **Savings Calculator** - Update goal values
- **Investment Tracker** - Update tickers
- **Net Worth Tracker** - Update account names

## Next Steps

After completing this exercise, move to:
- [24-cleanlist](../24-cleanlist/README.md) - Clean lists
- [25-countwords](../25-countwords/README.md) - Advanced word count