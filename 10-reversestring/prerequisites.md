# Prerequisites for ReverseString

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to convert string to rune slice**
   ```go
   runes := []rune("Hello")
   ```

2. **How to swap array elements**
   ```go
   for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
       runes[i], runes[j] = runes[j], runes[i]
   }
   ```

3. **How to convert rune slice back to string**
   ```go
   result := string(runes)
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Modify strings** by converting to runes
2. **Use two-pointer technique** for in-place operations
3. **Handle multi-byte UTF-8 characters**
4. **Build complex transformations**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Reverse expense descriptions
- **Savings Calculator** - Format historical data
- **Investment Tracker** - Display timeline backwards
- **Net Worth Tracker** - Show account history

## Next Steps

After completing this exercise, move to:
- [11-ispalindrome](../11-ispalindrome/README.md) - String normalization
- [12-removespaces](../12-removespaces/README.md) - String filtering