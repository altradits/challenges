# Prerequisites for CountAlpha

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check character ranges**
   ```go
   if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
       // Character is alphabetic
   }
   ```

2. **How to use OR operator**
   ```go
   // || combines conditions
   if condition1 || condition2 {
       // Either condition is true
   }
   ```

3. **How to count in a loop**
   ```go
   count := 0
   for _, c := range s {
       if isAlpha(c) {
           count++
       }
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Classify characters**
2. **Combine multiple conditions**
3. **Count specific character types**
4. **Build character analysis tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Count letters in category names
- **Savings Calculator** - Validate input format
- **Investment Tracker** - Check ticker format
- **Net Worth Tracker** - Validate account names

## Next Steps

After completing this exercise, move to:
- [08-checknumber](../08-checknumber/README.md) - Digit detection
- [09-countvowels](../09-countvowels/README.md) - Vowel counting