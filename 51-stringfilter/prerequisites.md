# Prerequisites for StringFilter

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to iterate over strings**
   ```go
   for i, char := range s {
       // Process each character
   }
   ```

2. **How to use conditional logic**
   ```go
   if shouldKeep(char) {
       // Add to result
   }
   ```

3. **How to build strings**
   ```go
   var result strings.Builder
   result.WriteRune(char)
   return result.String()
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Filter characters from strings**
2. **Build predicate functions**
3. **Create data cleaners**
4. **Process strings functionally**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Remove special characters
- **Savings Calculator** - Filter digits only
- **Investment Tracker** - Clean ticker input
- **Net Worth Tracker** - Remove formatting

## Next Steps

After completing this exercise, move to:
- [52-stringreduce](../52-stringreduce/README.md) - Reduce strings
- [53-stringformat](../53-stringformat/README.md) - Format strings