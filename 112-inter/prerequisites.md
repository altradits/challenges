# Prerequisites for Intersection

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to build a set from a string**
   ```go
   set := make(map[rune]bool)
   for _, c := range s2 {
       set[c] = true
   }
   ```

2. **How to check set membership**
   ```go
   if set[c] {
       // c is in the set
   }
   ```

3. **How to track seen characters**
   ```go
   seen := make(map[rune]bool)
   for _, c := range s1 {
       if set[c] && !seen[c] {
           // c is in both strings and not yet added
           seen[c] = true
       }
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Find common characters**
2. **Use maps as sets**
3. **Preserve order while deduplicating**
4. **Build set operations**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find common categories
- **Savings Calculator** - Find shared accounts
- **Investment Tracker** - Find common tickers
- **Net Worth Tracker** - Find shared assets

## Next Steps

After completing this exercise, move to:
- [38-stringbuilder](../38-stringbuilder/README.md) - Efficient building
- [39-stringsplit](../39-stringsplit/README.md) - Split words