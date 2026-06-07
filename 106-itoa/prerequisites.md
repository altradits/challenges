# Prerequisites for Itoa

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to extract digits**
   ```go
   for n > 0 {
       digit := n % 10
       n = n / 10
   }
   ```

2. **How to convert numbers to characters**
   ```go
   char := '0' + digit  // '0' + 5 = '5'
   ```

3. **How to reverse a string**
   ```go
   // Build digits in reverse, then reverse the result
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Convert numbers to strings**
2. **Handle negative numbers**
3. **Build strings from digits**
4. **Work with number bases**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Format currency amounts
- **Savings Calculator** - Display months/years
- **Investment Tracker** - Format prices
- **Net Worth Tracker** - Display totals

## Next Steps

After completing this exercise, move to:
- [107-thirdchar](../107-thirdchar/README.md) - Thirdchar
- [108-zipstring](../108-zipstring/README.md) - Zipstring