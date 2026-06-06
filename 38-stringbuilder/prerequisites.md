# Prerequisites for StringBuilder

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use strings.Builder**
   ```go
   var b strings.Builder
   b.Grow(100)  // Pre-allocate capacity
   b.WriteRune('H')
   b.WriteString("ello")
   ```

2. **How to check for vowels**
   ```go
   vowels := "aeiouAEIOU"
   for _, r := range s {
       if !strings.ContainsRune(vowels, r) {
           // r is not a vowel
       }
   }
   ```

3. **How to build strings efficiently**
   ```go
   // Use Builder instead of += for better performance
   // Builder: O(n) total
   // +=: O(n²) total
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Build strings efficiently**
2. **Use Builder methods**
3. **Pre-allocate memory**
4. **Create high-performance tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Build expense reports
- **Savings Calculator** - Format multi-line output
- **Investment Tracker** - Build performance strings
- **Net Worth Tracker** - Combine account data

## Next Steps

After completing this exercise, move to:
- [39-stringsplit](../39-stringsplit/README.md) - Split words
- [40-stringjoin](../40-stringjoin/README.md) - Join strings