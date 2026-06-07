# Prerequisites for StringTrim

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to find start position**
   ```go
   start := 0
   for start < len(s) && unicode.IsSpace(rune(s[start])) {
       start++
   }
   ```

2. **How to find end position**
   ```go
   end := len(s)
   for end > start && unicode.IsSpace(rune(s[end-1])) {
       end--
   }
   ```

3. **How to extract substring**
   ```go
   return s[start:end]
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Remove leading/trailing whitespace**
2. **Use unicode package**
3. **Handle all whitespace types**
4. **Build string cleaning tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Clean expense input
- **Savings Calculator** - Trim user input
- **Investment Tracker** - Clean ticker symbols
- **Net Worth Tracker** - Clean account names

## Next Steps

After completing this exercise, move to:
- [119-stringcontains](../119-stringcontains/README.md) - Stringcontains
- [120-stringindex](../120-stringindex/README.md) - Stringindex