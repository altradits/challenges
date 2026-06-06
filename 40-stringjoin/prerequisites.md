# Prerequisites for StringJoin

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for empty slice**
   ```go
   if len(elems) == 0 {
       return ""
   }
   ```

2. **How to iterate over slice**
   ```go
   for i, s := range elems {
       // i is index, s is string
   }
   ```

3. **How to use strings.Builder**
   ```go
   var b strings.Builder
   b.WriteString(elems[0])
   for i := 1; i < len(elems); i++ {
       b.WriteString(sep)
       b.WriteString(elems[i])
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Combine strings with separator**
2. **Handle empty inputs**
3. **Build efficient concatenation**
4. **Create list formatting tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Join expense list
- **Savings Calculator** - Format output
- **Investment Tracker** - Combine tickers
- **Net Worth Tracker** - Build account list

## Next Steps

After completing this exercise, move to:
- [30-cameltosnakecase](../30-cameltosnakecase/README.md) - Naming conversion
- [31-itoa](../31-itoa/README.md) - Number to string