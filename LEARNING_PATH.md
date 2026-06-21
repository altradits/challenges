# Go Mastery: Complete Learning Path

Follow the directories **in number order (01 → 139)**. Each number is exactly where you should be.

```
Phase 1  (01–05):  Hello World — print basics
Phase 2  (06–12):  Go Foundations — language building blocks
Phase 3  (13–36):  Beginner Practice — apply the basics one concept at a time
Phase 4  (37–65):  Strings Package Mastery — the full strings toolkit
Phase 5  (66–129): Challenge Problems — combine everything, harder problems
Phase 6  (130–139): Capstone Projects — real-world applications
```

---

## Phase 1: Hello World (01–05)

Learn the four-line Go skeleton by heart. Write each program from scratch without copying.

| # | Directory | What You Practice |
|---|-----------|-------------------|
| 01 | `01-only1` | `package main`, `import "fmt"`, `fmt.Println`, print `1` |
| 02 | `02-onlya` | Same skeleton, print `a` |
| 03 | `03-onlyb` | Same skeleton, print `b` |
| 04 | `04-onlyf` | Same skeleton, print `f` |
| 05 | `05-onlyz` | Same skeleton, print `z` — muscle memory test |

---

## Phase 2: Go Language Foundations (06–12)

One concept per lesson. Study each one before doing challenges.

| # | Directory | Concept Introduced |
|---|-----------|-------------------|
| 06 | `06-functions` | Function syntax, parameters, return types, `package piscine` |
| 07 | `07-forloops` | All for loop patterns: index, while-style, range, break/continue |
| 08 | `08-stringspackage` | The `strings` package: Fields, Contains, Builder, Split, etc. |
| 09 | `09-slicesintro` | Slices: `make`, `append`, sub-slicing, range, nil vs empty |
| 10 | `10-mapsintro` | Maps: `make`, key-value, zero value trick, comma-ok, sets |
| 11 | `11-arrays` | Fixed-size arrays: `[N]T`, value semantics, arrays vs slices |
| 12 | `12-recursion` | Recursion: base case, call stack trace, factorial, GCD |

---

## Phase 3: Beginner Practice (13–36)

Short focused exercises. Each one teaches exactly one string or function technique.
Read the `skills.md` first, then solve the challenge.

| # | Directory | Skill |
|---|-----------|-------|
| 13 | `13-stringlength` | `len()` and manual character counting |
| 14 | `14-firstchar` | First character: `s[0]`, `string(s[0])` |
| 15 | `15-lastchar` | Last character: `s[len(s)-1]` |
| 16 | `16-isempty` | Empty string check: `s == ""` or `len(s) == 0` |
| 17 | `17-toupper` | `strings.ToUpper` |
| 18 | `18-tolower` | `strings.ToLower` |
| 19 | `19-countalpha` | `unicode.IsLetter`, counting with for...range |
| 20 | `20-checknumber` | `unicode.IsDigit`, detecting digit characters |
| 21 | `21-countvowels` | `strings.ContainsRune`, vowel detection |
| 22 | `22-reversestring` | Reverse a string using `[]rune` |
| 23 | `23-ispalindrome` | Palindrome check: reverse + compare |
| 24 | `24-removespaces` | `strings.ReplaceAll`, filtering characters |
| 25 | `25-countrepeats` | `map[rune]int` for character frequency |
| 26 | `26-retainfirsthalf` | String slicing: `s[:len(s)/2]` |
| 27 | `27-wordcount` | `strings.Fields` for word splitting |
| 28 | `28-findchar` | `strings.Index`, manual loop search |
| 29 | `29-countchar` | `strings.Count` for single character |
| 30 | `30-findlastchar` | `strings.LastIndex` |
| 31 | `31-replacechar` | `strings.ReplaceAll` for character replacement |
| 32 | `32-printif` | `if/else`, `len()`, `||` operator |
| 33 | `33-printifnot` | Inverted conditions, De Morgan's Law |
| 34 | `34-longestword` | Running maximum with `strings.Fields` |
| 35 | `35-searchreplace` | `strings.Index` + slice concatenation |
| 36 | `36-cleanlist` | `map[string]bool` as a set for deduplication |

---

## Phase 4: Strings Package Mastery (37–65)

Deep dive into every function in the `strings` and `fmt` packages.

| # | Directory | Function / Concept |
|---|-----------|-------------------|
| 37 | `37-countwords` | Advanced word counting with `strings.Fields` |
| 38 | `38-findsubstring` | `strings.Index`, sliding window |
| 39 | `39-replaceall` | `strings.ReplaceAll` |
| 40 | `40-split` | `strings.Split` |
| 41 | `41-join` | `strings.Join` |
| 42 | `42-cameltosnakecase` | `unicode`, `strings.Builder`, case detection |
| 43 | `43-itoa` | `strconv.Itoa`, `strconv.Atoi` |
| 44 | `44-thirdchar` | Index step patterns |
| 45 | `45-zipstring` | Run-length encoding |
| 46 | `46-saveandmiss` | Periodic skip pattern |
| 47 | `47-reversestrcap` | Reverse + case toggle |
| 48 | `48-union` | Set union with `map[rune]bool` |
| 49 | `49-inter` | Set intersection |
| 50 | `50-stringbuilder` | `strings.Builder` deep dive |
| 51 | `51-stringsplit` | `strings.Split` variants |
| 52 | `52-stringjoin` | `strings.Join` |
| 53 | `53-stringrepeat` | `strings.Repeat` |
| 54 | `54-stringreplace` | `strings.Replace` (count-limited) |
| 55 | `55-stringtrim` | `strings.TrimSpace`, `Trim`, `TrimPrefix`, `TrimSuffix` |
| 56 | `56-stringcontains` | `strings.Contains` |
| 57 | `57-stringindex` | `strings.Index` detail |
| 58 | `58-stringcount` | `strings.Count` |
| 59 | `59-stringprefix` | `strings.HasPrefix`, `TrimPrefix` |
| 60 | `60-stringsuffix` | `strings.HasSuffix`, `TrimSuffix` |
| 61 | `61-stringfield` | `strings.Fields` vs `Split` vs `FieldsFunc` |
| 62 | `62-stringmap` | `strings.Map`, higher-order `func(rune) rune` |
| 63 | `63-stringfilter` | Filter pattern with `strings.Builder` |
| 64 | `64-stringreduce` | Fold / accumulator pattern over runes |
| 65 | `65-stringformat` | `fmt.Sprintf` format verbs: `%s %d %f %v %T %q %x` |

---

## Phase 5: Challenge Problems (66–129)

These are the original piscine challenges. You now have all the tools — apply them.
Each one is harder than the Phase 3/4 equivalents. Some require multiple concepts together.

| # | Directory | Challenge |
|---|-----------|-----------|
| 66 | `66-checknumber` | Check if a string contains any digit |
| 67 | `67-count-character` | Count occurrences of a specific character |
| 68 | `68-missing-pieces` | Supplementary concept reference (len vs index, etc.) |
| 69 | `69-findsubstring` | Find substring index manually |
| 70 | `70-printif` | Conditional print with length logic |
| 71 | `71-printifnot` | Inverted condition print |
| 72 | `72-removespaces` | Remove all spaces from a string |
| 73 | `73-retainfirsthalf` | Keep only the first half of a string |
| 74 | `74-reversestring` | Reverse a string |
| 75 | `75-titlecase` | Convert string to Title Case |
| 76 | `76-wordcount` | Count words in a string |
| 77 | `77-cameltosnakecase` | Convert camelCase to snake_case |
| 78 | `78-countrepeats` | Count characters that appear more than once |
| 79 | `79-digitlen` | Count how many digits are in a string |
| 80 | `80-firstword` | Extract the first word |
| 81 | `81-gcd` | Greatest Common Divisor (recursive Euclidean algorithm) |
| 82 | `82-hashcode` | Generate a simple hash code from a string |
| 83 | `83-lastword` | Extract the last word |
| 84 | `84-longestword` | Find the longest word |
| 85 | `85-replaceall` | Replace all occurrences of a substring |
| 86 | `86-searchreplace` | Search and replace the first occurrence |
| 87 | `87-splitjoin` | Split a string then join with a different separator |
| 88 | `88-wordanatomy` | Analyze prefix/suffix structure of words |
| 89 | `89-wordanatomy2` | Advanced word anatomy with custom prefix/suffix lists |
| 90 | `90-clean-the-list` | Normalize a grocery list (trim, capitalize, ensure "milk") |
| 91 | `91-cleanstr` | Remove extra whitespace between words |
| 92 | `92-expandstr` | Add spaces between each character |
| 93 | `93-findprevprime` | Find the largest prime ≤ n |
| 94 | `94-fromto` | Generate a range string like "1→5" |
| 95 | `95-iscapitalized` | Check if every word starts with a capital letter |
| 96 | `96-itoa-manual` | Implement integer-to-string without `strconv` |
| 97 | `97-printmemory` | Print `[10]byte` array as hex + ASCII (like `xxd`) |
| 98 | `98-printrevcomb` | Print reversed combinations of digits |
| 99 | `99-thirdtimeisacharm` | Extract every third character |
| 100 | `100-weareunique` | Count unique characters in a string |
| 101 | `101-zipstring` | Run-length encode: "aaa" → "a3" |
| 102 | `102-addprimesum` | Sum all primes up to n |
| 103 | `103-canjump` | Jump game — can you reach the last index? |
| 104 | `104-chunk` | Split a slice into chunks of size k |
| 105 | `105-concatalternate` | Interleave characters from two strings |
| 106 | `106-concatslice` | Concatenate two slices |
| 107 | `107-fprime` | Print prime factors of n |
| 108 | `108-hiddenp` | Check if string A is hidden inside concatenation of string B with itself |
| 109 | `109-inter` | Intersection of two strings (unique shared chars) |
| 110 | `110-reversestrcap` | Reverse + toggle case of each character |
| 111 | `111-saveandmiss` | Keep every 2nd character, skip every 3rd |
| 112 | `112-union` | Union of two strings (sorted unique chars) |
| 113 | `113-wdmatch` | Check if string A is a scattered subsequence of B |
| 114 | `114-fifthandskip` | Extract every 5th character starting from the 5th |
| 115 | `115-notdecimal` | Check if a string is NOT a decimal number |
| 116 | `116-revconcatalternate` | Reverse interleave characters |
| 117 | `117-slice` | Implement slice operations (remove, insert) |
| 118 | `118-findpairs` | Find all pairs in a slice summing to a target |
| 119 | `119-revwstr` | Reverse the words in a sentence |
| 120 | `120-rostring` | Rotate string by n positions |
| 121 | `121-wordflip` | Reverse each word in a sentence |
| 122 | `122-itoabase` | Convert integer to string in any base (2–36) |
| 123 | `123-options` | Parse command-line flags like `--flag=value` |
| 124 | `124-piglatin` | Convert words to Pig Latin |
| 125 | `125-romannumbers` | Convert integer to Roman numerals |
| 126 | `126-brackets` | Validate bracket balance with a stack |
| 127 | `127-rpncalc` | Reverse Polish Notation calculator |
| 128 | `128-brainfuck` | Brainfuck interpreter |
| 129 | `129-grouping` | Regex-like group matching with `|` alternation |

---

## Phase 6: Capstone Projects (130–139)

Apply everything in real projects.

| # | Directory | Project |
|---|-----------|---------|
| 130 | `130-financial-freedom-api` | REST API: financial freedom calculator |
| 131 | `131-index-fund-tracker` | Index fund portfolio tracker |
| 132 | `132-savings-runway-calculator` | Savings runway calculator |
| 133 | `133-portfolio-website` | Personal portfolio website |
| 134 | `134-professional` | Professional Go project + Bitcoin OSS contribution path |
| 135 | `135-python-data` | Python data analysis course |
| 136 | `136-api-lesson` | API design lessons |
| 137 | `137-capstone` | Main capstone project |
| 138 | `138-js-database` | JavaScript + database course |
| 139 | `139-js-frontend` | JavaScript frontend course |

---

## Quick Reference: What's in Each Phase

```
01-05   Print one character   →  Hello World muscle memory
06-12   Foundation lessons    →  Functions, loops, strings pkg, slices, maps, arrays, recursion
13-36   Beginner practice     →  One skill per exercise, guided by README + skills.md
37-65   Strings mastery       →  Every strings/fmt/strconv function with examples
66-129  Challenge problems    →  Hard piscine-style challenges, combine multiple concepts
130-139 Capstone projects     →  Full applications
```

**Always read `skills.md` before you attempt the challenge in `README.md`.**
**Use `prerequisites.md` if you're stuck — it tells you exactly what to review.**
