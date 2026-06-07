# Prerequisites for IsPalindrome

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to clean strings**
   ```go
   cleaned := ""
   for _, c := range s {
       if isAlphanumeric(c) {
           cleaned += string(c)
       }
   }
   ```

2. **How to convert to lowercase**
   ```go
   if c >= 'A' && c <= 'Z' {
       c = c + 32  // Convert to lowercase
   }
   ```

3. **How to compare strings**
   ```go
   if s == reversed {
       // Strings are equal
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Normalize strings** for comparison
2. **Build palindrome detection**
3. **Handle case-insensitive matching**
4. **Create text validation tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Validate symmetric data
- **Savings Calculator** - Check input patterns
- **Investment Tracker** - Validate ticker patterns
- **Net Worth Tracker** - Check account symmetry

## Next Steps

After completing this exercise, move to:
- [87-removespaces](../87-removespaces/README.md) - Removespaces
- [88-countrepeats](../88-countrepeats/README.md) - Countrepeats