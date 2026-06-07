# Prerequisites for StringSplit

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use strings.Fields**
   ```go
   words := strings.Fields(s)  // Split by whitespace
   ```

2. **How to copy slices**
   ```go
   result := make([]string, len(words))
   copy(result, words)
   ```

3. **How to handle empty strings**
   ```go
   if len(words) == 0 {
       return []string{}
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Split strings into words**
2. **Use standard library functions**
3. **Handle whitespace correctly**
4. **Create text tokenization tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Split expense description
- **Savings Calculator** - Parse input
- **Investment Tracker** - Split fund names
- **Net Worth Tracker** - Split account descriptions

## Next Steps

After completing this exercise, move to:
- [40-stringjoin](../40-stringjoin/README.md) - Join strings
- [41-stringrepeat](../41-stringrepeat/README.md) - Repeat strings