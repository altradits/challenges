# Prerequisites for FindSubstring

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use nested loops**
   ```go
   for i := 0; i <= len(text)-len(substring); i++ {
       match := true
       for j := 0; j < len(substring); j++ {
           if text[i+j] != substring[j] {
               match = false
               break
           }
       }
   }
   ```

2. **How to check substring match**
   ```go
   if text[i:i+len(substring)] == substring {
       // Found match
   }
   ```

3. **How to handle edge cases**
   ```go
   if len(substring) == 0 {
       return 0
   }
   if len(text) < len(substring) {
       return -1
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Search for substrings**
2. **Use sliding window technique**
3. **Handle pattern matching**
4. **Build search algorithms**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find category in input
- **Savings Calculator** - Find goal keyword
- **Investment Tracker** - Find ticker in text
- **Net Worth Tracker** - Find account in list

## Next Steps

After completing this exercise, move to:
- [102-replaceall](../102-replaceall/README.md) - Replaceall
- [103-split](../103-split/README.md) - Split