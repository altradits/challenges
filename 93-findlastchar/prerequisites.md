# Prerequisites for FindLastChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to track last position**
   ```go
   lastIndex := -1
   for i, c := range s {
       if c == target {
           lastIndex = i
       }
   }
   return lastIndex
   ```

2. **How to initialize with -1**
   ```go
   // -1 means "not found"
   ```

3. **How to update on match**
   ```go
   // Keep updating until the last match
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Find last occurrence**
2. **Track positions**
3. **Handle not found case**
4. **Build reverse search tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find last separator
- **Savings Calculator** - Find last decimal
- **Investment Tracker** - Find last ticker
- **Net Worth Tracker** - Find last account

## Next Steps

After completing this exercise, move to:
- [94-replacechar](../94-replacechar/README.md) - Replacechar
- [95-printif](../95-printif/README.md) - Printif