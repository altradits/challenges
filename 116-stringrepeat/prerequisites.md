# Prerequisites for StringRepeat

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use loops for repetition**
   ```go
   for i := 0; i < n; i++ {
       // Repeat n times
   }
   ```

2. **How to use strings.Builder**
   ```go
   var b strings.Builder
   for i := 0; i < n; i++ {
       b.WriteString(s)
   }
   return b.String()
   ```

3. **How to handle edge cases**
   ```go
   if n <= 0 {
       return ""
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Repeat strings efficiently**
2. **Use Builder for performance**
3. **Handle repetition edge cases**
4. **Build string generation tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Repeat separator lines
- **Savings Calculator** - Format output
- **Investment Tracker** - Create charts
- **Net Worth Tracker** - Build reports

## Next Steps

After completing this exercise, move to:
- [117-stringreplace](../117-stringreplace/README.md) - Stringreplace
- [118-stringtrim](../118-stringtrim/README.md) - Stringtrim