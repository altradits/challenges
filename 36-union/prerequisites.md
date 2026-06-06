# Prerequisites for Union

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use maps for tracking**
   ```go
   seen := make(map[rune]bool)
   if !seen[c] {
       // First time seeing this character
   }
   ```

2. **How to iterate over two strings**
   ```go
   for _, c := range s1 + s2 {
       // Process all characters
   }
   ```

3. **How to build rune slices**
   ```go
   var result []rune
   result = append(result, c)
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Combine unique characters**
2. **Use maps for deduplication**
3. **Preserve order while removing duplicates**
4. **Build set operations**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Combine unique categories
- **Savings Calculator** - Merge account types
- **Investment Tracker** - Combine ticker symbols
- **Net Worth Tracker** - Merge account names

## Next Steps

After completing this exercise, move to:
- [37-inter](../37-inter/README.md) - Intersection
- [38-stringbuilder](../38-stringbuilder/README.md) - Efficient building