# Prerequisites for CamelToSnakeCase

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for uppercase**
   ```go
   if c >= 'A' && c <= 'Z' {
       // Character is uppercase
   }
   ```

2. **How to validate input**
   ```go
   for _, c := range s {
       if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
           return s  // Invalid character
       }
   }
   ```

3. **How to build strings**
   ```go
   var result []byte
   for i := 0; i < len(s); i++ {
       if s[i] >= 'A' && s[i] <= 'Z' && i > 0 {
           result = append(result, '_')
       }
       result = append(result, s[i])
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Validate string format**
2. **Transform naming conventions**
3. **Handle edge cases**
4. **Build string converters**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Convert field names
- **Savings Calculator** - Format output
- **Investment Tracker** - Normalize tickers
- **Net Worth Tracker** - Standardize names

## Next Steps

After completing this exercise, move to:
- [31-itoa](../31-itoa/README.md) - Number to string
- [32-thirdchar](../32-thirdchar/README.md) - Every 3rd character