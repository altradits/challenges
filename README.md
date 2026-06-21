# String Manipulation Mastery Course

Welcome to your complete string manipulation practice course! This collection of **53 exercises** is designed to take you from **complete beginner** to **expert** level in Go string handling.

## How to Use

Each exercise has its own numbered folder containing:
- `README.md` - Comprehensive theory, examples, and instructions
- `main.go` - Starter code with TODO function for you to implement

**Your task:** Open each `main.go`, study the README to learn the concepts, then implement the function. Run `go run .` to test your solution.

## Course Structure

### Beginner Level (01-08) - Foundations

These exercises teach you the absolute fundamentals of string handling in Go.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 01 | [String Length](13-stringlength/README.md) | `13-stringlength/` | `for...range` loops, counting, UTF-8 basics |
| 02 | [First Character](14-firstchar/README.md) | `14-firstchar/` | String indexing, rune vs byte, type conversion |
| 03 | [Last Character](15-lastchar/README.md) | `15-lastchar/` | Index calculation, `len()-1`, boundary checks |
| 04 | [Is Empty](16-isempty/README.md) | `16-isempty/` | Boolean returns, early return pattern, iteration |
| 05 | [To Upper](17-toupper/README.md) | `17-toupper/` | ASCII math, case conversion, character ranges |
| 06 | [To Lower](18-tolower/README.md) | `18-tolower/` | Reverse case conversion, symmetry of operations |
| 07 | [Count Alpha](19-countalpha/README.md) | `19-countalpha/` | Character classification, OR conditions |
| 08 | [Check Number](20-checknumber/README.md) | `20-checknumber/` | Digit detection, early return for efficiency |

**Beginner Tips:**
- Use `for...range` to iterate over strings (it handles UTF-8 correctly)
- Remember: `'a'` (97) and `'A'` (65) differ by 32 in ASCII
- Always check for empty strings first!
- Strings are immutable - build new strings, don't modify

---

### Beginner-Intermediate Level (09-14) - Building Patterns

These exercises build on basics and introduce more complex logic patterns.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 09 | [Count Vowels](21-countvowels/README.md) | `21-countvowels/` | Multi-condition logic, case-insensitive matching |
| 10 | [Reverse String](22-reversestring/README.md) | `22-reversestring/` | `[]rune` conversion, two-pointer swap technique |
| 11 | [Is Palindrome](23-ispalindrome/README.md) | `23-ispalindrome/` | String normalization, two-pointer comparison |
| 12 | [Remove Spaces](24-removespaces/README.md) | `24-removespaces/` | String filtering, selective character removal |
| 13 | [Count Repeats](25-countrepeats/README.md) | `25-countrepeats/` | State tracking, consecutive pattern detection |
| 14 | [Retain First Half](26-retainfirsthalf/README.md) | `26-retainfirsthalf/` | String slicing, integer division, edge cases |

**Beginner-Intermediate Tips:**
- Convert strings to `[]rune` when you need to modify or reverse them
- Build new strings by appending characters rather than modifying in place
- Think about edge cases: empty strings, single characters, all spaces
- Use state variables (flags) to track conditions across iterations

---

### Intermediate Level (15-22) - Real-World Processing

These exercises simulate real-world text processing tasks.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 15 | [Word Count](27-wordcount/README.md) | `27-wordcount/` | Word boundary detection, state tracking |
| 16 | [Find Char](28-findchar/README.md) | `28-findchar/` | Character search, index return, not-found handling |
| 17 | [Count Char](29-countchar/README.md) | `29-countchar/` | Accumulator pattern, counting specific values |
| 18 | [Find Last Char](30-findlastchar/README.md) | `30-findlastchar/` | Reverse search, tracking last occurrence |
| 19 | [Replace Char](31-replacechar/README.md) | `31-replacechar/` | Character substitution, building with replacements |
| 20 | [Print If](32-printif/README.md) | `32-printif/` | Conditional output, length-based logic |
| 21 | [Print If Not](33-printifnot/README.md) | `33-printifnot/` | Inverse conditions, logical inversion |
| 22 | [Longest Word](34-longestword/README.md) | `34-longestword/` | Finding maximums, word extraction |

**Intermediate Tips:**
- Use a "state" variable to track whether you're inside a word or between words
- For substring search, try matching from each possible starting position
- When building strings, use `[]byte` or `[]rune` for efficiency
- Test with edge cases: empty strings, multiple spaces, no matches

---

### Advanced Level (23-30) - Expert Techniques

These exercises combine multiple concepts for complex string operations.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 23 | [Search Replace](35-searchreplace/README.md) | `35-searchreplace/` | Substring search and replacement, slicing |
| 24 | [Clean List](36-cleanlist/README.md) | `36-cleanlist/` | Processing collections, map-filter pattern |
| 25 | [Count Words Advanced](37-countwords/README.md) | `37-countwords/` | Advanced word boundaries, punctuation handling |
| 26 | [Find Substring](38-findsubstring/README.md) | `38-findsubstring/` | Pattern matching, sliding window, nested loops |
| 27 | [Replace All](39-replaceall/README.md) | `39-replaceall/` | Global replacement, scan-and-build pattern |
| 28 | [Split](40-split/README.md) | `40-split/` | Tokenizing strings, slice manipulation |
| 29 | [Join](41-join/README.md) | `41-join/` | Combining strings, separator handling |
| 30 | [Camel to Snake](42-cameltosnakecase/README.md) | `42-cameltosnakecase/` | Validation, transformation, naming conventions |

### Expert Level (31-37) - Mastery

These exercises integrate all previous concepts for advanced string mastery.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 31 | [Integer to String](43-itoa/README.md) | `43-itoa/` | Number-to-string conversion, digit extraction |
| 32 | [Every Third Char](32-thirdchar/README.md) | `32-thirdchar/` | Modular indexing, periodic extraction |
| 33 | [Zip String](45-zipstring/README.md) | `45-zipstring/` | Run-length encoding, pattern compression |
| 34 | [Save and Miss](46-saveandmiss/README.md) | `46-saveandmiss/` | Periodic selection, chunk processing |
| 35 | [Reverse Str Cap](47-reversestrcap/README.md) | `47-reversestrcap/` | Selective case transformation |
| 36 | [Union](48-union/README.md) | `48-union/` | Set union, duplicate removal, order preservation |
| 37 | [Intersection](49-inter/README.md) | `49-inter/` | Set intersection, common character finding |

**Advanced Tips:**
- Combine multiple operations: validate → transform → return
- Think about the inverse operation (split ↔ join)
- Always validate input before transforming
- Consider what happens with edge cases in complex operations

---

### String Utilities Level (38-53) - Practical Tools

These exercises teach practical string utility functions you'll use in real projects.

| # | Exercise | Folder | Key Concepts |
|---|----------|--------|--------------|
| 38 | [String Builder](50-stringbuilder/README.md) | `50-stringbuilder/` | `strings.Builder`, efficient concatenation |
| 39 | [String Split](51-stringsplit/README.md) | `51-stringsplit/` | Word splitting, whitespace handling |
| 40 | [String Join](52-stringjoin/README.md) | `52-stringjoin/` | Joining slices, separator logic |
| 41 | [String Repeat](53-stringrepeat/README.md) | `53-stringrepeat/` | Repetition patterns, loop-based building |
| 42 | [String Replace](54-stringreplace/README.md) | `54-stringreplace/` | Substring replacement, incremental building |
| 43 | [String Trim](55-stringtrim/README.md) | `55-stringtrim/` | Whitespace removal, boundary detection |
| 44 | [String Contains](56-stringcontains/README.md) | `56-stringcontains/` | Substring search, boolean logic |
| 45 | [String Index](57-stringindex/README.md) | `57-stringindex/` | Finding positions, index return |
| 46 | [String Count](58-stringcount/README.md) | `58-stringcount/` | Counting occurrences, loop patterns |
| 47 | [String Prefix](59-stringprefix/README.md) | `59-stringprefix/` | Prefix checking, slice comparison |
| 48 | [String Suffix](60-stringsuffix/README.md) | `60-stringsuffix/` | Suffix checking, end-of-string logic |
| 49 | [String Field](61-stringfield/README.md) | `61-stringfield/` | Delimiter splitting, field extraction |
| 50 | [String Map](62-stringmap/README.md) | `62-stringmap/` | Functional mapping, rune transformation |
| 51 | [String Filter](63-stringfilter/README.md) | `63-stringfilter/` | Predicate filtering, character selection |
| 52 | [String Reduce](64-stringreduce/README.md) | `64-stringreduce/` | Cumulative reduction, fold operation |
| 53 | [String Format](65-stringformat/README.md) | `65-stringformat/` | Template formatting, placeholder replacement |

**Utility Tips:**
- `strings.Builder` is more efficient than `+=` for building strings
- Learn the standard library `strings` package functions
- These utilities form the building blocks of text processing
- Practice implementing them to understand how they work internally

---

## Key Go String Concepts

### Strings Are Immutable

In Go, strings cannot be modified in place. When you need to "change" a string, you build a new one.

```go
// WRONG - this won't compile
s[0] = 'H'

// CORRECT - build a new string
result := ""
for _, c := range s {
    result += string(c)
}
```

### Use `for...range` for Iteration

`for...range` iterates over runes (Unicode code points), not bytes:

```go
for i, char := range "Hello" {
    fmt.Println(i, string(char))
}
// Output:
// 0 H
// 1 e
// 2 l
// 3 l
// 4 o
```

### Convert to `[]rune` for Modification

If you need to modify characters or reverse a string:

```go
runes := []rune("Hello")
runes[0] = 'J'
result := string(runes) // "Jello"
```

### ASCII Math

You can convert cases using ASCII values:

```go
// 'a' (97) to 'A' (65): subtract 32
// 'A' (65) to 'a' (97): add 32
lower := 'A' + 32  // 'a'
upper := 'a' - 32  // 'A'
```

## Progress Tracking

Work through the exercises in order. Each one builds on skills from the previous ones:

```
Beginner:        Length → First/Last → Empty → Case Conversion → Count/Check
                 ↓
Beginner-Inter:  Vowels → Reverse → Palindrome → Filter → Repeats → Slice
                 ↓
Intermediate:    Word Count → Find/Count → Replace → Conditionals → Max
                 ↓
Advanced:        Search/Replace → Collections → Split/Join → Transform
                 ↓
Expert:          ITOA → Periodic → RLE → Set Operations
                 ↓
Utilities:       Builder → Split/Join → Replace/Trim → Map/Filter/Reduce
```

## Getting Help

If you're stuck:
1. **Re-read the README** - It contains all the theory you need
2. **Look at the hints** in `main.go`
3. **Draw the logic** on paper - trace through examples manually
4. **Test with simple input** first, then edge cases
5. **Check similar exercises** - patterns repeat across the course

## Additional Challenges

These are additional challenges in your workspace that complement the course:

| Challenge | Description |
|-----------|-------------|
| `cleanstr/` | Clean whitespace from strings |
| `expandstr/` | Expand spaces between words |
| `findprevprime/` | Find previous prime number |
| `fromto/` | Generate number range strings |
| `iscapitalized/` | Check if words are capitalized |
| `itoa/` | Integer to string conversion |
| `printmemory/` | Display memory as hex and ASCII |
| `printrevcomb/` | Print reversed digit combinations |
| `thirdtimeisacharm/` | Extract every third character |
| `weareunique/` | Count unique characters between strings |
| `zipstring/` | Run-length encode strings |
| `addprimesum/` | Sum of primes up to n |
| `canjump/` | Jump game array problem |
| `chunk/` | Chunk slice into sub-slices |
| `concatalternate/` | Alternate concatenation of slices |
| `concatslice/` | Concatenate two slices |
| `fprime/` | Prime factorization |
| `hiddenp/` | Check if string is hidden in another |
| `inter/` | Intersection of characters |
| `reversestrcap/` | Reverse capitalization of words |
| `saveandmiss/` | Periodic save/skip pattern |
| `union/` | Union of characters |
| `wdmatch/` | Word match checker |
| `fifthandskip/` | Extract every 5th character |
| `notdecimal/` | Check if string is not a decimal number |
| `revconcatalternate/` | Reverse alternate concatenation |
| `slice/` | Slice manipulation |
| `findpairs/` | Find pairs in array |
| `revwstr/` | Reverse words in string |
| `rostring/` | Rotate string |
| `wordflip/` | Flip words in string |
| `itoabase/` | Integer to any base conversion |
| `options/` | Parse command-line options |
| `piglatin/` | Convert to Pig Latin |
| `romannumbers/` | Roman numeral conversion |
| `brackets/` | Bracket validation |
| `rpncalc/` | Reverse Polish notation calculator |
| `brainfuck/` | Brainfuck interpreter |
| `grouping/` | Regex-like grouping |

## 3 Backend Capstones

After mastering string manipulation, apply your skills to these backend projects:

### 1. Financial Freedom API Engine

Build a REST API that helps users track their path to financial independence.

**Features:**
- Calculate savings runway (how many months until financial freedom)
- Track expense patterns and categorize spending
- Generate investment projection strings
- Format currency and percentage reports
- Validate financial data inputs

**String Skills Applied:**
- Parsing and formatting currency strings
- Validating account numbers and transaction IDs
- Generating human-readable financial reports
- Processing CSV export/import strings

### 2. Index Fund Tracker

Create a backend service for tracking index fund performance.

**Features:**
- Parse fund ticker symbols and names
- Format performance reports with percentage changes
- Validate and normalize fund data
- Generate comparison strings between funds
- Process historical data strings

**String Skills Applied:**
- Parsing ticker symbols (e.g., "VTI", "VOO")
- Formatting percentage changes with signs
- Normalizing fund names for comparison
- Building comparison report strings

### 3. Savings Runway Calculator

Build a service that calculates how long savings will last.

**Features:**
- Parse expense descriptions and amounts
- Format runway calculations in human-readable form
- Validate date strings and periods
- Generate savings milestone messages
- Process budget category strings

**String Skills Applied:**
- Parsing expense descriptions
- Formatting time periods ("X years Y months")
- Validating date formats
- Building milestone notification strings

## What You'll Achieve

After completing all 53 exercises, you will be able to:
- Iterate over strings safely using `for...range`
- Convert between strings, runes, and byte slices
- Search for characters and substrings
- Replace and transform strings
- Split and join strings
- Validate string patterns
- Process collections of strings
- Build complex text processing pipelines
- Implement practical string utility functions
- Apply string skills to real-world backend problems

## Next Steps

After completing this course:
1. Revisit your existing challenges with new confidence
2. Build your own string utility library
3. Tackle more complex text processing problems
4. Explore the `strings` package standard library
5. Start the backend capstone projects

**Happy coding! Start with `13-stringlength` and work your way up.**
