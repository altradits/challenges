# Go Mastery: Complete Learning Path

Follow the directories **in number order (01 → 150)**. Each number is exactly where you should be.

```
Phase 1  (01–05):  Hello World — the four-line Go skeleton
Phase 2  (06–27):  Go Foundations — every core language concept
Phase 3  (28–51):  Beginner Practice — apply the basics one concept at a time
Phase 4  (52–80):  Strings Package Mastery — the full strings/fmt/strconv toolkit
Phase 5  (81–144): Challenge Problems — combine everything, harder piscine-style problems
Phase 6  (145–150): Capstone Projects — real-world Go applications + Bitcoin open source
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

## Phase 2: Go Language Foundations (06–27)

One concept per lesson. Every concept here is needed to write production Go.
Study each `skills.md` before doing the challenge.

| # | Directory | Concept Introduced |
|---|-----------|-------------------|
| 06 | `06-functions` | Function syntax, parameters, return types, `package piscine` |
| 07 | `07-forloops` | All for loop patterns: index, while-style, range, break/continue |
| 08 | `08-stringspackage` | The `strings` package: Fields, Contains, Builder, Split, etc. |
| 09 | `09-slicesintro` | Slices: `make`, `append`, sub-slicing, range, nil vs empty |
| 10 | `10-mapsintro` | Maps: `make`, key-value, zero value trick, comma-ok, sets |
| 11 | `11-arrays` | Fixed-size arrays: `[N]T`, value semantics, arrays vs slices |
| 12 | `12-recursion` | Recursion: base case, call stack trace, factorial, GCD |
| 13 | `13-structs` | Struct types, field access, constructors, value semantics |
| 14 | `14-pointers` | `&`, `*`, mutation through pointers, nil, pointer-to-struct |
| 15 | `15-methods` | Value vs pointer receivers, method sets, when to use each |
| 16 | `16-interfaces` | Interface definition, implicit satisfaction, `Stringer`, `io.Reader` |
| 17 | `17-errorhandling` | `error` interface, `errors.New`, `fmt.Errorf %w`, `errors.Is/As`, `panic/recover` |
| 18 | `18-typeassertions` | Type assertions `.(T)`, type switches, `any` |
| 19 | `19-fileio` | `os.ReadFile/WriteFile`, `os.Open`, `defer`, `bufio.Scanner`, `io.Reader/Writer` |
| 20 | `20-regexp` | `regexp.MustCompile`, `MatchString`, `FindAllString`, `ReplaceAllString` |
| 21 | `21-sort-and-math` | `sort.Ints/Strings/Slice`, `sort.Search`, `math.Sqrt/Pow/Abs/Pi` |
| 22 | `22-modules-and-tooling` | `go mod init/get/tidy`, `go fmt/vet/doc/build`, `golangci-lint`, `init()`, `_` |
| 23 | `23-testing-basics` | `*_test.go`, `testing.T`, table-driven tests, `t.Run`, benchmarks |
| 24 | `24-goroutines` | `go` keyword, `sync.WaitGroup`, loop variable capture, `-race` flag |
| 25 | `25-channels` | Unbuffered/buffered channels, `close`, `range`, pipeline pattern |
| 26 | `26-select-and-sync` | `select`, `time.After`, `sync.Mutex`, `sync.RWMutex`, `sync.Once` |
| 27 | `27-context` | `context.WithCancel/Timeout/Value`, `ctx.Done()`, cancellation propagation |

---

## Phase 3: Beginner Practice (28–51)

Short focused exercises. One skill per challenge.
Read `skills.md` first, then solve the challenge in `README.md`.

| # | Directory | Skill |
|---|-----------|-------|
| 28 | `28-stringlength` | `len()` and manual character counting |
| 29 | `29-firstchar` | First character: `s[0]`, `string(s[0])` |
| 30 | `30-lastchar` | Last character: `s[len(s)-1]` |
| 31 | `31-isempty` | Empty string check: `s == ""` or `len(s) == 0` |
| 32 | `32-toupper` | `strings.ToUpper` |
| 33 | `33-tolower` | `strings.ToLower` |
| 34 | `34-countalpha` | `unicode.IsLetter`, counting with for...range |
| 35 | `35-checknumber` | `unicode.IsDigit`, detecting digit characters |
| 36 | `36-countvowels` | `strings.ContainsRune`, vowel detection |
| 37 | `37-reversestring` | Reverse a string using `[]rune` |
| 38 | `38-ispalindrome` | Palindrome check: reverse + compare |
| 39 | `39-removespaces` | `strings.ReplaceAll`, filtering characters |
| 40 | `40-countrepeats` | `map[rune]int` for character frequency |
| 41 | `41-retainfirsthalf` | String slicing: `s[:len(s)/2]` |
| 42 | `42-wordcount` | `strings.Fields` for word splitting |
| 43 | `43-findchar` | `strings.Index`, manual loop search |
| 44 | `44-countchar` | `strings.Count` for single character |
| 45 | `45-findlastchar` | `strings.LastIndex` |
| 46 | `46-replacechar` | `strings.ReplaceAll` for character replacement |
| 47 | `47-printif` | `if/else`, `len()`, `||` operator |
| 48 | `48-printifnot` | Inverted conditions, De Morgan's Law |
| 49 | `49-longestword` | Running maximum with `strings.Fields` |
| 50 | `50-searchreplace` | `strings.Index` + slice concatenation |
| 51 | `51-cleanlist` | `map[string]bool` as a set for deduplication |

---

## Phase 4: Strings Package Mastery (52–80)

Deep dive into every function in the `strings`, `fmt`, and `strconv` packages.

| # | Directory | Function / Concept |
|---|-----------|-------------------|
| 52 | `52-countwords` | Advanced word counting with `strings.Fields` |
| 53 | `53-findsubstring` | `strings.Index`, sliding window |
| 54 | `54-replaceall` | `strings.ReplaceAll` |
| 55 | `55-split` | `strings.Split` |
| 56 | `56-join` | `strings.Join` |
| 57 | `57-cameltosnakecase` | `unicode`, `strings.Builder`, case detection |
| 58 | `58-itoa` | `strconv.Itoa`, `strconv.Atoi` |
| 59 | `59-thirdchar` | Index step patterns |
| 60 | `60-zipstring` | Run-length encoding |
| 61 | `61-saveandmiss` | Periodic skip pattern |
| 62 | `62-reversestrcap` | Reverse + case toggle |
| 63 | `63-union` | Set union with `map[rune]bool` |
| 64 | `64-inter` | Set intersection |
| 65 | `65-stringbuilder` | `strings.Builder` deep dive |
| 66 | `66-stringsplit` | `strings.Split` variants |
| 67 | `67-stringjoin` | `strings.Join` |
| 68 | `68-stringrepeat` | `strings.Repeat` |
| 69 | `69-stringreplace` | `strings.Replace` (count-limited) |
| 70 | `70-stringtrim` | `strings.TrimSpace`, `Trim`, `TrimPrefix`, `TrimSuffix` |
| 71 | `71-stringcontains` | `strings.Contains` |
| 72 | `72-stringindex` | `strings.Index` detail |
| 73 | `73-stringcount` | `strings.Count` |
| 74 | `74-stringprefix` | `strings.HasPrefix`, `TrimPrefix` |
| 75 | `75-stringsuffix` | `strings.HasSuffix`, `TrimSuffix` |
| 76 | `76-stringfield` | `strings.Fields` vs `Split` vs `FieldsFunc` |
| 77 | `77-stringmap` | `strings.Map`, higher-order `func(rune) rune` |
| 78 | `78-stringfilter` | Filter pattern with `strings.Builder` |
| 79 | `79-stringreduce` | Fold / accumulator pattern over runes |
| 80 | `80-stringformat` | `fmt.Sprintf` format verbs: `%s %d %f %v %T %q %x` |

---

## Phase 5: Challenge Problems (81–144)

The original piscine challenges. You now have all the tools — apply them.
Harder than Phase 3/4. Most require combining multiple concepts.

| # | Directory | Challenge |
|---|-----------|-----------|
| 81 | `81-checknumber` | Check if a string contains any digit |
| 82 | `82-count-character` | Count occurrences of a specific character |
| 83 | `83-missing-pieces` | Supplementary concept reference (len vs index, etc.) |
| 84 | `84-findsubstring` | Find substring index manually |
| 85 | `85-printif` | Conditional print with length logic |
| 86 | `86-printifnot` | Inverted condition print |
| 87 | `87-removespaces` | Remove all spaces from a string |
| 88 | `88-retainfirsthalf` | Keep only the first half of a string |
| 89 | `89-reversestring` | Reverse a string |
| 90 | `90-titlecase` | Convert string to Title Case |
| 91 | `91-wordcount` | Count words in a string |
| 92 | `92-cameltosnakecase` | Convert camelCase to snake_case |
| 93 | `93-countrepeats` | Count characters that appear more than once |
| 94 | `94-digitlen` | Count how many digits are in a string |
| 95 | `95-firstword` | Extract the first word |
| 96 | `96-gcd` | Greatest Common Divisor (recursive Euclidean algorithm) |
| 97 | `97-hashcode` | Generate a simple hash code from a string |
| 98 | `98-lastword` | Extract the last word |
| 99 | `99-longestword` | Find the longest word |
| 100 | `100-replaceall` | Replace all occurrences of a substring |
| 101 | `101-searchreplace` | Search and replace the first occurrence |
| 102 | `102-splitjoin` | Split a string then join with a different separator |
| 103 | `103-wordanatomy` | Analyze prefix/suffix structure of words |
| 104 | `104-wordanatomy2` | Advanced word anatomy with custom prefix/suffix lists |
| 105 | `105-clean-the-list` | Normalize a grocery list (trim, capitalize, ensure "milk") |
| 106 | `106-cleanstr` | Remove extra whitespace between words |
| 107 | `107-expandstr` | Add spaces between each character |
| 108 | `108-findprevprime` | Find the largest prime ≤ n |
| 109 | `109-fromto` | Generate a range string like "1→5" |
| 110 | `110-iscapitalized` | Check if every word starts with a capital letter |
| 111 | `111-itoa-manual` | Implement integer-to-string without `strconv` |
| 112 | `112-printmemory` | Print `[10]byte` array as hex + ASCII (like `xxd`) |
| 113 | `113-printrevcomb` | Print reversed combinations of digits |
| 114 | `114-thirdtimeisacharm` | Extract every third character |
| 115 | `115-weareunique` | Count unique characters in a string |
| 116 | `116-zipstring` | Run-length encode: "aaa" → "a3" |
| 117 | `117-addprimesum` | Sum all primes up to n |
| 118 | `118-canjump` | Jump game — can you reach the last index? |
| 119 | `119-chunk` | Split a slice into chunks of size k |
| 120 | `120-concatalternate` | Interleave characters from two strings |
| 121 | `121-concatslice` | Concatenate two slices |
| 122 | `122-fprime` | Print prime factors of n |
| 123 | `123-hiddenp` | Check if string A is hidden inside concatenation of string B with itself |
| 124 | `124-inter` | Intersection of two strings (unique shared chars) |
| 125 | `125-reversestrcap` | Reverse + toggle case of each character |
| 126 | `126-saveandmiss` | Keep every 2nd character, skip every 3rd |
| 127 | `127-union` | Union of two strings (sorted unique chars) |
| 128 | `128-wdmatch` | Check if string A is a scattered subsequence of B |
| 129 | `129-fifthandskip` | Extract every 5th character starting from the 5th |
| 130 | `130-notdecimal` | Check if a string is NOT a decimal number |
| 131 | `131-revconcatalternate` | Reverse interleave characters |
| 132 | `132-slice` | Implement slice operations (remove, insert) |
| 133 | `133-findpairs` | Find all pairs in a slice summing to a target |
| 134 | `134-revwstr` | Reverse the words in a sentence |
| 135 | `135-rostring` | Rotate string by n positions |
| 136 | `136-wordflip` | Reverse each word in a sentence |
| 137 | `137-itoabase` | Convert integer to string in any base (2–36) |
| 138 | `138-options` | Parse command-line flags like `--flag=value` |
| 139 | `139-piglatin` | Convert words to Pig Latin |
| 140 | `140-romannumbers` | Convert integer to Roman numerals |
| 141 | `141-brackets` | Validate bracket balance with a stack |
| 142 | `142-rpncalc` | Reverse Polish Notation calculator |
| 143 | `143-brainfuck` | Brainfuck interpreter |
| 144 | `144-grouping` | Regex-like group matching with `|` alternation |

---

## Phase 6: Capstone Projects (145–150)

Apply everything in real Go projects. Each one builds on all prior phases.

| # | Directory | Project |
|---|-----------|---------|
| 145 | `145-financial-freedom-api` | REST API: financial freedom calculator |
| 146 | `146-index-fund-tracker` | Index fund portfolio tracker |
| 147 | `147-savings-runway-calculator` | Savings runway calculator |
| 148 | `148-api-lesson` | Go API development: HTTP, routing, REST, JSON, auth, testing (10 exercises) |
| 149 | `149-capstone` | String tools analytics dashboard — integrates all your tools |
| 150 | `150-bitcoin-oss` | Your first merged PR in btcd or LND (Go Bitcoin open source) |

---

## Quick Reference

```
01-05   Hello World          →  The four-line Go skeleton
06-12   Core language        →  Functions, loops, strings, slices, maps, arrays, recursion
13-27   Object model         →  Structs, pointers, methods, interfaces, errors, types,
                                file I/O, regexp, sort/math, modules/tooling,
                                testing, goroutines, channels, select/sync, context
28-51   Beginner practice    →  One skill per exercise — build speed and confidence
52-80   Strings mastery      →  Every strings/fmt/strconv function with examples
81-144  Challenge problems   →  Hard piscine-style challenges, combine multiple concepts
145-150 Capstone projects    →  Real Go applications and Bitcoin open source contribution
```

**Always read `skills.md` before you attempt the challenge in `README.md`.**
**Use `prerequisites.md` if you're stuck — it tells you exactly what to review.**
