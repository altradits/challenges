# Prerequisites for StringContains

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check substring existence**
   ```go
   for i := 0; i <= len(s)-len(substr); i++ {
       if s[i:i+len(substr)] == substr {
           return true
       }
   }
   return false
   ```

2. **How to use strings.Index**
   ```go
   if strings.Index(s, substr) != -1 {
       // Found
   }
   ```

3. **How to handle empty substring**
   ```go
   if substr == "" {
       return true  // Empty string is always contained
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Check for substring presence**
2. **Use standard library functions**
3. **Handle edge cases**
4. **Build search tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Check for category
- **Savings Calculator** - Validate input
- **Investment Tracker** - Find ticker
- **Net Worth Tracker** - Check account type

## Next Steps

After completing this exercise, move to:
- [45-stringindex](../45-stringindex/README.md) - Find position
- [46-stringcount](../46-stringcount/README.md) - Count occurrences