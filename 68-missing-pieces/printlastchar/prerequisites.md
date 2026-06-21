# Prerequisites for printlastchar

## Basic Skills Needed

Before starting this challenge, you should have completed:

1. **strlenindex/** — You need to understand that len() counts from 1 and indexes start from 0.

   The key idea you must carry in:
   ```go
   text := "stano"
   len(text)      // 5  — counts from 1
   text[4]        // 'o' — the last character, at index 4 (= len-1)
   ```

2. **How to use len()**
   ```go
   len("hello")  // 5
   ```

3. **String indexing**
   ```go
   text := "Bitcoin"
   text[0]  // 'B' (as a byte: 66)
   text[6]  // 'n' (as a byte: 110)
   ```

4. **Basic Go program structure**
   ```go
   package main

   import "fmt"

   func main() {
       // your code
   }
   ```

## The Problem This Challenge Solves

If you hardcode the index:
```go
fmt.Println(text[4])  // works for "stano"
fmt.Println(text[4])  // wrong for "Bitcoin" — 'o', not 'n'
```

You get the wrong answer for different-length strings. The challenge is to write code that works for any string by computing the last index instead of guessing it.

## Skills You Will Gain

After completing this challenge, you will be able to:

1. **Calculate the last index dynamically** using `len(str) - 1`
2. **Retrieve the last byte** using `str[len(str)-1]`
3. **Convert a byte to a string** using `string(byte)`
4. **Explain why hardcoding indexes is wrong** for general-purpose code
5. **Use the three-layer pattern**: last index → byte value → string character

## How This Connects to Bitcoin Development

Bitcoin addresses, transaction IDs, and script opcodes are all strings or byte slices. Code that reads the last byte of a script, checks the last character of an address, or processes the end of a buffer all uses exactly this pattern:

```go
script[len(script)-1]         // last opcode byte
txid[len(txid)-1]             // last byte of transaction ID
address[len(address)-1]       // last character of address
```

Mastering this now means you will recognise it instantly in production Go Bitcoin code.

## What Comes Next

After completing this challenge:
- [69-findsubstring](../../69-findsubstring/README.md) — Find a substring within a larger string
- [strlenindex](../strlenindex/prerequisites.md) — Review if you need to strengthen the len/index relationship
