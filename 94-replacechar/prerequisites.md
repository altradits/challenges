# Prerequisites for ReplaceChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to build strings conditionally**
   ```go
   result := ""
   for _, c := range s {
       if c == old {
           result += string(new)
       } else {
           result += string(c)
       }
   }
   ```

2. **How to use strings.Builder**
   ```go
   var b strings.Builder
   for _, c := range s {
       if c == old {
           b.WriteRune(new)
       } else {
           b.WriteRune(c)
       }
   }
   ```

3. **How to handle rune parameters**
   ```go
   func ReplaceChar(s string, old, new rune) string {
       // old and new are single characters
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Replace characters selectively**
2. **Build strings efficiently**
3. **Handle single character replacement**
4. **Create text transformation tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Replace currency symbols
- **Savings Calculator** - Clean input
- **Investment Tracker** - Normalize tickers
- **Net Worth Tracker** - Update account names

## Next Steps

After completing this exercise, move to:
- [20-printif](../20-printif/README.md) - Conditional printing
- [21-printifnot](../21-printifnot/README.md) - Inverse condition