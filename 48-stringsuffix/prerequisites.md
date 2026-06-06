# Prerequisites for StringSuffix

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check string length**
   ```go
   if len(s) < len(suffix) {
       return false
   }
   ```

2. **How to extract substrings**
   ```go
   s[len(s)-len(suffix):]  // Last N characters
   ```

3. **How to compare strings**
   ```go
   if s[len(s)-len(suffix):] == suffix {
       return true
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Check string suffixes**
2. **Handle edge cases**
3. **Build validation functions**
4. **Create file extension checkers**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Validate file extensions
- **Savings Calculator** - Check for percentage suffix
- **Investment Tracker** - Validate currency codes
- **Net Worth Tracker** - Check for "USD" suffix

## Next Steps

After completing this exercise, move to:
- [49-stringfield](../49-stringfield/README.md) - Split by delimiter
- [50-stringmap](../50-stringmap/README.md) - Transform strings