# Prerequisites for CleanList

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to trim spaces**
   ```go
   func trimSpace(s string) string {
       start := 0
       end := len(s)
       for start < end && (s[start] == ' ' || s[start] == '\t') {
           start++
       }
       for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
           end--
       }
       return s[start:end]
   }
   ```

2. **How to check for empty strings**
   ```go
   if trimmed == "" {
       // Skip this item
   }
   ```

3. **How to build string slices**
   ```go
   result := make([]string, 0)
   result = append(result, trimmed)
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Clean string lists**
2. **Remove empty entries**
3. **Trim whitespace**
4. **Build data sanitization tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Clean expense list
- **Savings Calculator** - Clean input data
- **Investment Tracker** - Clean ticker list
- **Net Worth Tracker** - Clean account list

## Next Steps

After completing this exercise, move to:
- [25-countwords](../25-countwords/README.md) - Advanced word count
- [26-findsubstring](../26-findsubstring/README.md) - Find substring