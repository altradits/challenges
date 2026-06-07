# Prerequisites for FindChar

## Basic Skills Needed

Before starting this exercise, you should know:

1. **How to use for...range with index**
   ```go
   for i, c := range s {
       // i is the index, c is the character
   }
   ```

2. **How to return early**
   ```go
   for i, c := range s {
       if c == target {
           return i  // Return immediately when found
       }
   }
   return -1  // Not found
   ```

3. **How to handle not found case**
   ```go
   // Return -1 if character not found
   ```

## Skills You'll Learn

After completing this exercise, you'll be able to:

1. **Search for characters**
2. **Return positions**
3. **Use early return pattern**
4. **Build search functions**

## How This Helps Your Capstone

This skill is used in:
- **Budget Planner** - Find separator in input
- **Savings Calculator** - Find decimal point
- **Investment Tracker** - Find ticker separator
- **Net Worth Tracker** - Find account delimiter

## Next Steps

After completing this exercise, move to:
- [17-countchar](../17-countchar/README.md) - Count occurrences
- [18-findlastchar](../18-findlastchar/README.md) - Find last occurrence