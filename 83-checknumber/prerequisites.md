# Prerequisites for CheckNumber

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for digits**
   ```go
   if c >= '0' && c <= '9' {
       // Character is a digit
   }
   ```

2. **How to return early**
   ```go
   for _, c := range s {
       if c >= '0' && c <= '9' {
           return true  // Found a digit, return immediately
       }
   }
   return false
   ```

3. **How to use boolean return**
   ```go
   func CheckNumber(s string) bool {
       // Return true if found, false otherwise
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Detect digits in strings**
2. **Use early return for efficiency**
3. **Build validation functions**
4. **Create input sanitization tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Check for amounts in input
- **Savings Calculator** - Validate numeric input
- **Investment Tracker** - Check for prices
- **Net Worth Tracker** - Validate account values

## Next Steps

After completing this exercise, move to:
- [84-countvowels](../84-countvowels/README.md) - Countvowels
- [85-reversestring](../85-reversestring/README.md) - Reversestring