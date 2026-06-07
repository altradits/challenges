# Prerequisites for StringPrefix

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check string length**
   ```go
   if len(s) < len(prefix) {
       return false
   }
   ```

2. **How to extract substrings**
   ```go
   s[:len(prefix)]  // First N characters
   ```

3. **How to compare strings**
   ```go
   if s[:len(prefix)] == prefix {
       return true
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Check string prefixes**
2. **Handle edge cases**
3. **Build validation functions**
4. **Create pattern matchers**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Validate "USD" or "EUR" prefix
- **Savings Calculator** - Check for "$" prefix
- **Investment Tracker** - Validate ticker format
- **Net Worth Tracker** - Check account type prefix

## Next Steps

After completing this exercise, move to:
- [48-stringsuffix](../48-stringsuffix/README.md) - Check suffix
- [49-stringfield](../49-stringfield/README.md) - Split by delimiter