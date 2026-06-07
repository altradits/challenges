# Prerequisites for CountChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to compare runes**
   ```go
   for _, r := range s {
       if r == target {
           // Found match
       }
   }
   ```

2. **How to count in a loop**
   ```go
   count := 0
   for _, r := range s {
       if r == target {
           count++
       }
   }
   ```

3. **How to handle empty strings**
   ```go
   if len(s) == 0 {
       return 0
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Count specific characters**
2. **Use rune comparison**
3. **Build character frequency tools**
4. **Create text analysis functions**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Count commas in expense list
- **Savings Calculator** - Count decimal places
- **Investment Tracker** - Count special characters
- **Net Worth Tracker** - Count account separators

## Next Steps

After completing this exercise, move to:
- [93-findlastchar](../93-findlastchar/README.md) - Findlastchar
- [94-replacechar](../94-replacechar/README.md) - Replacechar