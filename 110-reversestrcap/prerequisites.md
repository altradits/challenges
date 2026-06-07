# Prerequisites for ReverseStrCap

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to split into words**
   ```go
   words := strings.Fields(s)
   ```

2. **How to convert case**
   ```go
   // Lowercase: c + 32
   // Uppercase: c - 32
   ```

3. **How to work with rune slices**
   ```go
   runes := []rune(word)
   runes[len(runes)-1] = uppercase(runes[len(runes)-1])
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Transform specific characters**
2. **Work with word boundaries**
3. **Apply selective case changes**
4. **Build text transformation tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Format category names
- **Savings Calculator** - Format output
- **Investment Tracker** - Format ticker display
- **Net Worth Tracker** - Format account names

## Next Steps

After completing this exercise, move to:
- [111-union](../111-union/README.md) - Union
- [112-inter](../112-inter/README.md) - Inter