# Prerequisites for StringCount

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to find all occurrences**
   ```go
   start := 0
   for {
       idx := strings.Index(s[start:], substr)
       if idx == -1 {
           break
       }
       // Found at position start + idx
       start += idx + 1  // +1 for overlapping
   }
   ```

2. **How to handle empty substring**
   ```go
   if substr == "" {
       return 0  // Can't count empty substring
   }
   ```

3. **How to use strings.Index**
   ```go
   // Returns -1 if not found, index if found
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Count substring occurrences**
2. **Handle overlapping matches**
3. **Build search tools**
4. **Create text analysis functions**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Count category occurrences
- **Savings Calculator** - Count decimal places
- **Investment Tracker** - Count ticker patterns
- **Net Worth Tracker** - Count account types

## Next Steps

After completing this exercise, move to:
- [47-stringprefix](../47-stringprefix/README.md) - Check prefix
- [48-stringsuffix](../48-stringsuffix/README.md) - Check suffix