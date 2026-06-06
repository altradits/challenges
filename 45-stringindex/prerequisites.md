# Prerequisites for StringIndex

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to iterate with index**
   ```go
   for i := 0; i < len(s); i++ {
       // i is the byte index
   }
   ```

2. **How to extract substrings**
   ```go
   if s[i:i+len(substr)] == substr {
       // Found match at index i
   }
   ```

3. **How to handle edge cases**
   ```go
   if len(substr) == 0 {
       return 0  // Empty substring found at start
   }
   if len(s) < len(substr) {
       return -1  // Not found
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Find substring positions**
2. **Use sliding window technique**
3. **Handle search boundaries**
4. **Build search algorithms**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find category in input
- **Savings Calculator** - Find goal keyword
- **Investment Tracker** - Find ticker in text
- **Net Worth Tracker** - Find account in list

## Next Steps

After completing this exercise, move to:
- [46-stringcount](../46-stringcount/README.md) - Count occurrences
- [47-stringprefix](../47-stringprefix/README.md) - Check prefix