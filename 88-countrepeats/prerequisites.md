# Prerequisites for CountRepeats

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to track state**
   ```go
   inRepeat := false
   for i := 1; i < len(s); i++ {
       if s[i] == s[i-1] {
           // Consecutive same character
       }
   }
   ```

2. **How to detect transitions**
   ```go
   if s[i] == s[i-1] && !inRepeat {
       // Entering a repeat group
       inRepeat = true
       count++
   }
   ```

3. **How to handle edge cases**
   ```go
   if len(s) == 0 {
       return 0
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Detect consecutive patterns**
2. **Track state across iterations**
3. **Count groups, not individual characters**
4. **Build pattern analysis tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Detect repeated categories
- **Savings Calculator** - Check for duplicate entries
- **Investment Tracker** - Find repeated patterns
- **Net Worth Tracker** - Identify duplicate accounts

## Next Steps

After completing this exercise, move to:
- [89-retainfirsthalf](../89-retainfirsthalf/README.md) - Retainfirsthalf
- [90-wordcount](../90-wordcount/README.md) - Wordcount