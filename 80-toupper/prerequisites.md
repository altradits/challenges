# Prerequisites for ToUpper

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for lowercase**
   ```go
   if c >= 'a' && c <= 'z' {
       // Character is lowercase
   }
   ```

2. **How to convert to uppercase**
   ```go
   c - 32  // 'a' (97) - 32 = 'A' (65)
   ```

3. **How to build strings**
   ```go
   result := ""
   for _, c := range s {
       if c >= 'a' && c <= 'z' {
           result += string(c - 32)
       } else {
           result += string(c)
       }
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Convert character cases**
2. **Handle non-alphabetic characters**
3. **Build transformation functions**
4. **Create text normalization tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Normalize category names
- **Savings Calculator** - Clean input
- **Investment Tracker** - Format ticker symbols
- **Net Worth Tracker** - Standardize account names

## Next Steps

After completing this exercise, move to:
- [81-tolower](../81-tolower/README.md) - Tolower
- [82-countalpha](../82-countalpha/README.md) - Countalpha