# Prerequisites for FirstChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for empty strings**
   ```go
   if len(s) == 0 {
       return ""
   }
   ```

2. **How to use for...range**
   ```go
   for i, c := range s {
       // i is index, c is rune
       return string(c)  // First character
   }
   ```

3. **How to convert rune to string**
   ```go
   string(r)  // Convert rune to string
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Extract specific characters**
2. **Handle empty string edge cases**
3. **Use rune type correctly**
4. **Build character extraction tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Get first letter of category
- **Savings Calculator** - Check first character of input
- **Investment Tracker** - Validate ticker first letter
- **Net Worth Tracker** - Get account type

## Next Steps

After completing this exercise, move to:
- [03-lastchar](../03-lastchar/README.md) - Get last character
- [04-isempty](../04-isempty/README.md) - Check emptiness