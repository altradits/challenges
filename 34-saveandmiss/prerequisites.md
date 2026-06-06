# Prerequisites for SaveAndMiss

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use modulo for cycles**
   ```go
   for i := 0; i < len(s); i++ {
       pos := i % (2 * n)  // 2n cycle
       if pos < n {
           // Save phase
       } else {
           // Miss phase
       }
   }
   ```

2. **How to build strings conditionally**
   ```go
   result := ""
   for i, c := range s {
       if shouldSave(i, n) {
           result += string(c)
       }
   }
   ```

3. **How to handle edge cases**
   ```go
   if n <= 0 {
       return ""  // Nothing to save
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Implement periodic selection**
2. **Use modulo arithmetic**
3. **Build pattern-based tools**
4. **Create selective extraction**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Sample expenses
- **Savings Calculator** - Periodic reports
- **Investment Tracker** - Sample data points
- **Net Worth Tracker** - Snapshot selection

## Next Steps

After completing this exercise, move to:
- [35-reversestrcap](../35-reversestrcap/README.md) - Reverse capitalization
- [36-union](../36-union/README.md) - String union