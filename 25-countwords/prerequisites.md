# Prerequisites for CountWordsAdvanced

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for alphanumeric characters**
   ```go
   if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
       // Character is alphanumeric
   }
   ```

2. **How to track word boundaries**
   ```go
   inWord := false
   for _, c := range s {
       isWordChar := isAlphanumeric(c)
       if isWordChar && !inWord {
           count++
           inWord = true
       } else if !isWordChar {
           inWord = false
       }
   }
   ```

3. **How to handle punctuation**
   ```go
   // Treat punctuation as word boundaries
   // "Hello, World!" = 2 words
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Parse complex text** with punctuation
2. **Handle multiple word separators**
3. **Build robust text parsers**
4. **Create word extraction logic**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Parse "rent:1200, groceries:300"
- **Savings Calculator** - Parse multi-line input
- **Expense Splitter** - Parse people list
- **Net Worth Tracker** - Parse account descriptions

## Next Steps

After completing this exercise, move to:
- [26-findsubstring](../26-findsubstring/README.md) - Pattern matching
- [27-replaceall](../27-replaceall/README.md) - Global replacement