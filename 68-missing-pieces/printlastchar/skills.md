# Skills for printlastchar

## What You'll Learn

**Previous:** [strlenindex](../strlenindex/skills.md)

**Challenge:** Print the last character of any string using len() — without hardcoding the index

## New Concepts Explained

### 1. Getting the last index with len(str) - 1

You cannot hardcode the index of the last character, because different strings have different lengths. The general formula is:

```go
text := "stano"
lastIndex := len(text) - 1  // 4
```

For any string, `len(str) - 1` gives you the last valid index.

### 2. The byte value at the last index

Once you have the last index, you can retrieve the byte at that position:

```go
text := "stano"
lastByte := text[len(text)-1]  // 111 (byte value of 'o')
```

`text[index]` always returns a byte (type `uint8`), not a character. For ASCII characters like English letters, one byte = one character. For multi-byte Unicode characters, this gets more complex — but for now, ASCII is the focus.

### 3. Converting a byte to a string with string()

To display the character rather than its numeric byte value, wrap it in `string()`:

```go
text := "stano"
last := string(text[len(text)-1])  // "o"
fmt.Println(last)                   // prints: o
```

`string(byte)` converts the byte to its UTF-8 string representation.

### 4. Putting it all together

```go
sample := "Bitcoin"

// Step 1: the last index
fmt.Println(len(sample) - 1)           // 6

// Step 2: the byte at that index
fmt.Println(sample[len(sample)-1])     // 110  (byte value of 'n')

// Step 3: the string value of the last character
fmt.Println(string(sample[len(sample)-1]))  // n
```

Each step peels back one layer: from position, to byte, to character.

### 5. Three approaches side by side

The code in `main.go` shows three ways to approach this:

**Approach 1: Hardcoded index (wrong for production)**
```go
fmt.Printf("The last character is: %c\n", text[4])
```
Works for "stano" but breaks for any other length string.

**Approach 2: Correct general solution**
```go
last := string(text[len(text)-1])
fmt.Println(last)
```
Works for any string.

**Approach 3: Showing all three layers**
```go
fmt.Println(len(sample)-1)                   // last index: 6
fmt.Println(sample[len(sample)-1])           // byte value: 110
fmt.Println(string(sample[len(sample)-1]))   // character: n
```

This three-step breakdown is the mental model you carry into all future string work.

## The Big Idea

Every string operation that touches the last element of a sequence uses this pattern: `collection[len(collection)-1]`. You will see this in loops, in sliding window algorithms, in parsers, and in Bitcoin transaction processing code.

**Next:** Continue to the main challenge series — [69-findsubstring](../../69-findsubstring/skills.md)
