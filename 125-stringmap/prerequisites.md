# Prerequisites for StringMap

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to iterate over strings**
   ```go
   for i, char := range s {
       // Process each character
   }
   ```

2. **How to build strings**
   ```go
   var result strings.Builder
   result.WriteString(string(char))
   return result.String()
   ```

3. **How to use functions as parameters**
   ```go
   func transform(s string, f func(rune) rune) string {
       // Apply f to each character
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Transform each character**
2. **Build functional utilities**
3. **Create mapping functions**
4. **Process strings functionally**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Format currency symbols
- **Savings Calculator** - Transform input formats
- **Investment Tracker** - Normalize ticker symbols
- **Net Worth Tracker** - Clean account names

## Next Steps

After completing this exercise, move to:
- [126-stringfilter](../126-stringfilter/README.md) - Stringfilter
- [127-stringreduce](../127-stringreduce/README.md) - Stringreduce