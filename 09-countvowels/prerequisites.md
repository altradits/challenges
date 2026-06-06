# Prerequisites for CountVowels

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for specific characters**
   ```go
   vowels := "aeiouAEIOU"
   for _, c := range s {
       if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
           // c is a vowel
       }
   }
   ```

2. **How to handle case-insensitive matching**
   ```go
   // Check both uppercase and lowercase
   if c == 'a' || c == 'A' || c == 'e' || c == 'E' {
       // ...
   }
   ```

3. **How to count in a loop**
   ```go
   count := 0
   for _, c := range s {
       if isVowel(c) {
           count++
       }
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Detect multiple character types**
2. **Handle case-insensitive matching**
3. **Build character classifiers**
4. **Create text analysis tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Analyze expense descriptions
- **Savings Calculator** - Validate input text
- **Investment Tracker** - Check ticker readability
- **Net Worth Tracker** - Analyze account names

## Next Steps

After completing this exercise, move to:
- [10-reversestring](../10-reversestring/README.md) - String reversal
- [11-ispalindrome](../11-ispalindrome/README.md) - Palindrome detection