# Prerequisites for RemoveSpaces

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to check for specific characters**
   ```go
   if s[i] == ' ' {
       // This is a space
   }
   ```

2. **How to build strings selectively**
   ```go
   result := ""
   for i := 0; i < len(s); i++ {
       if s[i] != ' ' {
           result += string(s[i])
       }
   }
   ```

3. **How to use strings.Builder**
   ```go
   var b strings.Builder
   for _, r := range s {
       if r != ' ' {
           b.WriteRune(r)
       }
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Filter characters** from strings
2. **Build selective output**
3. **Handle multiple space types** (space, tab, newline)
4. **Create clean string output**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Clean expense input
- **Savings Calculator** - Remove extra spaces
- **Investment Tracker** - Clean ticker input
- **Currency Converter** - Remove currency symbols

## Next Steps

After completing this exercise, move to:
- [13-countrepeats](../13-countrepeats/README.md) - Consecutive pattern detection
- [14-retainfirsthalf](../14-retainfirsthalf/README.md) - String slicing