# Prerequisites for ZipString

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to track consecutive characters**
   ```go
   current := rune(0)
   count := 0
   for _, r := range s {
       if r == current {
           count++
       } else {
           // New character, output previous
           current = r
           count = 1
       }
   }
   ```

2. **How to convert numbers to strings**
   ```go
   result := fmt.Sprint(count) + string(r)
   ```

3. **How to build strings**
   ```go
   var output string
   output += fmt.Sprint(count) + string(r)
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Detect consecutive patterns**
2. **Build run-length encoding**
3. **Track state across iterations**
4. **Create compression tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Compress expense data
- **Savings Calculator** - Format repeated values
- **Investment Tracker** - Compress historical data
- **Net Worth Tracker** - Summarize account history

## Next Steps

After completing this exercise, move to:
- [109-saveandmiss](../109-saveandmiss/README.md) - Saveandmiss
- [110-reversestrcap](../110-reversestrcap/README.md) - Reversestrcap