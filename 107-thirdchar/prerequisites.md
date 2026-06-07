# Prerequisites for ThirdChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use modulo for periodic selection**
   ```go
   for i, c := range s {
       if (i+1) % 3 == 0 {
           // Every 3rd character (1-indexed)
       }
   }
   ```

2. **How to build strings**
   ```go
   result := ""
   for i, c := range s {
       if (i+1) % 3 == 0 {
           result += string(c)
       }
   }
   ```

3. **How to handle empty strings**
   ```go
   if len(s) == 0 {
       return ""
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Extract periodic characters**
2. **Use modulo arithmetic**
3. **Build selection tools**
4. **Create pattern extractors**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Extract every 3rd expense
- **Savings Calculator** - Get periodic values
- **Investment Tracker** - Extract periodic prices
- **Net Worth Tracker** - Get periodic snapshots

## Next Steps

After completing this exercise, move to:
- [108-zipstring](../108-zipstring/README.md) - Zipstring
- [109-saveandmiss](../109-saveandmiss/README.md) - Saveandmiss