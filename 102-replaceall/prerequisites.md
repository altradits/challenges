# Prerequisites for ReplaceAll

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to find all occurrences**
   ```go
   start := 0
   for {
       idx := strings.Index(s[start:], old)
       if idx == -1 {
           break
       }
       // Found at position start + idx
       start += idx + len(old)
   }
   ```

2. **How to build strings incrementally**
   ```go
   var b strings.Builder
   b.WriteString(s[:idx])
   b.WriteString(new)
   ```

3. **How to handle empty old string**
   ```go
   if old == "" {
       return s  // Nothing to replace
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Replace all occurrences**
2. **Build strings efficiently**
3. **Handle multiple replacements**
4. **Create text substitution tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Replace category names
- **Savings Calculator** - Update values
- **Investment Tracker** - Update tickers
- **Net Worth Tracker** - Update account names

## Next Steps

After completing this exercise, move to:
- [103-split](../103-split/README.md) - Split
- [104-join](../104-join/README.md) - Join