# Prerequisites for LongestWord

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to split into words**
   ```go
   words := []string{}
   current := ""
   for i := 0; i < len(s); i++ {
       if s[i] == ' ' {
           if current != "" {
               words = append(words, current)
               current = ""
           }
       } else {
           current += string(s[i])
       }
   }
   ```

2. **How to find maximum**
   ```go
   longest := words[0]
   for _, w := range words {
       if len(w) > len(longest) {
           longest = w
       }
   }
   ```

3. **How to handle empty input**
   ```go
   if len(words) == 0 {
       return ""
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Extract words from text**
2. **Find maximum values**
3. **Handle multiple words**
4. **Build text analysis tools**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find longest category
- **Savings Calculator** - Find longest description
- **Investment Tracker** - Find longest ticker
- **Net Worth Tracker** - Find longest account

## Next Steps

After completing this exercise, move to:
- [98-searchreplace](../98-searchreplace/README.md) - Searchreplace
- [99-cleanlist](../99-cleanlist/README.md) - Cleanlist