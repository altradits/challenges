# Prerequisites for StringField

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to split strings**
   ```go
   parts := strings.Split(s, ",")
   ```

2. **How to access array elements**
   ```go
   if index < len(parts) {
       return parts[index]
   }
   ```

3. **How to handle out of bounds**
   ```go
   if index >= len(parts) {
       return ""
   }
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Extract fields from delimited strings**
2. **Handle CSV-like data**
3. **Build parsers**
4. **Create data extractors**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Parse CSV budget data
- **Savings Calculator** - Extract parameters from config
- **Investment Tracker** - Parse transaction records
- **Net Worth Tracker** - Parse account data

## Next Steps

After completing this exercise, move to:
- [125-stringmap](../125-stringmap/README.md) - Stringmap
- [126-stringfilter](../126-stringfilter/README.md) - Stringfilter