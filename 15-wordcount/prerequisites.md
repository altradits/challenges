# Prerequisites for WordCount

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to track word state**
   ```go
   inWord := false
   for i := 0; i < len(s); i++ {
       isWordChar := s[i] != ' ' && s[i] != '\t' && s[i] != '\n'
   }
   ```

2. **How to count transitions**
   ```go
   if isWordChar && !inWord {
       count++
       inWord = true
   } else if !isWordChar {
       inWord = false
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

1. **Count words in text**
2. **Track state across iterations**
3. **Handle multiple whitespace types**
4. **Build text analysis tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Count expense items
- **Savings Calculator** - Count input fields
- **Investment Tracker** - Count fund names
- **Net Worth Tracker** - Count accounts

## Next Steps

After completing this exercise, move to:
- [16-findchar](../16-findchar/README.md) - Find character
- [17-countchar](../17-countchar/README.md) - Count character