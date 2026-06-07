# Prerequisites for Split

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to find separator positions**
   ```go
   start := 0
   for i := 0; i < len(s); i++ {
       if s[i] == sep {
           // Found separator at position i
       }
   }
   ```

2. **How to extract substrings**
   ```go
   part := s[start:i]
   ```

3. **How to build string slices**
   ```go
   var result []string
   result = append(result, part)
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Tokenize strings**
2. **Handle multiple separators**
3. **Build parsers**
4. **Create CSV-like processors**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Split "amount,category,date"
- **Savings Calculator** - Parse multi-line input
- **Investment Tracker** - Split fund data
- **Net Worth Tracker** - Split account data

## Next Steps

After completing this exercise, move to:
- [104-join](../104-join/README.md) - Join
- [105-cameltosnakecase](../105-cameltosnakecase/README.md) - Cameltosnakecase